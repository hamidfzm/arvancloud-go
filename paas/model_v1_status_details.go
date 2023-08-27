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

// checks if the V1StatusDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1StatusDetails{}

// V1StatusDetails StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
type V1StatusDetails struct {
	// The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Causes []V1StatusCause `json:"causes,omitempty"`
	// The group attribute of the resource associated with the status StatusReason.
	Group *string `json:"group,omitempty"`
	// The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Name *string `json:"name,omitempty"`
	// If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	RetryAfterSeconds *int32 `json:"retryAfterSeconds,omitempty"`
	// UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	Uid *string `json:"uid,omitempty"`
}

// NewV1StatusDetails instantiates a new V1StatusDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1StatusDetails() *V1StatusDetails {
	this := V1StatusDetails{}
	return &this
}

// NewV1StatusDetailsWithDefaults instantiates a new V1StatusDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1StatusDetailsWithDefaults() *V1StatusDetails {
	this := V1StatusDetails{}
	return &this
}

// GetCauses returns the Causes field value if set, zero value otherwise.
func (o *V1StatusDetails) GetCauses() []V1StatusCause {
	if o == nil || IsNil(o.Causes) {
		var ret []V1StatusCause
		return ret
	}
	return o.Causes
}

// GetCausesOk returns a tuple with the Causes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatusDetails) GetCausesOk() ([]V1StatusCause, bool) {
	if o == nil || IsNil(o.Causes) {
		return nil, false
	}
	return o.Causes, true
}

// HasCauses returns a boolean if a field has been set.
func (o *V1StatusDetails) HasCauses() bool {
	if o != nil && !IsNil(o.Causes) {
		return true
	}

	return false
}

// SetCauses gets a reference to the given []V1StatusCause and assigns it to the Causes field.
func (o *V1StatusDetails) SetCauses(v []V1StatusCause) {
	o.Causes = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1StatusDetails) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatusDetails) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1StatusDetails) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1StatusDetails) SetGroup(v string) {
	o.Group = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1StatusDetails) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatusDetails) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1StatusDetails) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1StatusDetails) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1StatusDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatusDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1StatusDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1StatusDetails) SetName(v string) {
	o.Name = &v
}

// GetRetryAfterSeconds returns the RetryAfterSeconds field value if set, zero value otherwise.
func (o *V1StatusDetails) GetRetryAfterSeconds() int32 {
	if o == nil || IsNil(o.RetryAfterSeconds) {
		var ret int32
		return ret
	}
	return *o.RetryAfterSeconds
}

// GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatusDetails) GetRetryAfterSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.RetryAfterSeconds) {
		return nil, false
	}
	return o.RetryAfterSeconds, true
}

// HasRetryAfterSeconds returns a boolean if a field has been set.
func (o *V1StatusDetails) HasRetryAfterSeconds() bool {
	if o != nil && !IsNil(o.RetryAfterSeconds) {
		return true
	}

	return false
}

// SetRetryAfterSeconds gets a reference to the given int32 and assigns it to the RetryAfterSeconds field.
func (o *V1StatusDetails) SetRetryAfterSeconds(v int32) {
	o.RetryAfterSeconds = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *V1StatusDetails) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatusDetails) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *V1StatusDetails) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *V1StatusDetails) SetUid(v string) {
	o.Uid = &v
}

func (o V1StatusDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1StatusDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Causes) {
		toSerialize["causes"] = o.Causes
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RetryAfterSeconds) {
		toSerialize["retryAfterSeconds"] = o.RetryAfterSeconds
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableV1StatusDetails struct {
	value *V1StatusDetails
	isSet bool
}

func (v NullableV1StatusDetails) Get() *V1StatusDetails {
	return v.value
}

func (v *NullableV1StatusDetails) Set(val *V1StatusDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StatusDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StatusDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StatusDetails(val *V1StatusDetails) *NullableV1StatusDetails {
	return &NullableV1StatusDetails{value: val, isSet: true}
}

func (v NullableV1StatusDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StatusDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


