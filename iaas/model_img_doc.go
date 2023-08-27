/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ImgDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImgDoc{}

// ImgDoc struct for ImgDoc
type ImgDoc struct {
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewImgDoc instantiates a new ImgDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImgDoc() *ImgDoc {
	this := ImgDoc{}
	return &this
}

// NewImgDocWithDefaults instantiates a new ImgDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImgDocWithDefaults() *ImgDoc {
	this := ImgDoc{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ImgDoc) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImgDoc) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ImgDoc) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *ImgDoc) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImgDoc) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImgDoc) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImgDoc) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImgDoc) SetName(v string) {
	o.Name = &v
}

func (o ImgDoc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImgDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableImgDoc struct {
	value *ImgDoc
	isSet bool
}

func (v NullableImgDoc) Get() *ImgDoc {
	return v.value
}

func (v *NullableImgDoc) Set(val *ImgDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableImgDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableImgDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImgDoc(val *ImgDoc) *NullableImgDoc {
	return &NullableImgDoc{value: val, isSet: true}
}

func (v NullableImgDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImgDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


