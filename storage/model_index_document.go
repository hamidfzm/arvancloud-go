/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
)

// checks if the IndexDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexDocument{}

// IndexDocument Container for the <code>Suffix</code> element.
type IndexDocument struct {
	Suffix string `json:"Suffix"`
}

// NewIndexDocument instantiates a new IndexDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexDocument(suffix string) *IndexDocument {
	this := IndexDocument{}
	this.Suffix = suffix
	return &this
}

// NewIndexDocumentWithDefaults instantiates a new IndexDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexDocumentWithDefaults() *IndexDocument {
	this := IndexDocument{}
	return &this
}

// GetSuffix returns the Suffix field value
func (o *IndexDocument) GetSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value
// and a boolean to check if the value has been set.
func (o *IndexDocument) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suffix, true
}

// SetSuffix sets field value
func (o *IndexDocument) SetSuffix(v string) {
	o.Suffix = v
}

func (o IndexDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Suffix"] = o.Suffix
	return toSerialize, nil
}

type NullableIndexDocument struct {
	value *IndexDocument
	isSet bool
}

func (v NullableIndexDocument) Get() *IndexDocument {
	return v.value
}

func (v *NullableIndexDocument) Set(val *IndexDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexDocument(val *IndexDocument) *NullableIndexDocument {
	return &NullableIndexDocument{value: val, isSet: true}
}

func (v NullableIndexDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

