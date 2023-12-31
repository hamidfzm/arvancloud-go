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

// checks if the CORSRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CORSRule{}

// CORSRule Specifies a cross-origin access rule for an ArvanCloud S3 bucket.
type CORSRule struct {
	ID *string `json:"ID,omitempty"`
	AllowedHeaders *Array `json:"AllowedHeaders,omitempty"`
	AllowedMethods Array `json:"AllowedMethods"`
	AllowedOrigins Array `json:"AllowedOrigins"`
	ExposeHeaders *Array `json:"ExposeHeaders,omitempty"`
	MaxAgeSeconds *int32 `json:"MaxAgeSeconds,omitempty"`
}

// NewCORSRule instantiates a new CORSRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCORSRule(allowedMethods Array, allowedOrigins Array) *CORSRule {
	this := CORSRule{}
	this.AllowedMethods = allowedMethods
	this.AllowedOrigins = allowedOrigins
	return &this
}

// NewCORSRuleWithDefaults instantiates a new CORSRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCORSRuleWithDefaults() *CORSRule {
	this := CORSRule{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CORSRule) GetID() string {
	if o == nil || IsNil(o.ID) {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CORSRule) GetIDOk() (*string, bool) {
	if o == nil || IsNil(o.ID) {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CORSRule) HasID() bool {
	if o != nil && !IsNil(o.ID) {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CORSRule) SetID(v string) {
	o.ID = &v
}

// GetAllowedHeaders returns the AllowedHeaders field value if set, zero value otherwise.
func (o *CORSRule) GetAllowedHeaders() Array {
	if o == nil || IsNil(o.AllowedHeaders) {
		var ret Array
		return ret
	}
	return *o.AllowedHeaders
}

// GetAllowedHeadersOk returns a tuple with the AllowedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CORSRule) GetAllowedHeadersOk() (*Array, bool) {
	if o == nil || IsNil(o.AllowedHeaders) {
		return nil, false
	}
	return o.AllowedHeaders, true
}

// HasAllowedHeaders returns a boolean if a field has been set.
func (o *CORSRule) HasAllowedHeaders() bool {
	if o != nil && !IsNil(o.AllowedHeaders) {
		return true
	}

	return false
}

// SetAllowedHeaders gets a reference to the given Array and assigns it to the AllowedHeaders field.
func (o *CORSRule) SetAllowedHeaders(v Array) {
	o.AllowedHeaders = &v
}

// GetAllowedMethods returns the AllowedMethods field value
func (o *CORSRule) GetAllowedMethods() Array {
	if o == nil {
		var ret Array
		return ret
	}

	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value
// and a boolean to check if the value has been set.
func (o *CORSRule) GetAllowedMethodsOk() (*Array, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedMethods, true
}

// SetAllowedMethods sets field value
func (o *CORSRule) SetAllowedMethods(v Array) {
	o.AllowedMethods = v
}

// GetAllowedOrigins returns the AllowedOrigins field value
func (o *CORSRule) GetAllowedOrigins() Array {
	if o == nil {
		var ret Array
		return ret
	}

	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value
// and a boolean to check if the value has been set.
func (o *CORSRule) GetAllowedOriginsOk() (*Array, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowedOrigins, true
}

// SetAllowedOrigins sets field value
func (o *CORSRule) SetAllowedOrigins(v Array) {
	o.AllowedOrigins = v
}

// GetExposeHeaders returns the ExposeHeaders field value if set, zero value otherwise.
func (o *CORSRule) GetExposeHeaders() Array {
	if o == nil || IsNil(o.ExposeHeaders) {
		var ret Array
		return ret
	}
	return *o.ExposeHeaders
}

// GetExposeHeadersOk returns a tuple with the ExposeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CORSRule) GetExposeHeadersOk() (*Array, bool) {
	if o == nil || IsNil(o.ExposeHeaders) {
		return nil, false
	}
	return o.ExposeHeaders, true
}

// HasExposeHeaders returns a boolean if a field has been set.
func (o *CORSRule) HasExposeHeaders() bool {
	if o != nil && !IsNil(o.ExposeHeaders) {
		return true
	}

	return false
}

// SetExposeHeaders gets a reference to the given Array and assigns it to the ExposeHeaders field.
func (o *CORSRule) SetExposeHeaders(v Array) {
	o.ExposeHeaders = &v
}

// GetMaxAgeSeconds returns the MaxAgeSeconds field value if set, zero value otherwise.
func (o *CORSRule) GetMaxAgeSeconds() int32 {
	if o == nil || IsNil(o.MaxAgeSeconds) {
		var ret int32
		return ret
	}
	return *o.MaxAgeSeconds
}

// GetMaxAgeSecondsOk returns a tuple with the MaxAgeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CORSRule) GetMaxAgeSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAgeSeconds) {
		return nil, false
	}
	return o.MaxAgeSeconds, true
}

// HasMaxAgeSeconds returns a boolean if a field has been set.
func (o *CORSRule) HasMaxAgeSeconds() bool {
	if o != nil && !IsNil(o.MaxAgeSeconds) {
		return true
	}

	return false
}

// SetMaxAgeSeconds gets a reference to the given int32 and assigns it to the MaxAgeSeconds field.
func (o *CORSRule) SetMaxAgeSeconds(v int32) {
	o.MaxAgeSeconds = &v
}

func (o CORSRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CORSRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ID) {
		toSerialize["ID"] = o.ID
	}
	if !IsNil(o.AllowedHeaders) {
		toSerialize["AllowedHeaders"] = o.AllowedHeaders
	}
	toSerialize["AllowedMethods"] = o.AllowedMethods
	toSerialize["AllowedOrigins"] = o.AllowedOrigins
	if !IsNil(o.ExposeHeaders) {
		toSerialize["ExposeHeaders"] = o.ExposeHeaders
	}
	if !IsNil(o.MaxAgeSeconds) {
		toSerialize["MaxAgeSeconds"] = o.MaxAgeSeconds
	}
	return toSerialize, nil
}

type NullableCORSRule struct {
	value *CORSRule
	isSet bool
}

func (v NullableCORSRule) Get() *CORSRule {
	return v.value
}

func (v *NullableCORSRule) Set(val *CORSRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCORSRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCORSRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCORSRule(val *CORSRule) *NullableCORSRule {
	return &NullableCORSRule{value: val, isSet: true}
}

func (v NullableCORSRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCORSRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


