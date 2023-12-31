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

// checks if the RoutingRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingRulesInner{}

// RoutingRulesInner struct for RoutingRulesInner
type RoutingRulesInner struct {
	Condition *RoutingRuleCondition `json:"Condition,omitempty"`
	Redirect RoutingRuleRedirect `json:"Redirect"`
}

// NewRoutingRulesInner instantiates a new RoutingRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingRulesInner(redirect RoutingRuleRedirect) *RoutingRulesInner {
	this := RoutingRulesInner{}
	this.Redirect = redirect
	return &this
}

// NewRoutingRulesInnerWithDefaults instantiates a new RoutingRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingRulesInnerWithDefaults() *RoutingRulesInner {
	this := RoutingRulesInner{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RoutingRulesInner) GetCondition() RoutingRuleCondition {
	if o == nil || IsNil(o.Condition) {
		var ret RoutingRuleCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRulesInner) GetConditionOk() (*RoutingRuleCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RoutingRulesInner) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given RoutingRuleCondition and assigns it to the Condition field.
func (o *RoutingRulesInner) SetCondition(v RoutingRuleCondition) {
	o.Condition = &v
}

// GetRedirect returns the Redirect field value
func (o *RoutingRulesInner) GetRedirect() RoutingRuleRedirect {
	if o == nil {
		var ret RoutingRuleRedirect
		return ret
	}

	return o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value
// and a boolean to check if the value has been set.
func (o *RoutingRulesInner) GetRedirectOk() (*RoutingRuleRedirect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redirect, true
}

// SetRedirect sets field value
func (o *RoutingRulesInner) SetRedirect(v RoutingRuleRedirect) {
	o.Redirect = v
}

func (o RoutingRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Condition) {
		toSerialize["Condition"] = o.Condition
	}
	toSerialize["Redirect"] = o.Redirect
	return toSerialize, nil
}

type NullableRoutingRulesInner struct {
	value *RoutingRulesInner
	isSet bool
}

func (v NullableRoutingRulesInner) Get() *RoutingRulesInner {
	return v.value
}

func (v *NullableRoutingRulesInner) Set(val *RoutingRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingRulesInner(val *RoutingRulesInner) *NullableRoutingRulesInner {
	return &NullableRoutingRulesInner{value: val, isSet: true}
}

func (v NullableRoutingRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


