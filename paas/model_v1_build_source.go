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

// checks if the V1BuildSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BuildSource{}

// V1BuildSource BuildSource is the SCM used for the build.
type V1BuildSource struct {
	Binary *V1BinaryBuildSource `json:"binary,omitempty"`
	// configMaps represents a list of configMaps and their destinations that will be used for the build.
	ConfigMaps []V1ConfigMapBuildSource `json:"configMaps,omitempty"`
	// contextDir specifies the sub-directory where the source code for the application exists. This allows to have buildable sources in directory other than root of repository.
	ContextDir *string `json:"contextDir,omitempty"`
	// dockerfile is the raw contents of a Dockerfile which should be built. When this option is specified, the FROM may be modified based on your strategy base image and additional ENV stanzas from your strategy environment will be added after the FROM, but before the rest of your Dockerfile stanzas. The Dockerfile source type may be used with other options like git - in those cases the Git repo will have any innate Dockerfile replaced in the context dir.
	Dockerfile *string `json:"dockerfile,omitempty"`
	Git *V1GitBuildSource `json:"git,omitempty"`
	// images describes a set of images to be used to provide source for the build
	Images []V1ImageSource `json:"images,omitempty"`
	// secrets represents a list of secrets and their destinations that will be used only for the build.
	Secrets []V1SecretBuildSource `json:"secrets,omitempty"`
	SourceSecret *V1LocalObjectReference `json:"sourceSecret,omitempty"`
	// type of build input to accept
	Type string `json:"type"`
}

// NewV1BuildSource instantiates a new V1BuildSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BuildSource(type_ string) *V1BuildSource {
	this := V1BuildSource{}
	this.Type = type_
	return &this
}

// NewV1BuildSourceWithDefaults instantiates a new V1BuildSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BuildSourceWithDefaults() *V1BuildSource {
	this := V1BuildSource{}
	return &this
}

// GetBinary returns the Binary field value if set, zero value otherwise.
func (o *V1BuildSource) GetBinary() V1BinaryBuildSource {
	if o == nil || IsNil(o.Binary) {
		var ret V1BinaryBuildSource
		return ret
	}
	return *o.Binary
}

// GetBinaryOk returns a tuple with the Binary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetBinaryOk() (*V1BinaryBuildSource, bool) {
	if o == nil || IsNil(o.Binary) {
		return nil, false
	}
	return o.Binary, true
}

// HasBinary returns a boolean if a field has been set.
func (o *V1BuildSource) HasBinary() bool {
	if o != nil && !IsNil(o.Binary) {
		return true
	}

	return false
}

// SetBinary gets a reference to the given V1BinaryBuildSource and assigns it to the Binary field.
func (o *V1BuildSource) SetBinary(v V1BinaryBuildSource) {
	o.Binary = &v
}

// GetConfigMaps returns the ConfigMaps field value if set, zero value otherwise.
func (o *V1BuildSource) GetConfigMaps() []V1ConfigMapBuildSource {
	if o == nil || IsNil(o.ConfigMaps) {
		var ret []V1ConfigMapBuildSource
		return ret
	}
	return o.ConfigMaps
}

// GetConfigMapsOk returns a tuple with the ConfigMaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetConfigMapsOk() ([]V1ConfigMapBuildSource, bool) {
	if o == nil || IsNil(o.ConfigMaps) {
		return nil, false
	}
	return o.ConfigMaps, true
}

// HasConfigMaps returns a boolean if a field has been set.
func (o *V1BuildSource) HasConfigMaps() bool {
	if o != nil && !IsNil(o.ConfigMaps) {
		return true
	}

	return false
}

// SetConfigMaps gets a reference to the given []V1ConfigMapBuildSource and assigns it to the ConfigMaps field.
func (o *V1BuildSource) SetConfigMaps(v []V1ConfigMapBuildSource) {
	o.ConfigMaps = v
}

// GetContextDir returns the ContextDir field value if set, zero value otherwise.
func (o *V1BuildSource) GetContextDir() string {
	if o == nil || IsNil(o.ContextDir) {
		var ret string
		return ret
	}
	return *o.ContextDir
}

// GetContextDirOk returns a tuple with the ContextDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetContextDirOk() (*string, bool) {
	if o == nil || IsNil(o.ContextDir) {
		return nil, false
	}
	return o.ContextDir, true
}

