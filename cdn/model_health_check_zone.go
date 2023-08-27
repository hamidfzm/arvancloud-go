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

// checks if the HealthCheckZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckZone{}

// HealthCheckZone struct for HealthCheckZone
type HealthCheckZone struct {
	Id *string `json:"id,omitempty"`
	MonitoringLevel *string `json:"monitoring_level,omitempty"`
}

// NewHealthCheckZone instantiates a new HealthCheckZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckZone() *HealthCheckZone {
	this := HealthCheckZone{}
	return &this
}

// NewHealthCheckZoneWithDefaults instantiates a new HealthCheckZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckZoneWithDefaults() *HealthCheckZone {
	this := HealthCheckZone{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HealthCheckZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HealthCheckZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HealthCheckZone) SetId(v string) {
	o.Id = &v
}

// GetMonitoringLevel returns the MonitoringLevel field value if set, zero value otherwise.
func (o *HealthCheckZone) GetMonitoringLevel() string {
	if o == nil || IsNil(o.MonitoringLevel) {
		var ret string
		return ret
	}
	return *o.MonitoringLevel
}

// GetMonitoringLevelOk returns a tuple with the MonitoringLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckZone) GetMonitoringLevelOk() (*string, bool) {
	if o == nil || IsNil(o.MonitoringLevel) {
		return nil, false
	}
	return o.MonitoringLevel, true
}

// HasMonitoringLevel returns a boolean if a field has been set.
func (o *HealthCheckZone) HasMonitoringLevel() bool {
	if o != nil && !IsNil(o.MonitoringLevel) {
		return true
	}

	return false
}

// SetMonitoringLevel gets a reference to the given string and assigns it to the MonitoringLevel field.
func (o *HealthCheckZone) SetMonitoringLevel(v string) {
	o.MonitoringLevel = &v
}

func (o HealthCheckZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MonitoringLevel) {
		toSerialize["monitoring_level"] = o.MonitoringLevel
	}
	return toSerialize, nil
}

type NullableHealthCheckZone struct {
	value *HealthCheckZone
	isSet bool
}

func (v NullableHealthCheckZone) Get() *HealthCheckZone {
	return v.value
}

func (v *NullableHealthCheckZone) Set(val *HealthCheckZone) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckZone) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckZone(val *HealthCheckZone) *NullableHealthCheckZone {
	return &NullableHealthCheckZone{value: val, isSet: true}
}

func (v NullableHealthCheckZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

