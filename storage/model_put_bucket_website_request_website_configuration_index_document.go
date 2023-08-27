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

// checks if the PutBucketWebsiteRequestWebsiteConfigurationIndexDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketWebsiteRequestWebsiteConfigurationIndexDocument{}

// PutBucketWebsiteRequestWebsiteConfigurationIndexDocument struct for PutBucketWebsiteRequestWebsiteConfigurationIndexDocument
type PutBucketWebsiteRequestWebsiteConfigurationIndexDocument struct {
	Suffix string `json:"Suffix"`
}

// NewPutBucketWebsiteRequestWebsiteConfigurationIndexDocument instantiates a new PutBucketWebsiteRequestWebsiteConfigurationIndexDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketWebsiteRequestWebsiteConfigurationIndexDocument(suffix string) *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument {
	this := PutBucketWebsiteRequestWebsiteConfigurationIndexDocument{}
	this.Suffix = suffix
	return &this
}

// NewPutBucketWebsiteRequestWebsiteConfigurationIndexDocumentWithDefaults instantiates a new PutBucketWebsiteRequestWebsiteConfigurationIndexDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketWebsiteRequestWebsiteConfigurationIndexDocumentWithDefaults() *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument {
	this := PutBucketWebsiteRequestWebsiteConfigurationIndexDocument{}
	return &this
}

// GetSuffix returns the Suffix field value
func (o *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) GetSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value
// and a boolean to check if the value has been set.
func (o *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suffix, true
}

// SetSuffix sets field value
func (o *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) SetSuffix(v string) {
	o.Suffix = v
}

func (o PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Suffix"] = o.Suffix
	return toSerialize, nil
}

type NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument struct {
	value *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument
	isSet bool
}

func (v NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument) Get() *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument {
	return v.value
}

func (v *NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument) Set(val *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument(val *PutBucketWebsiteRequestWebsiteConfigurationIndexDocument) *NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument {
	return &NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument{value: val, isSet: true}
}

func (v NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketWebsiteRequestWebsiteConfigurationIndexDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

