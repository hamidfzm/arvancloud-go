/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the SRVRecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SRVRecordValue{}

// SRVRecordValue struct for SRVRecordValue
type SRVRecordValue struct {
	Target string `json:"target"`
	Port NullableInt32 `json:"port"`
	Weight NullableInt32 `json:"weight,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
}

// NewSRVRecordValue instantiates a new SRVRecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSRVRecordValue(target string, port NullableInt32) *SRVRecordValue {
	this := SRVRecordValue{}
	this.Target = target
	this.Port = port
	return &this
}

// NewSRVRecordValueWithDefaults instantiates a new SRVRecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSRVRecordValueWithDefaults() *SRVRecordValue {
	this := SRVRecordValue{}
	return &this
}

// GetTarget returns the Target field value
func (o *SRVRecordValue) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SRVRecordValue) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *SRVRecordValue) SetTarget(v string) {
	o.Target = v
}

// GetPort returns the Port field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *SRVRecordValue) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SRVRecordValue) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// SetPort sets field value
func (o *SRVRecordValue) SetPort(v int32) {
	o.Port.Set(&v)
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SRVRecordValue) GetWeight() int32 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret int32
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SRVRecordValue) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *SRVRecordValue) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableInt32 and assigns it to the Weight field.
func (o *SRVRecordValue) SetWeight(v int32) {
	o.Weight.Set(&v)
}
// SetWeightNil sets the value for Weight to be an explicit nil
func (o *SRVRecordValue) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *SRVRecordValue) UnsetWeight() {
	o.Weight.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SRVRecordValue) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SRVRecordValue) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *SRVRecordValue) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *SRVRecordValue) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *SRVRecordValue) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *SRVRecordValue) UnsetPriority() {
	o.Priority.Unset()
}

func (o SRVRecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SRVRecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target"] = o.Target
	toSerialize["port"] = o.Port.Get()
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	return toSerialize, nil
}

type NullableSRVRecordValue struct {
	value *SRVRecordValue
	isSet bool
}

func (v NullableSRVRecordValue) Get() *SRVRecordValue {
	return v.value
}

func (v *NullableSRVRecordValue) Set(val *SRVRecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSRVRecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSRVRecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSRVRecordValue(val *SRVRecordValue) *NullableSRVRecordValue {
	return &NullableSRVRecordValue{value: val, isSet: true}
}

func (v NullableSRVRecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSRVRecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

