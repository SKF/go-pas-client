// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsSetPointAlarmThresholdRequest models set point alarm threshold request
//
// swagger:model models.SetPointAlarmThresholdRequest
type ModelsSetPointAlarmThresholdRequest struct {

	// band alarms
	BandAlarms []*ModelsBandAlarm `json:"bandAlarms"`

	// full scale
	FullScale float64 `json:"fullScale,omitempty"`

	// hal alarms
	HalAlarms []*ModelsHALAlarm `json:"halAlarms"`

	// inspection
	Inspection *ModelsInspection `json:"inspection,omitempty"`

	// origin
	Origin *ModelsOrigin `json:"origin,omitempty"`

	// overall
	Overall *ModelsOverall `json:"overall,omitempty"`

	// rate of change
	RateOfChange *ModelsRateOfChange `json:"rateOfChange,omitempty"`

	// The type values are available [here](/v1/docs/service#threshold-type).
	// Example: 2
	// Required: true
	// Maximum: 3
	// Minimum: 0
	ThresholdType *int32 `json:"thresholdType"`
}

// Validate validates this models set point alarm threshold request
func (m *ModelsSetPointAlarmThresholdRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandAlarms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHalAlarms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInspection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateOfChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholdType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateBandAlarms(formats strfmt.Registry) error {
	if swag.IsZero(m.BandAlarms) { // not required
		return nil
	}

	for i := 0; i < len(m.BandAlarms); i++ {
		if swag.IsZero(m.BandAlarms[i]) { // not required
			continue
		}

		if m.BandAlarms[i] != nil {
			if err := m.BandAlarms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bandAlarms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bandAlarms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateHalAlarms(formats strfmt.Registry) error {
	if swag.IsZero(m.HalAlarms) { // not required
		return nil
	}

	for i := 0; i < len(m.HalAlarms); i++ {
		if swag.IsZero(m.HalAlarms[i]) { // not required
			continue
		}

		if m.HalAlarms[i] != nil {
			if err := m.HalAlarms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("halAlarms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("halAlarms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateInspection(formats strfmt.Registry) error {
	if swag.IsZero(m.Inspection) { // not required
		return nil
	}

	if m.Inspection != nil {
		if err := m.Inspection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inspection")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateOverall(formats strfmt.Registry) error {
	if swag.IsZero(m.Overall) { // not required
		return nil
	}

	if m.Overall != nil {
		if err := m.Overall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overall")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateRateOfChange(formats strfmt.Registry) error {
	if swag.IsZero(m.RateOfChange) { // not required
		return nil
	}

	if m.RateOfChange != nil {
		if err := m.RateOfChange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateOfChange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rateOfChange")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) validateThresholdType(formats strfmt.Registry) error {

	if err := validate.Required("thresholdType", "body", m.ThresholdType); err != nil {
		return err
	}

	if err := validate.MinimumInt("thresholdType", "body", int64(*m.ThresholdType), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("thresholdType", "body", int64(*m.ThresholdType), 3, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this models set point alarm threshold request based on the context it is used
func (m *ModelsSetPointAlarmThresholdRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBandAlarms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHalAlarms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInspection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRateOfChange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) contextValidateBandAlarms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BandAlarms); i++ {

		if m.BandAlarms[i] != nil {
			if err := m.BandAlarms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bandAlarms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bandAlarms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) contextValidateHalAlarms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HalAlarms); i++ {

		if m.HalAlarms[i] != nil {
			if err := m.HalAlarms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("halAlarms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("halAlarms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) contextValidateInspection(ctx context.Context, formats strfmt.Registry) error {

	if m.Inspection != nil {
		if err := m.Inspection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inspection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inspection")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {
		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) contextValidateOverall(ctx context.Context, formats strfmt.Registry) error {

	if m.Overall != nil {
		if err := m.Overall.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overall")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsSetPointAlarmThresholdRequest) contextValidateRateOfChange(ctx context.Context, formats strfmt.Registry) error {

	if m.RateOfChange != nil {
		if err := m.RateOfChange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateOfChange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rateOfChange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsSetPointAlarmThresholdRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsSetPointAlarmThresholdRequest) UnmarshalBinary(b []byte) error {
	var res ModelsSetPointAlarmThresholdRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
