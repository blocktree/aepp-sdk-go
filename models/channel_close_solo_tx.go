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

// ChannelCloseSoloTx channel close solo tx
// swagger:model ChannelCloseSoloTx
type ChannelCloseSoloTx struct {

	// channel id
	// Required: true
	ChannelID EncodedHash `json:"channel_id"`

	// fee
	// Required: true
	// Minimum: 0
	Fee *int64 `json:"fee"`

	// from
	// Required: true
	From EncodedHash `json:"from"`

	// nonce
	// Minimum: 0
	Nonce *int64 `json:"nonce,omitempty"`

	// payload
	// Required: true
	Payload *string `json:"payload"`

	// Proof of inclusion containing information for closing the channel
	// Required: true
	Poi *string `json:"poi"`

	// ttl
	// Minimum: 0
	TTL *int64 `json:"ttl,omitempty"`
}

// Validate validates this channel close solo tx
func (m *ChannelCloseSoloTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePoi(formats); err != nil {
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

func (m *ChannelCloseSoloTx) validateChannelID(formats strfmt.Registry) error {

	if err := m.ChannelID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("channel_id")
		}
		return err
	}

	return nil
}

func (m *ChannelCloseSoloTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if err := validate.MinimumInt("fee", "body", int64(*m.Fee), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseSoloTx) validateFrom(formats strfmt.Registry) error {

	if err := m.From.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("from")
		}
		return err
	}

	return nil
}

func (m *ChannelCloseSoloTx) validateNonce(formats strfmt.Registry) error {

	if swag.IsZero(m.Nonce) { // not required
		return nil
	}

	if err := validate.MinimumInt("nonce", "body", int64(*m.Nonce), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseSoloTx) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseSoloTx) validatePoi(formats strfmt.Registry) error {

	if err := validate.Required("poi", "body", m.Poi); err != nil {
		return err
	}

	return nil
}

func (m *ChannelCloseSoloTx) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelCloseSoloTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelCloseSoloTx) UnmarshalBinary(b []byte) error {
	var res ChannelCloseSoloTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
