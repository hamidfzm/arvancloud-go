# V1ContainerStateRunning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | Pointer to **string** | Time at which the container was last (re-)started | [optional] 

## Methods

### NewV1ContainerStateRunning

`func NewV1ContainerStateRunning() *V1ContainerStateRunning`

NewV1ContainerStateRunning instantiates a new V1ContainerStateRunning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerStateRunningWithDefaults

`func NewV1ContainerStateRunningWithDefaults() *V1ContainerStateRunning`

NewV1ContainerStateRunningWithDefaults instantiates a new V1ContainerStateRunning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *V1ContainerStateRunning) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V1ContainerStateRunning) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V1ContainerStateRunning) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V1ContainerStateRunning) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


