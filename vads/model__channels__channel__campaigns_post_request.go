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

// checks if the ChannelsChannelCampaignsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelsChannelCampaignsPostRequest{}

// ChannelsChannelCampaignsPostRequest struct for ChannelsChannelCampaignsPostRequest
type ChannelsChannelCampaignsPostRequest struct {
	// Title of campaign
	Title string `json:"title"`
	// Description of campaign
	Description *string `json:"description,omitempty"`
	// Skip type - If ad skip_type is follow_campaign then this will be use
	SkipType string `json:"skip_type"`
	// Skip offset in seconds (required if skip type is allow)
	SkipOffset *int32 `json:"skip_offset,omitempty"`
	// If false then vast not working
	Active bool `json:"active"`
}

// NewChannelsChannelCampaignsPostRequest instantiates a new ChannelsChannelCampaignsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelsChannelCampaignsPostRequest(title string, skipType string, active bool) *ChannelsChannelCampaignsPostRequest {
	this := ChannelsChannelCampaignsPostRequest{}
	this.Title = title
	this.SkipType = skipType
	this.Active = active
	return &this
}

// NewChannelsChannelCampaignsPostRequestWithDefaults instantiates a new ChannelsChannelCampaignsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelsChannelCampaignsPostRequestWithDefaults() *ChannelsChannelCampaignsPostRequest {
	this := ChannelsChannelCampaignsPostRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *ChannelsChannelCampaignsPostRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ChannelsChannelCampaignsPostRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ChannelsChannelCampaignsPostRequest) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChannelsChannelCampaignsPostRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelsChannelCampaignsPostRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChannelsChannelCampaignsPostRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChannelsChannelCampaignsPostRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSkipType returns the SkipType field value
func (o *ChannelsChannelCampaignsPostRequest) GetSkipType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SkipType
}

// GetSkipTypeOk returns a tuple with the SkipType field value
// and a boolean to check if the value has been set.
func (o *ChannelsChannelCampaignsPostRequest) GetSkipTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkipType, true
}

// SetSkipType sets field value
func (o *ChannelsChannelCampaignsPostRequest) SetSkipType(v string) {
	o.SkipType = v
}

// GetSkipOffset returns the SkipOffset field value if set, zero value otherwise.
func (o *ChannelsChannelCampaignsPostRequest) GetSkipOffset() int32 {
	if o == nil || IsNil(o.SkipOffset) {
		var ret int32
		return ret
	}
	return *o.SkipOffset
}

// GetSkipOffsetOk returns a tuple with the SkipOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelsChannelCampaignsPostRequest) GetSkipOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.SkipOffset) {
		return nil, false
	}
	return o.SkipOffset, true
}

// HasSkipOffset returns a boolean if a field has been set.
func (o *ChannelsChannelCampaignsPostRequest) HasSkipOffset() bool {
	if o != nil && !IsNil(o.SkipOffset) {
		return true
	}

	return false
}

// SetSkipOffset gets a reference to the given int32 and assigns it to the SkipOffset field.
func (o *ChannelsChannelCampaignsPostRequest) SetSkipOffset(v int32) {
	o.SkipOffset = &v
}

// GetActive returns the Active field value
func (o *ChannelsChannelCampaignsPostRequest) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ChannelsChannelCampaignsPostRequest) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ChannelsChannelCampaignsPostRequest) SetActive(v bool) {
	o.Active = v
}

func (o ChannelsChannelCampaignsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelsChannelCampaignsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["skip_type"] = o.SkipType
	if !IsNil(o.SkipOffset) {
		toSerialize["skip_offset"] = o.SkipOffset
	}
	toSerialize["active"] = o.Active
	return toSerialize, nil
}

type NullableChannelsChannelCampaignsPostRequest struct {
	value *ChannelsChannelCampaignsPostRequest
	isSet bool
}

func (v NullableChannelsChannelCampaignsPostRequest) Get() *ChannelsChannelCampaignsPostRequest {
	return v.value
}

func (v *NullableChannelsChannelCampaignsPostRequest) Set(val *ChannelsChannelCampaignsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelsChannelCampaignsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelsChannelCampaignsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelsChannelCampaignsPostRequest(val *ChannelsChannelCampaignsPostRequest) *NullableChannelsChannelCampaignsPostRequest {
	return &NullableChannelsChannelCampaignsPostRequest{value: val, isSet: true}
}

func (v NullableChannelsChannelCampaignsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelsChannelCampaignsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


