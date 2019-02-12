// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContractCallTxJSON contract call tx JSON
// swagger:model ContractCallTxJSON
type ContractCallTxJSON struct {
	versionField *uint64

	ContractCallTx
}

// Type gets the type of this subtype
func (m *ContractCallTxJSON) Type() string {
	return "ContractCallTxJSON"
}

// SetType sets the type of this subtype
func (m *ContractCallTxJSON) SetType(val string) {

}

// Version gets the version of this subtype
func (m *ContractCallTxJSON) Version() *uint64 {
	return m.versionField
}

// SetVersion sets the version of this subtype
func (m *ContractCallTxJSON) SetVersion(val *uint64) {
	m.versionField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ContractCallTxJSON) UnmarshalJSON(raw []byte) error {
	var data struct {
		ContractCallTx
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`

		Version *uint64 `json:"version"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ContractCallTxJSON

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.versionField = base.Version

	result.ContractCallTx = data.ContractCallTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ContractCallTxJSON) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		ContractCallTx
	}{

		ContractCallTx: m.ContractCallTx,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`

		Version *uint64 `json:"version"`
	}{

		Type: m.Type(),

		Version: m.Version(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this contract call tx JSON
func (m *ContractCallTxJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with ContractCallTx
	if err := m.ContractCallTx.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContractCallTxJSON) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractCallTxJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractCallTxJSON) UnmarshalBinary(b []byte) error {
	var res ContractCallTxJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
