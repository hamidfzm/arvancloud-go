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

// checks if the V1Binding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Binding{}

// V1Binding Binding ties one object to another; for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.
type V1Binding struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Target V1ObjectReference `json:"target"`
}

// NewV1Binding instantiates a new V1Binding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Binding(target V1ObjectReference) *V1Binding {
	this := V1Binding{}
	this.Target = target
	return &this
}

// NewV1BindingWithDefaults instantiates a new V1Binding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BindingWithDefaults() *V1Binding {
	this := V1Binding{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1Binding) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Binding) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1Binding) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1Binding) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1Binding) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Binding) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1Binding) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1Binding) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1Binding) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Binding) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1Binding) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1Binding) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetTarget returns the Target field value
func (o *V1Binding) GetTarget() V1ObjectReference {
	if o == nil {
		var ret V1ObjectReference
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *V1Binding) GetTargetOk() (*V1ObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *V1Binding) SetTarget(v V1ObjectReference) {
	o.Target = v
}

func (o V1Binding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Binding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

type NullableV1Binding struct {
	value *V1Binding
	isSet bool
}

func (v NullableV1Binding) Get() *V1Binding {
	return v.value
}

func (v *NullableV1Binding) Set(val *V1Binding) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Binding) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Binding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Binding(val *V1Binding) *NullableV1Binding {
	return &NullableV1Binding{value: val, isSet: true}
}

func (v NullableV1Binding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Binding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


