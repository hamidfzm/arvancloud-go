# V1BuildStatusOutputTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageDigest** | Pointer to **string** | imageDigest is the digest of the built Docker image. The digest uniquely identifies the image in the registry to which it was pushed.  Please note that this field may not always be set even if the push completes successfully - e.g. when the registry returns no digest or returns it in a format that the builder doesn&#39;t understand. | [optional] 

## Methods

### NewV1BuildStatusOutputTo

`func NewV1BuildStatusOutputTo() *V1BuildStatusOutputTo`

NewV1BuildStatusOutputTo instantiates a new V1BuildStatusOutputTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildStatusOutputToWithDefaults

`func NewV1BuildStatusOutputToWithDefaults() *V1BuildStatusOutputTo`

NewV1BuildStatusOutputToWithDefaults instantiates a new V1BuildStatusOutputTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageDigest

`func (o *V1BuildStatusOutputTo) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *V1BuildStatusOutputTo) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *V1BuildStatusOutputTo) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *V1BuildStatusOutputTo) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


