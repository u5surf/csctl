// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TemplateConfiguration template configuration
// swagger:model TemplateConfiguration
type TemplateConfiguration struct {

	// Provider-specific cloud resources
	// Required: true
	Resource *TemplateResource `json:"resource"`

	// Provider-agnostic cloud resources
	// Required: true
	Variable TemplateVariableMap `json:"variable"`
}

// Validate validates this template configuration
func (m *TemplateConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateConfiguration) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *TemplateConfiguration) validateVariable(formats strfmt.Registry) error {

	if err := m.Variable.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("variable")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateConfiguration) UnmarshalBinary(b []byte) error {
	var res TemplateConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
