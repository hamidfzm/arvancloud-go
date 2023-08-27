/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the BulkReportsTrafficsTotalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReportsTrafficsTotalRequest{}

// BulkReportsTrafficsTotalRequest struct for BulkReportsTrafficsTotalRequest
type BulkReportsTrafficsTotalRequest struct {
	// List of domains' IDs
	Domains []string `json:"domains,omitempty"`
	// Whether to include sub-domains or report only root domain traffic
	ExcludeSubdomains *bool `json:"excludeSubdomains,omitempty"`
}

// NewBulkReportsTrafficsTotalRequest instantiates a new BulkReportsTrafficsTotalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReportsTrafficsTotalRequest() *BulkReportsTrafficsTotalRequest {
	this := BulkReportsTrafficsTotalRequest{}
	return &this
}

// NewBulkReportsTrafficsTotalRequestWithDefaults instantiates a new BulkReportsTrafficsTotalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReportsTrafficsTotalRequestWithDefaults() *BulkReportsTrafficsTotalRequest {
	this := BulkReportsTrafficsTotalRequest{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *BulkReportsTrafficsTotalRequest) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReportsTrafficsTotalRequest) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *BulkReportsTrafficsTotalRequest) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *BulkReportsTrafficsTotalRequest) SetDomains(v []string) {
	o.Domains = v
}

// GetExcludeSubdomains returns the ExcludeSubdomains field value if set, zero value otherwise.
func (o *BulkReportsTrafficsTotalRequest) GetExcludeSubdomains() bool {
	if o == nil || IsNil(o.ExcludeSubdomains) {
		var ret bool
		return ret
	}
	return *o.ExcludeSubdomains
}

// GetExcludeSubdomainsOk returns a tuple with the ExcludeSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReportsTrafficsTotalRequest) GetExcludeSubdomainsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExcludeSubdomains) {
		return nil, false
	}
	return o.ExcludeSubdomains, true
}

// HasExcludeSubdomains returns a boolean if a field has been set.
func (o *BulkReportsTrafficsTotalRequest) HasExcludeSubdomains() bool {
	if o != nil && !IsNil(o.ExcludeSubdomains) {
		return true
	}

	return false
}

// SetExcludeSubdomains gets a reference to the given bool and assigns it to the ExcludeSubdomains field.
func (o *BulkReportsTrafficsTotalRequest) SetExcludeSubdomains(v bool) {
	o.ExcludeSubdomains = &v
}

func (o BulkReportsTrafficsTotalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReportsTrafficsTotalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.ExcludeSubdomains) {
		toSerialize["excludeSubdomains"] = o.ExcludeSubdomains
	}
	return toSerialize, nil
}

type NullableBulkReportsTrafficsTotalRequest struct {
	value *BulkReportsTrafficsTotalRequest
	isSet bool
}

func (v NullableBulkReportsTrafficsTotalRequest) Get() *BulkReportsTrafficsTotalRequest {
	return v.value
}

func (v *NullableBulkReportsTrafficsTotalRequest) Set(val *BulkReportsTrafficsTotalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReportsTrafficsTotalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReportsTrafficsTotalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReportsTrafficsTotalRequest(val *BulkReportsTrafficsTotalRequest) *NullableBulkReportsTrafficsTotalRequest {
	return &NullableBulkReportsTrafficsTotalRequest{value: val, isSet: true}
}

func (v NullableBulkReportsTrafficsTotalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReportsTrafficsTotalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


