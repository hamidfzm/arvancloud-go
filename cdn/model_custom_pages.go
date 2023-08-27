/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the CustomPages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomPages{}

// CustomPages struct for CustomPages
type CustomPages struct {
	UnderConstruction *CustomPage `json:"under_construction,omitempty"`
	FirewallError *CustomPage `json:"firewall_error,omitempty"`
	WafProtection *CustomPage `json:"waf_protection,omitempty"`
	RateLimitExceeded *CustomPage `json:"rate_limit_exceeded,omitempty"`
	SecureLinkExpired *CustomPage `json:"secure_link_expired,omitempty"`
	SecureLinkInvalid *CustomPage `json:"secure_link_invalid,omitempty"`
	Error500 *CustomPage `json:"error_500,omitempty"`
	DdosJs *CustomPage `json:"ddos_js,omitempty"`
	DdosCaptcha *CustomPage `json:"ddos_captcha,omitempty"`
}

// NewCustomPages instantiates a new CustomPages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomPages() *CustomPages {
	this := CustomPages{}
	return &this
}

// NewCustomPagesWithDefaults instantiates a new CustomPages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomPagesWithDefaults() *CustomPages {
	this := CustomPages{}
	return &this
}

// GetUnderConstruction returns the UnderConstruction field value if set, zero value otherwise.
func (o *CustomPages) GetUnderConstruction() CustomPage {
	if o == nil || IsNil(o.UnderConstruction) {
		var ret CustomPage
		return ret
	}
	return *o.UnderConstruction
}

// GetUnderConstructionOk returns a tuple with the UnderConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetUnderConstructionOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.UnderConstruction) {
		return nil, false
	}
	return o.UnderConstruction, true
}

// HasUnderConstruction returns a boolean if a field has been set.
func (o *CustomPages) HasUnderConstruction() bool {
	if o != nil && !IsNil(o.UnderConstruction) {
		return true
	}

	return false
}

// SetUnderConstruction gets a reference to the given CustomPage and assigns it to the UnderConstruction field.
func (o *CustomPages) SetUnderConstruction(v CustomPage) {
	o.UnderConstruction = &v
}

// GetFirewallError returns the FirewallError field value if set, zero value otherwise.
func (o *CustomPages) GetFirewallError() CustomPage {
	if o == nil || IsNil(o.FirewallError) {
		var ret CustomPage
		return ret
	}
	return *o.FirewallError
}

// GetFirewallErrorOk returns a tuple with the FirewallError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetFirewallErrorOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.FirewallError) {
		return nil, false
	}
	return o.FirewallError, true
}

// HasFirewallError returns a boolean if a field has been set.
func (o *CustomPages) HasFirewallError() bool {
	if o != nil && !IsNil(o.FirewallError) {
		return true
	}

	return false
}

// SetFirewallError gets a reference to the given CustomPage and assigns it to the FirewallError field.
func (o *CustomPages) SetFirewallError(v CustomPage) {
	o.FirewallError = &v
}

// GetWafProtection returns the WafProtection field value if set, zero value otherwise.
func (o *CustomPages) GetWafProtection() CustomPage {
	if o == nil || IsNil(o.WafProtection) {
		var ret CustomPage
		return ret
	}
	return *o.WafProtection
}

// GetWafProtectionOk returns a tuple with the WafProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetWafProtectionOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.WafProtection) {
		return nil, false
	}
	return o.WafProtection, true
}

// HasWafProtection returns a boolean if a field has been set.
func (o *CustomPages) HasWafProtection() bool {
	if o != nil && !IsNil(o.WafProtection) {
		return true
	}

	return false
}

// SetWafProtection gets a reference to the given CustomPage and assigns it to the WafProtection field.
func (o *CustomPages) SetWafProtection(v CustomPage) {
	o.WafProtection = &v
}

// GetRateLimitExceeded returns the RateLimitExceeded field value if set, zero value otherwise.
func (o *CustomPages) GetRateLimitExceeded() CustomPage {
	if o == nil || IsNil(o.RateLimitExceeded) {
		var ret CustomPage
		return ret
	}
	return *o.RateLimitExceeded
}

// GetRateLimitExceededOk returns a tuple with the RateLimitExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetRateLimitExceededOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.RateLimitExceeded) {
		return nil, false
	}
	return o.RateLimitExceeded, true
}

// HasRateLimitExceeded returns a boolean if a field has been set.
func (o *CustomPages) HasRateLimitExceeded() bool {
	if o != nil && !IsNil(o.RateLimitExceeded) {
		return true
	}

	return false
}

// SetRateLimitExceeded gets a reference to the given CustomPage and assigns it to the RateLimitExceeded field.
func (o *CustomPages) SetRateLimitExceeded(v CustomPage) {
	o.RateLimitExceeded = &v
}

// GetSecureLinkExpired returns the SecureLinkExpired field value if set, zero value otherwise.
func (o *CustomPages) GetSecureLinkExpired() CustomPage {
	if o == nil || IsNil(o.SecureLinkExpired) {
		var ret CustomPage
		return ret
	}
	return *o.SecureLinkExpired
}

// GetSecureLinkExpiredOk returns a tuple with the SecureLinkExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetSecureLinkExpiredOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.SecureLinkExpired) {
		return nil, false
	}
	return o.SecureLinkExpired, true
}

// HasSecureLinkExpired returns a boolean if a field has been set.
func (o *CustomPages) HasSecureLinkExpired() bool {
	if o != nil && !IsNil(o.SecureLinkExpired) {
		return true
	}

	return false
}

