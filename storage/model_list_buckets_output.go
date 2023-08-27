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

// checks if the ListBucketsOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBucketsOutput{}

// ListBucketsOutput struct for ListBucketsOutput
type ListBucketsOutput struct {
	Buckets *Array `json:"Buckets,omitempty"`
	Owner *Owner `json:"Owner,omitempty"`
}

// NewListBucketsOutput instantiates a new ListBucketsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBucketsOutput() *ListBucketsOutput {
	this := ListBucketsOutput{}
	return &this
}

// NewListBucketsOutputWithDefaults instantiates a new ListBucketsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBucketsOutputWithDefaults() *ListBucketsOutput {
	this := ListBucketsOutput{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *ListBucketsOutput) GetBuckets() Array {
	if o == nil || IsNil(o.Buckets) {
		var ret Array
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBucketsOutput) GetBucketsOk() (*Array, bool) {
	if o == nil || IsNil(o.Buckets) {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *ListBucketsOutput) HasBuckets() bool {
	if o != nil && !IsNil(o.Buckets) {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given Array and assigns it to the Buckets field.
func (o *ListBucketsOutput) SetBuckets(v Array) {
	o.Buckets = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ListBucketsOutput) GetOwner() Owner {
	if o == nil || IsNil(o.Owner) {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBucketsOutput) GetOwnerOk() (*Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ListBucketsOutput) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *ListBucketsOutput) SetOwner(v Owner) {
	o.Owner = &v
}

func (o ListBucketsOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBucketsOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Buckets) {
		toSerialize["Buckets"] = o.Buckets
	}
	if !IsNil(o.Owner) {
		toSerialize["Owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableListBucketsOutput struct {
	value *ListBucketsOutput
	isSet bool
}

func (v NullableListBucketsOutput) Get() *ListBucketsOutput {
	return v.value
}

func (v *NullableListBucketsOutput) Set(val *ListBucketsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListBucketsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListBucketsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBucketsOutput(val *ListBucketsOutput) *NullableListBucketsOutput {
	return &NullableListBucketsOutput{value: val, isSet: true}
}

func (v NullableListBucketsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBucketsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

