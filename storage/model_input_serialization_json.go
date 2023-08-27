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

// checks if the InputSerializationJSON type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InputSerializationJSON{}

// InputSerializationJSON struct for InputSerializationJSON
type InputSerializationJSON struct {
	Type *JSONType `json:"Type,omitempty"`
}

// NewInputSerializationJSON instantiates a new InputSerializationJSON object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputSerializationJSON() *InputSerializationJSON {
	this := InputSerializationJSON{}
	return &this
}

// NewInputSerializationJSONWithDefaults instantiates a new InputSerializationJSON object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputSerializationJSONWithDefaults() *InputSerializationJSON {
	this := InputSerializationJSON{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InputSerializationJSON) GetType() JSONType {
	if o == nil || IsNil(o.Type) {
		var ret JSONType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputSerializationJSON) GetTypeOk() (*JSONType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InputSerializationJSON) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given JSONType and assigns it to the Type field.
func (o *InputSerializationJSON) SetType(v JSONType) {
	o.Type = &v
}

func (o InputSerializationJSON) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InputSerializationJSON) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInputSerializationJSON struct {
	value *InputSerializationJSON
	isSet bool
}

func (v NullableInputSerializationJSON) Get() *InputSerializationJSON {
	return v.value
}

func (v *NullableInputSerializationJSON) Set(val *InputSerializationJSON) {
	v.value = val
	v.isSet = true
}

func (v NullableInputSerializationJSON) IsSet() bool {
	return v.isSet
}

func (v *NullableInputSerializationJSON) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputSerializationJSON(val *InputSerializationJSON) *NullableInputSerializationJSON {
	return &NullableInputSerializationJSON{value: val, isSet: true}
}

func (v NullableInputSerializationJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputSerializationJSON) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

