# V1GitSourceRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**V1SourceControlUser**](V1SourceControlUser.md) |  | [optional] 
**Commit** | Pointer to **string** | commit is the commit hash identifying a specific commit | [optional] 
**Committer** | Pointer to [**V1SourceControlUser**](V1SourceControlUser.md) |  | [optional] 
**Message** | Pointer to **string** | message is the description of a specific commit | [optional] 

## Methods

### NewV1GitSourceRevision

`func NewV1GitSourceRevision() *V1GitSourceRevision`

NewV1GitSourceRevision instantiates a new V1GitSourceRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GitSourceRevisionWithDefaults

`func NewV1GitSourceRevisionWithDefaults() *V1GitSourceRevision`

NewV1GitSourceRevisionWithDefaults instantiates a new V1GitSourceRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *V1GitSourceRevision) GetAuthor() V1SourceControlUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *V1GitSourceRevision) GetAuthorOk() (*V1SourceControlUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *V1GitSourceRevision) SetAuthor(v V1SourceControlUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *V1GitSourceRevision) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommit

`func (o *V1GitSourceRevision) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *V1GitSourceRevision) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *V1GitSourceRevision) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *V1GitSourceRevision) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetCommitter

`func (o *V1GitSourceRevision) GetCommitter() V1SourceControlUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *V1GitSourceRevision) GetCommitterOk() (*V1SourceControlUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *V1GitSourceRevision) SetCommitter(v V1SourceControlUser)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *V1GitSourceRevision) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetMessage

`func (o *V1GitSourceRevision) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1GitSourceRevision) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1GitSourceRevision) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1GitSourceRevision) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


