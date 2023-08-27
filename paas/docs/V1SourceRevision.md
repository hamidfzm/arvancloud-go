# V1SourceRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Git** | Pointer to [**V1GitSourceRevision**](V1GitSourceRevision.md) |  | [optional] 
**Type** | **string** | type of the build source, may be one of &#39;Source&#39;, &#39;Dockerfile&#39;, &#39;Binary&#39;, or &#39;Images&#39; | 

## Methods

### NewV1SourceRevision

`func NewV1SourceRevision(type_ string, ) *V1SourceRevision`

NewV1SourceRevision instantiates a new V1SourceRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SourceRevisionWithDefaults

`func NewV1SourceRevisionWithDefaults() *V1SourceRevision`

NewV1SourceRevisionWithDefaults instantiates a new V1SourceRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGit

`func (o *V1SourceRevision) GetGit() V1GitSourceRevision`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *V1SourceRevision) GetGitOk() (*V1GitSourceRevision, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *V1SourceRevision) SetGit(v V1GitSourceRevision)`

SetGit sets Git field to given value.

### HasGit

`func (o *V1SourceRevision) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetType

`func (o *V1SourceRevision) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1SourceRevision) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1SourceRevision) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


