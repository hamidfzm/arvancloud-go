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

// checks if the V1RepositoryImportSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RepositoryImportSpec{}

// V1RepositoryImportSpec RepositoryImportSpec describes a request to import images from a Docker image repository.
type V1RepositoryImportSpec struct {
	From V1ObjectReference `json:"from"`
	ImportPolicy *V1TagImportPolicy `json:"importPolicy,omitempty"`
	// IncludeManifest determines if the manifest for each image is returned in the response
	IncludeManifest *bool `json:"includeManifest,omitempty"`
	ReferencePolicy *V1TagReferencePolicy `json:"referencePolicy,omitempty"`
}

// NewV1RepositoryImportSpec instantiates a new V1RepositoryImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RepositoryImportSpec(from V1ObjectReference) *V1RepositoryImportSpec {
	this := V1RepositoryImportSpec{}
	this.From = from
	return &this
}

// NewV1RepositoryImportSpecWithDefaults instantiates a new V1RepositoryImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RepositoryImportSpecWithDefaults() *V1RepositoryImportSpec {
	this := V1RepositoryImportSpec{}
	return &this
}

// GetFrom returns the From field value
func (o *V1RepositoryImportSpec) GetFrom() V1ObjectReference {
	if o == nil {
		var ret V1ObjectReference
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *V1RepositoryImportSpec) GetFromOk() (*V1ObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *V1RepositoryImportSpec) SetFrom(v V1ObjectReference) {
	o.From = v
}

// GetImportPolicy returns the ImportPolicy field value if set, zero value otherwise.
func (o *V1RepositoryImportSpec) GetImportPolicy() V1TagImportPolicy {
	if o == nil || IsNil(o.ImportPolicy) {
		var ret V1TagImportPolicy
		return ret
	}
	return *o.ImportPolicy
}

// GetImportPolicyOk returns a tuple with the ImportPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RepositoryImportSpec) GetImportPolicyOk() (*V1TagImportPolicy, bool) {
	if o == nil || IsNil(o.ImportPolicy) {
		return nil, false
	}
	return o.ImportPolicy, true
}

// HasImportPolicy returns a boolean if a field has been set.
func (o *V1RepositoryImportSpec) HasImportPolicy() bool {
	if o != nil && !IsNil(o.ImportPolicy) {
		return true
	}

	return false
}

// SetImportPolicy gets a reference to the given V1TagImportPolicy and assigns it to the ImportPolicy field.
func (o *V1RepositoryImportSpec) SetImportPolicy(v V1TagImportPolicy) {
	o.ImportPolicy = &v
}

// GetIncludeManifest returns the IncludeManifest field value if set, zero value otherwise.
func (o *V1RepositoryImportSpec) GetIncludeManifest() bool {
	if o == nil || IsNil(o.IncludeManifest) {
		var ret bool
		return ret
	}
	return *o.IncludeManifest
}

// GetIncludeManifestOk returns a tuple with the IncludeManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RepositoryImportSpec) GetIncludeManifestOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeManifest) {
		return nil, false
	}
	return o.IncludeManifest, true
}

// HasIncludeManifest returns a boolean if a field has been set.
func (o *V1RepositoryImportSpec) HasIncludeManifest() bool {
	if o != nil && !IsNil(o.IncludeManifest) {
		return true
	}

	return false
}

// SetIncludeManifest gets a reference to the given bool and assigns it to the IncludeManifest field.
func (o *V1RepositoryImportSpec) SetIncludeManifest(v bool) {
	o.IncludeManifest = &v
}

// GetReferencePolicy returns the ReferencePolicy field value if set, zero value otherwise.
func (o *V1RepositoryImportSpec) GetReferencePolicy() V1TagReferencePolicy {
	if o == nil || IsNil(o.ReferencePolicy) {
		var ret V1TagReferencePolicy
		return ret
	}
	return *o.ReferencePolicy
}

// GetReferencePolicyOk returns a tuple with the ReferencePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RepositoryImportSpec) GetReferencePolicyOk() (*V1TagReferencePolicy, bool) {
	if o == nil || IsNil(o.ReferencePolicy) {
		return nil, false
	}
	return o.ReferencePolicy, true
}

// HasReferencePolicy returns a boolean if a field has been set.
func (o *V1RepositoryImportSpec) HasReferencePolicy() bool {
	if o != nil && !IsNil(o.ReferencePolicy) {
		return true
	}

	return false
}

// SetReferencePolicy gets a reference to the given V1TagReferencePolicy and assigns it to the ReferencePolicy field.
func (o *V1RepositoryImportSpec) SetReferencePolicy(v V1TagReferencePolicy) {
	o.ReferencePolicy = &v
}

func (o V1RepositoryImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RepositoryImportSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	if !IsNil(o.ImportPolicy) {
		toSerialize["importPolicy"] = o.ImportPolicy
	}
	if !IsNil(o.IncludeManifest) {
		toSerialize["includeManifest"] = o.IncludeManifest
	}
	if !IsNil(o.ReferencePolicy) {
		toSerialize["referencePolicy"] = o.ReferencePolicy
	}
	return toSerialize, nil
}

type NullableV1RepositoryImportSpec struct {
	value *V1RepositoryImportSpec
	isSet bool
}

func (v NullableV1RepositoryImportSpec) Get() *V1RepositoryImportSpec {
	return v.value
}

func (v *NullableV1RepositoryImportSpec) Set(val *V1RepositoryImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RepositoryImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RepositoryImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RepositoryImportSpec(val *V1RepositoryImportSpec) *NullableV1RepositoryImportSpec {
	return &NullableV1RepositoryImportSpec{value: val, isSet: true}
}

func (v NullableV1RepositoryImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RepositoryImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

