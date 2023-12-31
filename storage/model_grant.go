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

// checks if the Grant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Grant{}

// Grant Container for grant information.
type Grant struct {
	Grantee *GrantGrantee `json:"Grantee,omitempty"`
	Permission *Permission `json:"Permission,omitempty"`
}

// NewGrant instantiates a new Grant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrant() *Grant {
	this := Grant{}
	return &this
}

// NewGrantWithDefaults instantiates a new Grant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantWithDefaults() *Grant {
	this := Grant{}
	return &this
}

// GetGrantee returns the Grantee field value if set, zero value otherwise.
func (o *Grant) GetGrantee() GrantGrantee {
	if o == nil || IsNil(o.Grantee) {
		var ret GrantGrantee
		return ret
	}
	return *o.Grantee
}

// GetGranteeOk returns a tuple with the Grantee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetGranteeOk() (*GrantGrantee, bool) {
	if o == nil || IsNil(o.Grantee) {
		return nil, false
	}
	return o.Grantee, true
}

// HasGrantee returns a boolean if a field has been set.
func (o *Grant) HasGrantee() bool {
	if o != nil && !IsNil(o.Grantee) {
		return true
	}

	return false
}

// SetGrantee gets a reference to the given GrantGrantee and assigns it to the Grantee field.
func (o *Grant) SetGrantee(v GrantGrantee) {
	o.Grantee = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *Grant) GetPermission() Permission {
	if o == nil || IsNil(o.Permission) {
		var ret Permission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetPermissionOk() (*Permission, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *Grant) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given Permission and assigns it to the Permission field.
func (o *Grant) SetPermission(v Permission) {
	o.Permission = &v
}

func (o Grant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Grant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grantee) {
		toSerialize["Grantee"] = o.Grantee
	}
	if !IsNil(o.Permission) {
		toSerialize["Permission"] = o.Permission
	}
	return toSerialize, nil
}

type NullableGrant struct {
	value *Grant
	isSet bool
}

func (v NullableGrant) Get() *Grant {
	return v.value
}

func (v *NullableGrant) Set(val *Grant) {
	v.value = val
	v.isSet = true
}

func (v NullableGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrant(val *Grant) *NullableGrant {
	return &NullableGrant{value: val, isSet: true}
}

func (v NullableGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


