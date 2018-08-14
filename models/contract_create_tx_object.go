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
)

// ContractCreateTxObject contract create tx object
// swagger:model ContractCreateTxObject
type ContractCreateTxObject struct {
	vsnField int64

	ContractCreateData
}

// DataSchema gets the data schema of this subtype
func (m *ContractCreateTxObject) DataSchema() string {
	return "ContractCreateTxObject"
}

// SetDataSchema sets the data schema of this subtype
func (m *ContractCreateTxObject) SetDataSchema(val string) {

}

// Vsn gets the vsn of this subtype
func (m *ContractCreateTxObject) Vsn() int64 {
	return m.vsnField
}

// SetVsn sets the vsn of this subtype
func (m *ContractCreateTxObject) SetVsn(val int64) {
	m.vsnField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ContractCreateTxObject) UnmarshalJSON(raw []byte) error {
	var data struct {
		ContractCreateData
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DataSchema string `json:"data_schema"`

		Vsn int64 `json:"vsn,omitempty"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ContractCreateTxObject

	if base.DataSchema != result.DataSchema() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid data_schema value: %q", base.DataSchema)
	}

	result.vsnField = base.Vsn

	result.ContractCreateData = data.ContractCreateData

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ContractCreateTxObject) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		ContractCreateData
	}{

		ContractCreateData: m.ContractCreateData,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DataSchema string `json:"data_schema"`

		Vsn int64 `json:"vsn,omitempty"`
	}{

		DataSchema: m.DataSchema(),

		Vsn: m.Vsn(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this contract create tx object
func (m *ContractCreateTxObject) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ContractCreateData
	if err := m.ContractCreateData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ContractCreateTxObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractCreateTxObject) UnmarshalBinary(b []byte) error {
	var res ContractCreateTxObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
