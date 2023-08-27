/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the DnsRecordData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRecordData{}

// DnsRecordData struct for DnsRecordData
type DnsRecordData struct {
	Data *DnsRecordGeneric `json:"data,omitempty"`
}

// NewDnsRecordData instantiates a new DnsRecordData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRecordData() *DnsRecordData {
	this := DnsRecordData{}
	return &this
}

// NewDnsRecordDataWithDefaults instantiates a new DnsRecordData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRecordDataWithDefaults() *DnsRecordData {
	this := DnsRecordData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DnsRecordData) GetData() DnsRecordGeneric {
	if o == nil || IsNil(o.Data) {
		var ret DnsRecordGeneric
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecordData) GetDataOk() (*DnsRecordGeneric, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DnsRecordData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DnsRecordGeneric and assigns it to the Data field.
func (o *DnsRecordData) SetData(v DnsRecordGeneric) {
	o.Data = &v
}

func (o DnsRecordData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRecordData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDnsRecordData struct {
	value *DnsRecordData
	isSet bool
}

func (v NullableDnsRecordData) Get() *DnsRecordData {
	return v.value
}

func (v *NullableDnsRecordData) Set(val *DnsRecordData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecordData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecordData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecordData(val *DnsRecordData) *NullableDnsRecordData {
	return &NullableDnsRecordData{value: val, isSet: true}
}

func (v NullableDnsRecordData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecordData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


