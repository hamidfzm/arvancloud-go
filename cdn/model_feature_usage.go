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

// checks if the FeatureUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureUsage{}

// FeatureUsage struct for FeatureUsage
type FeatureUsage struct {
	FeatureId *string `json:"feature_id,omitempty"`
	Pricing NullableFeaturePricing `json:"pricing,omitempty"`
	EstimatedCost *EstimatedCost `json:"estimated_cost,omitempty"`
	Usage *float32 `json:"usage,omitempty"`
}

// NewFeatureUsage instantiates a new FeatureUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureUsage() *FeatureUsage {
	this := FeatureUsage{}
	return &this
}

// NewFeatureUsageWithDefaults instantiates a new FeatureUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureUsageWithDefaults() *FeatureUsage {
	this := FeatureUsage{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *FeatureUsage) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureUsage) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *FeatureUsage) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *FeatureUsage) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureUsage) GetPricing() FeaturePricing {
	if o == nil || IsNil(o.Pricing.Get()) {
		var ret FeaturePricing
		return ret
	}
	return *o.Pricing.Get()
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureUsage) GetPricingOk() (*FeaturePricing, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pricing.Get(), o.Pricing.IsSet()
}

// HasPricing returns a boolean if a field has been set.
func (o *FeatureUsage) HasPricing() bool {
	if o != nil && o.Pricing.IsSet() {
		return true
	}

	return false
}

// SetPricing gets a reference to the given NullableFeaturePricing and assigns it to the Pricing field.
func (o *FeatureUsage) SetPricing(v FeaturePricing) {
	o.Pricing.Set(&v)
}
// SetPricingNil sets the value for Pricing to be an explicit nil
func (o *FeatureUsage) SetPricingNil() {
	o.Pricing.Set(nil)
}

// UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
func (o *FeatureUsage) UnsetPricing() {
	o.Pricing.Unset()
}

// GetEstimatedCost returns the EstimatedCost field value if set, zero value otherwise.
func (o *FeatureUsage) GetEstimatedCost() EstimatedCost {
	if o == nil || IsNil(o.EstimatedCost) {
		var ret EstimatedCost
		return ret
	}
	return *o.EstimatedCost
}

// GetEstimatedCostOk returns a tuple with the EstimatedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureUsage) GetEstimatedCostOk() (*EstimatedCost, bool) {
	if o == nil || IsNil(o.EstimatedCost) {
		return nil, false
	}
	return o.EstimatedCost, true
}

// HasEstimatedCost returns a boolean if a field has been set.
func (o *FeatureUsage) HasEstimatedCost() bool {
	if o != nil && !IsNil(o.EstimatedCost) {
		return true
	}

	return false
}

// SetEstimatedCost gets a reference to the given EstimatedCost and assigns it to the EstimatedCost field.
func (o *FeatureUsage) SetEstimatedCost(v EstimatedCost) {
	o.EstimatedCost = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *FeatureUsage) GetUsage() float32 {
	if o == nil || IsNil(o.Usage) {
		var ret float32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureUsage) GetUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *FeatureUsage) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given float32 and assigns it to the Usage field.
func (o *FeatureUsage) SetUsage(v float32) {
	o.Usage = &v
}

func (o FeatureUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if o.Pricing.IsSet() {
		toSerialize["pricing"] = o.Pricing.Get()
	}
	if !IsNil(o.EstimatedCost) {
		toSerialize["estimated_cost"] = o.EstimatedCost
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableFeatureUsage struct {
	value *FeatureUsage
	isSet bool
}

func (v NullableFeatureUsage) Get() *FeatureUsage {
	return v.value
}

func (v *NullableFeatureUsage) Set(val *FeatureUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureUsage(val *FeatureUsage) *NullableFeatureUsage {
	return &NullableFeatureUsage{value: val, isSet: true}
}

func (v NullableFeatureUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


