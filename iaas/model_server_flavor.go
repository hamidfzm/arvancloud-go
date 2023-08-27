/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ServerFlavor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerFlavor{}

// ServerFlavor struct for ServerFlavor
type ServerFlavor struct {
	Disk *int32 `json:"disk,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Ram *float64 `json:"ram,omitempty"`
	Swap *string `json:"swap,omitempty"`
	Vcpus *int32 `json:"vcpus,omitempty"`
}

// NewServerFlavor instantiates a new ServerFlavor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerFlavor() *ServerFlavor {
	this := ServerFlavor{}
	return &this
}

// NewServerFlavorWithDefaults instantiates a new ServerFlavor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerFlavorWithDefaults() *ServerFlavor {
	this := ServerFlavor{}
	return &this
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *ServerFlavor) GetDisk() int32 {
	if o == nil || IsNil(o.Disk) {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFlavor) GetDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *ServerFlavor) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *ServerFlavor) SetDisk(v int32) {
	o.Disk = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerFlavor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFlavor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerFlavor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerFlavor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerFlavor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFlavor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerFlavor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerFlavor) SetName(v string) {
	o.Name = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *ServerFlavor) GetRam() float64 {
	if o == nil || IsNil(o.Ram) {
		var ret float64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFlavor) GetRamOk() (*float64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *ServerFlavor) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given float64 and assigns it to the Ram field.
func (o *ServerFlavor) SetRam(v float64) {
	o.Ram = &v
}

// GetSwap returns the Swap field value if set, zero value otherwise.
func (o *ServerFlavor) GetSwap() string {
	if o == nil || IsNil(o.Swap) {
		var ret string
		return ret
	}
	return *o.Swap
}

// GetSwapOk returns a tuple with the Swap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFlavor) GetSwapOk() (*string, bool) {
	if o == nil || IsNil(o.Swap) {
		return nil, false
	}
	return o.Swap, true
}

// HasSwap returns a boolean if a field has been set.
func (o *ServerFlavor) HasSwap() bool {
	if o != nil && !IsNil(o.Swap) {
		return true
	}

	return false
}

// SetSwap gets a reference to the given string and assigns it to the Swap field.
func (o *ServerFlavor) SetSwap(v string) {
	o.Swap = &v
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise.
func (o *ServerFlavor) GetVcpus() int32 {
	if o == nil || IsNil(o.Vcpus) {
		var ret int32
		return ret
	}
	return *o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerFlavor) GetVcpusOk() (*int32, bool) {
	if o == nil || IsNil(o.Vcpus) {
		return nil, false
	}
	return o.Vcpus, true
}

// HasVcpus returns a boolean if a field has been set.
func (o *ServerFlavor) HasVcpus() bool {
	if o != nil && !IsNil(o.Vcpus) {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given int32 and assigns it to the Vcpus field.
func (o *ServerFlavor) SetVcpus(v int32) {
	o.Vcpus = &v
}

func (o ServerFlavor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerFlavor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Swap) {
		toSerialize["swap"] = o.Swap
	}
	if !IsNil(o.Vcpus) {
		toSerialize["vcpus"] = o.Vcpus
	}
	return toSerialize, nil
}

type NullableServerFlavor struct {
	value *ServerFlavor
	isSet bool
}

func (v NullableServerFlavor) Get() *ServerFlavor {
	return v.value
}

func (v *NullableServerFlavor) Set(val *ServerFlavor) {
	v.value = val
	v.isSet = true
}

func (v NullableServerFlavor) IsSet() bool {
	return v.isSet
}

func (v *NullableServerFlavor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerFlavor(val *ServerFlavor) *NullableServerFlavor {
	return &NullableServerFlavor{value: val, isSet: true}
}

func (v NullableServerFlavor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerFlavor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


