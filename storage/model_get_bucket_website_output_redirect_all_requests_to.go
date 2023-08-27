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

// checks if the GetBucketWebsiteOutputRedirectAllRequestsTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBucketWebsiteOutputRedirectAllRequestsTo{}

// GetBucketWebsiteOutputRedirectAllRequestsTo struct for GetBucketWebsiteOutputRedirectAllRequestsTo
type GetBucketWebsiteOutputRedirectAllRequestsTo struct {
	HostName string `json:"HostName"`
	Protocol *Protocol `json:"Protocol,omitempty"`
}

// NewGetBucketWebsiteOutputRedirectAllRequestsTo instantiates a new GetBucketWebsiteOutputRedirectAllRequestsTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketWebsiteOutputRedirectAllRequestsTo(hostName string) *GetBucketWebsiteOutputRedirectAllRequestsTo {
	this := GetBucketWebsiteOutputRedirectAllRequestsTo{}
	this.HostName = hostName
	return &this
}

// NewGetBucketWebsiteOutputRedirectAllRequestsToWithDefaults instantiates a new GetBucketWebsiteOutputRedirectAllRequestsTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketWebsiteOutputRedirectAllRequestsToWithDefaults() *GetBucketWebsiteOutputRedirectAllRequestsTo {
	this := GetBucketWebsiteOutputRedirectAllRequestsTo{}
	return &this
}

// GetHostName returns the HostName field value
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) SetHostName(v string) {
	o.HostName = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) GetProtocol() Protocol {
	if o == nil || IsNil(o.Protocol) {
		var ret Protocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) GetProtocolOk() (*Protocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *GetBucketWebsiteOutputRedirectAllRequestsTo) SetProtocol(v Protocol) {
	o.Protocol = &v
}

func (o GetBucketWebsiteOutputRedirectAllRequestsTo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBucketWebsiteOutputRedirectAllRequestsTo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["HostName"] = o.HostName
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	return toSerialize, nil
}

type NullableGetBucketWebsiteOutputRedirectAllRequestsTo struct {
	value *GetBucketWebsiteOutputRedirectAllRequestsTo
	isSet bool
}

func (v NullableGetBucketWebsiteOutputRedirectAllRequestsTo) Get() *GetBucketWebsiteOutputRedirectAllRequestsTo {
	return v.value
}

func (v *NullableGetBucketWebsiteOutputRedirectAllRequestsTo) Set(val *GetBucketWebsiteOutputRedirectAllRequestsTo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketWebsiteOutputRedirectAllRequestsTo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketWebsiteOutputRedirectAllRequestsTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketWebsiteOutputRedirectAllRequestsTo(val *GetBucketWebsiteOutputRedirectAllRequestsTo) *NullableGetBucketWebsiteOutputRedirectAllRequestsTo {
	return &NullableGetBucketWebsiteOutputRedirectAllRequestsTo{value: val, isSet: true}
}

func (v NullableGetBucketWebsiteOutputRedirectAllRequestsTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketWebsiteOutputRedirectAllRequestsTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

