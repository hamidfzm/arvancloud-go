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

// checks if the V1ServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ServiceAccount{}

// V1ServiceAccount ServiceAccount binds together: * a name, understood by users, and perhaps by peripheral systems, for an identity * a principal that can be authenticated and authorized * a set of secrets
type V1ServiceAccount struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// AutomountServiceAccountToken indicates whether pods running as this service account should have an API token automatically mounted. Can be overridden at the pod level.
	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty"`
	// ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod
	ImagePullSecrets []V1LocalObjectReference `json:"imagePullSecrets,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	// Secrets is the list of secrets allowed to be used by pods running using this ServiceAccount. More info: https://kubernetes.io/docs/concepts/configuration/secret
	Secrets []V1ObjectReference `json:"secrets,omitempty"`
}

// NewV1ServiceAccount instantiates a new V1ServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ServiceAccount() *V1ServiceAccount {
	this := V1ServiceAccount{}
	return &this
}

// NewV1ServiceAccountWithDefaults instantiates a new V1ServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ServiceAccountWithDefaults() *V1ServiceAccount {
	this := V1ServiceAccount{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1ServiceAccount) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ServiceAccount) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1ServiceAccount) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1ServiceAccount) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetAutomountServiceAccountToken returns the AutomountServiceAccountToken field value if set, zero value otherwise.
func (o *V1ServiceAccount) GetAutomountServiceAccountToken() bool {
	if o == nil || IsNil(o.AutomountServiceAccountToken) {
		var ret bool
		return ret
	}
	return *o.AutomountServiceAccountToken
}

// GetAutomountServiceAccountTokenOk returns a tuple with the AutomountServiceAccountToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ServiceAccount) GetAutomountServiceAccountTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomountServiceAccountToken) {
		return nil, false
	}
	return o.AutomountServiceAccountToken, true
}

// HasAutomountServiceAccountToken returns a boolean if a field has been set.
func (o *V1ServiceAccount) HasAutomountServiceAccountToken() bool {
	if o != nil && !IsNil(o.AutomountServiceAccountToken) {
		return true
	}

	return false
}

// SetAutomountServiceAccountToken gets a reference to the given bool and assigns it to the AutomountServiceAccountToken field.
func (o *V1ServiceAccount) SetAutomountServiceAccountToken(v bool) {
	o.AutomountServiceAccountToken = &v
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *V1ServiceAccount) GetImagePullSecrets() []V1LocalObjectReference {
	if o == nil || IsNil(o.ImagePullSecrets) {
		var ret []V1LocalObjectReference
		return ret
	}
	return o.ImagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ServiceAccount) GetImagePullSecretsOk() ([]V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.ImagePullSecrets) {
		return nil, false
	}
	return o.ImagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *V1ServiceAccount) HasImagePullSecrets() bool {
	if o != nil && !IsNil(o.ImagePullSecrets) {
		return true
	}

	return false
}

// SetImagePullSecrets gets a reference to the given []V1LocalObjectReference and assigns it to the ImagePullSecrets field.
func (o *V1ServiceAccount) SetImagePullSecrets(v []V1LocalObjectReference) {
	o.ImagePullSecrets = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1ServiceAccount) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ServiceAccount) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1ServiceAccount) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1ServiceAccount) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1ServiceAccount) GetMetadata() V1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ServiceAccount) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1ServiceAccount) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *V1ServiceAccount) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *V1ServiceAccount) GetSecrets() []V1ObjectReference {
	if o == nil || IsNil(o.Secrets) {
		var ret []V1ObjectReference
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ServiceAccount) GetSecretsOk() ([]V1ObjectReference, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *V1ServiceAccount) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []V1ObjectReference and assigns it to the Secrets field.
func (o *V1ServiceAccount) SetSecrets(v []V1ObjectReference) {
	o.Secrets = v
}

func (o V1ServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.AutomountServiceAccountToken) {
		toSerialize["automountServiceAccountToken"] = o.AutomountServiceAccountToken
	}
	if !IsNil(o.ImagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.ImagePullSecrets
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableV1ServiceAccount struct {
	value *V1ServiceAccount
	isSet bool
}

func (v NullableV1ServiceAccount) Get() *V1ServiceAccount {
	return v.value
}

func (v *NullableV1ServiceAccount) Set(val *V1ServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ServiceAccount(val *V1ServiceAccount) *NullableV1ServiceAccount {
	return &NullableV1ServiceAccount{value: val, isSet: true}
}

func (v NullableV1ServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


