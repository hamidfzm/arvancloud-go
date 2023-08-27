/*
ArvanCloud Video Advertising Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vads

import (
	"encoding/json"
)

// checks if the DomainPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainPostRequest{}

// DomainPostRequest struct for DomainPostRequest
type DomainPostRequest struct {
	// The subdomain. Only contain lower case letters and digits
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewDomainPostRequest instantiates a new DomainPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainPostRequest() *DomainPostRequest {
	this := DomainPostRequest{}
	return &this
}

// NewDomainPostRequestWithDefaults instantiates a new DomainPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainPostRequestWithDefaults() *DomainPostRequest {
	this := DomainPostRequest{}
	return &this
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DomainPostRequest) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainPostRequest) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DomainPostRequest) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DomainPostRequest) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o DomainPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	return toSerialize, nil
}

type NullableDomainPostRequest struct {
	value *DomainPostRequest
	isSet bool
}

func (v NullableDomainPostRequest) Get() *DomainPostRequest {
	return v.value
}

func (v *NullableDomainPostRequest) Set(val *DomainPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainPostRequest(val *DomainPostRequest) *NullableDomainPostRequest {
	return &NullableDomainPostRequest{value: val, isSet: true}
}

func (v NullableDomainPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


