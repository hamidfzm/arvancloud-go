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

// checks if the Certificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certificate{}

// Certificate struct for Certificate
type Certificate struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	KeyType NullableString `json:"key_type,omitempty"`
	DomainNames []string `json:"domain_names,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCertificate instantiates a new Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate() *Certificate {
	this := Certificate{}
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Certificate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Certificate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Certificate) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Certificate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Certificate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Certificate) SetType(v string) {
	o.Type = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Certificate) GetKeyType() string {
	if o == nil || IsNil(o.KeyType.Get()) {
		var ret string
		return ret
	}
	return *o.KeyType.Get()
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Certificate) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyType.Get(), o.KeyType.IsSet()
}

// HasKeyType returns a boolean if a field has been set.
func (o *Certificate) HasKeyType() bool {
	if o != nil && o.KeyType.IsSet() {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given NullableString and assigns it to the KeyType field.
func (o *Certificate) SetKeyType(v string) {
	o.KeyType.Set(&v)
}
// SetKeyTypeNil sets the value for KeyType to be an explicit nil
func (o *Certificate) SetKeyTypeNil() {
	o.KeyType.Set(nil)
}

// UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
func (o *Certificate) UnsetKeyType() {
	o.KeyType.Unset()
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *Certificate) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *Certificate) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *Certificate) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *Certificate) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *Certificate) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *Certificate) SetIssuer(v string) {
	o.Issuer = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Certificate) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Certificate) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *Certificate) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Certificate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Certificate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Certificate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Certificate) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Certificate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Certificate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.KeyType.IsSet() {
		toSerialize["key_type"] = o.KeyType.Get()
	}
	if !IsNil(o.DomainNames) {
		toSerialize["domain_names"] = o.DomainNames
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableCertificate struct {
	value *Certificate
	isSet bool
}

func (v NullableCertificate) Get() *Certificate {
	return v.value
}

func (v *NullableCertificate) Set(val *Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate(val *Certificate) *NullableCertificate {
	return &NullableCertificate{value: val, isSet: true}
}

func (v NullableCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


