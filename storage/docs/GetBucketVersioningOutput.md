# GetBucketVersioningOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**BucketVersioningStatus**](BucketVersioningStatus.md) |  | [optional] 
**MFADelete** | Pointer to [**MFADeleteStatus**](MFADeleteStatus.md) |  | [optional] 

## Methods

### NewGetBucketVersioningOutput

`func NewGetBucketVersioningOutput() *GetBucketVersioningOutput`

NewGetBucketVersioningOutput instantiates a new GetBucketVersioningOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBucketVersioningOutputWithDefaults

`func NewGetBucketVersioningOutputWithDefaults() *GetBucketVersioningOutput`

NewGetBucketVersioningOutputWithDefaults instantiates a new GetBucketVersioningOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetBucketVersioningOutput) GetStatus() BucketVersioningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBucketVersioningOutput) GetStatusOk() (*BucketVersioningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBucketVersioningOutput) SetStatus(v BucketVersioningStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBucketVersioningOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMFADelete

`func (o *GetBucketVersioningOutput) GetMFADelete() MFADeleteStatus`

GetMFADelete returns the MFADelete field if non-nil, zero value otherwise.

### GetMFADeleteOk

`func (o *GetBucketVersioningOutput) GetMFADeleteOk() (*MFADeleteStatus, bool)`

GetMFADeleteOk returns a tuple with the MFADelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMFADelete

`func (o *GetBucketVersioningOutput) SetMFADelete(v MFADeleteStatus)`

SetMFADelete sets MFADelete field to given value.

### HasMFADelete

`func (o *GetBucketVersioningOutput) HasMFADelete() bool`

HasMFADelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


