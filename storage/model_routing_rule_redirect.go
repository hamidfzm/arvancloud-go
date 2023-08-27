/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
)

// checks if the RoutingRuleRedirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingRuleRedirect{}

// RoutingRuleRedirect struct for RoutingRuleRedirect
type RoutingRuleRedirect struct {
	HostName *string `json:"HostName,omitempty"`
	HttpRedirectCode *string `json:"HttpRedirectCode,omitempty"`
	Protocol *Protocol `json:"Protocol,omitempty"`
	ReplaceKeyPrefixWith *string `json:"ReplaceKeyPrefixWith,omitempty"`
	ReplaceKeyWith *string `json:"ReplaceKeyWith,omitempty"`
}

// NewRoutingRuleRedirect instantiates a new RoutingRuleRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingRuleRedirect() *RoutingRuleRedirect {
	this := RoutingRuleRedirect{}
	return &this
}

// NewRoutingRuleRedirectWithDefaults instantiates a new RoutingRuleRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingRuleRedirectWithDefaults() *RoutingRuleRedirect {
	this := RoutingRuleRedirect{}
	return &this
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *RoutingRuleRedirect) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleRedirect) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *RoutingRuleRedirect) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *RoutingRuleRedirect) SetHostName(v string) {
	o.HostName = &v
}

// GetHttpRedirectCode returns the HttpRedirectCode field value if set, zero value otherwise.
func (o *RoutingRuleRedirect) GetHttpRedirectCode() string {
	if o == nil || IsNil(o.HttpRedirectCode) {
		var ret string
		return ret
	}
	return *o.HttpRedirectCode
}

// GetHttpRedirectCodeOk returns a tuple with the HttpRedirectCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleRedirect) GetHttpRedirectCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpRedirectCode) {
		return nil, false
	}
	return o.HttpRedirectCode, true
}

// HasHttpRedirectCode returns a boolean if a field has been set.
func (o *RoutingRuleRedirect) HasHttpRedirectCode() bool {
	if o != nil && !IsNil(o.HttpRedirectCode) {
		return true
	}

	return false
}

// SetHttpRedirectCode gets a reference to the given string and assigns it to the HttpRedirectCode field.
func (o *RoutingRuleRedirect) SetHttpRedirectCode(v string) {
	o.HttpRedirectCode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *RoutingRuleRedirect) GetProtocol() Protocol {
	if o == nil || IsNil(o.Protocol) {
		var ret Protocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleRedirect) GetProtocolOk() (*Protocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *RoutingRuleRedirect) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *RoutingRuleRedirect) SetProtocol(v Protocol) {
	o.Protocol = &v
}

// GetReplaceKeyPrefixWith returns the ReplaceKeyPrefixWith field value if set, zero value otherwise.
func (o *RoutingRuleRedirect) GetReplaceKeyPrefixWith() string {
	if o == nil || IsNil(o.ReplaceKeyPrefixWith) {
		var ret string
		return ret
	}
	return *o.ReplaceKeyPrefixWith
}

// GetReplaceKeyPrefixWithOk returns a tuple with the ReplaceKeyPrefixWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleRedirect) GetReplaceKeyPrefixWithOk() (*string, bool) {
	if o == nil || IsNil(o.ReplaceKeyPrefixWith) {
		return nil, false
	}
	return o.ReplaceKeyPrefixWith, true
}

// HasReplaceKeyPrefixWith returns a boolean if a field has been set.
func (o *RoutingRuleRedirect) HasReplaceKeyPrefixWith() bool {
	if o != nil && !IsNil(o.ReplaceKeyPrefixWith) {
		return true
	}

	return false
}

// SetReplaceKeyPrefixWith gets a reference to the given string and assigns it to the ReplaceKeyPrefixWith field.
func (o *RoutingRuleRedirect) SetReplaceKeyPrefixWith(v string) {
	o.ReplaceKeyPrefixWith = &v
}

// GetReplaceKeyWith returns the ReplaceKeyWith field value if set, zero value otherwise.
func (o *RoutingRuleRedirect) GetReplaceKeyWith() string {
	if o == nil || IsNil(o.ReplaceKeyWith) {
		var ret string
		return ret
	}
	return *o.ReplaceKeyWith
}

// GetReplaceKeyWithOk returns a tuple with the ReplaceKeyWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleRedirect) GetReplaceKeyWithOk() (*string, bool) {
	if o == nil || IsNil(o.ReplaceKeyWith) {
		return nil, false
	}
	return o.ReplaceKeyWith, true
}

// HasReplaceKeyWith returns a boolean if a field has been set.
func (o *RoutingRuleRedirect) HasReplaceKeyWith() bool {
	if o != nil && !IsNil(o.ReplaceKeyWith) {
		return true
	}

	return false
}

// SetReplaceKeyWith gets a reference to the given string and assigns it to the ReplaceKeyWith field.
func (o *RoutingRuleRedirect) SetReplaceKeyWith(v string) {
	o.ReplaceKeyWith = &v
}

func (o RoutingRuleRedirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingRuleRedirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HostName) {
		toSerialize["HostName"] = o.HostName
	}
	if !IsNil(o.HttpRedirectCode) {
		toSerialize["HttpRedirectCode"] = o.HttpRedirectCode
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.ReplaceKeyPrefixWith) {
		toSerialize["ReplaceKeyPrefixWith"] = o.ReplaceKeyPrefixWith
	}
	if !IsNil(o.ReplaceKeyWith) {
		toSerialize["ReplaceKeyWith"] = o.ReplaceKeyWith
	}
	return toSerialize, nil
}

type NullableRoutingRuleRedirect struct {
	value *RoutingRuleRedirect
	isSet bool
}

func (v NullableRoutingRuleRedirect) Get() *RoutingRuleRedirect {
	return v.value
}

func (v *NullableRoutingRuleRedirect) Set(val *RoutingRuleRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingRuleRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingRuleRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingRuleRedirect(val *RoutingRuleRedirect) *NullableRoutingRuleRedirect {
	return &NullableRoutingRuleRedirect{value: val, isSet: true}
}

func (v NullableRoutingRuleRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingRuleRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

