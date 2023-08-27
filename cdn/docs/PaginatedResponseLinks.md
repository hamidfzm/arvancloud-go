# PaginatedResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **string** |  | 
**Last** | Pointer to **NullableString** |  | [optional] 
**Prev** | **NullableString** |  | 
**Next** | **NullableString** |  | 

## Methods

### NewPaginatedResponseLinks

`func NewPaginatedResponseLinks(first string, prev NullableString, next NullableString, ) *PaginatedResponseLinks`

NewPaginatedResponseLinks instantiates a new PaginatedResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseLinksWithDefaults

`func NewPaginatedResponseLinksWithDefaults() *PaginatedResponseLinks`

NewPaginatedResponseLinksWithDefaults instantiates a new PaginatedResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *PaginatedResponseLinks) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *PaginatedResponseLinks) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *PaginatedResponseLinks) SetFirst(v string)`

SetFirst sets First field to given value.


### GetLast

`func (o *PaginatedResponseLinks) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginatedResponseLinks) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginatedResponseLinks) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginatedResponseLinks) HasLast() bool`

HasLast returns a boolean if a field has been set.

### SetLastNil

`func (o *PaginatedResponseLinks) SetLastNil(b bool)`

 SetLastNil sets the value for Last to be an explicit nil

### UnsetLast
`func (o *PaginatedResponseLinks) UnsetLast()`

UnsetLast ensures that no value is present for Last, not even an explicit nil
### GetPrev

`func (o *PaginatedResponseLinks) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginatedResponseLinks) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginatedResponseLinks) SetPrev(v string)`

SetPrev sets Prev field to given value.


### SetPrevNil

`func (o *PaginatedResponseLinks) SetPrevNil(b bool)`

 SetPrevNil sets the value for Prev to be an explicit nil

### UnsetPrev
`func (o *PaginatedResponseLinks) UnsetPrev()`

UnsetPrev ensures that no value is present for Prev, not even an explicit nil
### GetNext

`func (o *PaginatedResponseLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResponseLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResponseLinks) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PaginatedResponseLinks) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResponseLinks) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


