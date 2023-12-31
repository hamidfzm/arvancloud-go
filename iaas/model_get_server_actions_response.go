/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the GetServerActionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServerActionsResponse{}

// GetServerActionsResponse struct for GetServerActionsResponse
type GetServerActionsResponse struct {
	Action *string `json:"action,omitempty"`
	Message *string `json:"message,omitempty"`
	StartTime *string `json:"start_time,omitempty"`
}

// NewGetServerActionsResponse instantiates a new GetServerActionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServerActionsResponse() *GetServerActionsResponse {
	this := GetServerActionsResponse{}
	return &this
}

// NewGetServerActionsResponseWithDefaults instantiates a new GetServerActionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServerActionsResponseWithDefaults() *GetServerActionsResponse {
	this := GetServerActionsResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GetServerActionsResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerActionsResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GetServerActionsResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *GetServerActionsResponse) SetAction(v string) {
	o.Action = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetServerActionsResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerActionsResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetServerActionsResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetServerActionsResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetServerActionsResponse) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServerActionsResponse) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetServerActionsResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *GetServerActionsResponse) SetStartTime(v string) {
	o.StartTime = &v
}

func (o GetServerActionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServerActionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableGetServerActionsResponse struct {
	value *GetServerActionsResponse
	isSet bool
}

func (v NullableGetServerActionsResponse) Get() *GetServerActionsResponse {
	return v.value
}

func (v *NullableGetServerActionsResponse) Set(val *GetServerActionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServerActionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServerActionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServerActionsResponse(val *GetServerActionsResponse) *NullableGetServerActionsResponse {
	return &NullableGetServerActionsResponse{value: val, isSet: true}
}

func (v NullableGetServerActionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServerActionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


