# V1DeploymentTriggerImageChangeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automatic** | Pointer to **bool** | Automatic means that the detection of a new tag value should result in an image update inside the pod template. | [optional] 
**ContainerNames** | Pointer to **[]string** | ContainerNames is used to restrict tag updates to the specified set of container names in a pod. If multiple triggers point to the same containers, the resulting behavior is undefined. Future API versions will make this a validation error. If ContainerNames does not point to a valid container, the trigger will be ignored. Future API versions will make this a validation error. | [optional] 
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**LastTriggeredImage** | Pointer to **string** | LastTriggeredImage is the last image to be triggered. | [optional] 

## Methods

### NewV1DeploymentTriggerImageChangeParams

`func NewV1DeploymentTriggerImageChangeParams(from V1ObjectReference, ) *V1DeploymentTriggerImageChangeParams`

NewV1DeploymentTriggerImageChangeParams instantiates a new V1DeploymentTriggerImageChangeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentTriggerImageChangeParamsWithDefaults

`func NewV1DeploymentTriggerImageChangeParamsWithDefaults() *V1DeploymentTriggerImageChangeParams`

NewV1DeploymentTriggerImageChangeParamsWithDefaults instantiates a new V1DeploymentTriggerImageChangeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatic

`func (o *V1DeploymentTriggerImageChangeParams) GetAutomatic() bool`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *V1DeploymentTriggerImageChangeParams) GetAutomaticOk() (*bool, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *V1DeploymentTriggerImageChangeParams) SetAutomatic(v bool)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *V1DeploymentTriggerImageChangeParams) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetContainerNames

`func (o *V1DeploymentTriggerImageChangeParams) GetContainerNames() []string`

GetContainerNames returns the ContainerNames field if non-nil, zero value otherwise.

### GetContainerNamesOk

`func (o *V1DeploymentTriggerImageChangeParams) GetContainerNamesOk() (*[]string, bool)`

GetContainerNamesOk returns a tuple with the ContainerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNames

`func (o *V1DeploymentTriggerImageChangeParams) SetContainerNames(v []string)`

SetContainerNames sets ContainerNames field to given value.

### HasContainerNames

`func (o *V1DeploymentTriggerImageChangeParams) HasContainerNames() bool`

HasContainerNames returns a boolean if a field has been set.

### GetFrom

`func (o *V1DeploymentTriggerImageChangeParams) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1DeploymentTriggerImageChangeParams) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1DeploymentTriggerImageChangeParams) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetLastTriggeredImage

`func (o *V1DeploymentTriggerImageChangeParams) GetLastTriggeredImage() string`

GetLastTriggeredImage returns the LastTriggeredImage field if non-nil, zero value otherwise.

### GetLastTriggeredImageOk

`func (o *V1DeploymentTriggerImageChangeParams) GetLastTriggeredImageOk() (*string, bool)`

GetLastTriggeredImageOk returns a tuple with the LastTriggeredImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredImage

`func (o *V1DeploymentTriggerImageChangeParams) SetLastTriggeredImage(v string)`

SetLastTriggeredImage sets LastTriggeredImage field to given value.

### HasLastTriggeredImage

`func (o *V1DeploymentTriggerImageChangeParams) HasLastTriggeredImage() bool`

HasLastTriggeredImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


