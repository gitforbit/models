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

// HTTPResponseRule HTTP Response Rule
//
// HAProxy HTTP response rule configuration (corresponds to http-response directives)
// swagger:model http_response_rule
type HTTPResponseRule struct {

	// cond
	// Enum: [if unless]
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// hdr format
	// Pattern: ^[^\s]+$
	HdrFormat string `json:"hdr_format,omitempty"`

	// hdr match
	// Pattern: ^[^\s]+$
	HdrMatch string `json:"hdr_match,omitempty"`

	// hdr name
	// Pattern: ^[^\s]+$
	HdrName string `json:"hdr_name,omitempty"`

	// hdr value
	// Pattern: ^[^\s]+$
	HdrValue string `json:"hdr_value,omitempty"`

	// id
	// Required: true
	ID int64 `json:"id"`

	// log level
	// Enum: [emerg alert crit err warning notice info debug silent]
	LogLevel string `json:"log_level,omitempty"`

	// spoe engine
	// Pattern: ^[^\s]+$
	SpoeEngine string `json:"spoe_engine,omitempty"`

	// spoe group
	// Pattern: ^[^\s]+$
	SpoeGroup string `json:"spoe_group,omitempty"`

	// status group
	StatusGroup int64 `json:"status_group,omitempty"`

	// type
	// Required: true
	// Enum: [allow deny add-header set-header set-log-level set-var set-status send-spoe-group replace-header replace-value]
	Type string `json:"type"`

	// var name
	// Pattern: ^[^\s]+$
	VarName string `json:"var_name,omitempty"`

	// var pattern
	// Pattern: ^[^\s]+$
	VarPattern string `json:"var_pattern,omitempty"`
}

// Validate validates this http response rule
func (m *HTTPResponseRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdrFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdrMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdrName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdrValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpoeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVarPattern(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var httpResponseRuleTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeCondPropEnum = append(httpResponseRuleTypeCondPropEnum, v)
	}
}

const (

	// HTTPResponseRuleCondIf captures enum value "if"
	HTTPResponseRuleCondIf string = "if"

	// HTTPResponseRuleCondUnless captures enum value "unless"
	HTTPResponseRuleCondUnless string = "unless"
)

// prop value enum
func (m *HTTPResponseRule) validateCondEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeCondPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateCond(formats strfmt.Registry) error {

	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateHdrFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.HdrFormat) { // not required
		return nil
	}

	if err := validate.Pattern("hdr_format", "body", string(m.HdrFormat), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateHdrMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.HdrMatch) { // not required
		return nil
	}

	if err := validate.Pattern("hdr_match", "body", string(m.HdrMatch), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateHdrName(formats strfmt.Registry) error {

	if swag.IsZero(m.HdrName) { // not required
		return nil
	}

	if err := validate.Pattern("hdr_name", "body", string(m.HdrName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateHdrValue(formats strfmt.Registry) error {

	if swag.IsZero(m.HdrValue) { // not required
		return nil
	}

	if err := validate.Pattern("hdr_value", "body", string(m.HdrValue), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emerg","alert","crit","err","warning","notice","info","debug","silent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeLogLevelPropEnum = append(httpResponseRuleTypeLogLevelPropEnum, v)
	}
}

const (

	// HTTPResponseRuleLogLevelEmerg captures enum value "emerg"
	HTTPResponseRuleLogLevelEmerg string = "emerg"

	// HTTPResponseRuleLogLevelAlert captures enum value "alert"
	HTTPResponseRuleLogLevelAlert string = "alert"

	// HTTPResponseRuleLogLevelCrit captures enum value "crit"
	HTTPResponseRuleLogLevelCrit string = "crit"

	// HTTPResponseRuleLogLevelErr captures enum value "err"
	HTTPResponseRuleLogLevelErr string = "err"

	// HTTPResponseRuleLogLevelWarning captures enum value "warning"
	HTTPResponseRuleLogLevelWarning string = "warning"

	// HTTPResponseRuleLogLevelNotice captures enum value "notice"
	HTTPResponseRuleLogLevelNotice string = "notice"

	// HTTPResponseRuleLogLevelInfo captures enum value "info"
	HTTPResponseRuleLogLevelInfo string = "info"

	// HTTPResponseRuleLogLevelDebug captures enum value "debug"
	HTTPResponseRuleLogLevelDebug string = "debug"

	// HTTPResponseRuleLogLevelSilent captures enum value "silent"
	HTTPResponseRuleLogLevelSilent string = "silent"
)

// prop value enum
func (m *HTTPResponseRule) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeLogLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateLogLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.LogLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeEngine(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeEngine) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_engine", "body", string(m.SpoeEngine), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateSpoeGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SpoeGroup) { // not required
		return nil
	}

	if err := validate.Pattern("spoe_group", "body", string(m.SpoeGroup), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var httpResponseRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny","add-header","set-header","set-log-level","set-var","set-status","send-spoe-group","replace-header","replace-value"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpResponseRuleTypeTypePropEnum = append(httpResponseRuleTypeTypePropEnum, v)
	}
}

const (

	// HTTPResponseRuleTypeAllow captures enum value "allow"
	HTTPResponseRuleTypeAllow string = "allow"

	// HTTPResponseRuleTypeDeny captures enum value "deny"
	HTTPResponseRuleTypeDeny string = "deny"

	// HTTPResponseRuleTypeAddHeader captures enum value "add-header"
	HTTPResponseRuleTypeAddHeader string = "add-header"

	// HTTPResponseRuleTypeSetHeader captures enum value "set-header"
	HTTPResponseRuleTypeSetHeader string = "set-header"

	// HTTPResponseRuleTypeSetLogLevel captures enum value "set-log-level"
	HTTPResponseRuleTypeSetLogLevel string = "set-log-level"

	// HTTPResponseRuleTypeSetVar captures enum value "set-var"
	HTTPResponseRuleTypeSetVar string = "set-var"

	// HTTPResponseRuleTypeSetStatus captures enum value "set-status"
	HTTPResponseRuleTypeSetStatus string = "set-status"

	// HTTPResponseRuleTypeSendSpoeGroup captures enum value "send-spoe-group"
	HTTPResponseRuleTypeSendSpoeGroup string = "send-spoe-group"

	// HTTPResponseRuleTypeReplaceHeader captures enum value "replace-header"
	HTTPResponseRuleTypeReplaceHeader string = "replace-header"

	// HTTPResponseRuleTypeReplaceValue captures enum value "replace-value"
	HTTPResponseRuleTypeReplaceValue string = "replace-value"
)

// prop value enum
func (m *HTTPResponseRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, httpResponseRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HTTPResponseRule) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarName(formats strfmt.Registry) error {

	if swag.IsZero(m.VarName) { // not required
		return nil
	}

	if err := validate.Pattern("var_name", "body", string(m.VarName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *HTTPResponseRule) validateVarPattern(formats strfmt.Registry) error {

	if swag.IsZero(m.VarPattern) { // not required
		return nil
	}

	if err := validate.Pattern("var_pattern", "body", string(m.VarPattern), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HTTPResponseRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPResponseRule) UnmarshalBinary(b []byte) error {
	var res HTTPResponseRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
