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

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// OffChainCallContract off chain call contract
// swagger:model OffChainCallContract
type OffChainCallContract struct {

	// abi version
	// Required: true
	AbiVersion Uint16 `json:"abi_version"`

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// Contract call data
	// Required: true
	CallData EncodedByteArray `json:"call_data"`

	// Contract caller
	// Required: true
	Caller EncodedPubkey `json:"caller"`

	// Contract address
	// Required: true
	Contract EncodedPubkey `json:"contract"`

	// gas
	// Required: true
	Gas Uint64 `json:"gas"`

	// gas price
	// Required: true
	GasPrice utils.BigInt `json:"gas_price"`
}

// Op gets the op of this subtype
func (m *OffChainCallContract) Op() string {
	return "OffChainCallContract"
}

// SetOp sets the op of this subtype
func (m *OffChainCallContract) SetOp(val string) {

}

// AbiVersion gets the abi version of this subtype

// Amount gets the amount of this subtype

// CallData gets the call data of this subtype

// Caller gets the caller of this subtype

// Contract gets the contract of this subtype

// Gas gets the gas of this subtype

// GasPrice gets the gas price of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OffChainCallContract) UnmarshalJSON(raw []byte) error {
	var data struct {

		// abi version
		// Required: true
		AbiVersion Uint16 `json:"abi_version"`

		// amount
		// Required: true
		Amount utils.BigInt `json:"amount"`

		// Contract call data
		// Required: true
		CallData EncodedByteArray `json:"call_data"`

		// Contract caller
		// Required: true
		Caller EncodedPubkey `json:"caller"`

		// Contract address
		// Required: true
		Contract EncodedPubkey `json:"contract"`

		// gas
		// Required: true
		Gas Uint64 `json:"gas"`

		// gas price
		// Required: true
		GasPrice utils.BigInt `json:"gas_price"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Op string `json:"op"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result OffChainCallContract

	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.AbiVersion = data.AbiVersion

	result.Amount = data.Amount

	result.CallData = data.CallData

	result.Caller = data.Caller

	result.Contract = data.Contract

	result.Gas = data.Gas

	result.GasPrice = data.GasPrice

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OffChainCallContract) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// abi version
		// Required: true
		AbiVersion Uint16 `json:"abi_version"`

		// amount
		// Required: true
		Amount utils.BigInt `json:"amount"`

		// Contract call data
		// Required: true
		CallData EncodedByteArray `json:"call_data"`

		// Contract caller
		// Required: true
		Caller EncodedPubkey `json:"caller"`

		// Contract address
		// Required: true
		Contract EncodedPubkey `json:"contract"`

		// gas
		// Required: true
		Gas Uint64 `json:"gas"`

		// gas price
		// Required: true
		GasPrice utils.BigInt `json:"gas_price"`
	}{

		AbiVersion: m.AbiVersion,

		Amount: m.Amount,

		CallData: m.CallData,

		Caller: m.Caller,

		Contract: m.Contract,

		Gas: m.Gas,

		GasPrice: m.GasPrice,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Op string `json:"op"`
	}{

		Op: m.Op(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this off chain call contract
func (m *OffChainCallContract) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaller(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OffChainCallContract) validateAbiVersion(formats strfmt.Registry) error {

	if err := m.AbiVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("abi_version")
		}
		return err
	}

	return nil
}

func (m *OffChainCallContract) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *OffChainCallContract) validateCallData(formats strfmt.Registry) error {

	if err := m.CallData.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("call_data")
		}
		return err
	}

	return nil
}

func (m *OffChainCallContract) validateCaller(formats strfmt.Registry) error {

	if err := m.Caller.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("caller")
		}
		return err
	}

	return nil
}

func (m *OffChainCallContract) validateContract(formats strfmt.Registry) error {

	if err := m.Contract.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("contract")
		}
		return err
	}

	return nil
}

func (m *OffChainCallContract) validateGas(formats strfmt.Registry) error {

	if err := m.Gas.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas")
		}
		return err
	}

	return nil
}

func (m *OffChainCallContract) validateGasPrice(formats strfmt.Registry) error {

	if err := m.GasPrice.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("gas_price")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OffChainCallContract) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffChainCallContract) UnmarshalBinary(b []byte) error {
	var res OffChainCallContract
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
