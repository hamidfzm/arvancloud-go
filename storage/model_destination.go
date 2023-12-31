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

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination Specifies information about where to publish analysis or configuration results for an ArvanCloud S3 bucket and S3 Replication Time Control (S3 RTC).
type Destination struct {
	Bucket string `json:"Bucket"`
	Account *string `json:"Account,omitempty"`
	StorageClass *StorageClass `json:"StorageClass,omitempty"`
	AccessControlTranslation *DestinationAccessControlTranslation `json:"AccessControlTranslation,omitempty"`
	EncryptionConfiguration *DestinationEncryptionConfiguration `json:"EncryptionConfiguration,omitempty"`
	ReplicationTime *DestinationReplicationTime `json:"ReplicationTime,omitempty"`
	Metrics *DestinationMetrics `json:"Metrics,omitempty"`
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination(bucket string) *Destination {
	this := Destination{}
	this.Bucket = bucket
	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	return &this
}

// GetBucket returns the Bucket field value
func (o *Destination) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *Destination) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *Destination) SetBucket(v string) {
	o.Bucket = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Destination) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Destination) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Destination) SetAccount(v string) {
	o.Account = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *Destination) GetStorageClass() StorageClass {
	if o == nil || IsNil(o.StorageClass) {
		var ret StorageClass
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetStorageClassOk() (*StorageClass, bool) {
	if o == nil || IsNil(o.StorageClass) {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *Destination) HasStorageClass() bool {
	if o != nil && !IsNil(o.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given StorageClass and assigns it to the StorageClass field.
func (o *Destination) SetStorageClass(v StorageClass) {
	o.StorageClass = &v
}

// GetAccessControlTranslation returns the AccessControlTranslation field value if set, zero value otherwise.
func (o *Destination) GetAccessControlTranslation() DestinationAccessControlTranslation {
	if o == nil || IsNil(o.AccessControlTranslation) {
		var ret DestinationAccessControlTranslation
		return ret
	}
	return *o.AccessControlTranslation
}

// GetAccessControlTranslationOk returns a tuple with the AccessControlTranslation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetAccessControlTranslationOk() (*DestinationAccessControlTranslation, bool) {
	if o == nil || IsNil(o.AccessControlTranslation) {
		return nil, false
	}
	return o.AccessControlTranslation, true
}

// HasAccessControlTranslation returns a boolean if a field has been set.
func (o *Destination) HasAccessControlTranslation() bool {
	if o != nil && !IsNil(o.AccessControlTranslation) {
		return true
	}

	return false
}

// SetAccessControlTranslation gets a reference to the given DestinationAccessControlTranslation and assigns it to the AccessControlTranslation field.
func (o *Destination) SetAccessControlTranslation(v DestinationAccessControlTranslation) {
	o.AccessControlTranslation = &v
}

// GetEncryptionConfiguration returns the EncryptionConfiguration field value if set, zero value otherwise.
func (o *Destination) GetEncryptionConfiguration() DestinationEncryptionConfiguration {
	if o == nil || IsNil(o.EncryptionConfiguration) {
		var ret DestinationEncryptionConfiguration
		return ret
	}
	return *o.EncryptionConfiguration
}

// GetEncryptionConfigurationOk returns a tuple with the EncryptionConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetEncryptionConfigurationOk() (*DestinationEncryptionConfiguration, bool) {
	if o == nil || IsNil(o.EncryptionConfiguration) {
		return nil, false
	}
	return o.EncryptionConfiguration, true
}

// HasEncryptionConfiguration returns a boolean if a field has been set.
func (o *Destination) HasEncryptionConfiguration() bool {
	if o != nil && !IsNil(o.EncryptionConfiguration) {
		return true
	}

	return false
}

// SetEncryptionConfiguration gets a reference to the given DestinationEncryptionConfiguration and assigns it to the EncryptionConfiguration field.
func (o *Destination) SetEncryptionConfiguration(v DestinationEncryptionConfiguration) {
	o.EncryptionConfiguration = &v
}

// GetReplicationTime returns the ReplicationTime field value if set, zero value otherwise.
func (o *Destination) GetReplicationTime() DestinationReplicationTime {
	if o == nil || IsNil(o.ReplicationTime) {
		var ret DestinationReplicationTime
		return ret
	}
	return *o.ReplicationTime
}

// GetReplicationTimeOk returns a tuple with the ReplicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetReplicationTimeOk() (*DestinationReplicationTime, bool) {
	if o == nil || IsNil(o.ReplicationTime) {
		return nil, false
	}
	return o.ReplicationTime, true
}

// HasReplicationTime returns a boolean if a field has been set.
func (o *Destination) HasReplicationTime() bool {
	if o != nil && !IsNil(o.ReplicationTime) {
		return true
	}

	return false
}

// SetReplicationTime gets a reference to the given DestinationReplicationTime and assigns it to the ReplicationTime field.
func (o *Destination) SetReplicationTime(v DestinationReplicationTime) {
	o.ReplicationTime = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Destination) GetMetrics() DestinationMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret DestinationMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetMetricsOk() (*DestinationMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Destination) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given DestinationMetrics and assigns it to the Metrics field.
func (o *Destination) SetMetrics(v DestinationMetrics) {
	o.Metrics = &v
}

func (o Destination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Bucket"] = o.Bucket
	if !IsNil(o.Account) {
		toSerialize["Account"] = o.Account
	}
	if !IsNil(o.StorageClass) {
		toSerialize["StorageClass"] = o.StorageClass
	}
	if !IsNil(o.AccessControlTranslation) {
		toSerialize["AccessControlTranslation"] = o.AccessControlTranslation
	}
	if !IsNil(o.EncryptionConfiguration) {
		toSerialize["EncryptionConfiguration"] = o.EncryptionConfiguration
	}
	if !IsNil(o.ReplicationTime) {
		toSerialize["ReplicationTime"] = o.ReplicationTime
	}
	if !IsNil(o.Metrics) {
		toSerialize["Metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


