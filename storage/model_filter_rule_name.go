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

// FilterRuleName the model 'FilterRuleName'
type FilterRuleName string

// List of FilterRuleName
const (
	FILTERRULENAME_PREFIX FilterRuleName = "prefix"
	FILTERRULENAME_SUFFIX FilterRuleName = "suffix"
)

// All allowed values of FilterRuleName enum
var AllowedFilterRuleNameEnumValues = []FilterRuleName{
	"prefix",
	"suffix",
}

func (v *FilterRuleName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterRuleName(value)
	for _, existing := range AllowedFilterRuleNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterRuleName", value)
}

// NewFilterRuleNameFromValue returns a pointer to a valid FilterRuleName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterRuleNameFromValue(v string) (*FilterRuleName, error) {
	ev := FilterRuleName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterRuleName: valid values are %v", v, AllowedFilterRuleNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterRuleName) IsValid() bool {
	for _, existing := range AllowedFilterRuleNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterRuleName value
func (v FilterRuleName) Ptr() *FilterRuleName {
	return &v
}

type NullableFilterRuleName struct {
	value *FilterRuleName
	isSet bool
}

func (v NullableFilterRuleName) Get() *FilterRuleName {
	return v.value
}

func (v *NullableFilterRuleName) Set(val *FilterRuleName) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterRuleName) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterRuleName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterRuleName(val *FilterRuleName) *NullableFilterRuleName {
	return &NullableFilterRuleName{value: val, isSet: true}
}

func (v NullableFilterRuleName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterRuleName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
