/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the GroupInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInstance{}

// GroupInstance struct for GroupInstance
type GroupInstance struct {
	Id *string `json:"id,omitempty"`
	Ips []string `json:"ips,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewGroupInstance instantiates a new GroupInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInstance() *GroupInstance {
	this := GroupInstance{}
	return &this
}

// NewGroupInstanceWithDefaults instantiates a new GroupInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInstanceWithDefaults() *GroupInstance {
	this := GroupInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupInstance) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInstance) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupInstance) SetId(v string) {
	o.Id = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *GroupInstance) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInstance) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *GroupInstance) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *GroupInstance) SetIps(v []string) {
	o.Ips = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupInstance) SetName(v string) {
	o.Name = &v
}

func (o GroupInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGroupInstance struct {
	value *GroupInstance
	isSet bool
}

func (v NullableGroupInstance) Get() *GroupInstance {
	return v.value
}

func (v *NullableGroupInstance) Set(val *GroupInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInstance(val *GroupInstance) *NullableGroupInstance {
	return &NullableGroupInstance{value: val, isSet: true}
}

func (v NullableGroupInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


