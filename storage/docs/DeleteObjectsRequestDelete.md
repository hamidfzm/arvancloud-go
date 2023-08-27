# DeleteObjectsRequestDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**Array**](array.md) |  | 
**Quiet** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeleteObjectsRequestDelete

`func NewDeleteObjectsRequestDelete(objects Array, ) *DeleteObjectsRequestDelete`

NewDeleteObjectsRequestDelete instantiates a new DeleteObjectsRequestDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteObjectsRequestDeleteWithDefaults

`func NewDeleteObjectsRequestDeleteWithDefaults() *DeleteObjectsRequestDelete`

NewDeleteObjectsRequestDeleteWithDefaults instantiates a new DeleteObjectsRequestDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *DeleteObjectsRequestDelete) GetObjects() Array`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *DeleteObjectsRequestDelete) GetObjectsOk() (*Array, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *DeleteObjectsRequestDelete) SetObjects(v Array)`

SetObjects sets Objects field to given value.


### GetQuiet

`func (o *DeleteObjectsRequestDelete) GetQuiet() bool`

GetQuiet returns the Quiet field if non-nil, zero value otherwise.

### GetQuietOk

`func (o *DeleteObjectsRequestDelete) GetQuietOk() (*bool, bool)`

GetQuietOk returns a tuple with the Quiet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiet

`func (o *DeleteObjectsRequestDelete) SetQuiet(v bool)`

SetQuiet sets Quiet field to given value.

### HasQuiet

`func (o *DeleteObjectsRequestDelete) HasQuiet() bool`

HasQuiet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


