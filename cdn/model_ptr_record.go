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

// checks if the PTRRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PTRRecord{}

// PTRRecord struct for PTRRecord
type PTRRecord struct {
	Value *PTRRecordValue `json:"value,omitempty"`
	Type *string `json:"type,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	Cloud *bool `json:"cloud,omitempty"`
	UpstreamHttps *string `json:"upstream_https,omitempty"`
	IpFilterMode *DnsRecordIpFilterMode `json:"ip_filter_mode,omitempty"`
	// Protected records cannot be modified or deleted by user.
	IsProtected *bool `json:"is_protected,omitempty"`
	Usage []string `json:"usage,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPTRRecord instantiates a new PTRRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPTRRecord() *PTRRecord {
	this := PTRRecord{}
	var cloud bool = false
	this.Cloud = &cloud
	return &this
}

// NewPTRRecordWithDefaults instantiates a new PTRRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPTRRecordWithDefaults() *PTRRecord {
	this := PTRRecord{}
	var type_ string = "ptr"
	this.Type = &type_
	var cloud bool = false
	this.Cloud = &cloud
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PTRRecord) GetValue() PTRRecordValue {
	if o == nil || IsNil(o.Value) {
		var ret PTRRecordValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetValueOk() (*PTRRecordValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PTRRecord) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PTRRecordValue and assigns it to the Value field.
func (o *PTRRecord) SetValue(v PTRRecordValue) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PTRRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PTRRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PTRRecord) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PTRRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PTRRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PTRRecord) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PTRRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PTRRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PTRRecord) SetName(v string) {
	o.Name = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *PTRRecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *PTRRecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *PTRRecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *PTRRecord) GetCloud() bool {
	if o == nil || IsNil(o.Cloud) {
		var ret bool
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.Cloud) {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *PTRRecord) HasCloud() bool {
	if o != nil && !IsNil(o.Cloud) {
		return true
	}

	return false
}

// SetCloud gets a reference to the given bool and assigns it to the Cloud field.
func (o *PTRRecord) SetCloud(v bool) {
	o.Cloud = &v
}

// GetUpstreamHttps returns the UpstreamHttps field value if set, zero value otherwise.
func (o *PTRRecord) GetUpstreamHttps() string {
	if o == nil || IsNil(o.UpstreamHttps) {
		var ret string
		return ret
	}
	return *o.UpstreamHttps
}

// GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetUpstreamHttpsOk() (*string, bool) {
	if o == nil || IsNil(o.UpstreamHttps) {
		return nil, false
	}
	return o.UpstreamHttps, true
}

// HasUpstreamHttps returns a boolean if a field has been set.
func (o *PTRRecord) HasUpstreamHttps() bool {
	if o != nil && !IsNil(o.UpstreamHttps) {
		return true
	}

	return false
}

// SetUpstreamHttps gets a reference to the given string and assigns it to the UpstreamHttps field.
func (o *PTRRecord) SetUpstreamHttps(v string) {
	o.UpstreamHttps = &v
}

// GetIpFilterMode returns the IpFilterMode field value if set, zero value otherwise.
func (o *PTRRecord) GetIpFilterMode() DnsRecordIpFilterMode {
	if o == nil || IsNil(o.IpFilterMode) {
		var ret DnsRecordIpFilterMode
		return ret
	}
	return *o.IpFilterMode
}

// GetIpFilterModeOk returns a tuple with the IpFilterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool) {
	if o == nil || IsNil(o.IpFilterMode) {
		return nil, false
	}
	return o.IpFilterMode, true
}

// HasIpFilterMode returns a boolean if a field has been set.
func (o *PTRRecord) HasIpFilterMode() bool {
	if o != nil && !IsNil(o.IpFilterMode) {
		return true
	}

	return false
}

// SetIpFilterMode gets a reference to the given DnsRecordIpFilterMode and assigns it to the IpFilterMode field.
func (o *PTRRecord) SetIpFilterMode(v DnsRecordIpFilterMode) {
	o.IpFilterMode = &v
}

// GetIsProtected returns the IsProtected field value if set, zero value otherwise.
func (o *PTRRecord) GetIsProtected() bool {
	if o == nil || IsNil(o.IsProtected) {
		var ret bool
		return ret
	}
	return *o.IsProtected
}

// GetIsProtectedOk returns a tuple with the IsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetIsProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProtected) {
		return nil, false
	}
	return o.IsProtected, true
}

// HasIsProtected returns a boolean if a field has been set.
func (o *PTRRecord) HasIsProtected() bool {
	if o != nil && !IsNil(o.IsProtected) {
		return true
	}

	return false
}

// SetIsProtected gets a reference to the given bool and assigns it to the IsProtected field.
func (o *PTRRecord) SetIsProtected(v bool) {
	o.IsProtected = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *PTRRecord) GetUsage() []string {
	if o == nil || IsNil(o.Usage) {
		var ret []string
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetUsageOk() ([]string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *PTRRecord) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []string and assigns it to the Usage field.
func (o *PTRRecord) SetUsage(v []string) {
	o.Usage = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PTRRecord) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PTRRecord) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PTRRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PTRRecord) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecord) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PTRRecord) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PTRRecord) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PTRRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PTRRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Cloud) {
		toSerialize["cloud"] = o.Cloud
	}
	if !IsNil(o.UpstreamHttps) {
		toSerialize["upstream_https"] = o.UpstreamHttps
	}
	if !IsNil(o.IpFilterMode) {
		toSerialize["ip_filter_mode"] = o.IpFilterMode
	}
	if !IsNil(o.IsProtected) {
		toSerialize["is_protected"] = o.IsProtected
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePTRRecord struct {
	value *PTRRecord
	isSet bool
}

func (v NullablePTRRecord) Get() *PTRRecord {
	return v.value
}

func (v *NullablePTRRecord) Set(val *PTRRecord) {
	v.value = val
	v.isSet = true
}

func (v NullablePTRRecord) IsSet() bool {
	return v.isSet
}

func (v *NullablePTRRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePTRRecord(val *PTRRecord) *NullablePTRRecord {
	return &NullablePTRRecord{value: val, isSet: true}
}

func (v NullablePTRRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePTRRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


