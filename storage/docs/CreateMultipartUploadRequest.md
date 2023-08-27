# CreateMultipartUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XAmzMeta** | Pointer to **map[string]string** | A map of metadata to store with the object in S3. | [optional] 

## Methods

### NewCreateMultipartUploadRequest

`func NewCreateMultipartUploadRequest() *CreateMultipartUploadRequest`

NewCreateMultipartUploadRequest instantiates a new CreateMultipartUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMultipartUploadRequestWithDefaults

`func NewCreateMultipartUploadRequestWithDefaults() *CreateMultipartUploadRequest`

NewCreateMultipartUploadRequestWithDefaults instantiates a new CreateMultipartUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXAmzMeta

`func (o *CreateMultipartUploadRequest) GetXAmzMeta() map[string]string`

GetXAmzMeta returns the XAmzMeta field if non-nil, zero value otherwise.

### GetXAmzMetaOk

`func (o *CreateMultipartUploadRequest) GetXAmzMetaOk() (*map[string]string, bool)`

GetXAmzMetaOk returns a tuple with the XAmzMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAmzMeta

`func (o *CreateMultipartUploadRequest) SetXAmzMeta(v map[string]string)`

SetXAmzMeta sets XAmzMeta field to given value.

### HasXAmzMeta

`func (o *CreateMultipartUploadRequest) HasXAmzMeta() bool`

HasXAmzMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


