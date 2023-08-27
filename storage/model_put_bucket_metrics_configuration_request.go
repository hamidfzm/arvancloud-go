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

// checks if the PutBucketMetricsConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketMetricsConfigurationRequest{}

// PutBucketMetricsConfigurationRequest struct for PutBucketMetricsConfigurationRequest
type PutBucketMetricsConfigurationRequest struct {
	MetricsConfiguration PutBucketMetricsConfigurationRequestMetricsConfiguration `json:"MetricsConfiguration"`
}

// NewPutBucketMetricsConfigurationRequest instantiates a new PutBucketMetricsConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketMetricsConfigurationRequest(metricsConfiguration PutBucketMetricsConfigurationRequestMetricsConfiguration) *PutBucketMetricsConfigurationRequest {
	this := PutBucketMetricsConfigurationRequest{}
	this.MetricsConfiguration = metricsConfiguration
	return &this
}

// NewPutBucketMetricsConfigurationRequestWithDefaults instantiates a new PutBucketMetricsConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketMetricsConfigurationRequestWithDefaults() *PutBucketMetricsConfigurationRequest {
	this := PutBucketMetricsConfigurationRequest{}
	return &this
}

// GetMetricsConfiguration returns the MetricsConfiguration field value
func (o *PutBucketMetricsConfigurationRequest) GetMetricsConfiguration() PutBucketMetricsConfigurationRequestMetricsConfiguration {
	if o == nil {
		var ret PutBucketMetricsConfigurationRequestMetricsConfiguration
		return ret
	}

	return o.MetricsConfiguration
}

// GetMetricsConfigurationOk returns a tuple with the MetricsConfiguration field value
// and a boolean to check if the value has been set.
func (o *PutBucketMetricsConfigurationRequest) GetMetricsConfigurationOk() (*PutBucketMetricsConfigurationRequestMetricsConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricsConfiguration, true
}

// SetMetricsConfiguration sets field value
func (o *PutBucketMetricsConfigurationRequest) SetMetricsConfiguration(v PutBucketMetricsConfigurationRequestMetricsConfiguration) {
	o.MetricsConfiguration = v
}

func (o PutBucketMetricsConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketMetricsConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["MetricsConfiguration"] = o.MetricsConfiguration
	return toSerialize, nil
}

type NullablePutBucketMetricsConfigurationRequest struct {
	value *PutBucketMetricsConfigurationRequest
	isSet bool
}

func (v NullablePutBucketMetricsConfigurationRequest) Get() *PutBucketMetricsConfigurationRequest {
	return v.value
}

func (v *NullablePutBucketMetricsConfigurationRequest) Set(val *PutBucketMetricsConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketMetricsConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketMetricsConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketMetricsConfigurationRequest(val *PutBucketMetricsConfigurationRequest) *NullablePutBucketMetricsConfigurationRequest {
	return &NullablePutBucketMetricsConfigurationRequest{value: val, isSet: true}
}

func (v NullablePutBucketMetricsConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketMetricsConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


