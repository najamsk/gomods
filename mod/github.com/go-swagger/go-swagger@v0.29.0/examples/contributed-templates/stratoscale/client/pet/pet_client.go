// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the pet client
type API interface {
	/*
	   PetCreate adds a new pet to the store*/
	PetCreate(ctx context.Context, params *PetCreateParams) (*PetCreateCreated, error)
	/*
	   PetDelete deletes a pet*/
	PetDelete(ctx context.Context, params *PetDeleteParams) (*PetDeleteNoContent, error)
	/*
	   PetGet gets pet by it s ID*/
	PetGet(ctx context.Context, params *PetGetParams) (*PetGetOK, error)
	/*
	   PetList lists pets*/
	PetList(ctx context.Context, params *PetListParams) (*PetListOK, error)
	/*
	   PetUpdate updates an existing pet*/
	PetUpdate(ctx context.Context, params *PetUpdateParams) (*PetUpdateCreated, error)
	/*
	   PetUploadImage uploads an image*/
	PetUploadImage(ctx context.Context, params *PetUploadImageParams) (*PetUploadImageOK, error)
}

// New creates a new pet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for pet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
PetCreate adds a new pet to the store
*/
func (a *Client) PetCreate(ctx context.Context, params *PetCreateParams) (*PetCreateCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PetCreate",
		Method:             "POST",
		PathPattern:        "/pet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PetCreateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PetCreateCreated), nil

}

/*
PetDelete deletes a pet
*/
func (a *Client) PetDelete(ctx context.Context, params *PetDeleteParams) (*PetDeleteNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PetDelete",
		Method:             "DELETE",
		PathPattern:        "/pet/{petId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PetDeleteReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PetDeleteNoContent), nil

}

/*
PetGet gets pet by it s ID
*/
func (a *Client) PetGet(ctx context.Context, params *PetGetParams) (*PetGetOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PetGet",
		Method:             "GET",
		PathPattern:        "/pet/{petId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PetGetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PetGetOK), nil

}

/*
PetList lists pets
*/
func (a *Client) PetList(ctx context.Context, params *PetListParams) (*PetListOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PetList",
		Method:             "GET",
		PathPattern:        "/pet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PetListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PetListOK), nil

}

/*
PetUpdate updates an existing pet
*/
func (a *Client) PetUpdate(ctx context.Context, params *PetUpdateParams) (*PetUpdateCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PetUpdate",
		Method:             "PUT",
		PathPattern:        "/pet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PetUpdateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PetUpdateCreated), nil

}

/*
PetUploadImage uploads an image
*/
func (a *Client) PetUploadImage(ctx context.Context, params *PetUploadImageParams) (*PetUploadImageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PetUploadImage",
		Method:             "POST",
		PathPattern:        "/pet/{petId}/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PetUploadImageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PetUploadImageOK), nil

}