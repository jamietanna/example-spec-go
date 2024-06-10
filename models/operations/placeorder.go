// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/jamietanna/speakeasy-example-spec/models/components"
)

type PlaceOrderResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successful operation
	Order *components.Order
}

func (o *PlaceOrderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PlaceOrderResponse) GetOrder() *components.Order {
	if o == nil {
		return nil
	}
	return o.Order
}
