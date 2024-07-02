// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/jamietanna/speakeasy-example-spec/models/components"
)

type GetOrderByIDRequest struct {
	// ID of order that needs to be fetched
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderId"`
}

func (o *GetOrderByIDRequest) GetOrderID() int64 {
	if o == nil {
		return 0
	}
	return o.OrderID
}

type GetOrderByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successful operation
	Order *components.Order
}

func (o *GetOrderByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetOrderByIDResponse) GetOrder() *components.Order {
	if o == nil {
		return nil
	}
	return o.Order
}
