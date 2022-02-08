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

// ModelsOperation models operation
//
// swagger:model models.Operation
type ModelsOperation struct {

	// from
	// Example: /overall/innerHigh
	From string `json:"from,omitempty"`

	// op
	// Example: replace
	// Required: true
	// Enum: [add remove replace move copy test]
	Op *string `json:"op"`

	// path
	// Example: /overall/outerHigh
	// Required: true
	Path *string `json:"path"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this models operation
func (m *ModelsOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var modelsOperationTypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","remove","replace","move","copy","test"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modelsOperationTypeOpPropEnum = append(modelsOperationTypeOpPropEnum, v)
	}
}

const (

	// ModelsOperationOpAdd captures enum value "add"
	ModelsOperationOpAdd string = "add"

	// ModelsOperationOpRemove captures enum value "remove"
	ModelsOperationOpRemove string = "remove"

	// ModelsOperationOpReplace captures enum value "replace"
	ModelsOperationOpReplace string = "replace"

	// ModelsOperationOpMove captures enum value "move"
	ModelsOperationOpMove string = "move"

	// ModelsOperationOpCopy captures enum value "copy"
	ModelsOperationOpCopy string = "copy"

	// ModelsOperationOpTest captures enum value "test"
	ModelsOperationOpTest string = "test"
)

// prop value enum
func (m *ModelsOperation) validateOpEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, modelsOperationTypeOpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ModelsOperation) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpEnum("op", "body", *m.Op); err != nil {
		return err
	}

	return nil
}

func (m *ModelsOperation) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models operation based on context it is used
func (m *ModelsOperation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsOperation) UnmarshalBinary(b []byte) error {
	var res ModelsOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}