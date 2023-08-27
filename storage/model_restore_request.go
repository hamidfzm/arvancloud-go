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

// checks if the RestoreRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreRequest{}

// RestoreRequest Container for restore job parameters.
type RestoreRequest struct {
	Days *int32 `json:"Days,omitempty"`
	GlacierJobParameters *RestoreRequestGlacierJobParameters `json:"GlacierJobParameters,omitempty"`
	Type *RestoreRequestType `json:"Type,omitempty"`
	Tier *Tier `json:"Tier,omitempty"`
	Description *string `json:"Description,omitempty"`
	SelectParameters *RestoreRequestSelectParameters `json:"SelectParameters,omitempty"`
	OutputLocation *RestoreRequestOutputLocation `json:"OutputLocation,omitempty"`
}

// NewRestoreRequest instantiates a new RestoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreRequest() *RestoreRequest {
	this := RestoreRequest{}
	return &this
}

// NewRestoreRequestWithDefaults instantiates a new RestoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreRequestWithDefaults() *RestoreRequest {
	this := RestoreRequest{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *RestoreRequest) GetDays() int32 {
	if o == nil || IsNil(o.Days) {
		var ret int32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.Days) {
		return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *RestoreRequest) HasDays() bool {
	if o != nil && !IsNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given int32 and assigns it to the Days field.
func (o *RestoreRequest) SetDays(v int32) {
	o.Days = &v
}

// GetGlacierJobParameters returns the GlacierJobParameters field value if set, zero value otherwise.
func (o *RestoreRequest) GetGlacierJobParameters() RestoreRequestGlacierJobParameters {
	if o == nil || IsNil(o.GlacierJobParameters) {
		var ret RestoreRequestGlacierJobParameters
		return ret
	}
	return *o.GlacierJobParameters
}

// GetGlacierJobParametersOk returns a tuple with the GlacierJobParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetGlacierJobParametersOk() (*RestoreRequestGlacierJobParameters, bool) {
	if o == nil || IsNil(o.GlacierJobParameters) {
		return nil, false
	}
	return o.GlacierJobParameters, true
}

// HasGlacierJobParameters returns a boolean if a field has been set.
func (o *RestoreRequest) HasGlacierJobParameters() bool {
	if o != nil && !IsNil(o.GlacierJobParameters) {
		return true
	}

	return false
}

// SetGlacierJobParameters gets a reference to the given RestoreRequestGlacierJobParameters and assigns it to the GlacierJobParameters field.
func (o *RestoreRequest) SetGlacierJobParameters(v RestoreRequestGlacierJobParameters) {
	o.GlacierJobParameters = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestoreRequest) GetType() RestoreRequestType {
	if o == nil || IsNil(o.Type) {
		var ret RestoreRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetTypeOk() (*RestoreRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestoreRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RestoreRequestType and assigns it to the Type field.
func (o *RestoreRequest) SetType(v RestoreRequestType) {
	o.Type = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *RestoreRequest) GetTier() Tier {
	if o == nil || IsNil(o.Tier) {
		var ret Tier
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetTierOk() (*Tier, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *RestoreRequest) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given Tier and assigns it to the Tier field.
func (o *RestoreRequest) SetTier(v Tier) {
	o.Tier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RestoreRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RestoreRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RestoreRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSelectParameters returns the SelectParameters field value if set, zero value otherwise.
func (o *RestoreRequest) GetSelectParameters() RestoreRequestSelectParameters {
	if o == nil || IsNil(o.SelectParameters) {
		var ret RestoreRequestSelectParameters
		return ret
	}
	return *o.SelectParameters
}

// GetSelectParametersOk returns a tuple with the SelectParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetSelectParametersOk() (*RestoreRequestSelectParameters, bool) {
	if o == nil || IsNil(o.SelectParameters) {
		return nil, false
	}
	return o.SelectParameters, true
}

// HasSelectParameters returns a boolean if a field has been set.
func (o *RestoreRequest) HasSelectParameters() bool {
	if o != nil && !IsNil(o.SelectParameters) {
		return true
	}

	return false
}

// SetSelectParameters gets a reference to the given RestoreRequestSelectParameters and assigns it to the SelectParameters field.
func (o *RestoreRequest) SetSelectParameters(v RestoreRequestSelectParameters) {
	o.SelectParameters = &v
}

// GetOutputLocation returns the OutputLocation field value if set, zero value otherwise.
func (o *RestoreRequest) GetOutputLocation() RestoreRequestOutputLocation {
	if o == nil || IsNil(o.OutputLocation) {
		var ret RestoreRequestOutputLocation
		return ret
	}
	return *o.OutputLocation
}

// GetOutputLocationOk returns a tuple with the OutputLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequest) GetOutputLocationOk() (*RestoreRequestOutputLocation, bool) {
	if o == nil || IsNil(o.OutputLocation) {
		return nil, false
	}
	return o.OutputLocation, true
}

// HasOutputLocation returns a boolean if a field has been set.
func (o *RestoreRequest) HasOutputLocation() bool {
	if o != nil && !IsNil(o.OutputLocation) {
		return true
	}

	return false
}

// SetOutputLocation gets a reference to the given RestoreRequestOutputLocation and assigns it to the OutputLocation field.
func (o *RestoreRequest) SetOutputLocation(v RestoreRequestOutputLocation) {
	o.OutputLocation = &v
}

func (o RestoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Days) {
		toSerialize["Days"] = o.Days
	}
	if !IsNil(o.GlacierJobParameters) {
		toSerialize["GlacierJobParameters"] = o.GlacierJobParameters
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Tier) {
		toSerialize["Tier"] = o.Tier
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.SelectParameters) {
		toSerialize["SelectParameters"] = o.SelectParameters
	}
	if !IsNil(o.OutputLocation) {
		toSerialize["OutputLocation"] = o.OutputLocation
	}
	return toSerialize, nil
}

type NullableRestoreRequest struct {
	value *RestoreRequest
	isSet bool
}

func (v NullableRestoreRequest) Get() *RestoreRequest {
	return v.value
}

func (v *NullableRestoreRequest) Set(val *RestoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreRequest(val *RestoreRequest) *NullableRestoreRequest {
	return &NullableRestoreRequest{value: val, isSet: true}
}

func (v NullableRestoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


