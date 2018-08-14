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

// ChannelSnapshotSoloTx channel snapshot solo tx
// swagger:model ChannelSnapshotSoloTx
type ChannelSnapshotSoloTx struct {

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

	// ttl
	// Minimum: 0
	TTL *int64 `json:"ttl,omitempty"`
}

// Validate validates this channel snapshot solo tx
func (m *ChannelSnapshotSoloTx) Validate(formats strfmt.Registry) error {
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

	if err := m.validateTTL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelSnapshotSoloTx) validateChannelID(formats strfmt.Registry) error {

	if err := m.ChannelID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("channel_id")
		}
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validateFee(formats strfmt.Registry) error {

	if err := validate.Required("fee", "body", m.Fee); err != nil {
		return err
	}

	if err := validate.MinimumInt("fee", "body", int64(*m.Fee), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validateFrom(formats strfmt.Registry) error {

	if err := m.From.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("from")
		}
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validateNonce(formats strfmt.Registry) error {

	if swag.IsZero(m.Nonce) { // not required
		return nil
	}

	if err := validate.MinimumInt("nonce", "body", int64(*m.Nonce), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validatePayload(formats strfmt.Registry) error {

	if err := validate.Required("payload", "body", m.Payload); err != nil {
		return err
	}

	return nil
}

func (m *ChannelSnapshotSoloTx) validateTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.TTL) { // not required
		return nil
	}

	if err := validate.MinimumInt("ttl", "body", int64(*m.TTL), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelSnapshotSoloTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelSnapshotSoloTx) UnmarshalBinary(b []byte) error {
	var res ChannelSnapshotSoloTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
