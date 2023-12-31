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

// checks if the CreateMultipartUploadOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMultipartUploadOutput{}

// CreateMultipartUploadOutput struct for CreateMultipartUploadOutput
type CreateMultipartUploadOutput struct {
	Bucket *string `json:"Bucket,omitempty"`
	Key *string `json:"Key,omitempty"`
	UploadId *string `json:"UploadId,omitempty"`
}

// NewCreateMultipartUploadOutput instantiates a new CreateMultipartUploadOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMultipartUploadOutput() *CreateMultipartUploadOutput {
	this := CreateMultipartUploadOutput{}
	return &this
}

// NewCreateMultipartUploadOutputWithDefaults instantiates a new CreateMultipartUploadOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMultipartUploadOutputWithDefaults() *CreateMultipartUploadOutput {
	this := CreateMultipartUploadOutput{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *CreateMultipartUploadOutput) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMultipartUploadOutput) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *CreateMultipartUploadOutput) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *CreateMultipartUploadOutput) SetBucket(v string) {
	o.Bucket = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreateMultipartUploadOutput) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMultipartUploadOutput) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreateMultipartUploadOutput) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreateMultipartUploadOutput) SetKey(v string) {
	o.Key = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *CreateMultipartUploadOutput) GetUploadId() string {
	if o == nil || IsNil(o.UploadId) {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMultipartUploadOutput) GetUploadIdOk() (*string, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *CreateMultipartUploadOutput) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *CreateMultipartUploadOutput) SetUploadId(v string) {
	o.UploadId = &v
}

func (o CreateMultipartUploadOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMultipartUploadOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bucket) {
		toSerialize["Bucket"] = o.Bucket
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.UploadId) {
		toSerialize["UploadId"] = o.UploadId
	}
	return toSerialize, nil
}

type NullableCreateMultipartUploadOutput struct {
	value *CreateMultipartUploadOutput
	isSet bool
}

func (v NullableCreateMultipartUploadOutput) Get() *CreateMultipartUploadOutput {
	return v.value
}

func (v *NullableCreateMultipartUploadOutput) Set(val *CreateMultipartUploadOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMultipartUploadOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMultipartUploadOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMultipartUploadOutput(val *CreateMultipartUploadOutput) *NullableCreateMultipartUploadOutput {
	return &NullableCreateMultipartUploadOutput{value: val, isSet: true}
}

func (v NullableCreateMultipartUploadOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMultipartUploadOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


