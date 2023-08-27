# V1RollingUpdateDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSurge** | Pointer to **string** | The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods. | [optional] 
**MaxUnavailable** | Pointer to **string** | The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods. | [optional] 

## Methods

### NewV1RollingUpdateDeployment

`func NewV1RollingUpdateDeployment() *V1RollingUpdateDeployment`

NewV1RollingUpdateDeployment instantiates a new V1RollingUpdateDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RollingUpdateDeploymentWithDefaults

`func NewV1RollingUpdateDeploymentWithDefaults() *V1RollingUpdateDeployment`

NewV1RollingUpdateDeploymentWithDefaults instantiates a new V1RollingUpdateDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSurge

`func (o *V1RollingUpdateDeployment) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *V1RollingUpdateDeployment) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *V1RollingUpdateDeployment) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *V1RollingUpdateDeployment) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *V1RollingUpdateDeployment) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *V1RollingUpdateDeployment) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *V1RollingUpdateDeployment) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *V1RollingUpdateDeployment) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


