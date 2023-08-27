/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"time"
)

// checks if the Part type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Part{}

// Part Container for elements related to a part.
type Part struct {
	PartNumber *int32 `json:"PartNumber,omitempty"`
	LastModified *time.Time `json:"LastModified,omitempty"`
	ETag *string `json:"ETag,omitempty"`
	Size *int32 `json:"Size,omitempty"`
}

// NewPart instantiates a new Part object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPart() *Part {
	this := Part{}
	return &this
}

// NewPartWithDefaults instantiates a new Part object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartWithDefaults() *Part {
	this := Part{}
	return &this
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *Part) GetPartNumber() int32 {
	if o == nil || IsNil(o.PartNumber) {
		var ret int32
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetPartNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *Part) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given int32 and assigns it to the PartNumber field.
func (o *Part) SetPartNumber(v int32) {
	o.PartNumber = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *Part) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *Part) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *Part) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *Part) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *Part) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *Part) SetETag(v string) {
	o.ETag = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Part) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Part) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Part) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Part) SetSize(v int32) {
	o.Size = &v
}

func (o Part) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Part) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PartNumber) {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if !IsNil(o.LastModified) {
		toSerialize["LastModified"] = o.LastModified
	}
	if !IsNil(o.ETag) {
		toSerialize["ETag"] = o.ETag
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	return toSerialize, nil
}

type NullablePart struct {
	value *Part
	isSet bool
}

func (v NullablePart) Get() *Part {
	return v.value
}

func (v *NullablePart) Set(val *Part) {
	v.value = val
	v.isSet = true
}

func (v NullablePart) IsSet() bool {
	return v.isSet
}

func (v *NullablePart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePart(val *Part) *NullablePart {
	return &NullablePart{value: val, isSet: true}
}

func (v NullablePart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


