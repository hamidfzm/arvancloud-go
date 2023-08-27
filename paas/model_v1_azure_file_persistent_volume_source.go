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

// checks if the V1AzureFilePersistentVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1AzureFilePersistentVolumeSource{}

// V1AzureFilePersistentVolumeSource AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type V1AzureFilePersistentVolumeSource struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName"`
	// the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
	SecretNamespace string `json:"secretNamespace"`
	// Share Name
	ShareName string `json:"shareName"`
}

// NewV1AzureFilePersistentVolumeSource instantiates a new V1AzureFilePersistentVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AzureFilePersistentVolumeSource(secretName string, secretNamespace string, shareName string) *V1AzureFilePersistentVolumeSource {
	this := V1AzureFilePersistentVolumeSource{}
	this.SecretName = secretName
	this.SecretNamespace = secretNamespace
	this.ShareName = shareName
	return &this
}

// NewV1AzureFilePersistentVolumeSourceWithDefaults instantiates a new V1AzureFilePersistentVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AzureFilePersistentVolumeSourceWithDefaults() *V1AzureFilePersistentVolumeSource {
	this := V1AzureFilePersistentVolumeSource{}
	return &this
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1AzureFilePersistentVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AzureFilePersistentVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1AzureFilePersistentVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1AzureFilePersistentVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSecretName returns the SecretName field value
func (o *V1AzureFilePersistentVolumeSource) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *V1AzureFilePersistentVolumeSource) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value
func (o *V1AzureFilePersistentVolumeSource) SetSecretName(v string) {
	o.SecretName = v
}

// GetSecretNamespace returns the SecretNamespace field value
func (o *V1AzureFilePersistentVolumeSource) GetSecretNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretNamespace
}

// GetSecretNamespaceOk returns a tuple with the SecretNamespace field value
// and a boolean to check if the value has been set.
func (o *V1AzureFilePersistentVolumeSource) GetSecretNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretNamespace, true
}

// SetSecretNamespace sets field value
func (o *V1AzureFilePersistentVolumeSource) SetSecretNamespace(v string) {
	o.SecretNamespace = v
}

// GetShareName returns the ShareName field value
func (o *V1AzureFilePersistentVolumeSource) GetShareName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareName
}

// GetShareNameOk returns a tuple with the ShareName field value
// and a boolean to check if the value has been set.
func (o *V1AzureFilePersistentVolumeSource) GetShareNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareName, true
}

// SetShareName sets field value
func (o *V1AzureFilePersistentVolumeSource) SetShareName(v string) {
	o.ShareName = v
}

func (o V1AzureFilePersistentVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1AzureFilePersistentVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	toSerialize["secretName"] = o.SecretName
	toSerialize["secretNamespace"] = o.SecretNamespace
	toSerialize["shareName"] = o.ShareName
	return toSerialize, nil
}

type NullableV1AzureFilePersistentVolumeSource struct {
	value *V1AzureFilePersistentVolumeSource
	isSet bool
}

func (v NullableV1AzureFilePersistentVolumeSource) Get() *V1AzureFilePersistentVolumeSource {
	return v.value
}

func (v *NullableV1AzureFilePersistentVolumeSource) Set(val *V1AzureFilePersistentVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AzureFilePersistentVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AzureFilePersistentVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AzureFilePersistentVolumeSource(val *V1AzureFilePersistentVolumeSource) *NullableV1AzureFilePersistentVolumeSource {
	return &NullableV1AzureFilePersistentVolumeSource{value: val, isSet: true}
}

func (v NullableV1AzureFilePersistentVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AzureFilePersistentVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

