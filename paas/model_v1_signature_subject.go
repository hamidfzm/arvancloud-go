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

// checks if the V1SignatureSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SignatureSubject{}

// V1SignatureSubject SignatureSubject holds information about a person or entity who created the signature.
type V1SignatureSubject struct {
	// Common name (e.g. openshift-signing-service).
	CommonName *string `json:"commonName,omitempty"`
	// Organization name.
	Organization *string `json:"organization,omitempty"`
	// If present, it is a human readable key id of public key belonging to the subject used to verify image signature. It should contain at least 64 lowest bits of public key's fingerprint (e.g. 0x685ebe62bf278440).
	PublicKeyID string `json:"publicKeyID"`
}

// NewV1SignatureSubject instantiates a new V1SignatureSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SignatureSubject(publicKeyID string) *V1SignatureSubject {
	this := V1SignatureSubject{}
	this.PublicKeyID = publicKeyID
	return &this
}

// NewV1SignatureSubjectWithDefaults instantiates a new V1SignatureSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SignatureSubjectWithDefaults() *V1SignatureSubject {
	this := V1SignatureSubject{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *V1SignatureSubject) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SignatureSubject) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *V1SignatureSubject) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *V1SignatureSubject) SetCommonName(v string) {
	o.CommonName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *V1SignatureSubject) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SignatureSubject) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *V1SignatureSubject) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *V1SignatureSubject) SetOrganization(v string) {
	o.Organization = &v
}

// GetPublicKeyID returns the PublicKeyID field value
func (o *V1SignatureSubject) GetPublicKeyID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyID
}

// GetPublicKeyIDOk returns a tuple with the PublicKeyID field value
// and a boolean to check if the value has been set.
func (o *V1SignatureSubject) GetPublicKeyIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyID, true
}

// SetPublicKeyID sets field value
func (o *V1SignatureSubject) SetPublicKeyID(v string) {
	o.PublicKeyID = v
}

func (o V1SignatureSubject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SignatureSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonName) {
		toSerialize["commonName"] = o.CommonName
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	toSerialize["publicKeyID"] = o.PublicKeyID
	return toSerialize, nil
}

type NullableV1SignatureSubject struct {
	value *V1SignatureSubject
	isSet bool
}

func (v NullableV1SignatureSubject) Get() *V1SignatureSubject {
	return v.value
}

func (v *NullableV1SignatureSubject) Set(val *V1SignatureSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SignatureSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SignatureSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SignatureSubject(val *V1SignatureSubject) *NullableV1SignatureSubject {
	return &NullableV1SignatureSubject{value: val, isSet: true}
}

func (v NullableV1SignatureSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SignatureSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

