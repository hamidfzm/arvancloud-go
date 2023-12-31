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

// checks if the V1PodAffinityTerm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodAffinityTerm{}

// V1PodAffinityTerm Defines a set of pods (namely those matching the labelSelector relative to the given namespace(s)) that this pod should be co-located (affinity) or not co-located (anti-affinity) with, where co-located is defined as running on a node whose value of the label with key <topologyKey> matches that of any node on which a pod of the set of pods is running
type V1PodAffinityTerm struct {
	LabelSelector *V1LabelSelector `json:"labelSelector,omitempty"`
	// namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means \"this pod's namespace\"
	Namespaces []string `json:"namespaces,omitempty"`
	// This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
	TopologyKey string `json:"topologyKey"`
}

// NewV1PodAffinityTerm instantiates a new V1PodAffinityTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodAffinityTerm(topologyKey string) *V1PodAffinityTerm {
	this := V1PodAffinityTerm{}
	this.TopologyKey = topologyKey
	return &this
}

// NewV1PodAffinityTermWithDefaults instantiates a new V1PodAffinityTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodAffinityTermWithDefaults() *V1PodAffinityTerm {
	this := V1PodAffinityTerm{}
	return &this
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *V1PodAffinityTerm) GetLabelSelector() V1LabelSelector {
	if o == nil || IsNil(o.LabelSelector) {
		var ret V1LabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodAffinityTerm) GetLabelSelectorOk() (*V1LabelSelector, bool) {
	if o == nil || IsNil(o.LabelSelector) {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *V1PodAffinityTerm) HasLabelSelector() bool {
	if o != nil && !IsNil(o.LabelSelector) {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given V1LabelSelector and assigns it to the LabelSelector field.
func (o *V1PodAffinityTerm) SetLabelSelector(v V1LabelSelector) {
	o.LabelSelector = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *V1PodAffinityTerm) GetNamespaces() []string {
	if o == nil || IsNil(o.Namespaces) {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodAffinityTerm) GetNamespacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *V1PodAffinityTerm) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *V1PodAffinityTerm) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetTopologyKey returns the TopologyKey field value
func (o *V1PodAffinityTerm) GetTopologyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value
// and a boolean to check if the value has been set.
func (o *V1PodAffinityTerm) GetTopologyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyKey, true
}

// SetTopologyKey sets field value
func (o *V1PodAffinityTerm) SetTopologyKey(v string) {
	o.TopologyKey = v
}

func (o V1PodAffinityTerm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodAffinityTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelSelector) {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	toSerialize["topologyKey"] = o.TopologyKey
	return toSerialize, nil
}

type NullableV1PodAffinityTerm struct {
	value *V1PodAffinityTerm
	isSet bool
}

func (v NullableV1PodAffinityTerm) Get() *V1PodAffinityTerm {
	return v.value
}

func (v *NullableV1PodAffinityTerm) Set(val *V1PodAffinityTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodAffinityTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodAffinityTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodAffinityTerm(val *V1PodAffinityTerm) *NullableV1PodAffinityTerm {
	return &NullableV1PodAffinityTerm{value: val, isSet: true}
}

func (v NullableV1PodAffinityTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodAffinityTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


