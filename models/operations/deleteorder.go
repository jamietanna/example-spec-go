// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/jamietanna/speakeasy-example-spec/models/components"
)

type DeleteOrderRequest struct {
	// ID of the order that needs to be deleted
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderId"`
}

func (o *DeleteOrderRequest) GetOrderID() int64 {
	if o == nil {
		return 0
	}
	return o.OrderID
}

type DeleteOrderResponse struct {
	HTTPMeta components.HTTPMetadata
	// Order deleted
	Order *components.Order
}

func (o *DeleteOrderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteOrderResponse) GetOrder() *components.Order {
	if o == nil {
		return nil
	}
	return o.Order
}
