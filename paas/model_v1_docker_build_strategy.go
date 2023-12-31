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

// checks if the V1DockerBuildStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DockerBuildStrategy{}

// V1DockerBuildStrategy DockerBuildStrategy defines input parameters specific to Docker build.
type V1DockerBuildStrategy struct {
	// buildArgs contains build arguments that will be resolved in the Dockerfile.  See https://docs.docker.com/engine/reference/builder/#/arg for more details.
	BuildArgs []V1EnvVar `json:"buildArgs,omitempty"`
	// dockerfilePath is the path of the Dockerfile that will be used to build the Docker image, relative to the root of the context (contextDir).
	DockerfilePath *string `json:"dockerfilePath,omitempty"`
	// env contains additional environment variables you want to pass into a builder container.
	Env []V1EnvVar `json:"env,omitempty"`
	// forcePull describes if the builder should pull the images from registry prior to building.
	ForcePull *bool `json:"forcePull,omitempty"`
	From *V1ObjectReference `json:"from,omitempty"`
	ImageOptimizationPolicy *V1ImageOptimizationPolicy `json:"imageOptimizationPolicy,omitempty"`
	// noCache if set to true indicates that the docker build must be executed with the --no-cache=true flag
	NoCache *bool `json:"noCache,omitempty"`
	PullSecret *V1LocalObjectReference `json:"pullSecret,omitempty"`
}

// NewV1DockerBuildStrategy instantiates a new V1DockerBuildStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DockerBuildStrategy() *V1DockerBuildStrategy {
	this := V1DockerBuildStrategy{}
	return &this
}

// NewV1DockerBuildStrategyWithDefaults instantiates a new V1DockerBuildStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DockerBuildStrategyWithDefaults() *V1DockerBuildStrategy {
	this := V1DockerBuildStrategy{}
	return &this
}

// GetBuildArgs returns the BuildArgs field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetBuildArgs() []V1EnvVar {
	if o == nil || IsNil(o.BuildArgs) {
		var ret []V1EnvVar
		return ret
	}
	return o.BuildArgs
}

// GetBuildArgsOk returns a tuple with the BuildArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetBuildArgsOk() ([]V1EnvVar, bool) {
	if o == nil || IsNil(o.BuildArgs) {
		return nil, false
	}
	return o.BuildArgs, true
}

// HasBuildArgs returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasBuildArgs() bool {
	if o != nil && !IsNil(o.BuildArgs) {
		return true
	}

	return false
}

