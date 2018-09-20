// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetStickRequestRulesOKBody get stick request rules o k body
// swagger:model getStickRequestRulesOKBody
type GetStickRequestRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data StickRequestRules `json:"data"`
}

// Validate validates this get stick request rules o k body
func (m *GetStickRequestRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetStickRequestRulesOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if err := m.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetStickRequestRulesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetStickRequestRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetStickRequestRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
