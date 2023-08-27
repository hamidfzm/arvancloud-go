# BatchTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceList** | Pointer to **[]string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**TagList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBatchTagRequest

`func NewBatchTagRequest() *BatchTagRequest`

NewBatchTagRequest instantiates a new BatchTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchTagRequestWithDefaults

`func NewBatchTagRequestWithDefaults() *BatchTagRequest`

NewBatchTagRequestWithDefaults instantiates a new BatchTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceList

`func (o *BatchTagRequest) GetInstanceList() []string`

GetInstanceList returns the InstanceList field if non-nil, zero value otherwise.

### GetInstanceListOk

`func (o *BatchTagRequest) GetInstanceListOk() (*[]string, bool)`

GetInstanceListOk returns a tuple with the InstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceList

`func (o *BatchTagRequest) SetInstanceList(v []string)`

SetInstanceList sets InstanceList field to given value.

### HasInstanceList

`func (o *BatchTagRequest) HasInstanceList() bool`

HasInstanceList returns a boolean if a field has been set.

### GetInstanceType

`func (o *BatchTagRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BatchTagRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BatchTagRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BatchTagRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetTagList

`func (o *BatchTagRequest) GetTagList() []string`

GetTagList returns the TagList field if non-nil, zero value otherwise.

### GetTagListOk

`func (o *BatchTagRequest) GetTagListOk() (*[]string, bool)`

GetTagListOk returns a tuple with the TagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagList

`func (o *BatchTagRequest) SetTagList(v []string)`

SetTagList sets TagList field to given value.

### HasTagList

`func (o *BatchTagRequest) HasTagList() bool`

HasTagList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


