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

// checks if the LogForwarderDNSType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderDNSType{}

// LogForwarderDNSType Dns log type
type LogForwarderDNSType struct {
	Timestamp *bool `json:"timestamp,omitempty"`
	Uuid *bool `json:"uuid,omitempty"`
	Record *bool `json:"record,omitempty"`
	Type *bool `json:"type,omitempty"`
	Ip *bool `json:"ip,omitempty"`
	Country *bool `json:"country,omitempty"`
	Asn *bool `json:"asn,omitempty"`
	ResponseCode *bool `json:"response_code,omitempty"`
	ProcessTime *bool `json:"process_time,omitempty"`
}

// NewLogForwarderDNSType instantiates a new LogForwarderDNSType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderDNSType() *LogForwarderDNSType {
	this := LogForwarderDNSType{}
	return &this
}

// NewLogForwarderDNSTypeWithDefaults instantiates a new LogForwarderDNSType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderDNSTypeWithDefaults() *LogForwarderDNSType {
	this := LogForwarderDNSType{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetTimestamp() bool {
	if o == nil || IsNil(o.Timestamp) {
		var ret bool
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetTimestampOk() (*bool, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given bool and assigns it to the Timestamp field.
func (o *LogForwarderDNSType) SetTimestamp(v bool) {
	o.Timestamp = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetUuid() bool {
	if o == nil || IsNil(o.Uuid) {
		var ret bool
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetUuidOk() (*bool, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given bool and assigns it to the Uuid field.
func (o *LogForwarderDNSType) SetUuid(v bool) {
	o.Uuid = &v
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetRecord() bool {
	if o == nil || IsNil(o.Record) {
		var ret bool
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetRecordOk() (*bool, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given bool and assigns it to the Record field.
func (o *LogForwarderDNSType) SetRecord(v bool) {
	o.Record = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetType() bool {
	if o == nil || IsNil(o.Type) {
		var ret bool
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given bool and assigns it to the Type field.
func (o *LogForwarderDNSType) SetType(v bool) {
	o.Type = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetIp() bool {
	if o == nil || IsNil(o.Ip) {
		var ret bool
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetIpOk() (*bool, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given bool and assigns it to the Ip field.
func (o *LogForwarderDNSType) SetIp(v bool) {
	o.Ip = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetCountry() bool {
	if o == nil || IsNil(o.Country) {
		var ret bool
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetCountryOk() (*bool, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given bool and assigns it to the Country field.
func (o *LogForwarderDNSType) SetCountry(v bool) {
	o.Country = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetAsn() bool {
	if o == nil || IsNil(o.Asn) {
		var ret bool
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetAsnOk() (*bool, bool) {
	if o == nil || IsNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasAsn() bool {
	if o != nil && !IsNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given bool and assigns it to the Asn field.
func (o *LogForwarderDNSType) SetAsn(v bool) {
	o.Asn = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetResponseCode() bool {
	if o == nil || IsNil(o.ResponseCode) {
		var ret bool
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetResponseCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasResponseCode() bool {
	if o != nil && !IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given bool and assigns it to the ResponseCode field.
func (o *LogForwarderDNSType) SetResponseCode(v bool) {
	o.ResponseCode = &v
}

// GetProcessTime returns the ProcessTime field value if set, zero value otherwise.
func (o *LogForwarderDNSType) GetProcessTime() bool {
	if o == nil || IsNil(o.ProcessTime) {
		var ret bool
		return ret
	}
	return *o.ProcessTime
}

// GetProcessTimeOk returns a tuple with the ProcessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderDNSType) GetProcessTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.ProcessTime) {
		return nil, false
	}
	return o.ProcessTime, true
}

// HasProcessTime returns a boolean if a field has been set.
func (o *LogForwarderDNSType) HasProcessTime() bool {
	if o != nil && !IsNil(o.ProcessTime) {
		return true
	}

	return false
}

// SetProcessTime gets a reference to the given bool and assigns it to the ProcessTime field.
func (o *LogForwarderDNSType) SetProcessTime(v bool) {
	o.ProcessTime = &v
}

func (o LogForwarderDNSType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderDNSType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !IsNil(o.ResponseCode) {
		toSerialize["response_code"] = o.ResponseCode
	}
	if !IsNil(o.ProcessTime) {
		toSerialize["process_time"] = o.ProcessTime
	}
	return toSerialize, nil
}

type NullableLogForwarderDNSType struct {
	value *LogForwarderDNSType
	isSet bool
}

func (v NullableLogForwarderDNSType) Get() *LogForwarderDNSType {
	return v.value
}

func (v *NullableLogForwarderDNSType) Set(val *LogForwarderDNSType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderDNSType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderDNSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderDNSType(val *LogForwarderDNSType) *NullableLogForwarderDNSType {
	return &NullableLogForwarderDNSType{value: val, isSet: true}
}

func (v NullableLogForwarderDNSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderDNSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

