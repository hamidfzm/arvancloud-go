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

// checks if the PutBucketInventoryConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketInventoryConfigurationRequest{}

// PutBucketInventoryConfigurationRequest struct for PutBucketInventoryConfigurationRequest
type PutBucketInventoryConfigurationRequest struct {
	InventoryConfiguration PutBucketInventoryConfigurationRequestInventoryConfiguration `json:"InventoryConfiguration"`
}

// NewPutBucketInventoryConfigurationRequest instantiates a new PutBucketInventoryConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketInventoryConfigurationRequest(inventoryConfiguration PutBucketInventoryConfigurationRequestInventoryConfiguration) *PutBucketInventoryConfigurationRequest {
	this := PutBucketInventoryConfigurationRequest{}
	this.InventoryConfiguration = inventoryConfiguration
	return &this
}

// NewPutBucketInventoryConfigurationRequestWithDefaults instantiates a new PutBucketInventoryConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketInventoryConfigurationRequestWithDefaults() *PutBucketInventoryConfigurationRequest {
	this := PutBucketInventoryConfigurationRequest{}
	return &this
}

// GetInventoryConfiguration returns the InventoryConfiguration field value
func (o *PutBucketInventoryConfigurationRequest) GetInventoryConfiguration() PutBucketInventoryConfigurationRequestInventoryConfiguration {
	if o == nil {
		var ret PutBucketInventoryConfigurationRequestInventoryConfiguration
		return ret
	}

	return o.InventoryConfiguration
}

// GetInventoryConfigurationOk returns a tuple with the InventoryConfiguration field value
// and a boolean to check if the value has been set.
func (o *PutBucketInventoryConfigurationRequest) GetInventoryConfigurationOk() (*PutBucketInventoryConfigurationRequestInventoryConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryConfiguration, true
}

// SetInventoryConfiguration sets field value
func (o *PutBucketInventoryConfigurationRequest) SetInventoryConfiguration(v PutBucketInventoryConfigurationRequestInventoryConfiguration) {
	o.InventoryConfiguration = v
}

func (o PutBucketInventoryConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketInventoryConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["InventoryConfiguration"] = o.InventoryConfiguration
	return toSerialize, nil
}

type NullablePutBucketInventoryConfigurationRequest struct {
	value *PutBucketInventoryConfigurationRequest
	isSet bool
}

func (v NullablePutBucketInventoryConfigurationRequest) Get() *PutBucketInventoryConfigurationRequest {
	return v.value
}

func (v *NullablePutBucketInventoryConfigurationRequest) Set(val *PutBucketInventoryConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketInventoryConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketInventoryConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketInventoryConfigurationRequest(val *PutBucketInventoryConfigurationRequest) *NullablePutBucketInventoryConfigurationRequest {
	return &NullablePutBucketInventoryConfigurationRequest{value: val, isSet: true}
}

func (v NullablePutBucketInventoryConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketInventoryConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


