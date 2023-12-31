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

// checks if the UploadPartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadPartRequest{}

// UploadPartRequest struct for UploadPartRequest
type UploadPartRequest struct {
	// Object data.
	Body *string `json:"Body,omitempty"`
}

// NewUploadPartRequest instantiates a new UploadPartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadPartRequest() *UploadPartRequest {
	this := UploadPartRequest{}
	return &this
}

// NewUploadPartRequestWithDefaults instantiates a new UploadPartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadPartRequestWithDefaults() *UploadPartRequest {
	this := UploadPartRequest{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UploadPartRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadPartRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UploadPartRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UploadPartRequest) SetBody(v string) {
	o.Body = &v
}

func (o UploadPartRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadPartRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["Body"] = o.Body
	}
	return toSerialize, nil
}

type NullableUploadPartRequest struct {
	value *UploadPartRequest
	isSet bool
}

func (v NullableUploadPartRequest) Get() *UploadPartRequest {
	return v.value
}

func (v *NullableUploadPartRequest) Set(val *UploadPartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadPartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadPartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadPartRequest(val *UploadPartRequest) *NullableUploadPartRequest {
	return &NullableUploadPartRequest{value: val, isSet: true}
}

func (v NullableUploadPartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadPartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


