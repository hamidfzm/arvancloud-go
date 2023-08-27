/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1RecreateDeploymentStrategyParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RecreateDeploymentStrategyParams{}

// V1RecreateDeploymentStrategyParams RecreateDeploymentStrategyParams are the input to the Recreate deployment strategy.
type V1RecreateDeploymentStrategyParams struct {
	Mid *V1LifecycleHook `json:"mid,omitempty"`
	Post *V1LifecycleHook `json:"post,omitempty"`
	Pre *V1LifecycleHook `json:"pre,omitempty"`
	// TimeoutSeconds is the time to wait for updates before giving up. If the value is nil, a default will be used.
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
}

// NewV1RecreateDeploymentStrategyParams instantiates a new V1RecreateDeploymentStrategyParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RecreateDeploymentStrategyParams() *V1RecreateDeploymentStrategyParams {
	this := V1RecreateDeploymentStrategyParams{}
	return &this
}

// NewV1RecreateDeploymentStrategyParamsWithDefaults instantiates a new V1RecreateDeploymentStrategyParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RecreateDeploymentStrategyParamsWithDefaults() *V1RecreateDeploymentStrategyParams {
	this := V1RecreateDeploymentStrategyParams{}
	return &this
}

// GetMid returns the Mid field value if set, zero value otherwise.
func (o *V1RecreateDeploymentStrategyParams) GetMid() V1LifecycleHook {
	if o == nil || IsNil(o.Mid) {
		var ret V1LifecycleHook
		return ret
	}
	return *o.Mid
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RecreateDeploymentStrategyParams) GetMidOk() (*V1LifecycleHook, bool) {
	if o == nil || IsNil(o.Mid) {
		return nil, false
	}
	return o.Mid, true
}

// HasMid returns a boolean if a field has been set.
func (o *V1RecreateDeploymentStrategyParams) HasMid() bool {
	if o != nil && !IsNil(o.Mid) {
		return true
	}

	return false
}

// SetMid gets a reference to the given V1LifecycleHook and assigns it to the Mid field.
func (o *V1RecreateDeploymentStrategyParams) SetMid(v V1LifecycleHook) {
	o.Mid = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *V1RecreateDeploymentStrategyParams) GetPost() V1LifecycleHook {
	if o == nil || IsNil(o.Post) {
		var ret V1LifecycleHook
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RecreateDeploymentStrategyParams) GetPostOk() (*V1LifecycleHook, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *V1RecreateDeploymentStrategyParams) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given V1LifecycleHook and assigns it to the Post field.
func (o *V1RecreateDeploymentStrategyParams) SetPost(v V1LifecycleHook) {
	o.Post = &v
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *V1RecreateDeploymentStrategyParams) GetPre() V1LifecycleHook {
	if o == nil || IsNil(o.Pre) {
		var ret V1LifecycleHook
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RecreateDeploymentStrategyParams) GetPreOk() (*V1LifecycleHook, bool) {
	if o == nil || IsNil(o.Pre) {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *V1RecreateDeploymentStrategyParams) HasPre() bool {
	if o != nil && !IsNil(o.Pre) {
		return true
	}

	return false
}

// SetPre gets a reference to the given V1LifecycleHook and assigns it to the Pre field.
func (o *V1RecreateDeploymentStrategyParams) SetPre(v V1LifecycleHook) {
	o.Pre = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *V1RecreateDeploymentStrategyParams) GetTimeoutSeconds() int64 {
	if o == nil || IsNil(o.TimeoutSeconds) {
		var ret int64
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RecreateDeploymentStrategyParams) GetTimeoutSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeoutSeconds) {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *V1RecreateDeploymentStrategyParams) HasTimeoutSeconds() bool {
	if o != nil && !IsNil(o.TimeoutSeconds) {
		return true
	}

	return false
}

// SetTimeoutSeconds gets a reference to the given int64 and assigns it to the TimeoutSeconds field.
func (o *V1RecreateDeploymentStrategyParams) SetTimeoutSeconds(v int64) {
	o.TimeoutSeconds = &v
}

func (o V1RecreateDeploymentStrategyParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RecreateDeploymentStrategyParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mid) {
		toSerialize["mid"] = o.Mid
	}
	if !IsNil(o.Post) {
		toSerialize["post"] = o.Post
	}
	if !IsNil(o.Pre) {
		toSerialize["pre"] = o.Pre
	}
	if !IsNil(o.TimeoutSeconds) {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	return toSerialize, nil
}

type NullableV1RecreateDeploymentStrategyParams struct {
	value *V1RecreateDeploymentStrategyParams
	isSet bool
}

func (v NullableV1RecreateDeploymentStrategyParams) Get() *V1RecreateDeploymentStrategyParams {
	return v.value
}

func (v *NullableV1RecreateDeploymentStrategyParams) Set(val *V1RecreateDeploymentStrategyParams) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RecreateDeploymentStrategyParams) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RecreateDeploymentStrategyParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RecreateDeploymentStrategyParams(val *V1RecreateDeploymentStrategyParams) *NullableV1RecreateDeploymentStrategyParams {
	return &NullableV1RecreateDeploymentStrategyParams{value: val, isSet: true}
}

func (v NullableV1RecreateDeploymentStrategyParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RecreateDeploymentStrategyParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


