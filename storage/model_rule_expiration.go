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

// checks if the RuleExpiration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleExpiration{}

// RuleExpiration struct for RuleExpiration
type RuleExpiration struct {
	Date *time.Time `json:"Date,omitempty"`
	Days *int32 `json:"Days,omitempty"`
	ExpiredObjectDeleteMarker *bool `json:"ExpiredObjectDeleteMarker,omitempty"`
}

// NewRuleExpiration instantiates a new RuleExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleExpiration() *RuleExpiration {
	this := RuleExpiration{}
	return &this
}

// NewRuleExpirationWithDefaults instantiates a new RuleExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleExpirationWithDefaults() *RuleExpiration {
	this := RuleExpiration{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *RuleExpiration) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleExpiration) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *RuleExpiration) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *RuleExpiration) SetDate(v time.Time) {
	o.Date = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *RuleExpiration) GetDays() int32 {
	if o == nil || IsNil(o.Days) {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleExpiration) GetDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *RuleExpiration) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *RuleExpiration) SetDays(v int32) {
	o.Days = &v
}

// GetExpiredObjectDeleteMarker returns the ExpiredObjectDeleteMarker field value if set, zero value otherwise.
func (o *RuleExpiration) GetExpiredObjectDeleteMarker() bool {
	if o == nil || IsNil(o.ExpiredObjectDeleteMarker) {
		var ret bool
		return ret
	}
	return *o.ExpiredObjectDeleteMarker
}

// GetExpiredObjectDeleteMarkerOk returns a tuple with the ExpiredObjectDeleteMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleExpiration) GetExpiredObjectDeleteMarkerOk() (*bool, bool) {
	if o == nil || IsNil(o.ExpiredObjectDeleteMarker) {
		return nil, false
	}
	return o.ExpiredObjectDeleteMarker, true
}

// HasExpiredObjectDeleteMarker returns a boolean if a field has been set.
func (o *RuleExpiration) HasExpiredObjectDeleteMarker() bool {
	if o != nil && !IsNil(o.ExpiredObjectDeleteMarker) {
		return true
	}

	return false
}

// SetExpiredObjectDeleteMarker gets a reference to the given bool and assigns it to the ExpiredObjectDeleteMarker field.
func (o *RuleExpiration) SetExpiredObjectDeleteMarker(v bool) {
	o.ExpiredObjectDeleteMarker = &v
}

func (o RuleExpiration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleExpiration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.Days) {
		toSerialize["Days"] = o.Days
	}
	if !IsNil(o.ExpiredObjectDeleteMarker) {
		toSerialize["ExpiredObjectDeleteMarker"] = o.ExpiredObjectDeleteMarker
	}
	return toSerialize, nil
}

type NullableRuleExpiration struct {
	value *RuleExpiration
	isSet bool
}

func (v NullableRuleExpiration) Get() *RuleExpiration {
	return v.value
}

func (v *NullableRuleExpiration) Set(val *RuleExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleExpiration(val *RuleExpiration) *NullableRuleExpiration {
	return &NullableRuleExpiration{value: val, isSet: true}
}

func (v NullableRuleExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


