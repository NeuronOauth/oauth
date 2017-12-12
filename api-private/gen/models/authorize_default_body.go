// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuthorizeDefaultBody authorize default body
// swagger:model authorizeDefaultBody
type AuthorizeDefaultBody struct {

	// Error code
	Code string `json:"code,omitempty"`

	// errors
	Errors AuthorizeDefaultBodyErrors `json:"errors"`

	// Error message
	Message string `json:"message,omitempty"`

	// status
	Status *int32 `json:"status,omitempty"`
}

// Validate validates this authorize default body
func (m *AuthorizeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizeDefaultBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizeDefaultBody) UnmarshalBinary(b []byte) error {
	var res AuthorizeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
