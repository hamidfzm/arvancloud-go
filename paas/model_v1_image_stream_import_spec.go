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

// checks if the V1ImageStreamImportSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ImageStreamImportSpec{}

// V1ImageStreamImportSpec ImageStreamImportSpec defines what images should be imported.
type V1ImageStreamImportSpec struct {
	// Images are a list of individual images to import.
	Images []V1ImageImportSpec `json:"images,omitempty"`
	// Import indicates whether to perform an import - if so, the specified tags are set on the spec and status of the image stream defined by the type meta.
	Import bool `json:"import"`
	Repository *V1RepositoryImportSpec `json:"repository,omitempty"`
}

// NewV1ImageStreamImportSpec instantiates a new V1ImageStreamImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageStreamImportSpec(import_ bool) *V1ImageStreamImportSpec {
	this := V1ImageStreamImportSpec{}
	this.Import = import_
	return &this
}

// NewV1ImageStreamImportSpecWithDefaults instantiates a new V1ImageStreamImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageStreamImportSpecWithDefaults() *V1ImageStreamImportSpec {
	this := V1ImageStreamImportSpec{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *V1ImageStreamImportSpec) GetImages() []V1ImageImportSpec {
	if o == nil || IsNil(o.Images) {
		var ret []V1ImageImportSpec
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageStreamImportSpec) GetImagesOk() ([]V1ImageImportSpec, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *V1ImageStreamImportSpec) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []V1ImageImportSpec and assigns it to the Images field.
func (o *V1ImageStreamImportSpec) SetImages(v []V1ImageImportSpec) {
	o.Images = v
}

// GetImport returns the Import field value
func (o *V1ImageStreamImportSpec) GetImport() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Import
}

// GetImportOk returns a tuple with the Import field value
// and a boolean to check if the value has been set.
func (o *V1ImageStreamImportSpec) GetImportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Import, true
}

// SetImport sets field value
func (o *V1ImageStreamImportSpec) SetImport(v bool) {
	o.Import = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *V1ImageStreamImportSpec) GetRepository() V1RepositoryImportSpec {
	if o == nil || IsNil(o.Repository) {
		var ret V1RepositoryImportSpec
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageStreamImportSpec) GetRepositoryOk() (*V1RepositoryImportSpec, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *V1ImageStreamImportSpec) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given V1RepositoryImportSpec and assigns it to the Repository field.
func (o *V1ImageStreamImportSpec) SetRepository(v V1RepositoryImportSpec) {
	o.Repository = &v
}

func (o V1ImageStreamImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ImageStreamImportSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	toSerialize["import"] = o.Import
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableV1ImageStreamImportSpec struct {
	value *V1ImageStreamImportSpec
	isSet bool
}

func (v NullableV1ImageStreamImportSpec) Get() *V1ImageStreamImportSpec {
	return v.value
}

func (v *NullableV1ImageStreamImportSpec) Set(val *V1ImageStreamImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageStreamImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageStreamImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageStreamImportSpec(val *V1ImageStreamImportSpec) *NullableV1ImageStreamImportSpec {
	return &NullableV1ImageStreamImportSpec{value: val, isSet: true}
}

func (v NullableV1ImageStreamImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageStreamImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


