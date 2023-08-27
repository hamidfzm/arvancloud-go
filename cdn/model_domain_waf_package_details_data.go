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

// checks if the DomainWafPackageDetailsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainWafPackageDetailsData{}

// DomainWafPackageDetailsData struct for DomainWafPackageDetailsData
type DomainWafPackageDetailsData struct {
	Data *DomainWafPackageDetails `json:"data,omitempty"`
}

// NewDomainWafPackageDetailsData instantiates a new DomainWafPackageDetailsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainWafPackageDetailsData() *DomainWafPackageDetailsData {
	this := DomainWafPackageDetailsData{}
	return &this
}

// NewDomainWafPackageDetailsDataWithDefaults instantiates a new DomainWafPackageDetailsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWafPackageDetailsDataWithDefaults() *DomainWafPackageDetailsData {
	this := DomainWafPackageDetailsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DomainWafPackageDetailsData) GetData() DomainWafPackageDetails {
	if o == nil || IsNil(o.Data) {
		var ret DomainWafPackageDetails
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackageDetailsData) GetDataOk() (*DomainWafPackageDetails, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DomainWafPackageDetailsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DomainWafPackageDetails and assigns it to the Data field.
func (o *DomainWafPackageDetailsData) SetData(v DomainWafPackageDetails) {
	o.Data = &v
}

func (o DomainWafPackageDetailsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainWafPackageDetailsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDomainWafPackageDetailsData struct {
	value *DomainWafPackageDetailsData
	isSet bool
}

func (v NullableDomainWafPackageDetailsData) Get() *DomainWafPackageDetailsData {
	return v.value
}

func (v *NullableDomainWafPackageDetailsData) Set(val *DomainWafPackageDetailsData) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainWafPackageDetailsData) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainWafPackageDetailsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainWafPackageDetailsData(val *DomainWafPackageDetailsData) *NullableDomainWafPackageDetailsData {
	return &NullableDomainWafPackageDetailsData{value: val, isSet: true}
}

func (v NullableDomainWafPackageDetailsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainWafPackageDetailsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


