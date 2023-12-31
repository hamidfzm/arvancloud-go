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

// checks if the ChannelsChannelPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelsChannelPutRequest{}

// ChannelsChannelPutRequest struct for ChannelsChannelPutRequest
type ChannelsChannelPutRequest struct {
	// Title of channel
	Title *string `json:"title,omitempty"`
	// Description of channel
	Description *string `json:"description,omitempty"`
}

// NewChannelsChannelPutRequest instantiates a new ChannelsChannelPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelsChannelPutRequest() *ChannelsChannelPutRequest {
	this := ChannelsChannelPutRequest{}
	return &this
}

// NewChannelsChannelPutRequestWithDefaults instantiates a new ChannelsChannelPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelsChannelPutRequestWithDefaults() *ChannelsChannelPutRequest {
	this := ChannelsChannelPutRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ChannelsChannelPutRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelsChannelPutRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ChannelsChannelPutRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ChannelsChannelPutRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChannelsChannelPutRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelsChannelPutRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChannelsChannelPutRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChannelsChannelPutRequest) SetDescription(v string) {
	o.Description = &v
}

func (o ChannelsChannelPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelsChannelPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableChannelsChannelPutRequest struct {
	value *ChannelsChannelPutRequest
	isSet bool
}

func (v NullableChannelsChannelPutRequest) Get() *ChannelsChannelPutRequest {
	return v.value
}

func (v *NullableChannelsChannelPutRequest) Set(val *ChannelsChannelPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelsChannelPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelsChannelPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelsChannelPutRequest(val *ChannelsChannelPutRequest) *NullableChannelsChannelPutRequest {
	return &NullableChannelsChannelPutRequest{value: val, isSet: true}
}

func (v NullableChannelsChannelPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelsChannelPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


