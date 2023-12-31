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

// checks if the OutputLocationS3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputLocationS3{}

// OutputLocationS3 struct for OutputLocationS3
type OutputLocationS3 struct {
	BucketName string `json:"BucketName"`
	Prefix string `json:"Prefix"`
	Encryption *Encryption `json:"Encryption,omitempty"`
	CannedACL *ObjectCannedACL `json:"CannedACL,omitempty"`
	AccessControlList *Grants `json:"AccessControlList,omitempty"`
	Tagging *S3LocationTagging `json:"Tagging,omitempty"`
	UserMetadata *Array `json:"UserMetadata,omitempty"`
	StorageClass *StorageClass `json:"StorageClass,omitempty"`
}

// NewOutputLocationS3 instantiates a new OutputLocationS3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputLocationS3(bucketName string, prefix string) *OutputLocationS3 {
	this := OutputLocationS3{}
	this.BucketName = bucketName
	this.Prefix = prefix
	return &this
}

// NewOutputLocationS3WithDefaults instantiates a new OutputLocationS3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputLocationS3WithDefaults() *OutputLocationS3 {
	this := OutputLocationS3{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *OutputLocationS3) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *OutputLocationS3) SetBucketName(v string) {
	o.BucketName = v
}

// GetPrefix returns the Prefix field value
func (o *OutputLocationS3) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *OutputLocationS3) SetPrefix(v string) {
	o.Prefix = v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *OutputLocationS3) GetEncryption() Encryption {
	if o == nil || IsNil(o.Encryption) {
		var ret Encryption
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetEncryptionOk() (*Encryption, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *OutputLocationS3) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given Encryption and assigns it to the Encryption field.
func (o *OutputLocationS3) SetEncryption(v Encryption) {
	o.Encryption = &v
}

// GetCannedACL returns the CannedACL field value if set, zero value otherwise.
func (o *OutputLocationS3) GetCannedACL() ObjectCannedACL {
	if o == nil || IsNil(o.CannedACL) {
		var ret ObjectCannedACL
		return ret
	}
	return *o.CannedACL
}

// GetCannedACLOk returns a tuple with the CannedACL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetCannedACLOk() (*ObjectCannedACL, bool) {
	if o == nil || IsNil(o.CannedACL) {
		return nil, false
	}
	return o.CannedACL, true
}

// HasCannedACL returns a boolean if a field has been set.
func (o *OutputLocationS3) HasCannedACL() bool {
	if o != nil && !IsNil(o.CannedACL) {
		return true
	}

	return false
}

// SetCannedACL gets a reference to the given ObjectCannedACL and assigns it to the CannedACL field.
func (o *OutputLocationS3) SetCannedACL(v ObjectCannedACL) {
	o.CannedACL = &v
}

// GetAccessControlList returns the AccessControlList field value if set, zero value otherwise.
func (o *OutputLocationS3) GetAccessControlList() Grants {
	if o == nil || IsNil(o.AccessControlList) {
		var ret Grants
		return ret
	}
	return *o.AccessControlList
}

// GetAccessControlListOk returns a tuple with the AccessControlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetAccessControlListOk() (*Grants, bool) {
	if o == nil || IsNil(o.AccessControlList) {
		return nil, false
	}
	return o.AccessControlList, true
}

// HasAccessControlList returns a boolean if a field has been set.
func (o *OutputLocationS3) HasAccessControlList() bool {
	if o != nil && !IsNil(o.AccessControlList) {
		return true
	}

	return false
}

// SetAccessControlList gets a reference to the given Grants and assigns it to the AccessControlList field.
func (o *OutputLocationS3) SetAccessControlList(v Grants) {
	o.AccessControlList = &v
}

// GetTagging returns the Tagging field value if set, zero value otherwise.
func (o *OutputLocationS3) GetTagging() S3LocationTagging {
	if o == nil || IsNil(o.Tagging) {
		var ret S3LocationTagging
		return ret
	}
	return *o.Tagging
}

// GetTaggingOk returns a tuple with the Tagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetTaggingOk() (*S3LocationTagging, bool) {
	if o == nil || IsNil(o.Tagging) {
		return nil, false
	}
	return o.Tagging, true
}

// HasTagging returns a boolean if a field has been set.
func (o *OutputLocationS3) HasTagging() bool {
	if o != nil && !IsNil(o.Tagging) {
		return true
	}

	return false
}

// SetTagging gets a reference to the given S3LocationTagging and assigns it to the Tagging field.
func (o *OutputLocationS3) SetTagging(v S3LocationTagging) {
	o.Tagging = &v
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *OutputLocationS3) GetUserMetadata() Array {
	if o == nil || IsNil(o.UserMetadata) {
		var ret Array
		return ret
	}
	return *o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetUserMetadataOk() (*Array, bool) {
	if o == nil || IsNil(o.UserMetadata) {
		return nil, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *OutputLocationS3) HasUserMetadata() bool {
	if o != nil && !IsNil(o.UserMetadata) {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given Array and assigns it to the UserMetadata field.
func (o *OutputLocationS3) SetUserMetadata(v Array) {
	o.UserMetadata = &v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *OutputLocationS3) GetStorageClass() StorageClass {
	if o == nil || IsNil(o.StorageClass) {
		var ret StorageClass
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputLocationS3) GetStorageClassOk() (*StorageClass, bool) {
	if o == nil || IsNil(o.StorageClass) {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *OutputLocationS3) HasStorageClass() bool {
	if o != nil && !IsNil(o.StorageClass) {
		return true
	}

	return false
}

// SetStorageClass gets a reference to the given StorageClass and assigns it to the StorageClass field.
func (o *OutputLocationS3) SetStorageClass(v StorageClass) {
	o.StorageClass = &v
}

func (o OutputLocationS3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputLocationS3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BucketName"] = o.BucketName
	toSerialize["Prefix"] = o.Prefix
	if !IsNil(o.Encryption) {
		toSerialize["Encryption"] = o.Encryption
	}
	if !IsNil(o.CannedACL) {
		toSerialize["CannedACL"] = o.CannedACL
	}
	if !IsNil(o.AccessControlList) {
		toSerialize["AccessControlList"] = o.AccessControlList
	}
	if !IsNil(o.Tagging) {
		toSerialize["Tagging"] = o.Tagging
	}
	if !IsNil(o.UserMetadata) {
		toSerialize["UserMetadata"] = o.UserMetadata
	}
	if !IsNil(o.StorageClass) {
		toSerialize["StorageClass"] = o.StorageClass
	}
	return toSerialize, nil
}

type NullableOutputLocationS3 struct {
	value *OutputLocationS3
	isSet bool
}

func (v NullableOutputLocationS3) Get() *OutputLocationS3 {
	return v.value
}

func (v *NullableOutputLocationS3) Set(val *OutputLocationS3) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputLocationS3) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputLocationS3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputLocationS3(val *OutputLocationS3) *NullableOutputLocationS3 {
	return &NullableOutputLocationS3{value: val, isSet: true}
}

func (v NullableOutputLocationS3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputLocationS3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


