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

// checks if the NotificationConfigurationFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationConfigurationFilter{}

// NotificationConfigurationFilter Specifies object key name filtering rules.
type NotificationConfigurationFilter struct {
	Key *NotificationConfigurationFilterKey `json:"Key,omitempty"`
}

// NewNotificationConfigurationFilter instantiates a new NotificationConfigurationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationConfigurationFilter() *NotificationConfigurationFilter {
	this := NotificationConfigurationFilter{}
	return &this
}

// NewNotificationConfigurationFilterWithDefaults instantiates a new NotificationConfigurationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationConfigurationFilterWithDefaults() *NotificationConfigurationFilter {
	this := NotificationConfigurationFilter{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *NotificationConfigurationFilter) GetKey() NotificationConfigurationFilterKey {
	if o == nil || IsNil(o.Key) {
		var ret NotificationConfigurationFilterKey
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfigurationFilter) GetKeyOk() (*NotificationConfigurationFilterKey, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *NotificationConfigurationFilter) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given NotificationConfigurationFilterKey and assigns it to the Key field.
func (o *NotificationConfigurationFilter) SetKey(v NotificationConfigurationFilterKey) {
	o.Key = &v
}

func (o NotificationConfigurationFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationConfigurationFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	return toSerialize, nil
}

type NullableNotificationConfigurationFilter struct {
	value *NotificationConfigurationFilter
	isSet bool
}

func (v NullableNotificationConfigurationFilter) Get() *NotificationConfigurationFilter {
	return v.value
}

func (v *NullableNotificationConfigurationFilter) Set(val *NotificationConfigurationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationConfigurationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationConfigurationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationConfigurationFilter(val *NotificationConfigurationFilter) *NullableNotificationConfigurationFilter {
	return &NullableNotificationConfigurationFilter{value: val, isSet: true}
}

func (v NullableNotificationConfigurationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationConfigurationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


