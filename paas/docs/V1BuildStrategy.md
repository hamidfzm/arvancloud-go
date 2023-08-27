# V1BuildStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomStrategy** | Pointer to [**V1CustomBuildStrategy**](V1CustomBuildStrategy.md) |  | [optional] 
**DockerStrategy** | Pointer to [**V1DockerBuildStrategy**](V1DockerBuildStrategy.md) |  | [optional] 
**JenkinsPipelineStrategy** | Pointer to [**V1JenkinsPipelineBuildStrategy**](V1JenkinsPipelineBuildStrategy.md) |  | [optional] 
**SourceStrategy** | Pointer to [**V1SourceBuildStrategy**](V1SourceBuildStrategy.md) |  | [optional] 
**Type** | **string** | type is the kind of build strategy. | 

## Methods

### NewV1BuildStrategy

`func NewV1BuildStrategy(type_ string, ) *V1BuildStrategy`

NewV1BuildStrategy instantiates a new V1BuildStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildStrategyWithDefaults

`func NewV1BuildStrategyWithDefaults() *V1BuildStrategy`

NewV1BuildStrategyWithDefaults instantiates a new V1BuildStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomStrategy

`func (o *V1BuildStrategy) GetCustomStrategy() V1CustomBuildStrategy`

GetCustomStrategy returns the CustomStrategy field if non-nil, zero value otherwise.

### GetCustomStrategyOk

`func (o *V1BuildStrategy) GetCustomStrategyOk() (*V1CustomBuildStrategy, bool)`

GetCustomStrategyOk returns a tuple with the CustomStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStrategy

`func (o *V1BuildStrategy) SetCustomStrategy(v V1CustomBuildStrategy)`

SetCustomStrategy sets CustomStrategy field to given value.

### HasCustomStrategy

`func (o *V1BuildStrategy) HasCustomStrategy() bool`

HasCustomStrategy returns a boolean if a field has been set.

### GetDockerStrategy

`func (o *V1BuildStrategy) GetDockerStrategy() V1DockerBuildStrategy`

GetDockerStrategy returns the DockerStrategy field if non-nil, zero value otherwise.

### GetDockerStrategyOk

`func (o *V1BuildStrategy) GetDockerStrategyOk() (*V1DockerBuildStrategy, bool)`

GetDockerStrategyOk returns a tuple with the DockerStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerStrategy

`func (o *V1BuildStrategy) SetDockerStrategy(v V1DockerBuildStrategy)`

SetDockerStrategy sets DockerStrategy field to given value.

### HasDockerStrategy

`func (o *V1BuildStrategy) HasDockerStrategy() bool`

HasDockerStrategy returns a boolean if a field has been set.

### GetJenkinsPipelineStrategy

`func (o *V1BuildStrategy) GetJenkinsPipelineStrategy() V1JenkinsPipelineBuildStrategy`

GetJenkinsPipelineStrategy returns the JenkinsPipelineStrategy field if non-nil, zero value otherwise.

### GetJenkinsPipelineStrategyOk

`func (o *V1BuildStrategy) GetJenkinsPipelineStrategyOk() (*V1JenkinsPipelineBuildStrategy, bool)`

GetJenkinsPipelineStrategyOk returns a tuple with the JenkinsPipelineStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsPipelineStrategy

`func (o *V1BuildStrategy) SetJenkinsPipelineStrategy(v V1JenkinsPipelineBuildStrategy)`

SetJenkinsPipelineStrategy sets JenkinsPipelineStrategy field to given value.

### HasJenkinsPipelineStrategy

`func (o *V1BuildStrategy) HasJenkinsPipelineStrategy() bool`

HasJenkinsPipelineStrategy returns a boolean if a field has been set.

### GetSourceStrategy

`func (o *V1BuildStrategy) GetSourceStrategy() V1SourceBuildStrategy`

GetSourceStrategy returns the SourceStrategy field if non-nil, zero value otherwise.

### GetSourceStrategyOk

`func (o *V1BuildStrategy) GetSourceStrategyOk() (*V1SourceBuildStrategy, bool)`

GetSourceStrategyOk returns a tuple with the SourceStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStrategy

`func (o *V1BuildStrategy) SetSourceStrategy(v V1SourceBuildStrategy)`

SetSourceStrategy sets SourceStrategy field to given value.

### HasSourceStrategy

`func (o *V1BuildStrategy) HasSourceStrategy() bool`

HasSourceStrategy returns a boolean if a field has been set.

### GetType

`func (o *V1BuildStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1BuildStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1BuildStrategy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


