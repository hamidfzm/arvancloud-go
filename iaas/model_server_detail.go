/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ServerDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerDetail{}

// ServerDetail struct for ServerDetail
type ServerDetail struct {
	Addresses *map[string][]ServerAddress `json:"addresses,omitempty"`
	ArNext *string `json:"ar_next,omitempty"`
	Created *string `json:"created,omitempty"`
	Flavor *ServerFlavor `json:"flavor,omitempty"`
	HaEnabled *bool `json:"ha_enabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Image *ServerImage `json:"image,omitempty"`
	KeyName *string `json:"key_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
	TaskState *string `json:"task_state,omitempty"`
}

// NewServerDetail instantiates a new ServerDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerDetail() *ServerDetail {
	this := ServerDetail{}
	return &this
}

// NewServerDetailWithDefaults instantiates a new ServerDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerDetailWithDefaults() *ServerDetail {
	this := ServerDetail{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ServerDetail) GetAddresses() map[string][]ServerAddress {
	if o == nil || IsNil(o.Addresses) {
		var ret map[string][]ServerAddress
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetAddressesOk() (*map[string][]ServerAddress, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ServerDetail) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given map[string][]ServerAddress and assigns it to the Addresses field.
func (o *ServerDetail) SetAddresses(v map[string][]ServerAddress) {
	o.Addresses = &v
}

// GetArNext returns the ArNext field value if set, zero value otherwise.
func (o *ServerDetail) GetArNext() string {
	if o == nil || IsNil(o.ArNext) {
		var ret string
		return ret
	}
	return *o.ArNext
}

// GetArNextOk returns a tuple with the ArNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetArNextOk() (*string, bool) {
	if o == nil || IsNil(o.ArNext) {
		return nil, false
	}
	return o.ArNext, true
}

// HasArNext returns a boolean if a field has been set.
func (o *ServerDetail) HasArNext() bool {
	if o != nil && !IsNil(o.ArNext) {
		return true
	}

	return false
}

// SetArNext gets a reference to the given string and assigns it to the ArNext field.
func (o *ServerDetail) SetArNext(v string) {
	o.ArNext = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServerDetail) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServerDetail) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ServerDetail) SetCreated(v string) {
	o.Created = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise.
func (o *ServerDetail) GetFlavor() ServerFlavor {
	if o == nil || IsNil(o.Flavor) {
		var ret ServerFlavor
		return ret
	}
	return *o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetFlavorOk() (*ServerFlavor, bool) {
	if o == nil || IsNil(o.Flavor) {
		return nil, false
	}
	return o.Flavor, true
}

// HasFlavor returns a boolean if a field has been set.
func (o *ServerDetail) HasFlavor() bool {
	if o != nil && !IsNil(o.Flavor) {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given ServerFlavor and assigns it to the Flavor field.
func (o *ServerDetail) SetFlavor(v ServerFlavor) {
	o.Flavor = &v
}

// GetHaEnabled returns the HaEnabled field value if set, zero value otherwise.
func (o *ServerDetail) GetHaEnabled() bool {
	if o == nil || IsNil(o.HaEnabled) {
		var ret bool
		return ret
	}
	return *o.HaEnabled
}

// GetHaEnabledOk returns a tuple with the HaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetHaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HaEnabled) {
		return nil, false
	}
	return o.HaEnabled, true
}

// HasHaEnabled returns a boolean if a field has been set.
func (o *ServerDetail) HasHaEnabled() bool {
	if o != nil && !IsNil(o.HaEnabled) {
		return true
	}

	return false
}

// SetHaEnabled gets a reference to the given bool and assigns it to the HaEnabled field.
func (o *ServerDetail) SetHaEnabled(v bool) {
	o.HaEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerDetail) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerDetail) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ServerDetail) GetImage() ServerImage {
	if o == nil || IsNil(o.Image) {
		var ret ServerImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetImageOk() (*ServerImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ServerDetail) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ServerImage and assigns it to the Image field.
func (o *ServerDetail) SetImage(v ServerImage) {
	o.Image = &v
}

// GetKeyName returns the KeyName field value if set, zero value otherwise.
func (o *ServerDetail) GetKeyName() string {
	if o == nil || IsNil(o.KeyName) {
		var ret string
		return ret
	}
	return *o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.KeyName) {
		return nil, false
	}
	return o.KeyName, true
}

// HasKeyName returns a boolean if a field has been set.
func (o *ServerDetail) HasKeyName() bool {
	if o != nil && !IsNil(o.KeyName) {
		return true
	}

	return false
}

// SetKeyName gets a reference to the given string and assigns it to the KeyName field.
func (o *ServerDetail) SetKeyName(v string) {
	o.KeyName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerDetail) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ServerDetail) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ServerDetail) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ServerDetail) SetPassword(v string) {
	o.Password = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *ServerDetail) GetSecurityGroups() []SecurityGroup {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret []SecurityGroup
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetSecurityGroupsOk() ([]SecurityGroup, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *ServerDetail) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []SecurityGroup and assigns it to the SecurityGroups field.
func (o *ServerDetail) SetSecurityGroups(v []SecurityGroup) {
	o.SecurityGroups = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServerDetail) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServerDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServerDetail) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServerDetail) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServerDetail) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *ServerDetail) SetTags(v []Tag) {
	o.Tags = v
}

// GetTaskState returns the TaskState field value if set, zero value otherwise.
func (o *ServerDetail) GetTaskState() string {
	if o == nil || IsNil(o.TaskState) {
		var ret string
		return ret
	}
	return *o.TaskState
}

// GetTaskStateOk returns a tuple with the TaskState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerDetail) GetTaskStateOk() (*string, bool) {
	if o == nil || IsNil(o.TaskState) {
		return nil, false
	}
	return o.TaskState, true
}

// HasTaskState returns a boolean if a field has been set.
func (o *ServerDetail) HasTaskState() bool {
	if o != nil && !IsNil(o.TaskState) {
		return true
	}

	return false
}

// SetTaskState gets a reference to the given string and assigns it to the TaskState field.
func (o *ServerDetail) SetTaskState(v string) {
	o.TaskState = &v
}

func (o ServerDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.ArNext) {
		toSerialize["ar_next"] = o.ArNext
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Flavor) {
		toSerialize["flavor"] = o.Flavor
	}
	if !IsNil(o.HaEnabled) {
		toSerialize["ha_enabled"] = o.HaEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.KeyName) {
		toSerialize["key_name"] = o.KeyName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["security_groups"] = o.SecurityGroups
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TaskState) {
		toSerialize["task_state"] = o.TaskState
	}
	return toSerialize, nil
}

type NullableServerDetail struct {
	value *ServerDetail
	isSet bool
}

func (v NullableServerDetail) Get() *ServerDetail {
	return v.value
}

func (v *NullableServerDetail) Set(val *ServerDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableServerDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableServerDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerDetail(val *ServerDetail) *NullableServerDetail {
	return &NullableServerDetail{value: val, isSet: true}
}

func (v NullableServerDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


