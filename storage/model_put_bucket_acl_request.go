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

// checks if the PutBucketAclRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketAclRequest{}

// PutBucketAclRequest struct for PutBucketAclRequest
type PutBucketAclRequest struct {
	AccessControlPolicy *PutBucketAclRequestAccessControlPolicy `json:"AccessControlPolicy,omitempty"`
}

// NewPutBucketAclRequest instantiates a new PutBucketAclRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketAclRequest() *PutBucketAclRequest {
	this := PutBucketAclRequest{}
	return &this
}

// NewPutBucketAclRequestWithDefaults instantiates a new PutBucketAclRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketAclRequestWithDefaults() *PutBucketAclRequest {
	this := PutBucketAclRequest{}
	return &this
}

// GetAccessControlPolicy returns the AccessControlPolicy field value if set, zero value otherwise.
func (o *PutBucketAclRequest) GetAccessControlPolicy() PutBucketAclRequestAccessControlPolicy {
	if o == nil || IsNil(o.AccessControlPolicy) {
		var ret PutBucketAclRequestAccessControlPolicy
		return ret
	}
	return *o.AccessControlPolicy
}

// GetAccessControlPolicyOk returns a tuple with the AccessControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBucketAclRequest) GetAccessControlPolicyOk() (*PutBucketAclRequestAccessControlPolicy, bool) {
	if o == nil || IsNil(o.AccessControlPolicy) {
		return nil, false
	}
	return o.AccessControlPolicy, true
}

// HasAccessControlPolicy returns a boolean if a field has been set.
func (o *PutBucketAclRequest) HasAccessControlPolicy() bool {
	if o != nil && !IsNil(o.AccessControlPolicy) {
		return true
	}

	return false
}

// SetAccessControlPolicy gets a reference to the given PutBucketAclRequestAccessControlPolicy and assigns it to the AccessControlPolicy field.
func (o *PutBucketAclRequest) SetAccessControlPolicy(v PutBucketAclRequestAccessControlPolicy) {
	o.AccessControlPolicy = &v
}

func (o PutBucketAclRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketAclRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessControlPolicy) {
		toSerialize["AccessControlPolicy"] = o.AccessControlPolicy
	}
	return toSerialize, nil
}

type NullablePutBucketAclRequest struct {
	value *PutBucketAclRequest
	isSet bool
}

func (v NullablePutBucketAclRequest) Get() *PutBucketAclRequest {
	return v.value
}

func (v *NullablePutBucketAclRequest) Set(val *PutBucketAclRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketAclRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketAclRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketAclRequest(val *PutBucketAclRequest) *NullablePutBucketAclRequest {
	return &NullablePutBucketAclRequest{value: val, isSet: true}
}

func (v NullablePutBucketAclRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketAclRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


