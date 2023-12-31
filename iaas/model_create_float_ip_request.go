/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the CreateFloatIPRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFloatIPRequest{}

// CreateFloatIPRequest struct for CreateFloatIPRequest
type CreateFloatIPRequest struct {
	Description *string `json:"description,omitempty"`
}

// NewCreateFloatIPRequest instantiates a new CreateFloatIPRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFloatIPRequest() *CreateFloatIPRequest {
	this := CreateFloatIPRequest{}
	return &this
}

// NewCreateFloatIPRequestWithDefaults instantiates a new CreateFloatIPRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFloatIPRequestWithDefaults() *CreateFloatIPRequest {
	this := CreateFloatIPRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFloatIPRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFloatIPRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFloatIPRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFloatIPRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreateFloatIPRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFloatIPRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCreateFloatIPRequest struct {
	value *CreateFloatIPRequest
	isSet bool
}

func (v NullableCreateFloatIPRequest) Get() *CreateFloatIPRequest {
	return v.value
}

func (v *NullableCreateFloatIPRequest) Set(val *CreateFloatIPRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFloatIPRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFloatIPRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFloatIPRequest(val *CreateFloatIPRequest) *NullableCreateFloatIPRequest {
	return &NullableCreateFloatIPRequest{value: val, isSet: true}
}

func (v NullableCreateFloatIPRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFloatIPRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


