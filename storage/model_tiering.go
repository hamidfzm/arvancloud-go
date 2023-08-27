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

// checks if the Tiering type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tiering{}

// Tiering The S3 Intelligent-Tiering storage class is designed to optimize storage costs by automatically moving data to the most cost-effective storage access tier, without additional operational overhead.
type Tiering struct {
	Days int32 `json:"Days"`
	AccessTier IntelligentTieringAccessTier `json:"AccessTier"`
}

// NewTiering instantiates a new Tiering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTiering(days int32, accessTier IntelligentTieringAccessTier) *Tiering {
	this := Tiering{}
	this.Days = days
	this.AccessTier = accessTier
	return &this
}

// NewTieringWithDefaults instantiates a new Tiering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTieringWithDefaults() *Tiering {
	this := Tiering{}
	return &this
}

// GetDays returns the Days field value
func (o *Tiering) GetDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *Tiering) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *Tiering) SetDays(v int32) {
	o.Days = v
}

// GetAccessTier returns the AccessTier field value
func (o *Tiering) GetAccessTier() IntelligentTieringAccessTier {
	if o == nil {
		var ret IntelligentTieringAccessTier
		return ret
	}

	return o.AccessTier
}

// GetAccessTierOk returns a tuple with the AccessTier field value
// and a boolean to check if the value has been set.
func (o *Tiering) GetAccessTierOk() (*IntelligentTieringAccessTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessTier, true
}

// SetAccessTier sets field value
func (o *Tiering) SetAccessTier(v IntelligentTieringAccessTier) {
	o.AccessTier = v
}

func (o Tiering) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tiering) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Days"] = o.Days
	toSerialize["AccessTier"] = o.AccessTier
	return toSerialize, nil
}

type NullableTiering struct {
	value *Tiering
	isSet bool
}

func (v NullableTiering) Get() *Tiering {
	return v.value
}

func (v *NullableTiering) Set(val *Tiering) {
	v.value = val
	v.isSet = true
}

func (v NullableTiering) IsSet() bool {
	return v.isSet
}

func (v *NullableTiering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTiering(val *Tiering) *NullableTiering {
	return &NullableTiering{value: val, isSet: true}
}

func (v NullableTiering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTiering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

