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

// GenericBlock generic block
// swagger:model GenericBlock
type GenericBlock struct {
	Header

	// data schema
	// Required: true
	DataSchema string `json:"data_schema"`

	// hash
	Hash EncodedHash `json:"hash,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GenericBlock) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Header
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Header = aO0

	// AO1
	var dataAO1 struct {
		DataSchema string `json:"data_schema"`

		Hash EncodedHash `json:"hash,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DataSchema = dataAO1.DataSchema

	m.Hash = dataAO1.Hash

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GenericBlock) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Header)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DataSchema string `json:"data_schema"`

		Hash EncodedHash `json:"hash,omitempty"`
	}

	dataAO1.DataSchema = m.DataSchema

	dataAO1.Hash = m.Hash

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this generic block
func (m *GenericBlock) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Header
	if err := m.Header.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericBlock) validateDataSchema(formats strfmt.Registry) error {

	if err := validate.RequiredString("data_schema", "body", string(m.DataSchema)); err != nil {
		return err
	}

	return nil
}

func (m *GenericBlock) validateHash(formats strfmt.Registry) error {

	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if err := m.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hash")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericBlock) UnmarshalBinary(b []byte) error {
	var res GenericBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