// HasContextDir returns a boolean if a field has been set.
func (o *V1BuildSource) HasContextDir() bool {
	if o != nil && !IsNil(o.ContextDir) {
		return true
	}

	return false
}

// SetContextDir gets a reference to the given string and assigns it to the ContextDir field.
func (o *V1BuildSource) SetContextDir(v string) {
	o.ContextDir = &v
}

// GetDockerfile returns the Dockerfile field value if set, zero value otherwise.
func (o *V1BuildSource) GetDockerfile() string {
	if o == nil || IsNil(o.Dockerfile) {
		var ret string
		return ret
	}
	return *o.Dockerfile
}

// GetDockerfileOk returns a tuple with the Dockerfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetDockerfileOk() (*string, bool) {
	if o == nil || IsNil(o.Dockerfile) {
		return nil, false
	}
	return o.Dockerfile, true
}

// HasDockerfile returns a boolean if a field has been set.
func (o *V1BuildSource) HasDockerfile() bool {
	if o != nil && !IsNil(o.Dockerfile) {
		return true
	}

	return false
}

// SetDockerfile gets a reference to the given string and assigns it to the Dockerfile field.
func (o *V1BuildSource) SetDockerfile(v string) {
	o.Dockerfile = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *V1BuildSource) GetGit() V1GitBuildSource {
	if o == nil || IsNil(o.Git) {
		var ret V1GitBuildSource
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetGitOk() (*V1GitBuildSource, bool) {
	if o == nil || IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *V1BuildSource) HasGit() bool {
	if o != nil && !IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given V1GitBuildSource and assigns it to the Git field.
func (o *V1BuildSource) SetGit(v V1GitBuildSource) {
	o.Git = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *V1BuildSource) GetImages() []V1ImageSource {
	if o == nil || IsNil(o.Images) {
		var ret []V1ImageSource
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetImagesOk() ([]V1ImageSource, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *V1BuildSource) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []V1ImageSource and assigns it to the Images field.
func (o *V1BuildSource) SetImages(v []V1ImageSource) {
	o.Images = v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *V1BuildSource) GetSecrets() []V1SecretBuildSource {
	if o == nil || IsNil(o.Secrets) {
		var ret []V1SecretBuildSource
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetSecretsOk() ([]V1SecretBuildSource, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *V1BuildSource) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []V1SecretBuildSource and assigns it to the Secrets field.
func (o *V1BuildSource) SetSecrets(v []V1SecretBuildSource) {
	o.Secrets = v
}

// GetSourceSecret returns the SourceSecret field value if set, zero value otherwise.
func (o *V1BuildSource) GetSourceSecret() V1LocalObjectReference {
	if o == nil || IsNil(o.SourceSecret) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.SourceSecret
}

// GetSourceSecretOk returns a tuple with the SourceSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetSourceSecretOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.SourceSecret) {
		return nil, false
	}
	return o.SourceSecret, true
}

// HasSourceSecret returns a boolean if a field has been set.
func (o *V1BuildSource) HasSourceSecret() bool {
	if o != nil && !IsNil(o.SourceSecret) {
		return true
	}

	return false
}

// SetSourceSecret gets a reference to the given V1LocalObjectReference and assigns it to the SourceSecret field.
func (o *V1BuildSource) SetSourceSecret(v V1LocalObjectReference) {
	o.SourceSecret = &v
}

// GetType returns the Type field value
func (o *V1BuildSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1BuildSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1BuildSource) SetType(v string) {
	o.Type = v
}

func (o V1BuildSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BuildSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Binary) {
		toSerialize["binary"] = o.Binary
	}
	if !IsNil(o.ConfigMaps) {
		toSerialize["configMaps"] = o.ConfigMaps
	}
	if !IsNil(o.ContextDir) {
		toSerialize["contextDir"] = o.ContextDir
	}
	if !IsNil(o.Dockerfile) {
		toSerialize["dockerfile"] = o.Dockerfile
	}
	if !IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	if !IsNil(o.SourceSecret) {
		toSerialize["sourceSecret"] = o.SourceSecret
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableV1BuildSource struct {
	value *V1BuildSource
	isSet bool
}

func (v NullableV1BuildSource) Get() *V1BuildSource {
	return v.value
}

func (v *NullableV1BuildSource) Set(val *V1BuildSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BuildSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BuildSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BuildSource(val *V1BuildSource) *NullableV1BuildSource {
	return &NullableV1BuildSource{value: val, isSet: true}
}

func (v NullableV1BuildSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BuildSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


