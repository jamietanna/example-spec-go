// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// APIErrorNotFound - Not Found error
type APIErrorNotFound struct {
	Status int    `json:"status"`
	Error_ string `json:"error"`
	Code   string `json:"code"`
}

var _ error = &APIErrorNotFound{}

func (e *APIErrorNotFound) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}