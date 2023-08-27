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

// checks if the V1ReplicationControllerList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ReplicationControllerList{}

// V1ReplicationControllerList ReplicationControllerList is a collection of replication controllers.
type V1ReplicationControllerList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// List of replication controllers. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller
	Items []V1ReplicationController `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ListMeta `json:"metadata,omitempty"`
}

// NewV1ReplicationControllerList instantiates a new V1ReplicationControllerList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ReplicationControllerList(items []V1ReplicationController) *V1ReplicationControllerList {
	this := V1ReplicationControllerList{}
	this.Items = items
	return &this
}

// NewV1ReplicationControllerListWithDefaults instantiates a new V1ReplicationControllerList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ReplicationControllerListWithDefaults() *V1ReplicationControllerList {
	this := V1ReplicationControllerList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1ReplicationControllerList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1ReplicationControllerList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1ReplicationControllerList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *V1ReplicationControllerList) GetItems() []V1ReplicationController {
	if o == nil {
		var ret []V1ReplicationController
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerList) GetItemsOk() ([]V1ReplicationController, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *V1ReplicationControllerList) SetItems(v []V1ReplicationController) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1ReplicationControllerList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1ReplicationControllerList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1ReplicationControllerList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1ReplicationControllerList) GetMetadata() V1ListMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1ReplicationControllerList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *V1ReplicationControllerList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o V1ReplicationControllerList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ReplicationControllerList) ToMap() (map[string]interface{}, error) {
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

type NullableV1ReplicationControllerList struct {
	value *V1ReplicationControllerList
	isSet bool
}

func (v NullableV1ReplicationControllerList) Get() *V1ReplicationControllerList {
	return v.value
}

func (v *NullableV1ReplicationControllerList) Set(val *V1ReplicationControllerList) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ReplicationControllerList) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ReplicationControllerList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ReplicationControllerList(val *V1ReplicationControllerList) *NullableV1ReplicationControllerList {
	return &NullableV1ReplicationControllerList{value: val, isSet: true}
}

func (v NullableV1ReplicationControllerList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ReplicationControllerList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

