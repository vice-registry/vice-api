// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ManagementLayer management layer
// swagger:model ManagementLayer

type ManagementLayer struct {

	// software
	Software string `json:"software,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

/* polymorph ManagementLayer software false */

/* polymorph ManagementLayer type false */

/* polymorph ManagementLayer version false */

// Validate validates this management layer
func (m *ManagementLayer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var managementLayerTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic","cloudcomputing","containercluster","jobscheduler"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		managementLayerTypeTypePropEnum = append(managementLayerTypeTypePropEnum, v)
	}
}

const (
	// ManagementLayerTypeBasic captures enum value "basic"
	ManagementLayerTypeBasic string = "basic"
	// ManagementLayerTypeCloudcomputing captures enum value "cloudcomputing"
	ManagementLayerTypeCloudcomputing string = "cloudcomputing"
	// ManagementLayerTypeContainercluster captures enum value "containercluster"
	ManagementLayerTypeContainercluster string = "containercluster"
	// ManagementLayerTypeJobscheduler captures enum value "jobscheduler"
	ManagementLayerTypeJobscheduler string = "jobscheduler"
)

// prop value enum
func (m *ManagementLayer) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, managementLayerTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ManagementLayer) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementLayer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementLayer) UnmarshalBinary(b []byte) error {
	var res ManagementLayer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
