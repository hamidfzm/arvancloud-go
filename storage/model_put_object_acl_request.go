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

// checks if the PutObjectAclRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutObjectAclRequest{}

// PutObjectAclRequest struct for PutObjectAclRequest
type PutObjectAclRequest struct {
	AccessControlPolicy *PutBucketAclRequestAccessControlPolicy `json:"AccessControlPolicy,omitempty"`
}

// NewPutObjectAclRequest instantiates a new PutObjectAclRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutObjectAclRequest() *PutObjectAclRequest {
	this := PutObjectAclRequest{}
	return &this
}

// NewPutObjectAclRequestWithDefaults instantiates a new PutObjectAclRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutObjectAclRequestWithDefaults() *PutObjectAclRequest {
	this := PutObjectAclRequest{}
	return &this
}

// GetAccessControlPolicy returns the AccessControlPolicy field value if set, zero value otherwise.
func (o *PutObjectAclRequest) GetAccessControlPolicy() PutBucketAclRequestAccessControlPolicy {
	if o == nil || IsNil(o.AccessControlPolicy) {
		var ret PutBucketAclRequestAccessControlPolicy
		return ret
	}
	return *o.AccessControlPolicy
}

// GetAccessControlPolicyOk returns a tuple with the AccessControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutObjectAclRequest) GetAccessControlPolicyOk() (*PutBucketAclRequestAccessControlPolicy, bool) {
	if o == nil || IsNil(o.AccessControlPolicy) {
		return nil, false
	}
	return o.AccessControlPolicy, true
}

// HasAccessControlPolicy returns a boolean if a field has been set.
func (o *PutObjectAclRequest) HasAccessControlPolicy() bool {
	if o != nil && !IsNil(o.AccessControlPolicy) {
		return true
	}

	return false
}

// SetAccessControlPolicy gets a reference to the given PutBucketAclRequestAccessControlPolicy and assigns it to the AccessControlPolicy field.
func (o *PutObjectAclRequest) SetAccessControlPolicy(v PutBucketAclRequestAccessControlPolicy) {
	o.AccessControlPolicy = &v
}

func (o PutObjectAclRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutObjectAclRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessControlPolicy) {
		toSerialize["AccessControlPolicy"] = o.AccessControlPolicy
	}
	return toSerialize, nil
}

type NullablePutObjectAclRequest struct {
	value *PutObjectAclRequest
	isSet bool
}

func (v NullablePutObjectAclRequest) Get() *PutObjectAclRequest {
	return v.value
}

func (v *NullablePutObjectAclRequest) Set(val *PutObjectAclRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutObjectAclRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutObjectAclRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutObjectAclRequest(val *PutObjectAclRequest) *NullablePutObjectAclRequest {
	return &NullablePutObjectAclRequest{value: val, isSet: true}
}

func (v NullablePutObjectAclRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutObjectAclRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


