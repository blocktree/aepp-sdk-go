// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	utils "github.com/aeternity/aepp-sdk-go/utils"
)

// SpendTx spend tx
// swagger:model SpendTx
type SpendTx struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// nonce
	Nonce Uint64 `json:"nonce,omitempty"`

	// payload
	// Required: true
	Payload EncodedByteArray `json:"payload"`

	// recipient id
	// Required: true
	RecipientID EncodedPubkey `json:"recipient_id"`

	// sender id
	// Required: true
	SenderID EncodedPubkey `json:"sender_id"`

	// ttl
	TTL Uint64 `json:"ttl,omitempty"`
}

// Validate validates this spend tx
func (m *SpendTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpendTx) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validateNonce(formats strfmt.Registry) error {

	if swag.IsZero(m.Nonce) { // not required
		return nil
	}

	if err := m.Nonce.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("nonce")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validatePayload(formats strfmt.Registry) error {

	if err := m.Payload.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("payload")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validateRecipientID(formats strfmt.Registry) error {

	if err := m.RecipientID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("recipient_id")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validateSenderID(formats strfmt.Registry) error {

	if err := m.SenderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sender_id")
		}
		return err
	}

	return nil
}

func (m *SpendTx) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if err := m.TTL.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ttl")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpendTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpendTx) UnmarshalBinary(b []byte) error {
	var res SpendTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
