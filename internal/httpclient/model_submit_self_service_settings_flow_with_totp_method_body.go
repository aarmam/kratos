/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceSettingsFlowWithTOTPMethodBody struct for SubmitSelfServiceSettingsFlowWithTOTPMethodBody
type SubmitSelfServiceSettingsFlowWithTOTPMethodBody struct {
	// CSRFToken is the anti-CSRF token
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Method  Should be set to \"totp\" when trying to add, update, or remove a totp pairing.
	Method string `json:"method"`
	// UnlinkTOTP if true will remove the TOTP pairing, effectively removing the credential. This can be used to set up a new TOTP device.
	UnlinkTotp *bool `json:"unlink_totp,omitempty"`
	// ValidationTOTP must contain a valid TOTP based on the
	VerificationTotp *string `json:"verification_totp,omitempty"`
}

// NewSubmitSelfServiceSettingsFlowWithTOTPMethodBody instantiates a new SubmitSelfServiceSettingsFlowWithTOTPMethodBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceSettingsFlowWithTOTPMethodBody(method string) *SubmitSelfServiceSettingsFlowWithTOTPMethodBody {
	this := SubmitSelfServiceSettingsFlowWithTOTPMethodBody{}
	this.Method = method
	return &this
}

// NewSubmitSelfServiceSettingsFlowWithTOTPMethodBodyWithDefaults instantiates a new SubmitSelfServiceSettingsFlowWithTOTPMethodBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceSettingsFlowWithTOTPMethodBodyWithDefaults() *SubmitSelfServiceSettingsFlowWithTOTPMethodBody {
	this := SubmitSelfServiceSettingsFlowWithTOTPMethodBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetMethod returns the Method field value
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) SetMethod(v string) {
	o.Method = v
}

// GetUnlinkTotp returns the UnlinkTotp field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetUnlinkTotp() bool {
	if o == nil || o.UnlinkTotp == nil {
		var ret bool
		return ret
	}
	return *o.UnlinkTotp
}

// GetUnlinkTotpOk returns a tuple with the UnlinkTotp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetUnlinkTotpOk() (*bool, bool) {
	if o == nil || o.UnlinkTotp == nil {
		return nil, false
	}
	return o.UnlinkTotp, true
}

// HasUnlinkTotp returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) HasUnlinkTotp() bool {
	if o != nil && o.UnlinkTotp != nil {
		return true
	}

	return false
}

// SetUnlinkTotp gets a reference to the given bool and assigns it to the UnlinkTotp field.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) SetUnlinkTotp(v bool) {
	o.UnlinkTotp = &v
}

// GetVerificationTotp returns the VerificationTotp field value if set, zero value otherwise.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetVerificationTotp() string {
	if o == nil || o.VerificationTotp == nil {
		var ret string
		return ret
	}
	return *o.VerificationTotp
}

// GetVerificationTotpOk returns a tuple with the VerificationTotp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) GetVerificationTotpOk() (*string, bool) {
	if o == nil || o.VerificationTotp == nil {
		return nil, false
	}
	return o.VerificationTotp, true
}

// HasVerificationTotp returns a boolean if a field has been set.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) HasVerificationTotp() bool {
	if o != nil && o.VerificationTotp != nil {
		return true
	}

	return false
}

// SetVerificationTotp gets a reference to the given string and assigns it to the VerificationTotp field.
func (o *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) SetVerificationTotp(v string) {
	o.VerificationTotp = &v
}

func (o SubmitSelfServiceSettingsFlowWithTOTPMethodBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if o.UnlinkTotp != nil {
		toSerialize["unlink_totp"] = o.UnlinkTotp
	}
	if o.VerificationTotp != nil {
		toSerialize["verification_totp"] = o.VerificationTotp
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody struct {
	value *SubmitSelfServiceSettingsFlowWithTOTPMethodBody
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody) Get() *SubmitSelfServiceSettingsFlowWithTOTPMethodBody {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody) Set(val *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody(val *SubmitSelfServiceSettingsFlowWithTOTPMethodBody) *NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody {
	return &NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlowWithTOTPMethodBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}