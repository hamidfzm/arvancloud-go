/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the EmailForwardingRecipientsListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingRecipientsListData{}

// EmailForwardingRecipientsListData struct for EmailForwardingRecipientsListData
type EmailForwardingRecipientsListData struct {
	Data []EmailForwardingRecipientsListInner `json:"data,omitempty"`
}

// NewEmailForwardingRecipientsListData instantiates a new EmailForwardingRecipientsListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingRecipientsListData() *EmailForwardingRecipientsListData {
	this := EmailForwardingRecipientsListData{}
	return &this
}

// NewEmailForwardingRecipientsListDataWithDefaults instantiates a new EmailForwardingRecipientsListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingRecipientsListDataWithDefaults() *EmailForwardingRecipientsListData {
	this := EmailForwardingRecipientsListData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailForwardingRecipientsListData) GetData() []EmailForwardingRecipientsListInner {
	if o == nil || IsNil(o.Data) {
		var ret []EmailForwardingRecipientsListInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingRecipientsListData) GetDataOk() ([]EmailForwardingRecipientsListInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailForwardingRecipientsListData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailForwardingRecipientsListInner and assigns it to the Data field.
func (o *EmailForwardingRecipientsListData) SetData(v []EmailForwardingRecipientsListInner) {
	o.Data = v
}

func (o EmailForwardingRecipientsListData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingRecipientsListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailForwardingRecipientsListData struct {
	value *EmailForwardingRecipientsListData
	isSet bool
}

func (v NullableEmailForwardingRecipientsListData) Get() *EmailForwardingRecipientsListData {
	return v.value
}

func (v *NullableEmailForwardingRecipientsListData) Set(val *EmailForwardingRecipientsListData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingRecipientsListData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingRecipientsListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingRecipientsListData(val *EmailForwardingRecipientsListData) *NullableEmailForwardingRecipientsListData {
	return &NullableEmailForwardingRecipientsListData{value: val, isSet: true}
}

func (v NullableEmailForwardingRecipientsListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingRecipientsListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

