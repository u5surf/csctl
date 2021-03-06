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

// AutoscalingPolicy autoscaling policy
// swagger:model AutoscalingPolicy
type AutoscalingPolicy struct {

	// AutoscalingPolicy ID
	ID UUID `json:"id,omitempty"`

	// String representation of the target metric to monitor
	//
	// Available values are provided by the given MetricsBackend
	// Required: true
	Metric *string `json:"metric"`

	// Arbitrary configuration object used to configure the metric polling
	MetricConfiguration interface{} `json:"metric_configuration,omitempty"`

	// MetricsBackend name associated with the AutoscalingPolicy
	MetricsBackend string `json:"metrics_backend,omitempty"`

	// Name of this policy
	// Required: true
	Name *string `json:"name"`

	// Number of seconds between polling the associated MetricsBackend
	// Required: true
	PollInterval *int32 `json:"poll_interval"`

	// Number of seconds the AutoscalingPolicy must alert the threshold before the policy triggers a scale up or scale down action
	// Required: true
	SamplePeriod *int32 `json:"sample_period"`

	// Scaling policy
	// Required: true
	ScalingPolicy *ScalingPolicy `json:"scaling_policy"`
}

// Validate validates this autoscaling policy
func (m *AutoscalingPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePollInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamplePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScalingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoscalingPolicy) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *AutoscalingPolicy) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

func (m *AutoscalingPolicy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AutoscalingPolicy) validatePollInterval(formats strfmt.Registry) error {

	if err := validate.Required("poll_interval", "body", m.PollInterval); err != nil {
		return err
	}

	return nil
}

func (m *AutoscalingPolicy) validateSamplePeriod(formats strfmt.Registry) error {

	if err := validate.Required("sample_period", "body", m.SamplePeriod); err != nil {
		return err
	}

	return nil
}

func (m *AutoscalingPolicy) validateScalingPolicy(formats strfmt.Registry) error {

	if err := validate.Required("scaling_policy", "body", m.ScalingPolicy); err != nil {
		return err
	}

	if m.ScalingPolicy != nil {
		if err := m.ScalingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaling_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoscalingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscalingPolicy) UnmarshalBinary(b []byte) error {
	var res AutoscalingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
