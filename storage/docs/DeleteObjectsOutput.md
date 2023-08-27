# DeleteObjectsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to [**Array**](array.md) |  | [optional] 
**Errors** | Pointer to [**Array**](array.md) |  | [optional] 

## Methods

### NewDeleteObjectsOutput

`func NewDeleteObjectsOutput() *DeleteObjectsOutput`

NewDeleteObjectsOutput instantiates a new DeleteObjectsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteObjectsOutputWithDefaults

`func NewDeleteObjectsOutputWithDefaults() *DeleteObjectsOutput`

NewDeleteObjectsOutputWithDefaults instantiates a new DeleteObjectsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *DeleteObjectsOutput) GetDeleted() Array`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteObjectsOutput) GetDeletedOk() (*Array, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteObjectsOutput) SetDeleted(v Array)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteObjectsOutput) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetErrors

`func (o *DeleteObjectsOutput) GetErrors() Array`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DeleteObjectsOutput) GetErrorsOk() (*Array, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DeleteObjectsOutput) SetErrors(v Array)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DeleteObjectsOutput) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


