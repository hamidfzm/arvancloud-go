/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
)

// checks if the LambdaFunctionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LambdaFunctionConfiguration{}

// LambdaFunctionConfiguration A container for specifying the configuration for Lambda notifications.
type LambdaFunctionConfiguration struct {
	// An optional unique identifier for configurations in a notification configuration. If you don't provide one, ArvanCloud S3 will assign an ID.
	Id *string `json:"Id,omitempty"`
	LambdaFunctionArn string `json:"LambdaFunctionArn"`
	Events Array `json:"Events"`
	Filter *NotificationConfigurationFilter `json:"Filter,omitempty"`
}

// NewLambdaFunctionConfiguration instantiates a new LambdaFunctionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLambdaFunctionConfiguration(lambdaFunctionArn string, events Array) *LambdaFunctionConfiguration {
	this := LambdaFunctionConfiguration{}
	this.LambdaFunctionArn = lambdaFunctionArn
	this.Events = events
	return &this
}

// NewLambdaFunctionConfigurationWithDefaults instantiates a new LambdaFunctionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLambdaFunctionConfigurationWithDefaults() *LambdaFunctionConfiguration {
	this := LambdaFunctionConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LambdaFunctionConfiguration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LambdaFunctionConfiguration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LambdaFunctionConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LambdaFunctionConfiguration) SetId(v string) {
	o.Id = &v
}

// GetLambdaFunctionArn returns the LambdaFunctionArn field value
func (o *LambdaFunctionConfiguration) GetLambdaFunctionArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LambdaFunctionArn
}

// GetLambdaFunctionArnOk returns a tuple with the LambdaFunctionArn field value
// and a boolean to check if the value has been set.
func (o *LambdaFunctionConfiguration) GetLambdaFunctionArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LambdaFunctionArn, true
}

// SetLambdaFunctionArn sets field value
func (o *LambdaFunctionConfiguration) SetLambdaFunctionArn(v string) {
	o.LambdaFunctionArn = v
}

// GetEvents returns the Events field value
func (o *LambdaFunctionConfiguration) GetEvents() Array {
	if o == nil {
		var ret Array
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *LambdaFunctionConfiguration) GetEventsOk() (*Array, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *LambdaFunctionConfiguration) SetEvents(v Array) {
	o.Events = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LambdaFunctionConfiguration) GetFilter() NotificationConfigurationFilter {
	if o == nil || IsNil(o.Filter) {
		var ret NotificationConfigurationFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LambdaFunctionConfiguration) GetFilterOk() (*NotificationConfigurationFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LambdaFunctionConfiguration) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NotificationConfigurationFilter and assigns it to the Filter field.
func (o *LambdaFunctionConfiguration) SetFilter(v NotificationConfigurationFilter) {
	o.Filter = &v
}

func (o LambdaFunctionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LambdaFunctionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["LambdaFunctionArn"] = o.LambdaFunctionArn
	toSerialize["Events"] = o.Events
	if !IsNil(o.Filter) {
		toSerialize["Filter"] = o.Filter
	}
	return toSerialize, nil
}

type NullableLambdaFunctionConfiguration struct {
	value *LambdaFunctionConfiguration
	isSet bool
}

func (v NullableLambdaFunctionConfiguration) Get() *LambdaFunctionConfiguration {
	return v.value
}

func (v *NullableLambdaFunctionConfiguration) Set(val *LambdaFunctionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableLambdaFunctionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableLambdaFunctionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLambdaFunctionConfiguration(val *LambdaFunctionConfiguration) *NullableLambdaFunctionConfiguration {
	return &NullableLambdaFunctionConfiguration{value: val, isSet: true}
}

func (v NullableLambdaFunctionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLambdaFunctionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


