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

// MetadataDirective the model 'MetadataDirective'
type MetadataDirective string

// List of MetadataDirective
const (
	METADATADIRECTIVE_COPY MetadataDirective = "COPY"
	METADATADIRECTIVE_REPLACE MetadataDirective = "REPLACE"
)

// All allowed values of MetadataDirective enum
var AllowedMetadataDirectiveEnumValues = []MetadataDirective{
	"COPY",
	"REPLACE",
}

func (v *MetadataDirective) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataDirective(value)
	for _, existing := range AllowedMetadataDirectiveEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataDirective", value)
}

// NewMetadataDirectiveFromValue returns a pointer to a valid MetadataDirective
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataDirectiveFromValue(v string) (*MetadataDirective, error) {
	ev := MetadataDirective(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataDirective: valid values are %v", v, AllowedMetadataDirectiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataDirective) IsValid() bool {
	for _, existing := range AllowedMetadataDirectiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataDirective value
func (v MetadataDirective) Ptr() *MetadataDirective {
	return &v
}

type NullableMetadataDirective struct {
	value *MetadataDirective
	isSet bool
}

func (v NullableMetadataDirective) Get() *MetadataDirective {
	return v.value
}

func (v *NullableMetadataDirective) Set(val *MetadataDirective) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataDirective) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataDirective) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataDirective(val *MetadataDirective) *NullableMetadataDirective {
	return &NullableMetadataDirective{value: val, isSet: true}
}

func (v NullableMetadataDirective) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataDirective) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

