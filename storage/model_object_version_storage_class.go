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

// ObjectVersionStorageClass the model 'ObjectVersionStorageClass'
type ObjectVersionStorageClass string

// List of ObjectVersionStorageClass
const (
	OBJECTVERSIONSTORAGECLASS_STANDARD ObjectVersionStorageClass = "STANDARD"
)

// All allowed values of ObjectVersionStorageClass enum
var AllowedObjectVersionStorageClassEnumValues = []ObjectVersionStorageClass{
	"STANDARD",
}

func (v *ObjectVersionStorageClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjectVersionStorageClass(value)
	for _, existing := range AllowedObjectVersionStorageClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjectVersionStorageClass", value)
}

// NewObjectVersionStorageClassFromValue returns a pointer to a valid ObjectVersionStorageClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjectVersionStorageClassFromValue(v string) (*ObjectVersionStorageClass, error) {
	ev := ObjectVersionStorageClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjectVersionStorageClass: valid values are %v", v, AllowedObjectVersionStorageClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjectVersionStorageClass) IsValid() bool {
	for _, existing := range AllowedObjectVersionStorageClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectVersionStorageClass value
func (v ObjectVersionStorageClass) Ptr() *ObjectVersionStorageClass {
	return &v
}

type NullableObjectVersionStorageClass struct {
	value *ObjectVersionStorageClass
	isSet bool
}

func (v NullableObjectVersionStorageClass) Get() *ObjectVersionStorageClass {
	return v.value
}

func (v *NullableObjectVersionStorageClass) Set(val *ObjectVersionStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectVersionStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectVersionStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectVersionStorageClass(val *ObjectVersionStorageClass) *NullableObjectVersionStorageClass {
	return &NullableObjectVersionStorageClass{value: val, isSet: true}
}

func (v NullableObjectVersionStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectVersionStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

