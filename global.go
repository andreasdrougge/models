// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Global Global
//
// HAProxy global configuration
// swagger:model global
type Global struct {

	// daemon
	// Enum: [enabled disabled]
	Daemon string `json:"daemon,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// nbproc
	Nbproc int64 `json:"nbproc,omitempty"`

	// runtime apis
	RuntimeApis []*GlobalRuntimeApisItems `json:"runtime_apis"`

	// ssl default bind ciphers
	SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers,omitempty"`

	// ssl default bind options
	SslDefaultBindOptions string `json:"ssl_default_bind_options,omitempty"`

	// tune ssl default dh param
	TuneSslDefaultDhParam int64 `json:"tune_ssl_default_dh_param,omitempty"`
}

// Validate validates this global
func (m *Global) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaemon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimeApis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var globalTypeDaemonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalTypeDaemonPropEnum = append(globalTypeDaemonPropEnum, v)
	}
}

const (

	// GlobalDaemonEnabled captures enum value "enabled"
	GlobalDaemonEnabled string = "enabled"

	// GlobalDaemonDisabled captures enum value "disabled"
	GlobalDaemonDisabled string = "disabled"
)

// prop value enum
func (m *Global) validateDaemonEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, globalTypeDaemonPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Global) validateDaemon(formats strfmt.Registry) error {

	if swag.IsZero(m.Daemon) { // not required
		return nil
	}

	// value enum
	if err := m.validateDaemonEnum("daemon", "body", m.Daemon); err != nil {
		return err
	}

	return nil
}

func (m *Global) validateRuntimeApis(formats strfmt.Registry) error {

	if swag.IsZero(m.RuntimeApis) { // not required
		return nil
	}

	for i := 0; i < len(m.RuntimeApis); i++ {
		if swag.IsZero(m.RuntimeApis[i]) { // not required
			continue
		}

		if m.RuntimeApis[i] != nil {
			if err := m.RuntimeApis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runtime_apis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Global) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Global) UnmarshalBinary(b []byte) error {
	var res Global
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
