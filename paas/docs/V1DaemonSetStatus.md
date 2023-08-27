# V1DaemonSetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollisionCount** | Pointer to **int32** | Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. | [optional] 
**Conditions** | Pointer to [**[]V1DaemonSetCondition**](V1DaemonSetCondition.md) | Represents the latest available observations of a DaemonSet&#39;s current state. | [optional] 
**CurrentNumberScheduled** | **int32** | The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ | 
**DesiredNumberScheduled** | **int32** | The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ | 
**NumberAvailable** | Pointer to **int32** | The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds) | [optional] 
**NumberMisscheduled** | **int32** | The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/ | 
**NumberReady** | **int32** | The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready. | 
**NumberUnavailable** | Pointer to **int32** | The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds) | [optional] 
**ObservedGeneration** | Pointer to **int64** | The most recent generation observed by the daemon set controller. | [optional] 
**UpdatedNumberScheduled** | Pointer to **int32** | The total number of nodes that are running updated daemon pod | [optional] 

## Methods

### NewV1DaemonSetStatus

`func NewV1DaemonSetStatus(currentNumberScheduled int32, desiredNumberScheduled int32, numberMisscheduled int32, numberReady int32, ) *V1DaemonSetStatus`

NewV1DaemonSetStatus instantiates a new V1DaemonSetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DaemonSetStatusWithDefaults

`func NewV1DaemonSetStatusWithDefaults() *V1DaemonSetStatus`

NewV1DaemonSetStatusWithDefaults instantiates a new V1DaemonSetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollisionCount

`func (o *V1DaemonSetStatus) GetCollisionCount() int32`

GetCollisionCount returns the CollisionCount field if non-nil, zero value otherwise.

### GetCollisionCountOk

`func (o *V1DaemonSetStatus) GetCollisionCountOk() (*int32, bool)`

GetCollisionCountOk returns a tuple with the CollisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisionCount

`func (o *V1DaemonSetStatus) SetCollisionCount(v int32)`

SetCollisionCount sets CollisionCount field to given value.

### HasCollisionCount

`func (o *V1DaemonSetStatus) HasCollisionCount() bool`

HasCollisionCount returns a boolean if a field has been set.

### GetConditions

`func (o *V1DaemonSetStatus) GetConditions() []V1DaemonSetCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1DaemonSetStatus) GetConditionsOk() (*[]V1DaemonSetCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1DaemonSetStatus) SetConditions(v []V1DaemonSetCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1DaemonSetStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentNumberScheduled

`func (o *V1DaemonSetStatus) GetCurrentNumberScheduled() int32`

GetCurrentNumberScheduled returns the CurrentNumberScheduled field if non-nil, zero value otherwise.

### GetCurrentNumberScheduledOk

`func (o *V1DaemonSetStatus) GetCurrentNumberScheduledOk() (*int32, bool)`

GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumberScheduled

`func (o *V1DaemonSetStatus) SetCurrentNumberScheduled(v int32)`

SetCurrentNumberScheduled sets CurrentNumberScheduled field to given value.


### GetDesiredNumberScheduled

`func (o *V1DaemonSetStatus) GetDesiredNumberScheduled() int32`

GetDesiredNumberScheduled returns the DesiredNumberScheduled field if non-nil, zero value otherwise.

### GetDesiredNumberScheduledOk

`func (o *V1DaemonSetStatus) GetDesiredNumberScheduledOk() (*int32, bool)`

GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNumberScheduled

`func (o *V1DaemonSetStatus) SetDesiredNumberScheduled(v int32)`

SetDesiredNumberScheduled sets DesiredNumberScheduled field to given value.


### GetNumberAvailable

`func (o *V1DaemonSetStatus) GetNumberAvailable() int32`

GetNumberAvailable returns the NumberAvailable field if non-nil, zero value otherwise.

### GetNumberAvailableOk

`func (o *V1DaemonSetStatus) GetNumberAvailableOk() (*int32, bool)`

GetNumberAvailableOk returns a tuple with the NumberAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAvailable

`func (o *V1DaemonSetStatus) SetNumberAvailable(v int32)`

SetNumberAvailable sets NumberAvailable field to given value.

### HasNumberAvailable

`func (o *V1DaemonSetStatus) HasNumberAvailable() bool`

HasNumberAvailable returns a boolean if a field has been set.

### GetNumberMisscheduled

`func (o *V1DaemonSetStatus) GetNumberMisscheduled() int32`

GetNumberMisscheduled returns the NumberMisscheduled field if non-nil, zero value otherwise.

### GetNumberMisscheduledOk

`func (o *V1DaemonSetStatus) GetNumberMisscheduledOk() (*int32, bool)`

GetNumberMisscheduledOk returns a tuple with the NumberMisscheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberMisscheduled

`func (o *V1DaemonSetStatus) SetNumberMisscheduled(v int32)`

SetNumberMisscheduled sets NumberMisscheduled field to given value.


### GetNumberReady

`func (o *V1DaemonSetStatus) GetNumberReady() int32`

GetNumberReady returns the NumberReady field if non-nil, zero value otherwise.

### GetNumberReadyOk

`func (o *V1DaemonSetStatus) GetNumberReadyOk() (*int32, bool)`

GetNumberReadyOk returns a tuple with the NumberReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberReady

`func (o *V1DaemonSetStatus) SetNumberReady(v int32)`

SetNumberReady sets NumberReady field to given value.


### GetNumberUnavailable

`func (o *V1DaemonSetStatus) GetNumberUnavailable() int32`

GetNumberUnavailable returns the NumberUnavailable field if non-nil, zero value otherwise.

### GetNumberUnavailableOk

`func (o *V1DaemonSetStatus) GetNumberUnavailableOk() (*int32, bool)`

GetNumberUnavailableOk returns a tuple with the NumberUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberUnavailable

`func (o *V1DaemonSetStatus) SetNumberUnavailable(v int32)`

SetNumberUnavailable sets NumberUnavailable field to given value.

### HasNumberUnavailable

`func (o *V1DaemonSetStatus) HasNumberUnavailable() bool`

HasNumberUnavailable returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *V1DaemonSetStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *V1DaemonSetStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *V1DaemonSetStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *V1DaemonSetStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetUpdatedNumberScheduled

`func (o *V1DaemonSetStatus) GetUpdatedNumberScheduled() int32`

GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field if non-nil, zero value otherwise.

### GetUpdatedNumberScheduledOk

`func (o *V1DaemonSetStatus) GetUpdatedNumberScheduledOk() (*int32, bool)`

GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedNumberScheduled

`func (o *V1DaemonSetStatus) SetUpdatedNumberScheduled(v int32)`

SetUpdatedNumberScheduled sets UpdatedNumberScheduled field to given value.

### HasUpdatedNumberScheduled

`func (o *V1DaemonSetStatus) HasUpdatedNumberScheduled() bool`

HasUpdatedNumberScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


