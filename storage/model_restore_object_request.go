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

// checks if the RestoreObjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreObjectRequest{}

// RestoreObjectRequest struct for RestoreObjectRequest
type RestoreObjectRequest struct {
	RestoreRequest *RestoreRequest `json:"RestoreRequest,omitempty"`
}

// NewRestoreObjectRequest instantiates a new RestoreObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreObjectRequest() *RestoreObjectRequest {
	this := RestoreObjectRequest{}
	return &this
}

// NewRestoreObjectRequestWithDefaults instantiates a new RestoreObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreObjectRequestWithDefaults() *RestoreObjectRequest {
	this := RestoreObjectRequest{}
	return &this
}

// GetRestoreRequest returns the RestoreRequest field value if set, zero value otherwise.
func (o *RestoreObjectRequest) GetRestoreRequest() RestoreRequest {
	if o == nil || IsNil(o.RestoreRequest) {
		var ret RestoreRequest
		return ret
	}
	return *o.RestoreRequest
}

// GetRestoreRequestOk returns a tuple with the RestoreRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreObjectRequest) GetRestoreRequestOk() (*RestoreRequest, bool) {
	if o == nil || IsNil(o.RestoreRequest) {
		return nil, false
	}
	return o.RestoreRequest, true
}

// HasRestoreRequest returns a boolean if a field has been set.
func (o *RestoreObjectRequest) HasRestoreRequest() bool {
	if o != nil && !IsNil(o.RestoreRequest) {
		return true
	}

	return false
}

// SetRestoreRequest gets a reference to the given RestoreRequest and assigns it to the RestoreRequest field.
func (o *RestoreObjectRequest) SetRestoreRequest(v RestoreRequest) {
	o.RestoreRequest = &v
}

func (o RestoreObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreObjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RestoreRequest) {
		toSerialize["RestoreRequest"] = o.RestoreRequest
	}
	return toSerialize, nil
}

type NullableRestoreObjectRequest struct {
	value *RestoreObjectRequest
	isSet bool
}

func (v NullableRestoreObjectRequest) Get() *RestoreObjectRequest {
	return v.value
}

func (v *NullableRestoreObjectRequest) Set(val *RestoreObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreObjectRequest(val *RestoreObjectRequest) *NullableRestoreObjectRequest {
	return &NullableRestoreObjectRequest{value: val, isSet: true}
}

func (v NullableRestoreObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


