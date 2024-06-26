// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/jamietanna/speakeasy-example-spec/models/components"
)

type FindPetsByTagsRequest struct {
	// Tags to filter by
	Tags []string `queryParam:"style=form,explode=true,name=tags"`
}

func (o *FindPetsByTagsRequest) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type FindPetsByTagsResponse struct {
	HTTPMeta components.HTTPMetadata
	// successful operation
	Pets []components.Pet
}

func (o *FindPetsByTagsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *FindPetsByTagsResponse) GetPets() []components.Pet {
	if o == nil {
		return nil
	}
	return o.Pets
}
