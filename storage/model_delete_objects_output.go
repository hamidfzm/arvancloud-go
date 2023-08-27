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

// checks if the DeleteObjectsOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteObjectsOutput{}

// DeleteObjectsOutput struct for DeleteObjectsOutput
type DeleteObjectsOutput struct {
	Deleted *Array `json:"Deleted,omitempty"`
	Errors *Array `json:"Errors,omitempty"`
}

// NewDeleteObjectsOutput instantiates a new DeleteObjectsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteObjectsOutput() *DeleteObjectsOutput {
	this := DeleteObjectsOutput{}
	return &this
}

// NewDeleteObjectsOutputWithDefaults instantiates a new DeleteObjectsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteObjectsOutputWithDefaults() *DeleteObjectsOutput {
	this := DeleteObjectsOutput{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DeleteObjectsOutput) GetDeleted() Array {
	if o == nil || IsNil(o.Deleted) {
		var ret Array
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteObjectsOutput) GetDeletedOk() (*Array, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DeleteObjectsOutput) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given Array and assigns it to the Deleted field.
func (o *DeleteObjectsOutput) SetDeleted(v Array) {
	o.Deleted = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *DeleteObjectsOutput) GetErrors() Array {
	if o == nil || IsNil(o.Errors) {
		var ret Array
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteObjectsOutput) GetErrorsOk() (*Array, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *DeleteObjectsOutput) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given Array and assigns it to the Errors field.
func (o *DeleteObjectsOutput) SetErrors(v Array) {
	o.Errors = &v
}

func (o DeleteObjectsOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteObjectsOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deleted) {
		toSerialize["Deleted"] = o.Deleted
	}
	if !IsNil(o.Errors) {
		toSerialize["Errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableDeleteObjectsOutput struct {
	value *DeleteObjectsOutput
	isSet bool
}

func (v NullableDeleteObjectsOutput) Get() *DeleteObjectsOutput {
	return v.value
}

func (v *NullableDeleteObjectsOutput) Set(val *DeleteObjectsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteObjectsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteObjectsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteObjectsOutput(val *DeleteObjectsOutput) *NullableDeleteObjectsOutput {
	return &NullableDeleteObjectsOutput{value: val, isSet: true}
}

func (v NullableDeleteObjectsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteObjectsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


