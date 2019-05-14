// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StoryDataContainer story data container
// swagger:model StoryDataContainer
type StoryDataContainer struct {

	// The total number of results returned by this call.
	Count int32 `json:"count,omitempty"`

	// The requested result limit.
	Limit int32 `json:"limit,omitempty"`

	// The requested offset (number of skipped results) of the call.
	Offset int32 `json:"offset,omitempty"`

	// The list of stories returned by the call
	Results []*Story `json:"results"`

	// The total number of resources available given the current filter set.
	Total int32 `json:"total,omitempty"`
}

// Validate validates this story data container
func (m *StoryDataContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoryDataContainer) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoryDataContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryDataContainer) UnmarshalBinary(b []byte) error {
	var res StoryDataContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}