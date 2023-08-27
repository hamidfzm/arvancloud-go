# V1DeploymentTriggerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageChangeParams** | Pointer to [**V1DeploymentTriggerImageChangeParams**](V1DeploymentTriggerImageChangeParams.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the trigger | [optional] 

## Methods

### NewV1DeploymentTriggerPolicy

`func NewV1DeploymentTriggerPolicy() *V1DeploymentTriggerPolicy`

NewV1DeploymentTriggerPolicy instantiates a new V1DeploymentTriggerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentTriggerPolicyWithDefaults

`func NewV1DeploymentTriggerPolicyWithDefaults() *V1DeploymentTriggerPolicy`

NewV1DeploymentTriggerPolicyWithDefaults instantiates a new V1DeploymentTriggerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageChangeParams

`func (o *V1DeploymentTriggerPolicy) GetImageChangeParams() V1DeploymentTriggerImageChangeParams`

GetImageChangeParams returns the ImageChangeParams field if non-nil, zero value otherwise.

### GetImageChangeParamsOk

`func (o *V1DeploymentTriggerPolicy) GetImageChangeParamsOk() (*V1DeploymentTriggerImageChangeParams, bool)`

GetImageChangeParamsOk returns a tuple with the ImageChangeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageChangeParams

`func (o *V1DeploymentTriggerPolicy) SetImageChangeParams(v V1DeploymentTriggerImageChangeParams)`

SetImageChangeParams sets ImageChangeParams field to given value.

### HasImageChangeParams

`func (o *V1DeploymentTriggerPolicy) HasImageChangeParams() bool`

HasImageChangeParams returns a boolean if a field has been set.

### GetType

`func (o *V1DeploymentTriggerPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DeploymentTriggerPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DeploymentTriggerPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1DeploymentTriggerPolicy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


