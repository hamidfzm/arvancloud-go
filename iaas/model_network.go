/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network struct for Network
type Network struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	AvailabilityZoneHints []string `json:"availability_zone_hints,omitempty"`
	AvailabilityZones []string `json:"availability_zones,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	DhcpIp *string `json:"dhcp_ip,omitempty"`
	Id *string `json:"id,omitempty"`
	Ipv4AddressScope *string `json:"ipv4_address_scope,omitempty"`
	Ipv6AddressScope *string `json:"ipv6_address_scope,omitempty"`
	Mtu *int64 `json:"mtu,omitempty"`
	Name *string `json:"name,omitempty"`
	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty"`
	QosPolicyId *string `json:"qos_policy_id,omitempty"`
	RevisionNumber *int64 `json:"revision_number,omitempty"`
	Routerexternal *bool `json:"router:external,omitempty"`
	Shared *bool `json:"shared,omitempty"`
	Status *string `json:"status,omitempty"`
	Subnets []Subnet `json:"subnets,omitempty"`
	Tags []string `json:"tags,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetAdminStateUp returns the AdminStateUp field value if set, zero value otherwise.
func (o *Network) GetAdminStateUp() bool {
	if o == nil || IsNil(o.AdminStateUp) {
		var ret bool
		return ret
	}
	return *o.AdminStateUp
}

// GetAdminStateUpOk returns a tuple with the AdminStateUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetAdminStateUpOk() (*bool, bool) {
	if o == nil || IsNil(o.AdminStateUp) {
		return nil, false
	}
	return o.AdminStateUp, true
}

// HasAdminStateUp returns a boolean if a field has been set.
func (o *Network) HasAdminStateUp() bool {
	if o != nil && !IsNil(o.AdminStateUp) {
		return true
	}

	return false
}

// SetAdminStateUp gets a reference to the given bool and assigns it to the AdminStateUp field.
func (o *Network) SetAdminStateUp(v bool) {
	o.AdminStateUp = &v
}

// GetAvailabilityZoneHints returns the AvailabilityZoneHints field value if set, zero value otherwise.
func (o *Network) GetAvailabilityZoneHints() []string {
	if o == nil || IsNil(o.AvailabilityZoneHints) {
		var ret []string
		return ret
	}
	return o.AvailabilityZoneHints
}

// GetAvailabilityZoneHintsOk returns a tuple with the AvailabilityZoneHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetAvailabilityZoneHintsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailabilityZoneHints) {
		return nil, false
	}
	return o.AvailabilityZoneHints, true
}

// HasAvailabilityZoneHints returns a boolean if a field has been set.
func (o *Network) HasAvailabilityZoneHints() bool {
	if o != nil && !IsNil(o.AvailabilityZoneHints) {
		return true
	}

	return false
}

// SetAvailabilityZoneHints gets a reference to the given []string and assigns it to the AvailabilityZoneHints field.
func (o *Network) SetAvailabilityZoneHints(v []string) {
	o.AvailabilityZoneHints = v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *Network) GetAvailabilityZones() []string {
	if o == nil || IsNil(o.AvailabilityZones) {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetAvailabilityZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailabilityZones) {
		return nil, false
	}
	return o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *Network) HasAvailabilityZones() bool {
	if o != nil && !IsNil(o.AvailabilityZones) {
		return true
	}

	return false
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *Network) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Network) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Network) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Network) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Network) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Network) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Network) SetDescription(v string) {
	o.Description = &v
}

// GetDhcpIp returns the DhcpIp field value if set, zero value otherwise.
func (o *Network) GetDhcpIp() string {
	if o == nil || IsNil(o.DhcpIp) {
		var ret string
		return ret
	}
	return *o.DhcpIp
}

// GetDhcpIpOk returns a tuple with the DhcpIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetDhcpIpOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpIp) {
		return nil, false
	}
	return o.DhcpIp, true
}

// HasDhcpIp returns a boolean if a field has been set.
func (o *Network) HasDhcpIp() bool {
	if o != nil && !IsNil(o.DhcpIp) {
		return true
	}

	return false
}

// SetDhcpIp gets a reference to the given string and assigns it to the DhcpIp field.
func (o *Network) SetDhcpIp(v string) {
	o.DhcpIp = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Network) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Network) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Network) SetId(v string) {
	o.Id = &v
}

// GetIpv4AddressScope returns the Ipv4AddressScope field value if set, zero value otherwise.
func (o *Network) GetIpv4AddressScope() string {
	if o == nil || IsNil(o.Ipv4AddressScope) {
		var ret string
		return ret
	}
	return *o.Ipv4AddressScope
}

// GetIpv4AddressScopeOk returns a tuple with the Ipv4AddressScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIpv4AddressScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4AddressScope) {
		return nil, false
	}
	return o.Ipv4AddressScope, true
}

// HasIpv4AddressScope returns a boolean if a field has been set.
func (o *Network) HasIpv4AddressScope() bool {
	if o != nil && !IsNil(o.Ipv4AddressScope) {
		return true
	}

	return false
}

// SetIpv4AddressScope gets a reference to the given string and assigns it to the Ipv4AddressScope field.
func (o *Network) SetIpv4AddressScope(v string) {
	o.Ipv4AddressScope = &v
}

// GetIpv6AddressScope returns the Ipv6AddressScope field value if set, zero value otherwise.
func (o *Network) GetIpv6AddressScope() string {
	if o == nil || IsNil(o.Ipv6AddressScope) {
		var ret string
		return ret
	}
	return *o.Ipv6AddressScope
}

// GetIpv6AddressScopeOk returns a tuple with the Ipv6AddressScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIpv6AddressScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6AddressScope) {
		return nil, false
	}
	return o.Ipv6AddressScope, true
}

// HasIpv6AddressScope returns a boolean if a field has been set.
func (o *Network) HasIpv6AddressScope() bool {
	if o != nil && !IsNil(o.Ipv6AddressScope) {
		return true
	}

	return false
}

// SetIpv6AddressScope gets a reference to the given string and assigns it to the Ipv6AddressScope field.
func (o *Network) SetIpv6AddressScope(v string) {
	o.Ipv6AddressScope = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *Network) GetMtu() int64 {
	if o == nil || IsNil(o.Mtu) {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetMtuOk() (*int64, bool) {
	if o == nil || IsNil(o.Mtu) {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *Network) HasMtu() bool {
	if o != nil && !IsNil(o.Mtu) {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *Network) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Network) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Network) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Network) SetName(v string) {
	o.Name = &v
}

// GetPortSecurityEnabled returns the PortSecurityEnabled field value if set, zero value otherwise.
func (o *Network) GetPortSecurityEnabled() bool {
	if o == nil || IsNil(o.PortSecurityEnabled) {
		var ret bool
		return ret
	}
	return *o.PortSecurityEnabled
}

// GetPortSecurityEnabledOk returns a tuple with the PortSecurityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetPortSecurityEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PortSecurityEnabled) {
		return nil, false
	}
	return o.PortSecurityEnabled, true
}

// HasPortSecurityEnabled returns a boolean if a field has been set.
func (o *Network) HasPortSecurityEnabled() bool {
	if o != nil && !IsNil(o.PortSecurityEnabled) {
		return true
	}

	return false
}

// SetPortSecurityEnabled gets a reference to the given bool and assigns it to the PortSecurityEnabled field.
func (o *Network) SetPortSecurityEnabled(v bool) {
	o.PortSecurityEnabled = &v
}

// GetQosPolicyId returns the QosPolicyId field value if set, zero value otherwise.
func (o *Network) GetQosPolicyId() string {
	if o == nil || IsNil(o.QosPolicyId) {
		var ret string
		return ret
	}
	return *o.QosPolicyId
}

// GetQosPolicyIdOk returns a tuple with the QosPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetQosPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.QosPolicyId) {
		return nil, false
	}
	return o.QosPolicyId, true
}

// HasQosPolicyId returns a boolean if a field has been set.
func (o *Network) HasQosPolicyId() bool {
	if o != nil && !IsNil(o.QosPolicyId) {
		return true
	}

	return false
}

// SetQosPolicyId gets a reference to the given string and assigns it to the QosPolicyId field.
func (o *Network) SetQosPolicyId(v string) {
	o.QosPolicyId = &v
}

// GetRevisionNumber returns the RevisionNumber field value if set, zero value otherwise.
func (o *Network) GetRevisionNumber() int64 {
	if o == nil || IsNil(o.RevisionNumber) {
		var ret int64
		return ret
	}
	return *o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetRevisionNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.RevisionNumber) {
		return nil, false
	}
	return o.RevisionNumber, true
}

// HasRevisionNumber returns a boolean if a field has been set.
func (o *Network) HasRevisionNumber() bool {
	if o != nil && !IsNil(o.RevisionNumber) {
		return true
	}

	return false
}

// SetRevisionNumber gets a reference to the given int64 and assigns it to the RevisionNumber field.
func (o *Network) SetRevisionNumber(v int64) {
	o.RevisionNumber = &v
}

// GetRouterexternal returns the Routerexternal field value if set, zero value otherwise.
func (o *Network) GetRouterexternal() bool {
	if o == nil || IsNil(o.Routerexternal) {
		var ret bool
		return ret
	}
	return *o.Routerexternal
}

// GetRouterexternalOk returns a tuple with the Routerexternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetRouterexternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Routerexternal) {
		return nil, false
	}
	return o.Routerexternal, true
}

// HasRouterexternal returns a boolean if a field has been set.
func (o *Network) HasRouterexternal() bool {
	if o != nil && !IsNil(o.Routerexternal) {
		return true
	}

	return false
}

// SetRouterexternal gets a reference to the given bool and assigns it to the Routerexternal field.
func (o *Network) SetRouterexternal(v bool) {
	o.Routerexternal = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *Network) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *Network) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *Network) SetShared(v bool) {
	o.Shared = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Network) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Network) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Network) SetStatus(v string) {
	o.Status = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *Network) GetSubnets() []Subnet {
	if o == nil || IsNil(o.Subnets) {
		var ret []Subnet
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetSubnetsOk() ([]Subnet, bool) {
	if o == nil || IsNil(o.Subnets) {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *Network) HasSubnets() bool {
	if o != nil && !IsNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []Subnet and assigns it to the Subnets field.
func (o *Network) SetSubnets(v []Subnet) {
	o.Subnets = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Network) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Network) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Network) SetTags(v []string) {
	o.Tags = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *Network) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *Network) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *Network) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Network) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Network) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Network) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdminStateUp) {
		toSerialize["admin_state_up"] = o.AdminStateUp
	}
	if !IsNil(o.AvailabilityZoneHints) {
		toSerialize["availability_zone_hints"] = o.AvailabilityZoneHints
	}
	if !IsNil(o.AvailabilityZones) {
		toSerialize["availability_zones"] = o.AvailabilityZones
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DhcpIp) {
		toSerialize["dhcp_ip"] = o.DhcpIp
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ipv4AddressScope) {
		toSerialize["ipv4_address_scope"] = o.Ipv4AddressScope
	}
	if !IsNil(o.Ipv6AddressScope) {
		toSerialize["ipv6_address_scope"] = o.Ipv6AddressScope
	}
	if !IsNil(o.Mtu) {
		toSerialize["mtu"] = o.Mtu
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PortSecurityEnabled) {
		toSerialize["port_security_enabled"] = o.PortSecurityEnabled
	}
	if !IsNil(o.QosPolicyId) {
		toSerialize["qos_policy_id"] = o.QosPolicyId
	}
	if !IsNil(o.RevisionNumber) {
		toSerialize["revision_number"] = o.RevisionNumber
	}
	if !IsNil(o.Routerexternal) {
		toSerialize["router:external"] = o.Routerexternal
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

