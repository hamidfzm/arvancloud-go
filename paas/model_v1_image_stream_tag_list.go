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

// checks if the V1ImageStreamTagList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ImageStreamTagList{}

// V1ImageStreamTagList ImageStreamTagList is a list of ImageStreamTag objects.
type V1ImageStreamTagList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Items is the list of image stream tags
	Items []V1ImageStreamTag `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ListMeta `json:"metadata,omitempty"`
}

// NewV1ImageStreamTagList instantiates a new V1ImageStreamTagList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageStreamTagList(items []V1ImageStreamTag) *V1ImageStreamTagList {
	this := V1ImageStreamTagList{}
	this.Items = items
	return &this
}

// NewV1ImageStreamTagListWithDefaults instantiates a new V1ImageStreamTagList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageStreamTagListWithDefaults() *V1ImageStreamTagList {
	this := V1ImageStreamTagList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1ImageStreamTagList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageStreamTagList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1ImageStreamTagList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1ImageStreamTagList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *V1ImageStreamTagList) GetItems() []V1ImageStreamTag {
	if o == nil {
		var ret []V1ImageStreamTag
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *V1ImageStreamTagList) GetItemsOk() ([]V1ImageStreamTag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *V1ImageStreamTagList) SetItems(v []V1ImageStreamTag) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1ImageStreamTagList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageStreamTagList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1ImageStreamTagList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1ImageStreamTagList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1ImageStreamTagList) GetMetadata() V1ListMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageStreamTagList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1ImageStreamTagList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *V1ImageStreamTagList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o V1ImageStreamTagList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ImageStreamTagList) ToMap() (map[string]interface{}, error) {
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

type NullableV1ImageStreamTagList struct {
	value *V1ImageStreamTagList
	isSet bool
}

func (v NullableV1ImageStreamTagList) Get() *V1ImageStreamTagList {
	return v.value
}

func (v *NullableV1ImageStreamTagList) Set(val *V1ImageStreamTagList) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageStreamTagList) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageStreamTagList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageStreamTagList(val *V1ImageStreamTagList) *NullableV1ImageStreamTagList {
	return &NullableV1ImageStreamTagList{value: val, isSet: true}
}

func (v NullableV1ImageStreamTagList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageStreamTagList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


