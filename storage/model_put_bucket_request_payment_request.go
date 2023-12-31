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

// checks if the PutBucketRequestPaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketRequestPaymentRequest{}

// PutBucketRequestPaymentRequest struct for PutBucketRequestPaymentRequest
type PutBucketRequestPaymentRequest struct {
	RequestPaymentConfiguration PutBucketRequestPaymentRequestRequestPaymentConfiguration `json:"RequestPaymentConfiguration"`
}

// NewPutBucketRequestPaymentRequest instantiates a new PutBucketRequestPaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketRequestPaymentRequest(requestPaymentConfiguration PutBucketRequestPaymentRequestRequestPaymentConfiguration) *PutBucketRequestPaymentRequest {
	this := PutBucketRequestPaymentRequest{}
	this.RequestPaymentConfiguration = requestPaymentConfiguration
	return &this
}

// NewPutBucketRequestPaymentRequestWithDefaults instantiates a new PutBucketRequestPaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketRequestPaymentRequestWithDefaults() *PutBucketRequestPaymentRequest {
	this := PutBucketRequestPaymentRequest{}
	return &this
}

// GetRequestPaymentConfiguration returns the RequestPaymentConfiguration field value
func (o *PutBucketRequestPaymentRequest) GetRequestPaymentConfiguration() PutBucketRequestPaymentRequestRequestPaymentConfiguration {
	if o == nil {
		var ret PutBucketRequestPaymentRequestRequestPaymentConfiguration
		return ret
	}

	return o.RequestPaymentConfiguration
}

// GetRequestPaymentConfigurationOk returns a tuple with the RequestPaymentConfiguration field value
// and a boolean to check if the value has been set.
func (o *PutBucketRequestPaymentRequest) GetRequestPaymentConfigurationOk() (*PutBucketRequestPaymentRequestRequestPaymentConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestPaymentConfiguration, true
}

// SetRequestPaymentConfiguration sets field value
func (o *PutBucketRequestPaymentRequest) SetRequestPaymentConfiguration(v PutBucketRequestPaymentRequestRequestPaymentConfiguration) {
	o.RequestPaymentConfiguration = v
}

func (o PutBucketRequestPaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketRequestPaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RequestPaymentConfiguration"] = o.RequestPaymentConfiguration
	return toSerialize, nil
}

type NullablePutBucketRequestPaymentRequest struct {
	value *PutBucketRequestPaymentRequest
	isSet bool
}

func (v NullablePutBucketRequestPaymentRequest) Get() *PutBucketRequestPaymentRequest {
	return v.value
}

func (v *NullablePutBucketRequestPaymentRequest) Set(val *PutBucketRequestPaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketRequestPaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketRequestPaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketRequestPaymentRequest(val *PutBucketRequestPaymentRequest) *NullablePutBucketRequestPaymentRequest {
	return &NullablePutBucketRequestPaymentRequest{value: val, isSet: true}
}

func (v NullablePutBucketRequestPaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketRequestPaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


