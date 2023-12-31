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

// IntelligentTieringAccessTier the model 'IntelligentTieringAccessTier'
type IntelligentTieringAccessTier string

// List of IntelligentTieringAccessTier
const (
	INTELLIGENTTIERINGACCESSTIER_ARCHIVE_ACCESS IntelligentTieringAccessTier = "ARCHIVE_ACCESS"
	INTELLIGENTTIERINGACCESSTIER_DEEP_ARCHIVE_ACCESS IntelligentTieringAccessTier = "DEEP_ARCHIVE_ACCESS"
)

// All allowed values of IntelligentTieringAccessTier enum
var AllowedIntelligentTieringAccessTierEnumValues = []IntelligentTieringAccessTier{
	"ARCHIVE_ACCESS",
	"DEEP_ARCHIVE_ACCESS",
}

func (v *IntelligentTieringAccessTier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntelligentTieringAccessTier(value)
	for _, existing := range AllowedIntelligentTieringAccessTierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntelligentTieringAccessTier", value)
}

// NewIntelligentTieringAccessTierFromValue returns a pointer to a valid IntelligentTieringAccessTier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntelligentTieringAccessTierFromValue(v string) (*IntelligentTieringAccessTier, error) {
	ev := IntelligentTieringAccessTier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntelligentTieringAccessTier: valid values are %v", v, AllowedIntelligentTieringAccessTierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntelligentTieringAccessTier) IsValid() bool {
	for _, existing := range AllowedIntelligentTieringAccessTierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntelligentTieringAccessTier value
func (v IntelligentTieringAccessTier) Ptr() *IntelligentTieringAccessTier {
	return &v
}

type NullableIntelligentTieringAccessTier struct {
	value *IntelligentTieringAccessTier
	isSet bool
}

func (v NullableIntelligentTieringAccessTier) Get() *IntelligentTieringAccessTier {
	return v.value
}

func (v *NullableIntelligentTieringAccessTier) Set(val *IntelligentTieringAccessTier) {
	v.value = val
	v.isSet = true
}

func (v NullableIntelligentTieringAccessTier) IsSet() bool {
	return v.isSet
}

func (v *NullableIntelligentTieringAccessTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntelligentTieringAccessTier(val *IntelligentTieringAccessTier) *NullableIntelligentTieringAccessTier {
	return &NullableIntelligentTieringAccessTier{value: val, isSet: true}
}

func (v NullableIntelligentTieringAccessTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntelligentTieringAccessTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

