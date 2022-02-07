// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package talent_test

import (
	"context"

	talent "cloud.google.com/go/talent/apiv4beta1"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
)

func ExampleNewProfileClient() {
	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleProfileClient_ListProfiles() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.ListProfilesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListProfiles(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleProfileClient_CreateProfile() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.CreateProfileRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateProfile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProfileClient_GetProfile() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.GetProfileRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetProfile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProfileClient_UpdateProfile() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.UpdateProfileRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateProfile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProfileClient_DeleteProfile() {
	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.DeleteProfileRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteProfile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProfileClient_SearchProfiles() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewProfileClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.SearchProfilesRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SearchProfiles(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}