# V1Lifecycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostStart** | Pointer to [**V1Handler**](V1Handler.md) |  | [optional] 
**PreStop** | Pointer to [**V1Handler**](V1Handler.md) |  | [optional] 

## Methods

### NewV1Lifecycle

`func NewV1Lifecycle() *V1Lifecycle`

NewV1Lifecycle instantiates a new V1Lifecycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LifecycleWithDefaults

`func NewV1LifecycleWithDefaults() *V1Lifecycle`

NewV1LifecycleWithDefaults instantiates a new V1Lifecycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostStart

`func (o *V1Lifecycle) GetPostStart() V1Handler`

GetPostStart returns the PostStart field if non-nil, zero value otherwise.

### GetPostStartOk

`func (o *V1Lifecycle) GetPostStartOk() (*V1Handler, bool)`

GetPostStartOk returns a tuple with the PostStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostStart

`func (o *V1Lifecycle) SetPostStart(v V1Handler)`

SetPostStart sets PostStart field to given value.

### HasPostStart

`func (o *V1Lifecycle) HasPostStart() bool`

HasPostStart returns a boolean if a field has been set.

### GetPreStop

`func (o *V1Lifecycle) GetPreStop() V1Handler`

GetPreStop returns the PreStop field if non-nil, zero value otherwise.

### GetPreStopOk

`func (o *V1Lifecycle) GetPreStopOk() (*V1Handler, bool)`

GetPreStopOk returns a tuple with the PreStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreStop

`func (o *V1Lifecycle) SetPreStop(v V1Handler)`

SetPreStop sets PreStop field to given value.

### HasPreStop

`func (o *V1Lifecycle) HasPreStop() bool`

HasPreStop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


