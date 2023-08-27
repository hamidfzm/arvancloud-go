# ListBucketsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | Pointer to [**Array**](array.md) |  | [optional] 
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 

## Methods

### NewListBucketsOutput

`func NewListBucketsOutput() *ListBucketsOutput`

NewListBucketsOutput instantiates a new ListBucketsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBucketsOutputWithDefaults

`func NewListBucketsOutputWithDefaults() *ListBucketsOutput`

NewListBucketsOutputWithDefaults instantiates a new ListBucketsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *ListBucketsOutput) GetBuckets() Array`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ListBucketsOutput) GetBucketsOk() (*Array, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ListBucketsOutput) SetBuckets(v Array)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *ListBucketsOutput) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetOwner

`func (o *ListBucketsOutput) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListBucketsOutput) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListBucketsOutput) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListBucketsOutput) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


