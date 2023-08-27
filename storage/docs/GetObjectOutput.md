# GetObjectOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ModelMap**](map.md) |  | [optional] 

## Methods

### NewGetObjectOutput

`func NewGetObjectOutput() *GetObjectOutput`

NewGetObjectOutput instantiates a new GetObjectOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectOutputWithDefaults

`func NewGetObjectOutputWithDefaults() *GetObjectOutput`

NewGetObjectOutputWithDefaults instantiates a new GetObjectOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *GetObjectOutput) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetObjectOutput) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetObjectOutput) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetObjectOutput) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMetadata

`func (o *GetObjectOutput) GetMetadata() ModelMap`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetObjectOutput) GetMetadataOk() (*ModelMap, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetObjectOutput) SetMetadata(v ModelMap)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetObjectOutput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