// SetBuildArgs gets a reference to the given []V1EnvVar and assigns it to the BuildArgs field.
func (o *V1DockerBuildStrategy) SetBuildArgs(v []V1EnvVar) {
	o.BuildArgs = v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetDockerfilePath() string {
	if o == nil || IsNil(o.DockerfilePath) {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetDockerfilePathOk() (*string, bool) {
	if o == nil || IsNil(o.DockerfilePath) {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasDockerfilePath() bool {
	if o != nil && !IsNil(o.DockerfilePath) {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *V1DockerBuildStrategy) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetEnv() []V1EnvVar {
	if o == nil || IsNil(o.Env) {
		var ret []V1EnvVar
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetEnvOk() ([]V1EnvVar, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []V1EnvVar and assigns it to the Env field.
func (o *V1DockerBuildStrategy) SetEnv(v []V1EnvVar) {
	o.Env = v
}

// GetForcePull returns the ForcePull field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetForcePull() bool {
	if o == nil || IsNil(o.ForcePull) {
		var ret bool
		return ret
	}
	return *o.ForcePull
}

// GetForcePullOk returns a tuple with the ForcePull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetForcePullOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePull) {
		return nil, false
	}
	return o.ForcePull, true
}

// HasForcePull returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasForcePull() bool {
	if o != nil && !IsNil(o.ForcePull) {
		return true
	}

	return false
}

// SetForcePull gets a reference to the given bool and assigns it to the ForcePull field.
func (o *V1DockerBuildStrategy) SetForcePull(v bool) {
	o.ForcePull = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetFrom() V1ObjectReference {
	if o == nil || IsNil(o.From) {
		var ret V1ObjectReference
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetFromOk() (*V1ObjectReference, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given V1ObjectReference and assigns it to the From field.
func (o *V1DockerBuildStrategy) SetFrom(v V1ObjectReference) {
	o.From = &v
}

// GetImageOptimizationPolicy returns the ImageOptimizationPolicy field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetImageOptimizationPolicy() V1ImageOptimizationPolicy {
	if o == nil || IsNil(o.ImageOptimizationPolicy) {
		var ret V1ImageOptimizationPolicy
		return ret
	}
	return *o.ImageOptimizationPolicy
}

// GetImageOptimizationPolicyOk returns a tuple with the ImageOptimizationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetImageOptimizationPolicyOk() (*V1ImageOptimizationPolicy, bool) {
	if o == nil || IsNil(o.ImageOptimizationPolicy) {
		return nil, false
	}
	return o.ImageOptimizationPolicy, true
}

// HasImageOptimizationPolicy returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasImageOptimizationPolicy() bool {
	if o != nil && !IsNil(o.ImageOptimizationPolicy) {
		return true
	}

	return false
}

// SetImageOptimizationPolicy gets a reference to the given V1ImageOptimizationPolicy and assigns it to the ImageOptimizationPolicy field.
func (o *V1DockerBuildStrategy) SetImageOptimizationPolicy(v V1ImageOptimizationPolicy) {
	o.ImageOptimizationPolicy = &v
}

// GetNoCache returns the NoCache field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetNoCache() bool {
	if o == nil || IsNil(o.NoCache) {
		var ret bool
		return ret
	}
	return *o.NoCache
}

// GetNoCacheOk returns a tuple with the NoCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetNoCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.NoCache) {
		return nil, false
	}
	return o.NoCache, true
}

// HasNoCache returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasNoCache() bool {
	if o != nil && !IsNil(o.NoCache) {
		return true
	}

	return false
}

// SetNoCache gets a reference to the given bool and assigns it to the NoCache field.
func (o *V1DockerBuildStrategy) SetNoCache(v bool) {
	o.NoCache = &v
}

// GetPullSecret returns the PullSecret field value if set, zero value otherwise.
func (o *V1DockerBuildStrategy) GetPullSecret() V1LocalObjectReference {
	if o == nil || IsNil(o.PullSecret) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.PullSecret
}

// GetPullSecretOk returns a tuple with the PullSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DockerBuildStrategy) GetPullSecretOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.PullSecret) {
		return nil, false
	}
	return o.PullSecret, true
}

// HasPullSecret returns a boolean if a field has been set.
func (o *V1DockerBuildStrategy) HasPullSecret() bool {
	if o != nil && !IsNil(o.PullSecret) {
		return true
	}

	return false
}

// SetPullSecret gets a reference to the given V1LocalObjectReference and assigns it to the PullSecret field.
func (o *V1DockerBuildStrategy) SetPullSecret(v V1LocalObjectReference) {
	o.PullSecret = &v
}

func (o V1DockerBuildStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DockerBuildStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildArgs) {
		toSerialize["buildArgs"] = o.BuildArgs
	}
	if !IsNil(o.DockerfilePath) {
		toSerialize["dockerfilePath"] = o.DockerfilePath
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.ForcePull) {
		toSerialize["forcePull"] = o.ForcePull
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.ImageOptimizationPolicy) {
		toSerialize["imageOptimizationPolicy"] = o.ImageOptimizationPolicy
	}
	if !IsNil(o.NoCache) {
		toSerialize["noCache"] = o.NoCache
	}
	if !IsNil(o.PullSecret) {
		toSerialize["pullSecret"] = o.PullSecret
	}
	return toSerialize, nil
}

type NullableV1DockerBuildStrategy struct {
	value *V1DockerBuildStrategy
	isSet bool
}

func (v NullableV1DockerBuildStrategy) Get() *V1DockerBuildStrategy {
	return v.value
}

func (v *NullableV1DockerBuildStrategy) Set(val *V1DockerBuildStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DockerBuildStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DockerBuildStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DockerBuildStrategy(val *V1DockerBuildStrategy) *NullableV1DockerBuildStrategy {
	return &NullableV1DockerBuildStrategy{value: val, isSet: true}
}

func (v NullableV1DockerBuildStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DockerBuildStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


