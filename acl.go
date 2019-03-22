// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ACL ACL Rules
//
// The use of Access Control Lists (ACL) provides a flexible solution to perform
// content switching and generally to take decisions based on content extracted
// from the request, the response or any environmental status.
//
// swagger:model acl
type ACL struct {

	// acl name
	// Pattern: ^[^\s]+$
	ACLName string `json:"acl_name,omitempty"`

	// criterion
	// Pattern: ^[^\s]+$
	Criterion string `json:"criterion,omitempty"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this acl
func (m *ACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACLName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriterion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ACL) validateACLName(formats strfmt.Registry) error {

	if swag.IsZero(m.ACLName) { // not required
		return nil
	}

	if err := validate.Pattern("acl_name", "body", string(m.ACLName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ACL) validateCriterion(formats strfmt.Registry) error {

	if swag.IsZero(m.Criterion) { // not required
		return nil
	}

	if err := validate.Pattern("criterion", "body", string(m.Criterion), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ACL) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ACL) UnmarshalBinary(b []byte) error {
	var res ACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
