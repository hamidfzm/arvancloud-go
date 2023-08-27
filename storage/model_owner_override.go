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

// OwnerOverride the model 'OwnerOverride'
type OwnerOverride string

// List of OwnerOverride
const (
	OWNEROVERRIDE_DESTINATION OwnerOverride = "Destination"
)

// All allowed values of OwnerOverride enum
var AllowedOwnerOverrideEnumValues = []OwnerOverride{
	"Destination",
}

func (v *OwnerOverride) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OwnerOverride(value)
	for _, existing := range AllowedOwnerOverrideEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OwnerOverride", value)
}

// NewOwnerOverrideFromValue returns a pointer to a valid OwnerOverride
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOwnerOverrideFromValue(v string) (*OwnerOverride, error) {
	ev := OwnerOverride(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OwnerOverride: valid values are %v", v, AllowedOwnerOverrideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OwnerOverride) IsValid() bool {
	for _, existing := range AllowedOwnerOverrideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnerOverride value
func (v OwnerOverride) Ptr() *OwnerOverride {
	return &v
}

type NullableOwnerOverride struct {
	value *OwnerOverride
	isSet bool
}

func (v NullableOwnerOverride) Get() *OwnerOverride {
	return v.value
}

func (v *NullableOwnerOverride) Set(val *OwnerOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerOverride(val *OwnerOverride) *NullableOwnerOverride {
	return &NullableOwnerOverride{value: val, isSet: true}
}

func (v NullableOwnerOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

