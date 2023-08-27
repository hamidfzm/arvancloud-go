/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
	"time"
)

// checks if the LoadBalancerOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerOrigin{}

// LoadBalancerOrigin struct for LoadBalancerOrigin
type LoadBalancerOrigin struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *bool `json:"status,omitempty"`
	Address *string `json:"address,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	HostHeader *string `json:"host_header,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewLoadBalancerOrigin instantiates a new LoadBalancerOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerOrigin() *LoadBalancerOrigin {
	this := LoadBalancerOrigin{}
	var protocol string = "auto"
	this.Protocol = &protocol
	return &this
}

// NewLoadBalancerOriginWithDefaults instantiates a new LoadBalancerOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerOriginWithDefaults() *LoadBalancerOrigin {
	this := LoadBalancerOrigin{}
	var protocol string = "auto"
	this.Protocol = &protocol
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadBalancerOrigin) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerOrigin) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LoadBalancerOrigin) SetStatus(v bool) {
	o.Status = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *LoadBalancerOrigin) SetAddress(v string) {
	o.Address = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *LoadBalancerOrigin) SetPort(v int32) {
	o.Port = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *LoadBalancerOrigin) SetWeight(v int32) {
	o.Weight = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *LoadBalancerOrigin) SetProtocol(v string) {
	o.Protocol = &v
}

// GetHostHeader returns the HostHeader field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetHostHeader() string {
	if o == nil || IsNil(o.HostHeader) {
		var ret string
		return ret
	}
	return *o.HostHeader
}

// GetHostHeaderOk returns a tuple with the HostHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetHostHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.HostHeader) {
		return nil, false
	}
	return o.HostHeader, true
}

// HasHostHeader returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasHostHeader() bool {
	if o != nil && !IsNil(o.HostHeader) {
		return true
	}

	return false
}

// SetHostHeader gets a reference to the given string and assigns it to the HostHeader field.
func (o *LoadBalancerOrigin) SetHostHeader(v string) {
	o.HostHeader = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LoadBalancerOrigin) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LoadBalancerOrigin) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOrigin) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoadBalancerOrigin) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LoadBalancerOrigin) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o LoadBalancerOrigin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.HostHeader) {
		toSerialize["host_header"] = o.HostHeader
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableLoadBalancerOrigin struct {
	value *LoadBalancerOrigin
	isSet bool
}

func (v NullableLoadBalancerOrigin) Get() *LoadBalancerOrigin {
	return v.value
}

func (v *NullableLoadBalancerOrigin) Set(val *LoadBalancerOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerOrigin(val *LoadBalancerOrigin) *NullableLoadBalancerOrigin {
	return &NullableLoadBalancerOrigin{value: val, isSet: true}
}

func (v NullableLoadBalancerOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


