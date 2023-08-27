# V1ContainerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerID** | Pointer to **string** | Container&#39;s ID in the format &#39;docker://&lt;container_id&gt;&#39;. | [optional] 
**Image** | **string** | The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images | 
**ImageID** | **string** | ImageID of the container&#39;s image. | 
**LastState** | Pointer to [**V1ContainerState**](V1ContainerState.md) |  | [optional] 
**Name** | **string** | This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated. | 
**Ready** | **bool** | Specifies whether the container has passed its readiness probe. | 
**RestartCount** | **int32** | The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC. | 
**State** | Pointer to [**V1ContainerState**](V1ContainerState.md) |  | [optional] 

## Methods

### NewV1ContainerStatus

`func NewV1ContainerStatus(image string, imageID string, name string, ready bool, restartCount int32, ) *V1ContainerStatus`

NewV1ContainerStatus instantiates a new V1ContainerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerStatusWithDefaults

`func NewV1ContainerStatusWithDefaults() *V1ContainerStatus`

NewV1ContainerStatusWithDefaults instantiates a new V1ContainerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerID

`func (o *V1ContainerStatus) GetContainerID() string`

GetContainerID returns the ContainerID field if non-nil, zero value otherwise.

### GetContainerIDOk

`func (o *V1ContainerStatus) GetContainerIDOk() (*string, bool)`

GetContainerIDOk returns a tuple with the ContainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerID

`func (o *V1ContainerStatus) SetContainerID(v string)`

SetContainerID sets ContainerID field to given value.

### HasContainerID

`func (o *V1ContainerStatus) HasContainerID() bool`

HasContainerID returns a boolean if a field has been set.

### GetImage

`func (o *V1ContainerStatus) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1ContainerStatus) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1ContainerStatus) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageID

`func (o *V1ContainerStatus) GetImageID() string`

GetImageID returns the ImageID field if non-nil, zero value otherwise.

### GetImageIDOk

`func (o *V1ContainerStatus) GetImageIDOk() (*string, bool)`

GetImageIDOk returns a tuple with the ImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageID

`func (o *V1ContainerStatus) SetImageID(v string)`

SetImageID sets ImageID field to given value.


### GetLastState

`func (o *V1ContainerStatus) GetLastState() V1ContainerState`

GetLastState returns the LastState field if non-nil, zero value otherwise.

### GetLastStateOk

`func (o *V1ContainerStatus) GetLastStateOk() (*V1ContainerState, bool)`

GetLastStateOk returns a tuple with the LastState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastState

`func (o *V1ContainerStatus) SetLastState(v V1ContainerState)`

SetLastState sets LastState field to given value.

### HasLastState

`func (o *V1ContainerStatus) HasLastState() bool`

HasLastState returns a boolean if a field has been set.

### GetName

`func (o *V1ContainerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ContainerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ContainerStatus) SetName(v string)`

SetName sets Name field to given value.


### GetReady

`func (o *V1ContainerStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *V1ContainerStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *V1ContainerStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetRestartCount

`func (o *V1ContainerStatus) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *V1ContainerStatus) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *V1ContainerStatus) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.


### GetState

`func (o *V1ContainerStatus) GetState() V1ContainerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V1ContainerStatus) GetStateOk() (*V1ContainerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V1ContainerStatus) SetState(v V1ContainerState)`

SetState sets State field to given value.

### HasState

`func (o *V1ContainerStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


