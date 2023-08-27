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

// checks if the LogForwarderGeneric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderGeneric{}

// LogForwarderGeneric struct for LogForwarderGeneric
type LogForwarderGeneric struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ConnectionType *string `json:"connection_type,omitempty"`
	DataFormat map[string]interface{} `json:"data_format,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"`
	Status *bool `json:"status,omitempty"`
}

// NewLogForwarderGeneric instantiates a new LogForwarderGeneric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderGeneric() *LogForwarderGeneric {
	this := LogForwarderGeneric{}
	return &this
}

// NewLogForwarderGenericWithDefaults instantiates a new LogForwarderGeneric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderGenericWithDefaults() *LogForwarderGeneric {
	this := LogForwarderGeneric{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogForwarderGeneric) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogForwarderGeneric) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogForwarderGeneric) SetType(v string) {
	o.Type = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetConnectionType() string {
	if o == nil || IsNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *LogForwarderGeneric) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetDataFormat returns the DataFormat field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetDataFormat() map[string]interface{} {
	if o == nil || IsNil(o.DataFormat) {
		var ret map[string]interface{}
		return ret
	}
	return o.DataFormat
}

// GetDataFormatOk returns a tuple with the DataFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetDataFormatOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DataFormat) {
		return map[string]interface{}{}, false
	}
	return o.DataFormat, true
}

// HasDataFormat returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasDataFormat() bool {
	if o != nil && !IsNil(o.DataFormat) {
		return true
	}

	return false
}

// SetDataFormat gets a reference to the given map[string]interface{} and assigns it to the DataFormat field.
func (o *LogForwarderGeneric) SetDataFormat(v map[string]interface{}) {
	o.DataFormat = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *LogForwarderGeneric) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogForwarderGeneric) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderGeneric) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogForwarderGeneric) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LogForwarderGeneric) SetStatus(v bool) {
	o.Status = &v
}

func (o LogForwarderGeneric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderGeneric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ConnectionType) {
		toSerialize["connection_type"] = o.ConnectionType
	}
	if !IsNil(o.DataFormat) {
		toSerialize["data_format"] = o.DataFormat
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableLogForwarderGeneric struct {
	value *LogForwarderGeneric
	isSet bool
}

func (v NullableLogForwarderGeneric) Get() *LogForwarderGeneric {
	return v.value
}

func (v *NullableLogForwarderGeneric) Set(val *LogForwarderGeneric) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderGeneric) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderGeneric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderGeneric(val *LogForwarderGeneric) *NullableLogForwarderGeneric {
	return &NullableLogForwarderGeneric{value: val, isSet: true}
}

func (v NullableLogForwarderGeneric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderGeneric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

