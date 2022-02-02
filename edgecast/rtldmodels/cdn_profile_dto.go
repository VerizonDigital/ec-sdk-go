// Code generated by go-swagger; DO NOT EDIT.

package rtldmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdnProfileDto cdn profile dto
//
// swagger:model CdnProfileDto
type CdnProfileDto struct {
	BaseProfileDto

	// custom cookies
	CustomCookies []string `json:"custom_cookies"`

	// custom request headers
	CustomRequestHeaders []string `json:"custom_request_headers"`

	// custom response headers
	CustomResponseHeaders []string `json:"custom_response_headers"`

	// filters
	Filters *RtldFiltersDto `json:"filters,omitempty"`

	// platforms
	Platforms []string `json:"platforms"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CdnProfileDto) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseProfileDto
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseProfileDto = aO0

	// AO1
	var dataAO1 struct {
		CustomCookies []string `json:"custom_cookies"`

		CustomRequestHeaders []string `json:"custom_request_headers"`

		CustomResponseHeaders []string `json:"custom_response_headers"`

		Filters *RtldFiltersDto `json:"filters,omitempty"`

		Platforms []string `json:"platforms"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CustomCookies = dataAO1.CustomCookies

	m.CustomRequestHeaders = dataAO1.CustomRequestHeaders

	m.CustomResponseHeaders = dataAO1.CustomResponseHeaders

	m.Filters = dataAO1.Filters

	m.Platforms = dataAO1.Platforms

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CdnProfileDto) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseProfileDto)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		CustomCookies []string `json:"custom_cookies"`

		CustomRequestHeaders []string `json:"custom_request_headers"`

		CustomResponseHeaders []string `json:"custom_response_headers"`

		Filters *RtldFiltersDto `json:"filters,omitempty"`

		Platforms []string `json:"platforms"`
	}

	dataAO1.CustomCookies = m.CustomCookies

	dataAO1.CustomRequestHeaders = m.CustomRequestHeaders

	dataAO1.CustomResponseHeaders = m.CustomResponseHeaders

	dataAO1.Filters = m.Filters

	dataAO1.Platforms = m.Platforms

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this cdn profile dto
func (m *CdnProfileDto) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseProfileDto
	if err := m.BaseProfileDto.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnProfileDto) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {
		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cdn profile dto based on the context it is used
func (m *CdnProfileDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseProfileDto
	if err := m.BaseProfileDto.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnProfileDto) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if m.Filters != nil {
		if err := m.Filters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnProfileDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnProfileDto) UnmarshalBinary(b []byte) error {
	var res CdnProfileDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
