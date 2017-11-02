// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EndpointSettings Configuration for a network endpoint.
// swagger:model EndpointSettings

type EndpointSettings struct {

	// aliases
	Aliases []string `json:"Aliases"`

	// DriverOpts is a mapping of driver options and values. These options
	// are passed directly to the driver and are driver specific.
	//
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// Unique ID for the service endpoint in a Sandbox.
	//
	EndpointID string `json:"EndpointID,omitempty"`

	// Gateway address for this network.
	//
	Gateway string `json:"Gateway,omitempty"`

	// Global IPv6 address.
	//
	GlobalIPV6Address string `json:"GlobalIPv6Address,omitempty"`

	// Mask length of the global IPv6 address.
	//
	GlobalIPV6PrefixLen int64 `json:"GlobalIPv6PrefixLen,omitempty"`

	// IPv4 address.
	//
	IPAddress string `json:"IPAddress,omitempty"`

	// Mask length of the IPv4 address.
	//
	IPPrefixLen int64 `json:"IPPrefixLen,omitempty"`

	// IPv6 gateway address.
	//
	IPV6Gateway string `json:"IPv6Gateway,omitempty"`

	// links
	Links []string `json:"Links"`

	// MAC address for the endpoint on this network.
	//
	MacAddress string `json:"MacAddress,omitempty"`

	// Unique ID of the network.
	//
	NetworkID string `json:"NetworkID,omitempty"`
}

/* polymorph EndpointSettings Aliases false */

/* polymorph EndpointSettings DriverOpts false */

/* polymorph EndpointSettings EndpointID false */

/* polymorph EndpointSettings Gateway false */

/* polymorph EndpointSettings GlobalIPv6Address false */

/* polymorph EndpointSettings GlobalIPv6PrefixLen false */

/* polymorph EndpointSettings IPAddress false */

/* polymorph EndpointSettings IPPrefixLen false */

/* polymorph EndpointSettings IPv6Gateway false */

/* polymorph EndpointSettings Links false */

/* polymorph EndpointSettings MacAddress false */

/* polymorph EndpointSettings NetworkID false */

// Validate validates this endpoint settings
func (m *EndpointSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAliases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointSettings) validateAliases(formats strfmt.Registry) error {

	if swag.IsZero(m.Aliases) { // not required
		return nil
	}

	return nil
}

func (m *EndpointSettings) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointSettings) UnmarshalBinary(b []byte) error {
	var res EndpointSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}