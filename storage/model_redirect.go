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

// checks if the Redirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Redirect{}

// Redirect Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.
type Redirect struct {
	HostName *string `json:"HostName,omitempty"`
	HttpRedirectCode *string `json:"HttpRedirectCode,omitempty"`
	Protocol *Protocol `json:"Protocol,omitempty"`
	ReplaceKeyPrefixWith *string `json:"ReplaceKeyPrefixWith,omitempty"`
	ReplaceKeyWith *string `json:"ReplaceKeyWith,omitempty"`
}

// NewRedirect instantiates a new Redirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirect() *Redirect {
	this := Redirect{}
	return &this
}

// NewRedirectWithDefaults instantiates a new Redirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectWithDefaults() *Redirect {
	this := Redirect{}
	return &this
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *Redirect) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *Redirect) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *Redirect) SetHostName(v string) {
	o.HostName = &v
}

// GetHttpRedirectCode returns the HttpRedirectCode field value if set, zero value otherwise.
func (o *Redirect) GetHttpRedirectCode() string {
	if o == nil || IsNil(o.HttpRedirectCode) {
		var ret string
		return ret
	}
	return *o.HttpRedirectCode
}

// GetHttpRedirectCodeOk returns a tuple with the HttpRedirectCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetHttpRedirectCodeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpRedirectCode) {
		return nil, false
	}
	return o.HttpRedirectCode, true
}

// HasHttpRedirectCode returns a boolean if a field has been set.
func (o *Redirect) HasHttpRedirectCode() bool {
	if o != nil && !IsNil(o.HttpRedirectCode) {
		return true
	}

	return false
}

// SetHttpRedirectCode gets a reference to the given string and assigns it to the HttpRedirectCode field.
func (o *Redirect) SetHttpRedirectCode(v string) {
	o.HttpRedirectCode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Redirect) GetProtocol() Protocol {
	if o == nil || IsNil(o.Protocol) {
		var ret Protocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetProtocolOk() (*Protocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Redirect) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *Redirect) SetProtocol(v Protocol) {
	o.Protocol = &v
}

// GetReplaceKeyPrefixWith returns the ReplaceKeyPrefixWith field value if set, zero value otherwise.
func (o *Redirect) GetReplaceKeyPrefixWith() string {
	if o == nil || IsNil(o.ReplaceKeyPrefixWith) {
		var ret string
		return ret
	}
	return *o.ReplaceKeyPrefixWith
}

// GetReplaceKeyPrefixWithOk returns a tuple with the ReplaceKeyPrefixWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetReplaceKeyPrefixWithOk() (*string, bool) {
	if o == nil || IsNil(o.ReplaceKeyPrefixWith) {
		return nil, false
	}
	return o.ReplaceKeyPrefixWith, true
}

// HasReplaceKeyPrefixWith returns a boolean if a field has been set.
func (o *Redirect) HasReplaceKeyPrefixWith() bool {
	if o != nil && !IsNil(o.ReplaceKeyPrefixWith) {
		return true
	}

	return false
}

// SetReplaceKeyPrefixWith gets a reference to the given string and assigns it to the ReplaceKeyPrefixWith field.
func (o *Redirect) SetReplaceKeyPrefixWith(v string) {
	o.ReplaceKeyPrefixWith = &v
}

// GetReplaceKeyWith returns the ReplaceKeyWith field value if set, zero value otherwise.
func (o *Redirect) GetReplaceKeyWith() string {
	if o == nil || IsNil(o.ReplaceKeyWith) {
		var ret string
		return ret
	}
	return *o.ReplaceKeyWith
}

// GetReplaceKeyWithOk returns a tuple with the ReplaceKeyWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Redirect) GetReplaceKeyWithOk() (*string, bool) {
	if o == nil || IsNil(o.ReplaceKeyWith) {
		return nil, false
	}
	return o.ReplaceKeyWith, true
}

// HasReplaceKeyWith returns a boolean if a field has been set.
func (o *Redirect) HasReplaceKeyWith() bool {
	if o != nil && !IsNil(o.ReplaceKeyWith) {
		return true
	}

	return false
}

// SetReplaceKeyWith gets a reference to the given string and assigns it to the ReplaceKeyWith field.
func (o *Redirect) SetReplaceKeyWith(v string) {
	o.ReplaceKeyWith = &v
}

func (o Redirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Redirect) ToMap() (map[string]interface{}, error) {
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

type NullableRedirect struct {
	value *Redirect
	isSet bool
}

func (v NullableRedirect) Get() *Redirect {
	return v.value
}

func (v *NullableRedirect) Set(val *Redirect) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirect(val *Redirect) *NullableRedirect {
	return &NullableRedirect{value: val, isSet: true}
}

func (v NullableRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

