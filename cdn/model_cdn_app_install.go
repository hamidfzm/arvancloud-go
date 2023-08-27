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

// checks if the CdnAppInstall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdnAppInstall{}

// CdnAppInstall struct for CdnAppInstall
type CdnAppInstall struct {
	IsInstall *bool `json:"is_install,omitempty"`
}

// NewCdnAppInstall instantiates a new CdnAppInstall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnAppInstall() *CdnAppInstall {
	this := CdnAppInstall{}
	return &this
}

// NewCdnAppInstallWithDefaults instantiates a new CdnAppInstall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnAppInstallWithDefaults() *CdnAppInstall {
	this := CdnAppInstall{}
	return &this
}

// GetIsInstall returns the IsInstall field value if set, zero value otherwise.
func (o *CdnAppInstall) GetIsInstall() bool {
	if o == nil || IsNil(o.IsInstall) {
		var ret bool
		return ret
	}
	return *o.IsInstall
}

// GetIsInstallOk returns a tuple with the IsInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnAppInstall) GetIsInstallOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInstall) {
		return nil, false
	}
	return o.IsInstall, true
}

// HasIsInstall returns a boolean if a field has been set.
func (o *CdnAppInstall) HasIsInstall() bool {
	if o != nil && !IsNil(o.IsInstall) {
		return true
	}

	return false
}

// SetIsInstall gets a reference to the given bool and assigns it to the IsInstall field.
func (o *CdnAppInstall) SetIsInstall(v bool) {
	o.IsInstall = &v
}

func (o CdnAppInstall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdnAppInstall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsInstall) {
		toSerialize["is_install"] = o.IsInstall
	}
	return toSerialize, nil
}

type NullableCdnAppInstall struct {
	value *CdnAppInstall
	isSet bool
}

func (v NullableCdnAppInstall) Get() *CdnAppInstall {
	return v.value
}

func (v *NullableCdnAppInstall) Set(val *CdnAppInstall) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnAppInstall) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnAppInstall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnAppInstall(val *CdnAppInstall) *NullableCdnAppInstall {
	return &NullableCdnAppInstall{value: val, isSet: true}
}

func (v NullableCdnAppInstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnAppInstall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

