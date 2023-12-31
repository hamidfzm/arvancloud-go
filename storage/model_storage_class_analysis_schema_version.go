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

// StorageClassAnalysisSchemaVersion the model 'StorageClassAnalysisSchemaVersion'
type StorageClassAnalysisSchemaVersion string

// List of StorageClassAnalysisSchemaVersion
const (
	STORAGECLASSANALYSISSCHEMAVERSION_V_1 StorageClassAnalysisSchemaVersion = "V_1"
)

// All allowed values of StorageClassAnalysisSchemaVersion enum
var AllowedStorageClassAnalysisSchemaVersionEnumValues = []StorageClassAnalysisSchemaVersion{
	"V_1",
}

func (v *StorageClassAnalysisSchemaVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageClassAnalysisSchemaVersion(value)
	for _, existing := range AllowedStorageClassAnalysisSchemaVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageClassAnalysisSchemaVersion", value)
}

// NewStorageClassAnalysisSchemaVersionFromValue returns a pointer to a valid StorageClassAnalysisSchemaVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageClassAnalysisSchemaVersionFromValue(v string) (*StorageClassAnalysisSchemaVersion, error) {
	ev := StorageClassAnalysisSchemaVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageClassAnalysisSchemaVersion: valid values are %v", v, AllowedStorageClassAnalysisSchemaVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageClassAnalysisSchemaVersion) IsValid() bool {
	for _, existing := range AllowedStorageClassAnalysisSchemaVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageClassAnalysisSchemaVersion value
func (v StorageClassAnalysisSchemaVersion) Ptr() *StorageClassAnalysisSchemaVersion {
	return &v
}

type NullableStorageClassAnalysisSchemaVersion struct {
	value *StorageClassAnalysisSchemaVersion
	isSet bool
}

func (v NullableStorageClassAnalysisSchemaVersion) Get() *StorageClassAnalysisSchemaVersion {
	return v.value
}

func (v *NullableStorageClassAnalysisSchemaVersion) Set(val *StorageClassAnalysisSchemaVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassAnalysisSchemaVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassAnalysisSchemaVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassAnalysisSchemaVersion(val *StorageClassAnalysisSchemaVersion) *NullableStorageClassAnalysisSchemaVersion {
	return &NullableStorageClassAnalysisSchemaVersion{value: val, isSet: true}
}

func (v NullableStorageClassAnalysisSchemaVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassAnalysisSchemaVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

