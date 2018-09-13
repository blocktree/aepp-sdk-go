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

// ContractCreateData contract create data
// swagger:model ContractCreateData
type ContractCreateData struct {

	// Amount
	// Required: true
	// Minimum: 0
	Amount *int64 `json:"amount"`

	// Contract call data
	// Required: true
	CallData *string `json:"call_data"`

	// Contract's code
	// Required: true
	Code *string `json:"code"`

	// Deposit
	// Required: true
	// Minimum: 0
	Deposit *int64 `json:"deposit"`

	// Transaction fee
	// Required: true
	// Minimum: 0
	Fee *int64 `json:"fee"`

	// Contract gas
	// Required: true
	// Minimum: 0
	Gas *int64 `json:"gas"`

	// Gas price
	// Required: true
	// Minimum: 0
	GasPrice *int64 `json:"gas_price"`

	// Owner's nonce
	Nonce int64 `json:"nonce,omitempty"`

	// Contract owner pub_key
	// Required: true
	OwnerID EncodedHash `json:"owner_id"`

	// Transaction TTL
	// Minimum: 0
	TTL *int64 `json:"ttl,omitempty"`

	// Virtual machine's version
	// Required: true
	// Maximum: 255
	// Minimum: 0
	VMVersion *int64 `json:"vm_version"`
}

// Validate validates this contract create data
func (m *ContractCreateData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
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

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
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

func (m *ContractCreateData) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	if err := validate.MinimumInt("amount", "body", int64(*m.Amount), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateCallData(formats strfmt.Registry) error {

	if err := validate.Required("call_data", "body", m.CallData); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateDeposit(formats strfmt.Registry) error {

	if err := validate.Required("deposit", "body", m.Deposit); err != nil {
		return err
	}

	if err := validate.MinimumInt("deposit", "body", int64(*m.Deposit), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if err := validate.MinimumInt("fee", "body", int64(*m.Fee), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateGas(formats strfmt.Registry) error {

	if err := validate.Required("gas", "body", m.Gas); err != nil {
		return err
	}

	if err := validate.MinimumInt("gas", "body", int64(*m.Gas), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateGasPrice(formats strfmt.Registry) error {

	if err := validate.Required("gas_price", "body", m.GasPrice); err != nil {
		return err
	}

	if err := validate.MinimumInt("gas_price", "body", int64(*m.GasPrice), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateOwnerID(formats strfmt.Registry) error {

	if err := m.OwnerID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("owner_id")
		}
		return err
	}

	return nil
}

func (m *ContractCreateData) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ContractCreateData) validateVMVersion(formats strfmt.Registry) error {

	if err := validate.Required("vm_version", "body", m.VMVersion); err != nil {
		return err
	}

	if err := validate.MinimumInt("vm_version", "body", int64(*m.VMVersion), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vm_version", "body", int64(*m.VMVersion), 255, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContractCreateData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContractCreateData) UnmarshalBinary(b []byte) error {
	var res ContractCreateData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
