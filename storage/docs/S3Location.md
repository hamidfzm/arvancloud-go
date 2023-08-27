# S3Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** |  | 
**Prefix** | **string** |  | 
**Encryption** | Pointer to [**Encryption**](Encryption.md) |  | [optional] 
**CannedACL** | Pointer to [**ObjectCannedACL**](ObjectCannedACL.md) |  | [optional] 
**AccessControlList** | Pointer to [**Grants**](Grants.md) |  | [optional] 
**Tagging** | Pointer to [**S3LocationTagging**](S3LocationTagging.md) |  | [optional] 
**UserMetadata** | Pointer to [**Array**](array.md) |  | [optional] 
**StorageClass** | Pointer to [**StorageClass**](StorageClass.md) |  | [optional] 

## Methods

### NewS3Location

`func NewS3Location(bucketName string, prefix string, ) *S3Location`

NewS3Location instantiates a new S3Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LocationWithDefaults

`func NewS3LocationWithDefaults() *S3Location`

NewS3LocationWithDefaults instantiates a new S3Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3Location) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3Location) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3Location) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetPrefix

`func (o *S3Location) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3Location) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3Location) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetEncryption

`func (o *S3Location) GetEncryption() Encryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *S3Location) GetEncryptionOk() (*Encryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *S3Location) SetEncryption(v Encryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *S3Location) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetCannedACL

`func (o *S3Location) GetCannedACL() ObjectCannedACL`

GetCannedACL returns the CannedACL field if non-nil, zero value otherwise.

### GetCannedACLOk

`func (o *S3Location) GetCannedACLOk() (*ObjectCannedACL, bool)`

GetCannedACLOk returns a tuple with the CannedACL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCannedACL

`func (o *S3Location) SetCannedACL(v ObjectCannedACL)`

SetCannedACL sets CannedACL field to given value.

### HasCannedACL

`func (o *S3Location) HasCannedACL() bool`

HasCannedACL returns a boolean if a field has been set.

### GetAccessControlList

`func (o *S3Location) GetAccessControlList() Grants`

GetAccessControlList returns the AccessControlList field if non-nil, zero value otherwise.

### GetAccessControlListOk

`func (o *S3Location) GetAccessControlListOk() (*Grants, bool)`

GetAccessControlListOk returns a tuple with the AccessControlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlList

`func (o *S3Location) SetAccessControlList(v Grants)`

SetAccessControlList sets AccessControlList field to given value.

### HasAccessControlList

`func (o *S3Location) HasAccessControlList() bool`

HasAccessControlList returns a boolean if a field has been set.

### GetTagging

`func (o *S3Location) GetTagging() S3LocationTagging`

GetTagging returns the Tagging field if non-nil, zero value otherwise.

### GetTaggingOk

`func (o *S3Location) GetTaggingOk() (*S3LocationTagging, bool)`

GetTaggingOk returns a tuple with the Tagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagging

`func (o *S3Location) SetTagging(v S3LocationTagging)`

SetTagging sets Tagging field to given value.

### HasTagging

`func (o *S3Location) HasTagging() bool`

HasTagging returns a boolean if a field has been set.

### GetUserMetadata

`func (o *S3Location) GetUserMetadata() Array`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *S3Location) GetUserMetadataOk() (*Array, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *S3Location) SetUserMetadata(v Array)`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *S3Location) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.

### GetStorageClass

`func (o *S3Location) GetStorageClass() StorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *S3Location) GetStorageClassOk() (*StorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *S3Location) SetStorageClass(v StorageClass)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *S3Location) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


