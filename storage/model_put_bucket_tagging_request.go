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

// checks if the PutBucketTaggingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketTaggingRequest{}

// PutBucketTaggingRequest struct for PutBucketTaggingRequest
type PutBucketTaggingRequest struct {
	Tagging PutBucketTaggingRequestTagging `json:"Tagging"`
}

// NewPutBucketTaggingRequest instantiates a new PutBucketTaggingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketTaggingRequest(tagging PutBucketTaggingRequestTagging) *PutBucketTaggingRequest {
	this := PutBucketTaggingRequest{}
	this.Tagging = tagging
	return &this
}

// NewPutBucketTaggingRequestWithDefaults instantiates a new PutBucketTaggingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketTaggingRequestWithDefaults() *PutBucketTaggingRequest {
	this := PutBucketTaggingRequest{}
	return &this
}

// GetTagging returns the Tagging field value
func (o *PutBucketTaggingRequest) GetTagging() PutBucketTaggingRequestTagging {
	if o == nil {
		var ret PutBucketTaggingRequestTagging
		return ret
	}

	return o.Tagging
}

// GetTaggingOk returns a tuple with the Tagging field value
// and a boolean to check if the value has been set.
func (o *PutBucketTaggingRequest) GetTaggingOk() (*PutBucketTaggingRequestTagging, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tagging, true
}

// SetTagging sets field value
func (o *PutBucketTaggingRequest) SetTagging(v PutBucketTaggingRequestTagging) {
	o.Tagging = v
}

func (o PutBucketTaggingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketTaggingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Tagging"] = o.Tagging
	return toSerialize, nil
}

type NullablePutBucketTaggingRequest struct {
	value *PutBucketTaggingRequest
	isSet bool
}

func (v NullablePutBucketTaggingRequest) Get() *PutBucketTaggingRequest {
	return v.value
}

func (v *NullablePutBucketTaggingRequest) Set(val *PutBucketTaggingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketTaggingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketTaggingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketTaggingRequest(val *PutBucketTaggingRequest) *NullablePutBucketTaggingRequest {
	return &NullablePutBucketTaggingRequest{value: val, isSet: true}
}

func (v NullablePutBucketTaggingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketTaggingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

