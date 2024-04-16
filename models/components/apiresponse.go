// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type APIResponse struct {
	Code    *int    `json:"code,omitempty"`
	Type    *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (o *APIResponse) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *APIResponse) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *APIResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
