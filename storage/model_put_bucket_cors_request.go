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

// checks if the PutBucketCorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketCorsRequest{}

// PutBucketCorsRequest struct for PutBucketCorsRequest
type PutBucketCorsRequest struct {
	CORSConfiguration PutBucketCorsRequestCORSConfiguration `json:"CORSConfiguration"`
}

// NewPutBucketCorsRequest instantiates a new PutBucketCorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketCorsRequest(cORSConfiguration PutBucketCorsRequestCORSConfiguration) *PutBucketCorsRequest {
	this := PutBucketCorsRequest{}
	this.CORSConfiguration = cORSConfiguration
	return &this
}

// NewPutBucketCorsRequestWithDefaults instantiates a new PutBucketCorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketCorsRequestWithDefaults() *PutBucketCorsRequest {
	this := PutBucketCorsRequest{}
	return &this
}

// GetCORSConfiguration returns the CORSConfiguration field value
func (o *PutBucketCorsRequest) GetCORSConfiguration() PutBucketCorsRequestCORSConfiguration {
	if o == nil {
		var ret PutBucketCorsRequestCORSConfiguration
		return ret
	}

	return o.CORSConfiguration
}

// GetCORSConfigurationOk returns a tuple with the CORSConfiguration field value
// and a boolean to check if the value has been set.
func (o *PutBucketCorsRequest) GetCORSConfigurationOk() (*PutBucketCorsRequestCORSConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CORSConfiguration, true
}

// SetCORSConfiguration sets field value
func (o *PutBucketCorsRequest) SetCORSConfiguration(v PutBucketCorsRequestCORSConfiguration) {
	o.CORSConfiguration = v
}

func (o PutBucketCorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketCorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CORSConfiguration"] = o.CORSConfiguration
	return toSerialize, nil
}

type NullablePutBucketCorsRequest struct {
	value *PutBucketCorsRequest
	isSet bool
}

func (v NullablePutBucketCorsRequest) Get() *PutBucketCorsRequest {
	return v.value
}

func (v *NullablePutBucketCorsRequest) Set(val *PutBucketCorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketCorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketCorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketCorsRequest(val *PutBucketCorsRequest) *NullablePutBucketCorsRequest {
	return &NullablePutBucketCorsRequest{value: val, isSet: true}
}

func (v NullablePutBucketCorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketCorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


