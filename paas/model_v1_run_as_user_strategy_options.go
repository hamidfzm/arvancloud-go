/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1RunAsUserStrategyOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RunAsUserStrategyOptions{}

// V1RunAsUserStrategyOptions RunAsUserStrategyOptions defines the strategy type and any options used to create the strategy.
type V1RunAsUserStrategyOptions struct {
	// Type is the strategy that will dictate what RunAsUser is used in the SecurityContext.
	Type *string `json:"type,omitempty"`
	// UID is the user id that containers must run as.  Required for the MustRunAs strategy if not using namespace/service account allocated uids.
	Uid *int64 `json:"uid,omitempty"`
	// UIDRangeMax defines the max value for a strategy that allocates by range.
	UidRangeMax *int64 `json:"uidRangeMax,omitempty"`
	// UIDRangeMin defines the min value for a strategy that allocates by range.
	UidRangeMin *int64 `json:"uidRangeMin,omitempty"`
}

// NewV1RunAsUserStrategyOptions instantiates a new V1RunAsUserStrategyOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RunAsUserStrategyOptions() *V1RunAsUserStrategyOptions {
	this := V1RunAsUserStrategyOptions{}
	return &this
}

// NewV1RunAsUserStrategyOptionsWithDefaults instantiates a new V1RunAsUserStrategyOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RunAsUserStrategyOptionsWithDefaults() *V1RunAsUserStrategyOptions {
	this := V1RunAsUserStrategyOptions{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1RunAsUserStrategyOptions) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RunAsUserStrategyOptions) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1RunAsUserStrategyOptions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1RunAsUserStrategyOptions) SetType(v string) {
	o.Type = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *V1RunAsUserStrategyOptions) GetUid() int64 {
	if o == nil || IsNil(o.Uid) {
		var ret int64
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RunAsUserStrategyOptions) GetUidOk() (*int64, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *V1RunAsUserStrategyOptions) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int64 and assigns it to the Uid field.
func (o *V1RunAsUserStrategyOptions) SetUid(v int64) {
	o.Uid = &v
}

// GetUidRangeMax returns the UidRangeMax field value if set, zero value otherwise.
func (o *V1RunAsUserStrategyOptions) GetUidRangeMax() int64 {
	if o == nil || IsNil(o.UidRangeMax) {
		var ret int64
		return ret
	}
	return *o.UidRangeMax
}

// GetUidRangeMaxOk returns a tuple with the UidRangeMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RunAsUserStrategyOptions) GetUidRangeMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.UidRangeMax) {
		return nil, false
	}
	return o.UidRangeMax, true
}

// HasUidRangeMax returns a boolean if a field has been set.
func (o *V1RunAsUserStrategyOptions) HasUidRangeMax() bool {
	if o != nil && !IsNil(o.UidRangeMax) {
		return true
	}

	return false
}

// SetUidRangeMax gets a reference to the given int64 and assigns it to the UidRangeMax field.
func (o *V1RunAsUserStrategyOptions) SetUidRangeMax(v int64) {
	o.UidRangeMax = &v
}

// GetUidRangeMin returns the UidRangeMin field value if set, zero value otherwise.
func (o *V1RunAsUserStrategyOptions) GetUidRangeMin() int64 {
	if o == nil || IsNil(o.UidRangeMin) {
		var ret int64
		return ret
	}
	return *o.UidRangeMin
}

// GetUidRangeMinOk returns a tuple with the UidRangeMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RunAsUserStrategyOptions) GetUidRangeMinOk() (*int64, bool) {
	if o == nil || IsNil(o.UidRangeMin) {
		return nil, false
	}
	return o.UidRangeMin, true
}

// HasUidRangeMin returns a boolean if a field has been set.
func (o *V1RunAsUserStrategyOptions) HasUidRangeMin() bool {
	if o != nil && !IsNil(o.UidRangeMin) {
		return true
	}

	return false
}

// SetUidRangeMin gets a reference to the given int64 and assigns it to the UidRangeMin field.
func (o *V1RunAsUserStrategyOptions) SetUidRangeMin(v int64) {
	o.UidRangeMin = &v
}

func (o V1RunAsUserStrategyOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RunAsUserStrategyOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.UidRangeMax) {
		toSerialize["uidRangeMax"] = o.UidRangeMax
	}
	if !IsNil(o.UidRangeMin) {
		toSerialize["uidRangeMin"] = o.UidRangeMin
	}
	return toSerialize, nil
}

type NullableV1RunAsUserStrategyOptions struct {
	value *V1RunAsUserStrategyOptions
	isSet bool
}

func (v NullableV1RunAsUserStrategyOptions) Get() *V1RunAsUserStrategyOptions {
	return v.value
}

func (v *NullableV1RunAsUserStrategyOptions) Set(val *V1RunAsUserStrategyOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RunAsUserStrategyOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RunAsUserStrategyOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RunAsUserStrategyOptions(val *V1RunAsUserStrategyOptions) *NullableV1RunAsUserStrategyOptions {
	return &NullableV1RunAsUserStrategyOptions{value: val, isSet: true}
}

func (v NullableV1RunAsUserStrategyOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RunAsUserStrategyOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

