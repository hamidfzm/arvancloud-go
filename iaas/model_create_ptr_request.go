/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the CreatePTRRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePTRRequest{}

// CreatePTRRequest struct for CreatePTRRequest
type CreatePTRRequest struct {
	Domain *string `json:"domain,omitempty"`
	Ip *string `json:"ip,omitempty"`
}

// NewCreatePTRRequest instantiates a new CreatePTRRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePTRRequest() *CreatePTRRequest {
	this := CreatePTRRequest{}
	return &this
}

// NewCreatePTRRequestWithDefaults instantiates a new CreatePTRRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePTRRequestWithDefaults() *CreatePTRRequest {
	this := CreatePTRRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CreatePTRRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePTRRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CreatePTRRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CreatePTRRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CreatePTRRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePTRRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CreatePTRRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CreatePTRRequest) SetIp(v string) {
	o.Ip = &v
}

func (o CreatePTRRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePTRRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

type NullableCreatePTRRequest struct {
	value *CreatePTRRequest
	isSet bool
}

func (v NullableCreatePTRRequest) Get() *CreatePTRRequest {
	return v.value
}

func (v *NullableCreatePTRRequest) Set(val *CreatePTRRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePTRRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePTRRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePTRRequest(val *CreatePTRRequest) *NullableCreatePTRRequest {
	return &NullableCreatePTRRequest{value: val, isSet: true}
}

func (v NullableCreatePTRRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePTRRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


