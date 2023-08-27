# V1LifecycleHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecNewPod** | Pointer to [**V1ExecNewPodHook**](V1ExecNewPodHook.md) |  | [optional] 
**FailurePolicy** | **string** | FailurePolicy specifies what action to take if the hook fails. | 
**TagImages** | Pointer to [**[]V1TagImageHook**](V1TagImageHook.md) | TagImages instructs the deployer to tag the current image referenced under a container onto an image stream tag. | [optional] 

## Methods

### NewV1LifecycleHook

`func NewV1LifecycleHook(failurePolicy string, ) *V1LifecycleHook`

NewV1LifecycleHook instantiates a new V1LifecycleHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LifecycleHookWithDefaults

`func NewV1LifecycleHookWithDefaults() *V1LifecycleHook`

NewV1LifecycleHookWithDefaults instantiates a new V1LifecycleHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecNewPod

`func (o *V1LifecycleHook) GetExecNewPod() V1ExecNewPodHook`

GetExecNewPod returns the ExecNewPod field if non-nil, zero value otherwise.

### GetExecNewPodOk

`func (o *V1LifecycleHook) GetExecNewPodOk() (*V1ExecNewPodHook, bool)`

GetExecNewPodOk returns a tuple with the ExecNewPod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecNewPod

`func (o *V1LifecycleHook) SetExecNewPod(v V1ExecNewPodHook)`

SetExecNewPod sets ExecNewPod field to given value.

### HasExecNewPod

`func (o *V1LifecycleHook) HasExecNewPod() bool`

HasExecNewPod returns a boolean if a field has been set.

### GetFailurePolicy

`func (o *V1LifecycleHook) GetFailurePolicy() string`

GetFailurePolicy returns the FailurePolicy field if non-nil, zero value otherwise.

### GetFailurePolicyOk

`func (o *V1LifecycleHook) GetFailurePolicyOk() (*string, bool)`

GetFailurePolicyOk returns a tuple with the FailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailurePolicy

`func (o *V1LifecycleHook) SetFailurePolicy(v string)`

SetFailurePolicy sets FailurePolicy field to given value.


### GetTagImages

`func (o *V1LifecycleHook) GetTagImages() []V1TagImageHook`

GetTagImages returns the TagImages field if non-nil, zero value otherwise.

### GetTagImagesOk

`func (o *V1LifecycleHook) GetTagImagesOk() (*[]V1TagImageHook, bool)`

GetTagImagesOk returns a tuple with the TagImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagImages

`func (o *V1LifecycleHook) SetTagImages(v []V1TagImageHook)`

SetTagImages sets TagImages field to given value.

### HasTagImages

`func (o *V1LifecycleHook) HasTagImages() bool`

HasTagImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


