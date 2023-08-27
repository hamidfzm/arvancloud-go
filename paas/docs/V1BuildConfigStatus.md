# V1BuildConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastVersion** | **int64** | lastVersion is used to inform about number of last triggered build. | 

## Methods

### NewV1BuildConfigStatus

`func NewV1BuildConfigStatus(lastVersion int64, ) *V1BuildConfigStatus`

NewV1BuildConfigStatus instantiates a new V1BuildConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildConfigStatusWithDefaults

`func NewV1BuildConfigStatusWithDefaults() *V1BuildConfigStatus`

NewV1BuildConfigStatusWithDefaults instantiates a new V1BuildConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastVersion

`func (o *V1BuildConfigStatus) GetLastVersion() int64`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *V1BuildConfigStatus) GetLastVersionOk() (*int64, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *V1BuildConfigStatus) SetLastVersion(v int64)`

SetLastVersion sets LastVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


