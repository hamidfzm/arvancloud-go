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

// checks if the V1CustomBuildStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CustomBuildStrategy{}

// V1CustomBuildStrategy CustomBuildStrategy defines input parameters specific to Custom build.
type V1CustomBuildStrategy struct {
	// buildAPIVersion is the requested API version for the Build object serialized and passed to the custom builder
	BuildAPIVersion *string `json:"buildAPIVersion,omitempty"`
	// env contains additional environment variables you want to pass into a builder container.
	Env []V1EnvVar `json:"env,omitempty"`
	// exposeDockerSocket will allow running Docker commands (and build Docker images) from inside the Docker container.
	ExposeDockerSocket *bool `json:"exposeDockerSocket,omitempty"`
	// forcePull describes if the controller should configure the build pod to always pull the images for the builder or only pull if it is not present locally
	ForcePull *bool `json:"forcePull,omitempty"`
	From V1ObjectReference `json:"from"`
	PullSecret *V1LocalObjectReference `json:"pullSecret,omitempty"`
	// secrets is a list of additional secrets that will be included in the build pod
	Secrets []V1SecretSpec `json:"secrets,omitempty"`
}

// NewV1CustomBuildStrategy instantiates a new V1CustomBuildStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CustomBuildStrategy(from V1ObjectReference) *V1CustomBuildStrategy {
	this := V1CustomBuildStrategy{}
	this.From = from
	return &this
}

// NewV1CustomBuildStrategyWithDefaults instantiates a new V1CustomBuildStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CustomBuildStrategyWithDefaults() *V1CustomBuildStrategy {
	this := V1CustomBuildStrategy{}
	return &this
}

// GetBuildAPIVersion returns the BuildAPIVersion field value if set, zero value otherwise.
func (o *V1CustomBuildStrategy) GetBuildAPIVersion() string {
	if o == nil || IsNil(o.BuildAPIVersion) {
		var ret string
		return ret
	}
	return *o.BuildAPIVersion
}

// GetBuildAPIVersionOk returns a tuple with the BuildAPIVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetBuildAPIVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildAPIVersion) {
		return nil, false
	}
	return o.BuildAPIVersion, true
}

// HasBuildAPIVersion returns a boolean if a field has been set.
func (o *V1CustomBuildStrategy) HasBuildAPIVersion() bool {
	if o != nil && !IsNil(o.BuildAPIVersion) {
		return true
	}

	return false
}

// SetBuildAPIVersion gets a reference to the given string and assigns it to the BuildAPIVersion field.
func (o *V1CustomBuildStrategy) SetBuildAPIVersion(v string) {
	o.BuildAPIVersion = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *V1CustomBuildStrategy) GetEnv() []V1EnvVar {
	if o == nil || IsNil(o.Env) {
		var ret []V1EnvVar
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetEnvOk() ([]V1EnvVar, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *V1CustomBuildStrategy) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []V1EnvVar and assigns it to the Env field.
func (o *V1CustomBuildStrategy) SetEnv(v []V1EnvVar) {
	o.Env = v
}

// GetExposeDockerSocket returns the ExposeDockerSocket field value if set, zero value otherwise.
func (o *V1CustomBuildStrategy) GetExposeDockerSocket() bool {
	if o == nil || IsNil(o.ExposeDockerSocket) {
		var ret bool
		return ret
	}
	return *o.ExposeDockerSocket
}

// GetExposeDockerSocketOk returns a tuple with the ExposeDockerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetExposeDockerSocketOk() (*bool, bool) {
	if o == nil || IsNil(o.ExposeDockerSocket) {
		return nil, false
	}
	return o.ExposeDockerSocket, true
}

// HasExposeDockerSocket returns a boolean if a field has been set.
func (o *V1CustomBuildStrategy) HasExposeDockerSocket() bool {
	if o != nil && !IsNil(o.ExposeDockerSocket) {
		return true
	}

	return false
}

// SetExposeDockerSocket gets a reference to the given bool and assigns it to the ExposeDockerSocket field.
func (o *V1CustomBuildStrategy) SetExposeDockerSocket(v bool) {
	o.ExposeDockerSocket = &v
}

// GetForcePull returns the ForcePull field value if set, zero value otherwise.
func (o *V1CustomBuildStrategy) GetForcePull() bool {
	if o == nil || IsNil(o.ForcePull) {
		var ret bool
		return ret
	}
	return *o.ForcePull
}

// GetForcePullOk returns a tuple with the ForcePull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetForcePullOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePull) {
		return nil, false
	}
	return o.ForcePull, true
}

