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

// TransitionStorageClass the model 'TransitionStorageClass'
type TransitionStorageClass string

// List of TransitionStorageClass
const (
	TRANSITIONSTORAGECLASS_GLACIER TransitionStorageClass = "GLACIER"
	TRANSITIONSTORAGECLASS_STANDARD_IA TransitionStorageClass = "STANDARD_IA"
	TRANSITIONSTORAGECLASS_ONEZONE_IA TransitionStorageClass = "ONEZONE_IA"
	TRANSITIONSTORAGECLASS_INTELLIGENT_TIERING TransitionStorageClass = "INTELLIGENT_TIERING"
	TRANSITIONSTORAGECLASS_DEEP_ARCHIVE TransitionStorageClass = "DEEP_ARCHIVE"
)

// All allowed values of TransitionStorageClass enum
var AllowedTransitionStorageClassEnumValues = []TransitionStorageClass{
	"GLACIER",
	"STANDARD_IA",
	"ONEZONE_IA",
	"INTELLIGENT_TIERING",
	"DEEP_ARCHIVE",
}

func (v *TransitionStorageClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransitionStorageClass(value)
	for _, existing := range AllowedTransitionStorageClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransitionStorageClass", value)
}

// NewTransitionStorageClassFromValue returns a pointer to a valid TransitionStorageClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransitionStorageClassFromValue(v string) (*TransitionStorageClass, error) {
	ev := TransitionStorageClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransitionStorageClass: valid values are %v", v, AllowedTransitionStorageClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransitionStorageClass) IsValid() bool {
	for _, existing := range AllowedTransitionStorageClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransitionStorageClass value
func (v TransitionStorageClass) Ptr() *TransitionStorageClass {
	return &v
}

type NullableTransitionStorageClass struct {
	value *TransitionStorageClass
	isSet bool
}

func (v NullableTransitionStorageClass) Get() *TransitionStorageClass {
	return v.value
}

func (v *NullableTransitionStorageClass) Set(val *TransitionStorageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableTransitionStorageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableTransitionStorageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransitionStorageClass(val *TransitionStorageClass) *NullableTransitionStorageClass {
	return &NullableTransitionStorageClass{value: val, isSet: true}
}

func (v NullableTransitionStorageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransitionStorageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

