// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/jamietanna/speakeasy-example-spec/models/components"
)

type DeletePetRequest struct {
	// Pet id to delete
	PetID  int64   `pathParam:"style=simple,explode=false,name=petId"`
	APIKey *string `header:"style=simple,explode=false,name=api_key"`
}

func (o *DeletePetRequest) GetPetID() int64 {
	if o == nil {
		return 0
	}
	return o.PetID
}

func (o *DeletePetRequest) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

type DeletePetResponse struct {
	HTTPMeta components.HTTPMetadata
	// Pet deleted
	Pet *components.Pet
}

func (o *DeletePetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeletePetResponse) GetPet() *components.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}
