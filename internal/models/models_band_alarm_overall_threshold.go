// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelsBandAlarmOverallThreshold models band alarm overall threshold
//
// swagger:model models.BandAlarmOverallThreshold
type ModelsBandAlarmOverallThreshold struct {

	// unit
	Unit string `json:"unit,omitempty"`

	// upper alert
	UpperAlert *ModelsBandAlarmThreshold `json:"upperAlert,omitempty"`

	// upper danger
	UpperDanger *ModelsBandAlarmThreshold `json:"upperDanger,omitempty"`
}

// Validate validates this models band alarm overall threshold
func (m *ModelsBandAlarmOverallThreshold) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpperAlert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpperDanger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsBandAlarmOverallThreshold) validateUpperAlert(formats strfmt.Registry) error {
	if swag.IsZero(m.UpperAlert) { // not required
		return nil
	}

	if m.UpperAlert != nil {
		if err := m.UpperAlert.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upperAlert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upperAlert")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsBandAlarmOverallThreshold) validateUpperDanger(formats strfmt.Registry) error {
	if swag.IsZero(m.UpperDanger) { // not required
		return nil
	}

	if m.UpperDanger != nil {
		if err := m.UpperDanger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upperDanger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upperDanger")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this models band alarm overall threshold based on the context it is used
func (m *ModelsBandAlarmOverallThreshold) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpperAlert(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpperDanger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsBandAlarmOverallThreshold) contextValidateUpperAlert(ctx context.Context, formats strfmt.Registry) error {

	if m.UpperAlert != nil {
		if err := m.UpperAlert.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upperAlert")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upperAlert")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsBandAlarmOverallThreshold) contextValidateUpperDanger(ctx context.Context, formats strfmt.Registry) error {

	if m.UpperDanger != nil {
		if err := m.UpperDanger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upperDanger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upperDanger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsBandAlarmOverallThreshold) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsBandAlarmOverallThreshold) UnmarshalBinary(b []byte) error {
	var res ModelsBandAlarmOverallThreshold
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
