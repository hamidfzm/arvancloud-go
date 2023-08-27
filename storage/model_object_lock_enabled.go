/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"fmt"
)

// ObjectLockEnabled the model 'ObjectLockEnabled'
type ObjectLockEnabled string

// List of ObjectLockEnabled
const (
	OBJECTLOCKENABLED_ENABLED ObjectLockEnabled = "Enabled"
)

// All allowed values of ObjectLockEnabled enum
var AllowedObjectLockEnabledEnumValues = []ObjectLockEnabled{
	"Enabled",
}

func (v *ObjectLockEnabled) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjectLockEnabled(value)
	for _, existing := range AllowedObjectLockEnabledEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjectLockEnabled", value)
}

// NewObjectLockEnabledFromValue returns a pointer to a valid ObjectLockEnabled
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjectLockEnabledFromValue(v string) (*ObjectLockEnabled, error) {
	ev := ObjectLockEnabled(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjectLockEnabled: valid values are %v", v, AllowedObjectLockEnabledEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjectLockEnabled) IsValid() bool {
	for _, existing := range AllowedObjectLockEnabledEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectLockEnabled value
func (v ObjectLockEnabled) Ptr() *ObjectLockEnabled {
	return &v
}

type NullableObjectLockEnabled struct {
	value *ObjectLockEnabled
	isSet bool
}

func (v NullableObjectLockEnabled) Get() *ObjectLockEnabled {
	return v.value
}

func (v *NullableObjectLockEnabled) Set(val *ObjectLockEnabled) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLockEnabled) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLockEnabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLockEnabled(val *ObjectLockEnabled) *NullableObjectLockEnabled {
	return &NullableObjectLockEnabled{value: val, isSet: true}
}

func (v NullableObjectLockEnabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLockEnabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
