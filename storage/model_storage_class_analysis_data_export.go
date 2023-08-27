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

// checks if the StorageClassAnalysisDataExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageClassAnalysisDataExport{}

// StorageClassAnalysisDataExport Container for data related to the storage class analysis for an ArvanCloud S3 bucket for export.
type StorageClassAnalysisDataExport struct {
	OutputSchemaVersion StorageClassAnalysisSchemaVersion `json:"OutputSchemaVersion"`
	Destination StorageClassAnalysisDataExportDestination `json:"Destination"`
}

// NewStorageClassAnalysisDataExport instantiates a new StorageClassAnalysisDataExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageClassAnalysisDataExport(outputSchemaVersion StorageClassAnalysisSchemaVersion, destination StorageClassAnalysisDataExportDestination) *StorageClassAnalysisDataExport {
	this := StorageClassAnalysisDataExport{}
	this.OutputSchemaVersion = outputSchemaVersion
	this.Destination = destination
	return &this
}

// NewStorageClassAnalysisDataExportWithDefaults instantiates a new StorageClassAnalysisDataExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageClassAnalysisDataExportWithDefaults() *StorageClassAnalysisDataExport {
	this := StorageClassAnalysisDataExport{}
	return &this
}

// GetOutputSchemaVersion returns the OutputSchemaVersion field value
func (o *StorageClassAnalysisDataExport) GetOutputSchemaVersion() StorageClassAnalysisSchemaVersion {
	if o == nil {
		var ret StorageClassAnalysisSchemaVersion
		return ret
	}

	return o.OutputSchemaVersion
}

// GetOutputSchemaVersionOk returns a tuple with the OutputSchemaVersion field value
// and a boolean to check if the value has been set.
func (o *StorageClassAnalysisDataExport) GetOutputSchemaVersionOk() (*StorageClassAnalysisSchemaVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputSchemaVersion, true
}

// SetOutputSchemaVersion sets field value
func (o *StorageClassAnalysisDataExport) SetOutputSchemaVersion(v StorageClassAnalysisSchemaVersion) {
	o.OutputSchemaVersion = v
}

// GetDestination returns the Destination field value
func (o *StorageClassAnalysisDataExport) GetDestination() StorageClassAnalysisDataExportDestination {
	if o == nil {
		var ret StorageClassAnalysisDataExportDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *StorageClassAnalysisDataExport) GetDestinationOk() (*StorageClassAnalysisDataExportDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *StorageClassAnalysisDataExport) SetDestination(v StorageClassAnalysisDataExportDestination) {
	o.Destination = v
}

func (o StorageClassAnalysisDataExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageClassAnalysisDataExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["OutputSchemaVersion"] = o.OutputSchemaVersion
	toSerialize["Destination"] = o.Destination
	return toSerialize, nil
}

type NullableStorageClassAnalysisDataExport struct {
	value *StorageClassAnalysisDataExport
	isSet bool
}

func (v NullableStorageClassAnalysisDataExport) Get() *StorageClassAnalysisDataExport {
	return v.value
}

func (v *NullableStorageClassAnalysisDataExport) Set(val *StorageClassAnalysisDataExport) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassAnalysisDataExport) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassAnalysisDataExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassAnalysisDataExport(val *StorageClassAnalysisDataExport) *NullableStorageClassAnalysisDataExport {
	return &NullableStorageClassAnalysisDataExport{value: val, isSet: true}
}

func (v NullableStorageClassAnalysisDataExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassAnalysisDataExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


