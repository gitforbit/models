// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SiteBackendsItems site backends items
// swagger:model siteBackendsItems
type SiteBackendsItems struct {

	// balance
	// Enum: [roundrobin least-connections hash-uri hash-source]
	Balance string `json:"balance,omitempty"`

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// http xff header insert
	// Enum: [enabled]
	HTTPXffHeaderInsert string `json:"http_xff_header_insert,omitempty"`

	// log
	// Enum: [enabled]
	Log string `json:"log,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name string `json:"name"`

	// protocol
	// Enum: [http tcp]
	Protocol string `json:"protocol,omitempty"`

	// servers
	Servers []*SiteBackendsItemsServersItems `json:"servers"`

	// use as
	// Required: true
	// Enum: [default conditional]
	UseAs string `json:"use_as"`
}

// Validate validates this site backends items
func (m *SiteBackendsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPXffHeaderInsert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseAs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var siteBackendsItemsTypeBalancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roundrobin","least-connections","hash-uri","hash-source"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsTypeBalancePropEnum = append(siteBackendsItemsTypeBalancePropEnum, v)
	}
}

const (

	// SiteBackendsItemsBalanceRoundrobin captures enum value "roundrobin"
	SiteBackendsItemsBalanceRoundrobin string = "roundrobin"

	// SiteBackendsItemsBalanceLeastConnections captures enum value "least-connections"
	SiteBackendsItemsBalanceLeastConnections string = "least-connections"

	// SiteBackendsItemsBalanceHashURI captures enum value "hash-uri"
	SiteBackendsItemsBalanceHashURI string = "hash-uri"

	// SiteBackendsItemsBalanceHashSource captures enum value "hash-source"
	SiteBackendsItemsBalanceHashSource string = "hash-source"
)

// prop value enum
func (m *SiteBackendsItems) validateBalanceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsTypeBalancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItems) validateBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	// value enum
	if err := m.validateBalanceEnum("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

var siteBackendsItemsTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsTypeCondPropEnum = append(siteBackendsItemsTypeCondPropEnum, v)
	}
}

const (

	// SiteBackendsItemsCondIf captures enum value "if"
	SiteBackendsItemsCondIf string = "if"

	// SiteBackendsItemsCondUnless captures enum value "unless"
	SiteBackendsItemsCondUnless string = "unless"
)

// prop value enum
func (m *SiteBackendsItems) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItems) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

var siteBackendsItemsTypeHTTPXffHeaderInsertPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsTypeHTTPXffHeaderInsertPropEnum = append(siteBackendsItemsTypeHTTPXffHeaderInsertPropEnum, v)
	}
}

const (

	// SiteBackendsItemsHTTPXffHeaderInsertEnabled captures enum value "enabled"
	SiteBackendsItemsHTTPXffHeaderInsertEnabled string = "enabled"
)

// prop value enum
func (m *SiteBackendsItems) validateHTTPXffHeaderInsertEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsTypeHTTPXffHeaderInsertPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItems) validateHTTPXffHeaderInsert(formats strfmt.Registry) error {

	if swag.IsZero(m.HTTPXffHeaderInsert) { // not required
		return nil
	}

	// value enum
	if err := m.validateHTTPXffHeaderInsertEnum("http_xff_header_insert", "body", m.HTTPXffHeaderInsert); err != nil {
		return err
	}

	return nil
}

var siteBackendsItemsTypeLogPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsTypeLogPropEnum = append(siteBackendsItemsTypeLogPropEnum, v)
	}
}

const (

	// SiteBackendsItemsLogEnabled captures enum value "enabled"
	SiteBackendsItemsLogEnabled string = "enabled"
)

// prop value enum
func (m *SiteBackendsItems) validateLogEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsTypeLogPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItems) validateLog(formats strfmt.Registry) error {

	if swag.IsZero(m.Log) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogEnum("log", "body", m.Log); err != nil {
		return err
	}

	return nil
}

func (m *SiteBackendsItems) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

var siteBackendsItemsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","tcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsTypeProtocolPropEnum = append(siteBackendsItemsTypeProtocolPropEnum, v)
	}
}

const (

	// SiteBackendsItemsProtocolHTTP captures enum value "http"
	SiteBackendsItemsProtocolHTTP string = "http"

	// SiteBackendsItemsProtocolTCP captures enum value "tcp"
	SiteBackendsItemsProtocolTCP string = "tcp"
)

// prop value enum
func (m *SiteBackendsItems) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItems) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *SiteBackendsItems) validateServers(formats strfmt.Registry) error {

	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	for i := 0; i < len(m.Servers); i++ {
		if swag.IsZero(m.Servers[i]) { // not required
			continue
		}

		if m.Servers[i] != nil {
			if err := m.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var siteBackendsItemsTypeUseAsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","conditional"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteBackendsItemsTypeUseAsPropEnum = append(siteBackendsItemsTypeUseAsPropEnum, v)
	}
}

const (

	// SiteBackendsItemsUseAsDefault captures enum value "default"
	SiteBackendsItemsUseAsDefault string = "default"

	// SiteBackendsItemsUseAsConditional captures enum value "conditional"
	SiteBackendsItemsUseAsConditional string = "conditional"
)

// prop value enum
func (m *SiteBackendsItems) validateUseAsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteBackendsItemsTypeUseAsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SiteBackendsItems) validateUseAs(formats strfmt.Registry) error {

	if err := validate.RequiredString("use_as", "body", string(m.UseAs)); err != nil {
		return err
	}

	// value enum
	if err := m.validateUseAsEnum("use_as", "body", m.UseAs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteBackendsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteBackendsItems) UnmarshalBinary(b []byte) error {
	var res SiteBackendsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
