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

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// OffChainNewContract off chain new contract
// swagger:model OffChainNewContract
type OffChainNewContract struct {

	// abi version
	// Required: true
	AbiVersion Uint16 `json:"abi_version"`

	// Contract call data
	// Required: true
	CallData EncodedByteArray `json:"call_data"`

	// code
	// Required: true
	Code *ByteCode `json:"code"`

	// deposit
	// Required: true
	Deposit utils.BigInt `json:"deposit"`

	// Contract owner
	// Required: true
	Owner EncodedPubkey `json:"owner"`

	// vm version
	// Required: true
	VMVersion Uint16 `json:"vm_version"`
}

// Op gets the op of this subtype
func (m *OffChainNewContract) Op() string {
	return "OffChainNewContract"
}

// SetOp sets the op of this subtype
func (m *OffChainNewContract) SetOp(val string) {

}

// AbiVersion gets the abi version of this subtype

// CallData gets the call data of this subtype

// Code gets the code of this subtype

// Deposit gets the deposit of this subtype

// Owner gets the owner of this subtype

// VMVersion gets the vm version of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *OffChainNewContract) UnmarshalJSON(raw []byte) error {
	var data struct {

		// abi version
		// Required: true
		AbiVersion Uint16 `json:"abi_version"`

		// Contract call data
		// Required: true
		CallData EncodedByteArray `json:"call_data"`

		// code
		// Required: true
		Code *ByteCode `json:"code"`

		// deposit
		// Required: true
		Deposit utils.BigInt `json:"deposit"`

		// Contract owner
		// Required: true
		Owner EncodedPubkey `json:"owner"`

		// vm version
		// Required: true
		VMVersion Uint16 `json:"vm_version"`
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

	var result OffChainNewContract

	if base.Op != result.Op() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid op value: %q", base.Op)
	}

	result.AbiVersion = data.AbiVersion

	result.CallData = data.CallData

	result.Code = data.Code

	result.Deposit = data.Deposit

	result.Owner = data.Owner

	result.VMVersion = data.VMVersion

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m OffChainNewContract) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// abi version
		// Required: true
		AbiVersion Uint16 `json:"abi_version"`

		// Contract call data
		// Required: true
		CallData EncodedByteArray `json:"call_data"`

		// code
		// Required: true
		Code *ByteCode `json:"code"`

		// deposit
		// Required: true
		Deposit utils.BigInt `json:"deposit"`

		// Contract owner
		// Required: true
		Owner EncodedPubkey `json:"owner"`

		// vm version
		// Required: true
		VMVersion Uint16 `json:"vm_version"`
	}{

		AbiVersion: m.AbiVersion,

		CallData: m.CallData,

		Code: m.Code,

		Deposit: m.Deposit,

		Owner: m.Owner,

		VMVersion: m.VMVersion,
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

// Validate validates this off chain new contract
func (m *OffChainNewContract) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbiVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCallData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OffChainNewContract) validateAbiVersion(formats strfmt.Registry) error {

	if err := m.AbiVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("abi_version")
		}
		return err
	}

	return nil
}

func (m *OffChainNewContract) validateCallData(formats strfmt.Registry) error {

	if err := m.CallData.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("call_data")
		}
		return err
	}

	return nil
}

func (m *OffChainNewContract) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if m.Code != nil {
		if err := m.Code.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("code")
			}
			return err
		}
	}

	return nil
}

func (m *OffChainNewContract) validateDeposit(formats strfmt.Registry) error {

	if err := m.Deposit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deposit")
		}
		return err
	}

	return nil
}

func (m *OffChainNewContract) validateOwner(formats strfmt.Registry) error {

	if err := m.Owner.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner")
		}
		return err
	}

	return nil
}

func (m *OffChainNewContract) validateVMVersion(formats strfmt.Registry) error {

	if err := m.VMVersion.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vm_version")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OffChainNewContract) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffChainNewContract) UnmarshalBinary(b []byte) error {
	var res OffChainNewContract
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
