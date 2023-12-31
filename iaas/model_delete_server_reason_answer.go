/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the DeleteServerReasonAnswer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteServerReasonAnswer{}

// DeleteServerReasonAnswer struct for DeleteServerReasonAnswer
type DeleteServerReasonAnswer struct {
	Answer *string `json:"answer,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewDeleteServerReasonAnswer instantiates a new DeleteServerReasonAnswer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteServerReasonAnswer() *DeleteServerReasonAnswer {
	this := DeleteServerReasonAnswer{}
	return &this
}

// NewDeleteServerReasonAnswerWithDefaults instantiates a new DeleteServerReasonAnswer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteServerReasonAnswerWithDefaults() *DeleteServerReasonAnswer {
	this := DeleteServerReasonAnswer{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *DeleteServerReasonAnswer) GetAnswer() string {
	if o == nil || IsNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteServerReasonAnswer) GetAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *DeleteServerReasonAnswer) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *DeleteServerReasonAnswer) SetAnswer(v string) {
	o.Answer = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteServerReasonAnswer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteServerReasonAnswer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteServerReasonAnswer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeleteServerReasonAnswer) SetId(v string) {
	o.Id = &v
}

func (o DeleteServerReasonAnswer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteServerReasonAnswer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeleteServerReasonAnswer struct {
	value *DeleteServerReasonAnswer
	isSet bool
}

func (v NullableDeleteServerReasonAnswer) Get() *DeleteServerReasonAnswer {
	return v.value
}

func (v *NullableDeleteServerReasonAnswer) Set(val *DeleteServerReasonAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteServerReasonAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteServerReasonAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteServerReasonAnswer(val *DeleteServerReasonAnswer) *NullableDeleteServerReasonAnswer {
	return &NullableDeleteServerReasonAnswer{value: val, isSet: true}
}

func (v NullableDeleteServerReasonAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteServerReasonAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


