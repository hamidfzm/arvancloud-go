/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1LabelSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1LabelSelector{}

// V1LabelSelector A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
type V1LabelSelector struct {
	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []V1LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.
	MatchLabels map[string]interface{} `json:"matchLabels,omitempty"`
}

// NewV1LabelSelector instantiates a new V1LabelSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1LabelSelector() *V1LabelSelector {
	this := V1LabelSelector{}
	return &this
}

// NewV1LabelSelectorWithDefaults instantiates a new V1LabelSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LabelSelectorWithDefaults() *V1LabelSelector {
	this := V1LabelSelector{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *V1LabelSelector) GetMatchExpressions() []V1LabelSelectorRequirement {
	if o == nil || IsNil(o.MatchExpressions) {
		var ret []V1LabelSelectorRequirement
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LabelSelector) GetMatchExpressionsOk() ([]V1LabelSelectorRequirement, bool) {
	if o == nil || IsNil(o.MatchExpressions) {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *V1LabelSelector) HasMatchExpressions() bool {
	if o != nil && !IsNil(o.MatchExpressions) {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []V1LabelSelectorRequirement and assigns it to the MatchExpressions field.
func (o *V1LabelSelector) SetMatchExpressions(v []V1LabelSelectorRequirement) {
	o.MatchExpressions = v
}

// GetMatchLabels returns the MatchLabels field value if set, zero value otherwise.
func (o *V1LabelSelector) GetMatchLabels() map[string]interface{} {
	if o == nil || IsNil(o.MatchLabels) {
		var ret map[string]interface{}
		return ret
	}
	return o.MatchLabels
}

// GetMatchLabelsOk returns a tuple with the MatchLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LabelSelector) GetMatchLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MatchLabels) {
		return map[string]interface{}{}, false
	}
	return o.MatchLabels, true
}

// HasMatchLabels returns a boolean if a field has been set.
func (o *V1LabelSelector) HasMatchLabels() bool {
	if o != nil && !IsNil(o.MatchLabels) {
		return true
	}

	return false
}

// SetMatchLabels gets a reference to the given map[string]interface{} and assigns it to the MatchLabels field.
func (o *V1LabelSelector) SetMatchLabels(v map[string]interface{}) {
	o.MatchLabels = v
}

func (o V1LabelSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1LabelSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchExpressions) {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if !IsNil(o.MatchLabels) {
		toSerialize["matchLabels"] = o.MatchLabels
	}
	return toSerialize, nil
}

type NullableV1LabelSelector struct {
	value *V1LabelSelector
	isSet bool
}

func (v NullableV1LabelSelector) Get() *V1LabelSelector {
	return v.value
}

func (v *NullableV1LabelSelector) Set(val *V1LabelSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableV1LabelSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableV1LabelSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1LabelSelector(val *V1LabelSelector) *NullableV1LabelSelector {
	return &NullableV1LabelSelector{value: val, isSet: true}
}

func (v NullableV1LabelSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1LabelSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


