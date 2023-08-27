# V1DeploymentConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | **int32** | AvailableReplicas is the total number of available pods targeted by this deployment config. | 
**Conditions** | Pointer to [**[]V1DeploymentCondition**](V1DeploymentCondition.md) | Conditions represents the latest available observations of a deployment config&#39;s current state. | [optional] 
**Details** | Pointer to [**V1DeploymentDetails**](V1DeploymentDetails.md) |  | [optional] 
**LatestVersion** | **int64** | LatestVersion is used to determine whether the current deployment associated with a deployment config is out of sync. | 
**ObservedGeneration** | **int64** | ObservedGeneration is the most recent generation observed by the deployment config controller. | 
**ReadyReplicas** | Pointer to **int32** | Total number of ready pods targeted by this deployment. | [optional] 
**Replicas** | **int32** | Replicas is the total number of pods targeted by this deployment config. | 
**UnavailableReplicas** | **int32** | UnavailableReplicas is the total number of unavailable pods targeted by this deployment config. | 
**UpdatedReplicas** | **int32** | UpdatedReplicas is the total number of non-terminated pods targeted by this deployment config that have the desired template spec. | 

## Methods

### NewV1DeploymentConfigStatus

`func NewV1DeploymentConfigStatus(availableReplicas int32, latestVersion int64, observedGeneration int64, replicas int32, unavailableReplicas int32, updatedReplicas int32, ) *V1DeploymentConfigStatus`

NewV1DeploymentConfigStatus instantiates a new V1DeploymentConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentConfigStatusWithDefaults

`func NewV1DeploymentConfigStatusWithDefaults() *V1DeploymentConfigStatus`

NewV1DeploymentConfigStatusWithDefaults instantiates a new V1DeploymentConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *V1DeploymentConfigStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *V1DeploymentConfigStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *V1DeploymentConfigStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.


### GetConditions

`func (o *V1DeploymentConfigStatus) GetConditions() []V1DeploymentCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1DeploymentConfigStatus) GetConditionsOk() (*[]V1DeploymentCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1DeploymentConfigStatus) SetConditions(v []V1DeploymentCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1DeploymentConfigStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDetails

`func (o *V1DeploymentConfigStatus) GetDetails() V1DeploymentDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *V1DeploymentConfigStatus) GetDetailsOk() (*V1DeploymentDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *V1DeploymentConfigStatus) SetDetails(v V1DeploymentDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *V1DeploymentConfigStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLatestVersion

`func (o *V1DeploymentConfigStatus) GetLatestVersion() int64`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *V1DeploymentConfigStatus) GetLatestVersionOk() (*int64, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *V1DeploymentConfigStatus) SetLatestVersion(v int64)`

SetLatestVersion sets LatestVersion field to given value.


### GetObservedGeneration

`func (o *V1DeploymentConfigStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *V1DeploymentConfigStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *V1DeploymentConfigStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.


### GetReadyReplicas

`func (o *V1DeploymentConfigStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *V1DeploymentConfigStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *V1DeploymentConfigStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *V1DeploymentConfigStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *V1DeploymentConfigStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1DeploymentConfigStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1DeploymentConfigStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetUnavailableReplicas

`func (o *V1DeploymentConfigStatus) GetUnavailableReplicas() int32`

GetUnavailableReplicas returns the UnavailableReplicas field if non-nil, zero value otherwise.

### GetUnavailableReplicasOk

`func (o *V1DeploymentConfigStatus) GetUnavailableReplicasOk() (*int32, bool)`

GetUnavailableReplicasOk returns a tuple with the UnavailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableReplicas

`func (o *V1DeploymentConfigStatus) SetUnavailableReplicas(v int32)`

SetUnavailableReplicas sets UnavailableReplicas field to given value.


### GetUpdatedReplicas

`func (o *V1DeploymentConfigStatus) GetUpdatedReplicas() int32`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *V1DeploymentConfigStatus) GetUpdatedReplicasOk() (*int32, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *V1DeploymentConfigStatus) SetUpdatedReplicas(v int32)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


