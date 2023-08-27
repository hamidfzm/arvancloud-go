# V1DeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**ExcludeTriggers** | Pointer to [**[]V1DeploymentTriggerType**](V1DeploymentTriggerType.md) | ExcludeTriggers instructs the instantiator to avoid processing the specified triggers. This field overrides the triggers from latest and allows clients to control specific logic. This field is ignored if not specified. | [optional] 
**Force** | **bool** | Force will try to force a new deployment to run. If the deployment config is paused, then setting this to true will return an Invalid error. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Latest** | **bool** | Latest will update the deployment config with the latest state from all triggers. | 
**Name** | **string** | Name of the deployment config for requesting a new deployment. | 

## Methods

### NewV1DeploymentRequest

`func NewV1DeploymentRequest(force bool, latest bool, name string, ) *V1DeploymentRequest`

NewV1DeploymentRequest instantiates a new V1DeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentRequestWithDefaults

`func NewV1DeploymentRequestWithDefaults() *V1DeploymentRequest`

NewV1DeploymentRequestWithDefaults instantiates a new V1DeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1DeploymentRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1DeploymentRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1DeploymentRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1DeploymentRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetExcludeTriggers

`func (o *V1DeploymentRequest) GetExcludeTriggers() []V1DeploymentTriggerType`

GetExcludeTriggers returns the ExcludeTriggers field if non-nil, zero value otherwise.

### GetExcludeTriggersOk

`func (o *V1DeploymentRequest) GetExcludeTriggersOk() (*[]V1DeploymentTriggerType, bool)`

GetExcludeTriggersOk returns a tuple with the ExcludeTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTriggers

`func (o *V1DeploymentRequest) SetExcludeTriggers(v []V1DeploymentTriggerType)`

SetExcludeTriggers sets ExcludeTriggers field to given value.

### HasExcludeTriggers

`func (o *V1DeploymentRequest) HasExcludeTriggers() bool`

HasExcludeTriggers returns a boolean if a field has been set.

### GetForce

`func (o *V1DeploymentRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *V1DeploymentRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *V1DeploymentRequest) SetForce(v bool)`

SetForce sets Force field to given value.


### GetKind

`func (o *V1DeploymentRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1DeploymentRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1DeploymentRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1DeploymentRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLatest

`func (o *V1DeploymentRequest) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *V1DeploymentRequest) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *V1DeploymentRequest) SetLatest(v bool)`

SetLatest sets Latest field to given value.


### GetName

`func (o *V1DeploymentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1DeploymentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1DeploymentRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


