# V1DeploymentCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageTrigger** | Pointer to [**V1DeploymentCauseImageTrigger**](V1DeploymentCauseImageTrigger.md) |  | [optional] 
**Type** | **string** | Type of the trigger that resulted in the creation of a new deployment | 

## Methods

### NewV1DeploymentCause

`func NewV1DeploymentCause(type_ string, ) *V1DeploymentCause`

NewV1DeploymentCause instantiates a new V1DeploymentCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentCauseWithDefaults

`func NewV1DeploymentCauseWithDefaults() *V1DeploymentCause`

NewV1DeploymentCauseWithDefaults instantiates a new V1DeploymentCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageTrigger

`func (o *V1DeploymentCause) GetImageTrigger() V1DeploymentCauseImageTrigger`

GetImageTrigger returns the ImageTrigger field if non-nil, zero value otherwise.

### GetImageTriggerOk

`func (o *V1DeploymentCause) GetImageTriggerOk() (*V1DeploymentCauseImageTrigger, bool)`

GetImageTriggerOk returns a tuple with the ImageTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTrigger

`func (o *V1DeploymentCause) SetImageTrigger(v V1DeploymentCauseImageTrigger)`

SetImageTrigger sets ImageTrigger field to given value.

### HasImageTrigger

`func (o *V1DeploymentCause) HasImageTrigger() bool`

HasImageTrigger returns a boolean if a field has been set.

### GetType

`func (o *V1DeploymentCause) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DeploymentCause) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DeploymentCause) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


