// Copyright 2021 DeepMap, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
)

const GinContextKey = "oapi-codegen/gin-context"
const UserDataKey = "oapi-codegen/user-data"

// Create validator middleware from a YAML file path
func OapiValidatorFromYamlFile(path string) (gin.HandlerFunc, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %s", path, err)
	}

	swagger, err := openapi3.NewLoader().LoadFromData(data)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s as Swagger YAML: %s",
			path, err)
	}
	return OapiRequestValidator(swagger), nil
}

// This is an gin middleware function which validates incoming HTTP requests
// to make sure that they conform to the given OAPI 3.0 specification. When
// OAPI validation fails on the request, we return an HTTP/400 with error message
func OapiRequestValidator(swagger *openapi3.T) gin.HandlerFunc {
	return OapiRequestValidatorWithOptions(swagger, nil)
}

// Options to customize request validation. These are passed through to
// openapi3filter.
type Options struct {
	Options      openapi3filter.Options
	ParamDecoder openapi3filter.ContentParameterDecoder
	UserData     interface{}
}

// Create a validator from a swagger object, with validation options
func OapiRequestValidatorWithOptions(swagger *openapi3.T, options *Options) gin.HandlerFunc {
	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		err := ValidateRequestFromContext(c, router, options)
		if err != nil {
			// note: i am not sure if this is the best way to handle this
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Next()
	}
}

// ValidateRequestFromContext is called from the middleware above and actually does the work
// of validating a request.
func ValidateRequestFromContext(c *gin.Context, router routers.Router, options *Options) error {
	req := c.Request
	route, pathParams, err := router.FindRoute(req)

	// We failed to find a matching route for the request.
	if err != nil {
		switch e := err.(type) {
		case *routers.RouteError:
			// We've got a bad request, the path requested doesn't match
			// either server, or path, or something.
			return errors.New(e.Reason)
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return fmt.Errorf("error validating route: %s", err.Error())
		}
	}

	validationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}

	// Pass the gin context into the request validator, so that any callbacks
	// which it invokes make it available.
	requestContext := context.WithValue(context.Background(), GinContextKey, c)

	if options != nil {
		validationInput.Options = &options.Options
		validationInput.ParamDecoder = options.ParamDecoder
		requestContext = context.WithValue(requestContext, UserDataKey, options.UserData)
	}

	err = openapi3filter.ValidateRequest(requestContext, validationInput)
	if err != nil {
		switch e := err.(type) {
		case *openapi3filter.RequestError:
			// We've got a bad request
			// Split up the verbose error by lines and return the first one
			// openapi errors seem to be multi-line with a decent message on the first
			errorLines := strings.Split(e.Error(), "\n")
			return fmt.Errorf("error in openapi3filter.RequestError: %s", errorLines[0])
		case *openapi3filter.SecurityRequirementsError:
			return fmt.Errorf("error in openapi3filter.SecurityRequirementsError: %s", e.Error())
		default:
			// This should never happen today, but if our upstream code changes,
			// we don't want to crash the server, so handle the unexpected error.
			return fmt.Errorf("error validating request: %s", err)
		}
	}
	return nil
}

// Helper function to get the echo context from within requests. It returns
// nil if not found or wrong type.
func GetGinContext(c context.Context) *gin.Context {
	iface := c.Value(GinContextKey)
	if iface == nil {
		return nil
	}
	ginCtx, ok := iface.(*gin.Context)
	if !ok {
		return nil
	}
	return ginCtx
}

func GetUserData(c context.Context) interface{} {
	return c.Value(UserDataKey)
}
