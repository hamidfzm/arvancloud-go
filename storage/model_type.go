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

// Type the model 'Type'
type Type string

// List of Type
const (
	TYPE_CANONICAL_USER Type = "CanonicalUser"
	TYPE_AMAZON_CUSTOMER_BY_EMAIL Type = "AmazonCustomerByEmail"
	TYPE_GROUP Type = "Group"
)

// All allowed values of Type enum
var AllowedTypeEnumValues = []Type{
	"CanonicalUser",
	"AmazonCustomerByEmail",
	"Group",
}

func (v *Type) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Type(value)
	for _, existing := range AllowedTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Type", value)
}

// NewTypeFromValue returns a pointer to a valid Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeFromValue(v string) (*Type, error) {
	ev := Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Type: valid values are %v", v, AllowedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Type) IsValid() bool {
	for _, existing := range AllowedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Type value
func (v Type) Ptr() *Type {
	return &v
}

type NullableType struct {
	value *Type
	isSet bool
}

func (v NullableType) Get() *Type {
	return v.value
}

func (v *NullableType) Set(val *Type) {
	v.value = val
	v.isSet = true
}

func (v NullableType) IsSet() bool {
	return v.isSet
}

func (v *NullableType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableType(val *Type) *NullableType {
	return &NullableType{value: val, isSet: true}
}

func (v NullableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
