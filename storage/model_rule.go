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

// checks if the Rule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rule{}

// Rule Specifies lifecycle rules for an ArvanCloud S3 bucket.
type Rule struct {
	Expiration *RuleExpiration `json:"Expiration,omitempty"`
	ID *string `json:"ID,omitempty"`
	Prefix string `json:"Prefix"`
	Status ExpirationStatus `json:"Status"`
	Transition *RuleTransition `json:"Transition,omitempty"`
	NoncurrentVersionTransition *NoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty"`
	NoncurrentVersionExpiration *NoncurrentVersionExpiration `json:"NoncurrentVersionExpiration,omitempty"`
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUpload `json:"AbortIncompleteMultipartUpload,omitempty"`
}

// NewRule instantiates a new Rule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRule(prefix string, status ExpirationStatus) *Rule {
	this := Rule{}
	this.Prefix = prefix
	this.Status = status
	return &this
}

// NewRuleWithDefaults instantiates a new Rule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleWithDefaults() *Rule {
	this := Rule{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *Rule) GetExpiration() RuleExpiration {
	if o == nil || IsNil(o.Expiration) {
		var ret RuleExpiration
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetExpirationOk() (*RuleExpiration, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *Rule) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given RuleExpiration and assigns it to the Expiration field.
func (o *Rule) SetExpiration(v RuleExpiration) {
	o.Expiration = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *Rule) GetID() string {
	if o == nil || IsNil(o.ID) {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetIDOk() (*string, bool) {
	if o == nil || IsNil(o.ID) {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *Rule) HasID() bool {
	if o != nil && !IsNil(o.ID) {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *Rule) SetID(v string) {
	o.ID = &v
}

// GetPrefix returns the Prefix field value
func (o *Rule) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *Rule) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *Rule) SetPrefix(v string) {
	o.Prefix = v
}

// GetStatus returns the Status field value
func (o *Rule) GetStatus() ExpirationStatus {
	if o == nil {
		var ret ExpirationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Rule) GetStatusOk() (*ExpirationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Rule) SetStatus(v ExpirationStatus) {
	o.Status = v
}

// GetTransition returns the Transition field value if set, zero value otherwise.
func (o *Rule) GetTransition() RuleTransition {
	if o == nil || IsNil(o.Transition) {
		var ret RuleTransition
		return ret
	}
	return *o.Transition
}

// GetTransitionOk returns a tuple with the Transition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetTransitionOk() (*RuleTransition, bool) {
	if o == nil || IsNil(o.Transition) {
		return nil, false
	}
	return o.Transition, true
}

// HasTransition returns a boolean if a field has been set.
func (o *Rule) HasTransition() bool {
	if o != nil && !IsNil(o.Transition) {
		return true
	}

	return false
}

// SetTransition gets a reference to the given RuleTransition and assigns it to the Transition field.
func (o *Rule) SetTransition(v RuleTransition) {
	o.Transition = &v
}

// GetNoncurrentVersionTransition returns the NoncurrentVersionTransition field value if set, zero value otherwise.
func (o *Rule) GetNoncurrentVersionTransition() NoncurrentVersionTransition {
	if o == nil || IsNil(o.NoncurrentVersionTransition) {
		var ret NoncurrentVersionTransition
		return ret
	}
	return *o.NoncurrentVersionTransition
}

// GetNoncurrentVersionTransitionOk returns a tuple with the NoncurrentVersionTransition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetNoncurrentVersionTransitionOk() (*NoncurrentVersionTransition, bool) {
	if o == nil || IsNil(o.NoncurrentVersionTransition) {
		return nil, false
	}
	return o.NoncurrentVersionTransition, true
}

// HasNoncurrentVersionTransition returns a boolean if a field has been set.
func (o *Rule) HasNoncurrentVersionTransition() bool {
	if o != nil && !IsNil(o.NoncurrentVersionTransition) {
		return true
	}

	return false
}

// SetNoncurrentVersionTransition gets a reference to the given NoncurrentVersionTransition and assigns it to the NoncurrentVersionTransition field.
func (o *Rule) SetNoncurrentVersionTransition(v NoncurrentVersionTransition) {
	o.NoncurrentVersionTransition = &v
}

// GetNoncurrentVersionExpiration returns the NoncurrentVersionExpiration field value if set, zero value otherwise.
func (o *Rule) GetNoncurrentVersionExpiration() NoncurrentVersionExpiration {
	if o == nil || IsNil(o.NoncurrentVersionExpiration) {
		var ret NoncurrentVersionExpiration
		return ret
	}
	return *o.NoncurrentVersionExpiration
}

// GetNoncurrentVersionExpirationOk returns a tuple with the NoncurrentVersionExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetNoncurrentVersionExpirationOk() (*NoncurrentVersionExpiration, bool) {
	if o == nil || IsNil(o.NoncurrentVersionExpiration) {
		return nil, false
	}
	return o.NoncurrentVersionExpiration, true
}

// HasNoncurrentVersionExpiration returns a boolean if a field has been set.
func (o *Rule) HasNoncurrentVersionExpiration() bool {
	if o != nil && !IsNil(o.NoncurrentVersionExpiration) {
		return true
	}

	return false
}

// SetNoncurrentVersionExpiration gets a reference to the given NoncurrentVersionExpiration and assigns it to the NoncurrentVersionExpiration field.
func (o *Rule) SetNoncurrentVersionExpiration(v NoncurrentVersionExpiration) {
	o.NoncurrentVersionExpiration = &v
}

// GetAbortIncompleteMultipartUpload returns the AbortIncompleteMultipartUpload field value if set, zero value otherwise.
func (o *Rule) GetAbortIncompleteMultipartUpload() AbortIncompleteMultipartUpload {
	if o == nil || IsNil(o.AbortIncompleteMultipartUpload) {
		var ret AbortIncompleteMultipartUpload
		return ret
	}
	return *o.AbortIncompleteMultipartUpload
}

// GetAbortIncompleteMultipartUploadOk returns a tuple with the AbortIncompleteMultipartUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetAbortIncompleteMultipartUploadOk() (*AbortIncompleteMultipartUpload, bool) {
	if o == nil || IsNil(o.AbortIncompleteMultipartUpload) {
		return nil, false
	}
	return o.AbortIncompleteMultipartUpload, true
}

// HasAbortIncompleteMultipartUpload returns a boolean if a field has been set.
func (o *Rule) HasAbortIncompleteMultipartUpload() bool {
	if o != nil && !IsNil(o.AbortIncompleteMultipartUpload) {
		return true
	}

	return false
}

// SetAbortIncompleteMultipartUpload gets a reference to the given AbortIncompleteMultipartUpload and assigns it to the AbortIncompleteMultipartUpload field.
func (o *Rule) SetAbortIncompleteMultipartUpload(v AbortIncompleteMultipartUpload) {
	o.AbortIncompleteMultipartUpload = &v
}

func (o Rule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["Expiration"] = o.Expiration
	}
	if !IsNil(o.ID) {
		toSerialize["ID"] = o.ID
	}
	toSerialize["Prefix"] = o.Prefix
	toSerialize["Status"] = o.Status
	if !IsNil(o.Transition) {
		toSerialize["Transition"] = o.Transition
	}
	if !IsNil(o.NoncurrentVersionTransition) {
		toSerialize["NoncurrentVersionTransition"] = o.NoncurrentVersionTransition
	}
	if !IsNil(o.NoncurrentVersionExpiration) {
		toSerialize["NoncurrentVersionExpiration"] = o.NoncurrentVersionExpiration
	}
	if !IsNil(o.AbortIncompleteMultipartUpload) {
		toSerialize["AbortIncompleteMultipartUpload"] = o.AbortIncompleteMultipartUpload
	}
	return toSerialize, nil
}

type NullableRule struct {
	value *Rule
	isSet bool
}

func (v NullableRule) Get() *Rule {
	return v.value
}

func (v *NullableRule) Set(val *Rule) {
	v.value = val
	v.isSet = true
}

func (v NullableRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRule(val *Rule) *NullableRule {
	return &NullableRule{value: val, isSet: true}
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


