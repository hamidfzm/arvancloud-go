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

// checks if the CdnAppLike type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdnAppLike{}

// CdnAppLike struct for CdnAppLike
type CdnAppLike struct {
	// True means she likes, False means she dislikes, null means she wants to get her vote back.
	Like NullableBool `json:"like,omitempty"`
}

// NewCdnAppLike instantiates a new CdnAppLike object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnAppLike() *CdnAppLike {
	this := CdnAppLike{}
	return &this
}

// NewCdnAppLikeWithDefaults instantiates a new CdnAppLike object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnAppLikeWithDefaults() *CdnAppLike {
	this := CdnAppLike{}
	return &this
}

// GetLike returns the Like field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CdnAppLike) GetLike() bool {
	if o == nil || IsNil(o.Like.Get()) {
		var ret bool
		return ret
	}
	return *o.Like.Get()
}

// GetLikeOk returns a tuple with the Like field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CdnAppLike) GetLikeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Like.Get(), o.Like.IsSet()
}

// HasLike returns a boolean if a field has been set.
func (o *CdnAppLike) HasLike() bool {
	if o != nil && o.Like.IsSet() {
		return true
	}

	return false
}

// SetLike gets a reference to the given NullableBool and assigns it to the Like field.
func (o *CdnAppLike) SetLike(v bool) {
	o.Like.Set(&v)
}
// SetLikeNil sets the value for Like to be an explicit nil
func (o *CdnAppLike) SetLikeNil() {
	o.Like.Set(nil)
}

// UnsetLike ensures that no value is present for Like, not even an explicit nil
func (o *CdnAppLike) UnsetLike() {
	o.Like.Unset()
}

func (o CdnAppLike) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdnAppLike) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Like.IsSet() {
		toSerialize["like"] = o.Like.Get()
	}
	return toSerialize, nil
}

type NullableCdnAppLike struct {
	value *CdnAppLike
	isSet bool
}

func (v NullableCdnAppLike) Get() *CdnAppLike {
	return v.value
}

func (v *NullableCdnAppLike) Set(val *CdnAppLike) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnAppLike) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnAppLike) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnAppLike(val *CdnAppLike) *NullableCdnAppLike {
	return &NullableCdnAppLike{value: val, isSet: true}
}

func (v NullableCdnAppLike) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnAppLike) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


