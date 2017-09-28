package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// RuntimeStats runtime stats
// swagger:model RuntimeStats
type RuntimeStats struct {

	// export worker
	ExportWorker int64 `json:"exportWorker,omitempty"`

	// exports pending
	ExportsPending int64 `json:"exportsPending,omitempty"`

	// import worker
	ImportWorker int64 `json:"importWorker,omitempty"`

	// imports pending
	ImportsPending int64 `json:"importsPending,omitempty"`
}

// Validate validates this runtime stats
func (m *RuntimeStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
