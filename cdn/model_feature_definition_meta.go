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

// checks if the FeatureDefinitionMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureDefinitionMeta{}

// FeatureDefinitionMeta struct for FeatureDefinitionMeta
type FeatureDefinitionMeta struct {
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewFeatureDefinitionMeta instantiates a new FeatureDefinitionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureDefinitionMeta() *FeatureDefinitionMeta {
	this := FeatureDefinitionMeta{}
	return &this
}

// NewFeatureDefinitionMetaWithDefaults instantiates a new FeatureDefinitionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureDefinitionMetaWithDefaults() *FeatureDefinitionMeta {
	this := FeatureDefinitionMeta{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FeatureDefinitionMeta) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDefinitionMeta) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FeatureDefinitionMeta) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FeatureDefinitionMeta) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureDefinitionMeta) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDefinitionMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureDefinitionMeta) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureDefinitionMeta) SetDescription(v string) {
	o.Description = &v
}

func (o FeatureDefinitionMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureDefinitionMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableFeatureDefinitionMeta struct {
	value *FeatureDefinitionMeta
	isSet bool
}

func (v NullableFeatureDefinitionMeta) Get() *FeatureDefinitionMeta {
	return v.value
}

func (v *NullableFeatureDefinitionMeta) Set(val *FeatureDefinitionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureDefinitionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureDefinitionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureDefinitionMeta(val *FeatureDefinitionMeta) *NullableFeatureDefinitionMeta {
	return &NullableFeatureDefinitionMeta{value: val, isSet: true}
}

func (v NullableFeatureDefinitionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureDefinitionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


