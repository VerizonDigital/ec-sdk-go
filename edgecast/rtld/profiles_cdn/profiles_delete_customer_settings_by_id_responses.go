// Code generated by go-swagger; DO NOT EDIT.

package profiles_cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProfilesDeleteCustomerSettingsByIDReader is a Reader for the ProfilesDeleteCustomerSettingsByID structure.
type ProfilesDeleteCustomerSettingsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilesDeleteCustomerSettingsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProfilesDeleteCustomerSettingsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProfilesDeleteCustomerSettingsByIDNoContent creates a ProfilesDeleteCustomerSettingsByIDNoContent with default headers values
func NewProfilesDeleteCustomerSettingsByIDNoContent() *ProfilesDeleteCustomerSettingsByIDNoContent {
	return &ProfilesDeleteCustomerSettingsByIDNoContent{}
}

/* ProfilesDeleteCustomerSettingsByIDNoContent describes a response with status code 204, with default header values.

Success
*/
type ProfilesDeleteCustomerSettingsByIDNoContent struct {
}

func (o *ProfilesDeleteCustomerSettingsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/cdn/profiles/{id}][%d] profilesDeleteCustomerSettingsByIdNoContent ", 204)
}

func (o *ProfilesDeleteCustomerSettingsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}