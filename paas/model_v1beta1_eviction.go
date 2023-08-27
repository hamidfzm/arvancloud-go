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

// checks if the V1beta1Eviction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1Eviction{}

// V1beta1Eviction Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A request to cause such an eviction is created by POSTing to .../pods/<pod name>/evictions.
type V1beta1Eviction struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	DeleteOptions *V1DeleteOptions `json:"deleteOptions,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
}

// NewV1beta1Eviction instantiates a new V1beta1Eviction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1Eviction() *V1beta1Eviction {
	this := V1beta1Eviction{}
	return &this
}

// NewV1beta1EvictionWithDefaults instantiates a new V1beta1Eviction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1EvictionWithDefaults() *V1beta1Eviction {
	this := V1beta1Eviction{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1beta1Eviction) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Eviction) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1beta1Eviction) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1beta1Eviction) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetDeleteOptions returns the DeleteOptions field value if set, zero value otherwise.
func (o *V1beta1Eviction) GetDeleteOptions() V1DeleteOptions {
	if o == nil || IsNil(o.DeleteOptions) {
		var ret V1DeleteOptions
		return ret
	}
	return *o.DeleteOptions
}

// GetDeleteOptionsOk returns a tuple with the DeleteOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Eviction) GetDeleteOptionsOk() (*V1DeleteOptions, bool) {
	if o == nil || IsNil(o.DeleteOptions) {
		return nil, false
	}
	return o.DeleteOptions, true
}

// HasDeleteOptions returns a boolean if a field has been set.
func (o *V1beta1Eviction) HasDeleteOptions() bool {
	if o != nil && !IsNil(o.DeleteOptions) {
		return true
	}

	return false
}

// SetDeleteOptions gets a reference to the given V1DeleteOptions and assigns it to the DeleteOptions field.
func (o *V1beta1Eviction) SetDeleteOptions(v V1DeleteOptions) {
	o.DeleteOptions = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1beta1Eviction) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Eviction) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1beta1Eviction) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1beta1Eviction) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1beta1Eviction) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1Eviction) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1beta1Eviction) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1beta1Eviction) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

func (o V1beta1Eviction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1Eviction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.DeleteOptions) {
		toSerialize["deleteOptions"] = o.DeleteOptions
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableV1beta1Eviction struct {
	value *V1beta1Eviction
	isSet bool
}

func (v NullableV1beta1Eviction) Get() *V1beta1Eviction {
	return v.value
}

func (v *NullableV1beta1Eviction) Set(val *V1beta1Eviction) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1Eviction) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1Eviction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1Eviction(val *V1beta1Eviction) *NullableV1beta1Eviction {
	return &NullableV1beta1Eviction{value: val, isSet: true}
}

func (v NullableV1beta1Eviction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1Eviction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

