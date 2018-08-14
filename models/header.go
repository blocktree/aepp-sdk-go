// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Header header
// swagger:model Header
type Header struct {

	// beneficiary
	Beneficiary EncodedHash `json:"beneficiary,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// miner
	Miner EncodedHash `json:"miner,omitempty"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// pow
	Pow Pow `json:"pow"`

	// prev hash
	PrevHash EncodedHash `json:"prev_hash,omitempty"`

	// state hash
	StateHash EncodedHash `json:"state_hash,omitempty"`

	// target
	Target int64 `json:"target,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// txs hash
	TxsHash EncodedHash `json:"txs_hash,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this header
func (m *Header) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMiner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxsHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Header) validateBeneficiary(formats strfmt.Registry) error {

	if swag.IsZero(m.Beneficiary) { // not required
		return nil
	}

	if err := m.Beneficiary.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("beneficiary")
		}
		return err
	}

	return nil
}

func (m *Header) validateMiner(formats strfmt.Registry) error {

	if swag.IsZero(m.Miner) { // not required
		return nil
	}

	if err := m.Miner.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("miner")
		}
		return err
	}

	return nil
}

func (m *Header) validatePow(formats strfmt.Registry) error {

	if swag.IsZero(m.Pow) { // not required
		return nil
	}

	if err := m.Pow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pow")
		}
		return err
	}

	return nil
}

func (m *Header) validatePrevHash(formats strfmt.Registry) error {

	if swag.IsZero(m.PrevHash) { // not required
		return nil
	}

	if err := m.PrevHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("prev_hash")
		}
		return err
	}

	return nil
}

func (m *Header) validateStateHash(formats strfmt.Registry) error {

	if swag.IsZero(m.StateHash) { // not required
		return nil
	}

	if err := m.StateHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state_hash")
		}
		return err
	}

	return nil
}

func (m *Header) validateTxsHash(formats strfmt.Registry) error {

	if swag.IsZero(m.TxsHash) { // not required
		return nil
	}

	if err := m.TxsHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("txs_hash")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Header) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Header) UnmarshalBinary(b []byte) error {
	var res Header
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
