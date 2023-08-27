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

// checks if the V1SecurityContextConstraintsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SecurityContextConstraintsList{}

// V1SecurityContextConstraintsList SecurityContextConstraintsList is a list of SecurityContextConstraints objects
type V1SecurityContextConstraintsList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// List of security context constraints.
	Items []V1SecurityContextConstraints `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ListMeta `json:"metadata,omitempty"`
}

// NewV1SecurityContextConstraintsList instantiates a new V1SecurityContextConstraintsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SecurityContextConstraintsList(items []V1SecurityContextConstraints) *V1SecurityContextConstraintsList {
	this := V1SecurityContextConstraintsList{}
	this.Items = items
	return &this
}

// NewV1SecurityContextConstraintsListWithDefaults instantiates a new V1SecurityContextConstraintsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SecurityContextConstraintsListWithDefaults() *V1SecurityContextConstraintsList {
	this := V1SecurityContextConstraintsList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1SecurityContextConstraintsList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContextConstraintsList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1SecurityContextConstraintsList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1SecurityContextConstraintsList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *V1SecurityContextConstraintsList) GetItems() []V1SecurityContextConstraints {
	if o == nil {
		var ret []V1SecurityContextConstraints
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *V1SecurityContextConstraintsList) GetItemsOk() ([]V1SecurityContextConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *V1SecurityContextConstraintsList) SetItems(v []V1SecurityContextConstraints) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1SecurityContextConstraintsList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContextConstraintsList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1SecurityContextConstraintsList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1SecurityContextConstraintsList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1SecurityContextConstraintsList) GetMetadata() V1ListMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContextConstraintsList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1SecurityContextConstraintsList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *V1SecurityContextConstraintsList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o V1SecurityContextConstraintsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SecurityContextConstraintsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableV1SecurityContextConstraintsList struct {
	value *V1SecurityContextConstraintsList
	isSet bool
}

func (v NullableV1SecurityContextConstraintsList) Get() *V1SecurityContextConstraintsList {
	return v.value
}

func (v *NullableV1SecurityContextConstraintsList) Set(val *V1SecurityContextConstraintsList) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SecurityContextConstraintsList) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SecurityContextConstraintsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SecurityContextConstraintsList(val *V1SecurityContextConstraintsList) *NullableV1SecurityContextConstraintsList {
	return &NullableV1SecurityContextConstraintsList{value: val, isSet: true}
}

func (v NullableV1SecurityContextConstraintsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SecurityContextConstraintsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