// SetSecureLinkExpired gets a reference to the given CustomPage and assigns it to the SecureLinkExpired field.
func (o *CustomPages) SetSecureLinkExpired(v CustomPage) {
	o.SecureLinkExpired = &v
}

// GetSecureLinkInvalid returns the SecureLinkInvalid field value if set, zero value otherwise.
func (o *CustomPages) GetSecureLinkInvalid() CustomPage {
	if o == nil || IsNil(o.SecureLinkInvalid) {
		var ret CustomPage
		return ret
	}
	return *o.SecureLinkInvalid
}

// GetSecureLinkInvalidOk returns a tuple with the SecureLinkInvalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetSecureLinkInvalidOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.SecureLinkInvalid) {
		return nil, false
	}
	return o.SecureLinkInvalid, true
}

// HasSecureLinkInvalid returns a boolean if a field has been set.
func (o *CustomPages) HasSecureLinkInvalid() bool {
	if o != nil && !IsNil(o.SecureLinkInvalid) {
		return true
	}

	return false
}

// SetSecureLinkInvalid gets a reference to the given CustomPage and assigns it to the SecureLinkInvalid field.
func (o *CustomPages) SetSecureLinkInvalid(v CustomPage) {
	o.SecureLinkInvalid = &v
}

// GetError500 returns the Error500 field value if set, zero value otherwise.
func (o *CustomPages) GetError500() CustomPage {
	if o == nil || IsNil(o.Error500) {
		var ret CustomPage
		return ret
	}
	return *o.Error500
}

// GetError500Ok returns a tuple with the Error500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetError500Ok() (*CustomPage, bool) {
	if o == nil || IsNil(o.Error500) {
		return nil, false
	}
	return o.Error500, true
}

// HasError500 returns a boolean if a field has been set.
func (o *CustomPages) HasError500() bool {
	if o != nil && !IsNil(o.Error500) {
		return true
	}

	return false
}

// SetError500 gets a reference to the given CustomPage and assigns it to the Error500 field.
func (o *CustomPages) SetError500(v CustomPage) {
	o.Error500 = &v
}

// GetDdosJs returns the DdosJs field value if set, zero value otherwise.
func (o *CustomPages) GetDdosJs() CustomPage {
	if o == nil || IsNil(o.DdosJs) {
		var ret CustomPage
		return ret
	}
	return *o.DdosJs
}

// GetDdosJsOk returns a tuple with the DdosJs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetDdosJsOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.DdosJs) {
		return nil, false
	}
	return o.DdosJs, true
}

// HasDdosJs returns a boolean if a field has been set.
func (o *CustomPages) HasDdosJs() bool {
	if o != nil && !IsNil(o.DdosJs) {
		return true
	}

	return false
}

// SetDdosJs gets a reference to the given CustomPage and assigns it to the DdosJs field.
func (o *CustomPages) SetDdosJs(v CustomPage) {
	o.DdosJs = &v
}

// GetDdosCaptcha returns the DdosCaptcha field value if set, zero value otherwise.
func (o *CustomPages) GetDdosCaptcha() CustomPage {
	if o == nil || IsNil(o.DdosCaptcha) {
		var ret CustomPage
		return ret
	}
	return *o.DdosCaptcha
}

// GetDdosCaptchaOk returns a tuple with the DdosCaptcha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomPages) GetDdosCaptchaOk() (*CustomPage, bool) {
	if o == nil || IsNil(o.DdosCaptcha) {
		return nil, false
	}
	return o.DdosCaptcha, true
}

// HasDdosCaptcha returns a boolean if a field has been set.
func (o *CustomPages) HasDdosCaptcha() bool {
	if o != nil && !IsNil(o.DdosCaptcha) {
		return true
	}

	return false
}

// SetDdosCaptcha gets a reference to the given CustomPage and assigns it to the DdosCaptcha field.
func (o *CustomPages) SetDdosCaptcha(v CustomPage) {
	o.DdosCaptcha = &v
}

func (o CustomPages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomPages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnderConstruction) {
		toSerialize["under_construction"] = o.UnderConstruction
	}
	if !IsNil(o.FirewallError) {
		toSerialize["firewall_error"] = o.FirewallError
	}
	if !IsNil(o.WafProtection) {
		toSerialize["waf_protection"] = o.WafProtection
	}
	if !IsNil(o.RateLimitExceeded) {
		toSerialize["rate_limit_exceeded"] = o.RateLimitExceeded
	}
	if !IsNil(o.SecureLinkExpired) {
		toSerialize["secure_link_expired"] = o.SecureLinkExpired
	}
	if !IsNil(o.SecureLinkInvalid) {
		toSerialize["secure_link_invalid"] = o.SecureLinkInvalid
	}
	if !IsNil(o.Error500) {
		toSerialize["error_500"] = o.Error500
	}
	if !IsNil(o.DdosJs) {
		toSerialize["ddos_js"] = o.DdosJs
	}
	if !IsNil(o.DdosCaptcha) {
		toSerialize["ddos_captcha"] = o.DdosCaptcha
	}
	return toSerialize, nil
}

type NullableCustomPages struct {
	value *CustomPages
	isSet bool
}

func (v NullableCustomPages) Get() *CustomPages {
	return v.value
}

func (v *NullableCustomPages) Set(val *CustomPages) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomPages) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomPages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomPages(val *CustomPages) *NullableCustomPages {
	return &NullableCustomPages{value: val, isSet: true}
}

func (v NullableCustomPages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomPages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