// HasForcePull returns a boolean if a field has been set.
func (o *V1CustomBuildStrategy) HasForcePull() bool {
	if o != nil && !IsNil(o.ForcePull) {
		return true
	}

	return false
}

// SetForcePull gets a reference to the given bool and assigns it to the ForcePull field.
func (o *V1CustomBuildStrategy) SetForcePull(v bool) {
	o.ForcePull = &v
}

// GetFrom returns the From field value
func (o *V1CustomBuildStrategy) GetFrom() V1ObjectReference {
	if o == nil {
		var ret V1ObjectReference
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetFromOk() (*V1ObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *V1CustomBuildStrategy) SetFrom(v V1ObjectReference) {
	o.From = v
}

// GetPullSecret returns the PullSecret field value if set, zero value otherwise.
func (o *V1CustomBuildStrategy) GetPullSecret() V1LocalObjectReference {
	if o == nil || IsNil(o.PullSecret) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.PullSecret
}

// GetPullSecretOk returns a tuple with the PullSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetPullSecretOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.PullSecret) {
		return nil, false
	}
	return o.PullSecret, true
}

// HasPullSecret returns a boolean if a field has been set.
func (o *V1CustomBuildStrategy) HasPullSecret() bool {
	if o != nil && !IsNil(o.PullSecret) {
		return true
	}

	return false
}

// SetPullSecret gets a reference to the given V1LocalObjectReference and assigns it to the PullSecret field.
func (o *V1CustomBuildStrategy) SetPullSecret(v V1LocalObjectReference) {
	o.PullSecret = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *V1CustomBuildStrategy) GetSecrets() []V1SecretSpec {
	if o == nil || IsNil(o.Secrets) {
		var ret []V1SecretSpec
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CustomBuildStrategy) GetSecretsOk() ([]V1SecretSpec, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *V1CustomBuildStrategy) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []V1SecretSpec and assigns it to the Secrets field.
func (o *V1CustomBuildStrategy) SetSecrets(v []V1SecretSpec) {
	o.Secrets = v
}

func (o V1CustomBuildStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CustomBuildStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildAPIVersion) {
		toSerialize["buildAPIVersion"] = o.BuildAPIVersion
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.ExposeDockerSocket) {
		toSerialize["exposeDockerSocket"] = o.ExposeDockerSocket
	}
	if !IsNil(o.ForcePull) {
		toSerialize["forcePull"] = o.ForcePull
	}
	toSerialize["from"] = o.From
	if !IsNil(o.PullSecret) {
		toSerialize["pullSecret"] = o.PullSecret
	}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}

type NullableV1CustomBuildStrategy struct {
	value *V1CustomBuildStrategy
	isSet bool
}

func (v NullableV1CustomBuildStrategy) Get() *V1CustomBuildStrategy {
	return v.value
}

func (v *NullableV1CustomBuildStrategy) Set(val *V1CustomBuildStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CustomBuildStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CustomBuildStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CustomBuildStrategy(val *V1CustomBuildStrategy) *NullableV1CustomBuildStrategy {
	return &NullableV1CustomBuildStrategy{value: val, isSet: true}
}

func (v NullableV1CustomBuildStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CustomBuildStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


