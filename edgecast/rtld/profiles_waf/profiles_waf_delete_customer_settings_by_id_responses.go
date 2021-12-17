// Code generated by go-swagger; DO NOT EDIT.

package profiles_waf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProfilesWafDeleteCustomerSettingsByIDReader is a Reader for the ProfilesWafDeleteCustomerSettingsByID structure.
type ProfilesWafDeleteCustomerSettingsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilesWafDeleteCustomerSettingsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProfilesWafDeleteCustomerSettingsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProfilesWafDeleteCustomerSettingsByIDNoContent creates a ProfilesWafDeleteCustomerSettingsByIDNoContent with default headers values
func NewProfilesWafDeleteCustomerSettingsByIDNoContent() *ProfilesWafDeleteCustomerSettingsByIDNoContent {
	return &ProfilesWafDeleteCustomerSettingsByIDNoContent{}
}

/* ProfilesWafDeleteCustomerSettingsByIDNoContent describes a response with status code 204, with default header values.

Success
*/
type ProfilesWafDeleteCustomerSettingsByIDNoContent struct {
}

func (o *ProfilesWafDeleteCustomerSettingsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/waf/profiles/{id}][%d] profilesWafDeleteCustomerSettingsByIdNoContent ", 204)
}

func (o *ProfilesWafDeleteCustomerSettingsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}