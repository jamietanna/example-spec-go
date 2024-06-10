// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/jamietanna/speakeasy-example-spec/models/components"
)

type UpdatePetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful operation
	Pet *components.Pet
}

func (o *UpdatePetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdatePetResponse) GetPet() *components.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}
