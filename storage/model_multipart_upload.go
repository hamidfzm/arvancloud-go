/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"time"
)

// checks if the MultipartUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipartUpload{}

// MultipartUpload Container for the <code>MultipartUpload</code> for the ArvanCloud S3 object.
type MultipartUpload struct {
	UploadId *string `json:"UploadId,omitempty"`
	Key *string `json:"Key,omitempty"`
	Initiated *time.Time `json:"Initiated,omitempty"`
	StorageClass *StorageClass `json:"StorageClass,omitempty"`
	Owner *Owner `json:"Owner,omitempty"`
	Initiator *MultipartUploadInitiator `json:"Initiator,omitempty"`
}

// NewMultipartUpload instantiates a new MultipartUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipartUpload() *MultipartUpload {
	this := MultipartUpload{}
	return &this
}

// NewMultipartUploadWithDefaults instantiates a new MultipartUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipartUploadWithDefaults() *MultipartUpload {
	this := MultipartUpload{}
	return &this
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *MultipartUpload) GetUploadId() string {
	if o == nil || IsNil(o.UploadId) {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartUpload) GetUploadIdOk() (*string, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *MultipartUpload) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *MultipartUpload) SetUploadId(v string) {
	o.UploadId = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MultipartUpload) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartUpload) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MultipartUpload) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MultipartUpload) SetKey(v string) {
	o.Key = &v
}

// GetInitiated returns the Initiated field value if set, zero value otherwise.
func (o *MultipartUpload) GetInitiated() time.Time {
	if o == nil || IsNil(o.Initiated) {
		var ret time.Time
		return ret
	}
	return *o.Initiated
}

// GetInitiatedOk returns a tuple with the Initiated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartUpload) GetInitiatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Initiated) {
		return nil, false
	}
	return o.Initiated, true
}

// HasInitiated returns a boolean if a field has been set.
func (o *MultipartUpload) HasInitiated() bool {
	if o != nil && !IsNil(o.Initiated) {
		return true
	}

	return false
}

// SetInitiated gets a reference to the given time.Time and assigns it to the Initiated field.
func (o *MultipartUpload) SetInitiated(v time.Time) {
	o.Initiated = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *MultipartUpload) GetStorageClass() StorageClass {
	if o == nil || IsNil(o.StorageClass) {
		var ret StorageClass
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartUpload) GetStorageClassOk() (*StorageClass, bool) {
	if o == nil || IsNil(o.StorageClass) {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *MultipartUpload) HasStorageClass() bool {
	if o != nil && !IsNil(o.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given StorageClass and assigns it to the StorageClass field.
func (o *MultipartUpload) SetStorageClass(v StorageClass) {
	o.StorageClass = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *MultipartUpload) GetOwner() Owner {
	if o == nil || IsNil(o.Owner) {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartUpload) GetOwnerOk() (*Owner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MultipartUpload) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *MultipartUpload) SetOwner(v Owner) {
	o.Owner = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *MultipartUpload) GetInitiator() MultipartUploadInitiator {
	if o == nil || IsNil(o.Initiator) {
		var ret MultipartUploadInitiator
		return ret
	}
	return *o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipartUpload) GetInitiatorOk() (*MultipartUploadInitiator, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *MultipartUpload) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given MultipartUploadInitiator and assigns it to the Initiator field.
func (o *MultipartUpload) SetInitiator(v MultipartUploadInitiator) {
	o.Initiator = &v
}

func (o MultipartUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipartUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadId) {
		toSerialize["UploadId"] = o.UploadId
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Initiated) {
		toSerialize["Initiated"] = o.Initiated
	}
	if !IsNil(o.StorageClass) {
		toSerialize["StorageClass"] = o.StorageClass
	}
	if !IsNil(o.Owner) {
		toSerialize["Owner"] = o.Owner
	}
	if !IsNil(o.Initiator) {
		toSerialize["Initiator"] = o.Initiator
	}
	return toSerialize, nil
}

type NullableMultipartUpload struct {
	value *MultipartUpload
	isSet bool
}

func (v NullableMultipartUpload) Get() *MultipartUpload {
	return v.value
}

func (v *NullableMultipartUpload) Set(val *MultipartUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipartUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipartUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipartUpload(val *MultipartUpload) *NullableMultipartUpload {
	return &NullableMultipartUpload{value: val, isSet: true}
}

func (v NullableMultipartUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipartUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


