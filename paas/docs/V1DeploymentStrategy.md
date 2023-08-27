# V1DeploymentStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpdate** | Pointer to [**V1RollingUpdateDeployment**](V1RollingUpdateDeployment.md) |  | [optional] 
**Type** | Pointer to **string** | Type of deployment. Can be \&quot;Recreate\&quot; or \&quot;RollingUpdate\&quot;. Default is RollingUpdate. | [optional] 

## Methods

### NewV1DeploymentStrategy

`func NewV1DeploymentStrategy() *V1DeploymentStrategy`

NewV1DeploymentStrategy instantiates a new V1DeploymentStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentStrategyWithDefaults

`func NewV1DeploymentStrategyWithDefaults() *V1DeploymentStrategy`

NewV1DeploymentStrategyWithDefaults instantiates a new V1DeploymentStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpdate

`func (o *V1DeploymentStrategy) GetRollingUpdate() V1RollingUpdateDeployment`

GetRollingUpdate returns the RollingUpdate field if non-nil, zero value otherwise.

### GetRollingUpdateOk

`func (o *V1DeploymentStrategy) GetRollingUpdateOk() (*V1RollingUpdateDeployment, bool)`

GetRollingUpdateOk returns a tuple with the RollingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpdate

`func (o *V1DeploymentStrategy) SetRollingUpdate(v V1RollingUpdateDeployment)`

SetRollingUpdate sets RollingUpdate field to given value.

### HasRollingUpdate

`func (o *V1DeploymentStrategy) HasRollingUpdate() bool`

HasRollingUpdate returns a boolean if a field has been set.

### GetType

`func (o *V1DeploymentStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DeploymentStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DeploymentStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1DeploymentStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


