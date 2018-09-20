// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Listener Listener
//
// HAProxy frontend listener configuration (corresponds to bind directives)
// swagger:model listener
type Listener struct {

	// address
	Address string `json:"address,omitempty"`

	// name
	// Required: true
	Name string `json:"name"`

	// port
	// Maximum: 65535
	// Minimum: 0
	Port *int64 `json:"port,omitempty"`

	// ssl
	// Enum: [enabled]
	Ssl string `json:"ssl,omitempty"`

	// ssl certificate
	// Enum: [default]
	SslCertificate string `json:"ssl_certificate,omitempty"`

	// tcp user timeout
	TCPUserTimeout *int64 `json:"tcp_user_timeout,omitempty"`

	// transparent
	// Enum: [enabled]
	Transparent string `json:"transparent,omitempty"`
}

// Validate validates this listener
func (m *Listener) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransparent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Listener) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Listener) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(*m.Port), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(*m.Port), 65535, false); err != nil {
		return err
	}

	return nil
}

var listenerTypeSslPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerTypeSslPropEnum = append(listenerTypeSslPropEnum, v)
	}
}

const (

	// ListenerSslEnabled captures enum value "enabled"
	ListenerSslEnabled string = "enabled"
)

// prop value enum
func (m *Listener) validateSslEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerTypeSslPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Listener) validateSsl(formats strfmt.Registry) error {

	if swag.IsZero(m.Ssl) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslEnum("ssl", "body", m.Ssl); err != nil {
		return err
	}

	return nil
}

var listenerTypeSslCertificatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerTypeSslCertificatePropEnum = append(listenerTypeSslCertificatePropEnum, v)
	}
}

const (

	// ListenerSslCertificateDefault captures enum value "default"
	ListenerSslCertificateDefault string = "default"
)

// prop value enum
func (m *Listener) validateSslCertificateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerTypeSslCertificatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Listener) validateSslCertificate(formats strfmt.Registry) error {

	if swag.IsZero(m.SslCertificate) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslCertificateEnum("ssl_certificate", "body", m.SslCertificate); err != nil {
		return err
	}

	return nil
}

var listenerTypeTransparentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listenerTypeTransparentPropEnum = append(listenerTypeTransparentPropEnum, v)
	}
}

const (

	// ListenerTransparentEnabled captures enum value "enabled"
	ListenerTransparentEnabled string = "enabled"
)

// prop value enum
func (m *Listener) validateTransparentEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, listenerTypeTransparentPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Listener) validateTransparent(formats strfmt.Registry) error {

	if swag.IsZero(m.Transparent) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransparentEnum("transparent", "body", m.Transparent); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Listener) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Listener) UnmarshalBinary(b []byte) error {
	var res Listener
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
