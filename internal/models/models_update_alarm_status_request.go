// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsUpdateAlarmStatusRequest models update alarm status request
//
// swagger:model models.UpdateAlarmStatusRequest
type ModelsUpdateAlarmStatusRequest struct {

	// content type
	// Example: DATA_POINT
	// Required: true
	// Enum: [DATA_POINT SPECTRUM QUESTION_ANSWERS]
	ContentType *string `json:"contentType"`

	// created at
	// Example: 2021-09-10T11:52:38.299Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// data point
	DataPoint *ModelsDataPoint `json:"dataPoint,omitempty"`

	// measurement Id
	// Example: 123e4567-e89b-12d3-a456-426614174000
	// Required: true
	// Format: uuid
	MeasurementID *strfmt.UUID `json:"measurementId"`

	// question answers
	// Example: ["yes","no","maybe"]
	QuestionAnswers []string `json:"questionAnswers,omitempty"`

	// rate of change
	// Example: 5
	RateOfChange *float64 `json:"rateOfChange,omitempty"`

	// spectrum
	Spectrum *ModelsSpectrum `json:"spectrum,omitempty"`

	// tags
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// Validate validates this models update alarm status request
func (m *ModelsUpdateAlarmStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeasurementID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpectrum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modelsUpdateAlarmStatusRequestTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_POINT","SPECTRUM","QUESTION_ANSWERS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelsUpdateAlarmStatusRequestTypeContentTypePropEnum = append(modelsUpdateAlarmStatusRequestTypeContentTypePropEnum, v)
	}
}

const (

	// ModelsUpdateAlarmStatusRequestContentTypeDATAPOINT captures enum value "DATA_POINT"
	ModelsUpdateAlarmStatusRequestContentTypeDATAPOINT string = "DATA_POINT"

	// ModelsUpdateAlarmStatusRequestContentTypeSPECTRUM captures enum value "SPECTRUM"
	ModelsUpdateAlarmStatusRequestContentTypeSPECTRUM string = "SPECTRUM"

	// ModelsUpdateAlarmStatusRequestContentTypeQUESTIONANSWERS captures enum value "QUESTION_ANSWERS"
	ModelsUpdateAlarmStatusRequestContentTypeQUESTIONANSWERS string = "QUESTION_ANSWERS"
)

// prop value enum
func (m *ModelsUpdateAlarmStatusRequest) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modelsUpdateAlarmStatusRequestTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("contentType", "body", m.ContentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) validateDataPoint(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPoint) { // not required
		return nil
	}

	if m.DataPoint != nil {
		if err := m.DataPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataPoint")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) validateMeasurementID(formats strfmt.Registry) error {

	if err := validate.Required("measurementId", "body", m.MeasurementID); err != nil {
		return err
	}

	if err := validate.FormatOf("measurementId", "body", "uuid", m.MeasurementID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) validateSpectrum(formats strfmt.Registry) error {
	if swag.IsZero(m.Spectrum) { // not required
		return nil
	}

	if m.Spectrum != nil {
		if err := m.Spectrum.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spectrum")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spectrum")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this models update alarm status request based on the context it is used
func (m *ModelsUpdateAlarmStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataPoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpectrum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) contextValidateDataPoint(ctx context.Context, formats strfmt.Registry) error {

	if m.DataPoint != nil {
		if err := m.DataPoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPoint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataPoint")
			}
			return err
		}
	}

	return nil
}

func (m *ModelsUpdateAlarmStatusRequest) contextValidateSpectrum(ctx context.Context, formats strfmt.Registry) error {

	if m.Spectrum != nil {
		if err := m.Spectrum.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spectrum")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spectrum")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsUpdateAlarmStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsUpdateAlarmStatusRequest) UnmarshalBinary(b []byte) error {
	var res ModelsUpdateAlarmStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
