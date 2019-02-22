// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AWSVPCConfiguration AWS VPC Configuration
// swagger:model AWSVPCConfiguration
type AWSVPCConfiguration struct {

	// CIDR block for this VPC
	CidrBlock string `json:"cidr_block,omitempty"`

	// AWS tags
	Tags interface{} `json:"tags,omitempty"`
}

// Validate validates this a w s v p c configuration
func (m *AWSVPCConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AWSVPCConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSVPCConfiguration) UnmarshalBinary(b []byte) error {
	var res AWSVPCConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}