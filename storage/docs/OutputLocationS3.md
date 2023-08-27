# OutputLocationS3

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

### NewOutputLocationS3

`func NewOutputLocationS3(bucketName string, prefix string, ) *OutputLocationS3`

NewOutputLocationS3 instantiates a new OutputLocationS3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLocationS3WithDefaults

`func NewOutputLocationS3WithDefaults() *OutputLocationS3`

NewOutputLocationS3WithDefaults instantiates a new OutputLocationS3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *OutputLocationS3) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *OutputLocationS3) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *OutputLocationS3) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetPrefix

`func (o *OutputLocationS3) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OutputLocationS3) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OutputLocationS3) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetEncryption

`func (o *OutputLocationS3) GetEncryption() Encryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *OutputLocationS3) GetEncryptionOk() (*Encryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *OutputLocationS3) SetEncryption(v Encryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *OutputLocationS3) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetCannedACL

`func (o *OutputLocationS3) GetCannedACL() ObjectCannedACL`

GetCannedACL returns the CannedACL field if non-nil, zero value otherwise.

### GetCannedACLOk

`func (o *OutputLocationS3) GetCannedACLOk() (*ObjectCannedACL, bool)`

GetCannedACLOk returns a tuple with the CannedACL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCannedACL

`func (o *OutputLocationS3) SetCannedACL(v ObjectCannedACL)`

SetCannedACL sets CannedACL field to given value.

### HasCannedACL

`func (o *OutputLocationS3) HasCannedACL() bool`

HasCannedACL returns a boolean if a field has been set.

### GetAccessControlList

`func (o *OutputLocationS3) GetAccessControlList() Grants`

GetAccessControlList returns the AccessControlList field if non-nil, zero value otherwise.

### GetAccessControlListOk

`func (o *OutputLocationS3) GetAccessControlListOk() (*Grants, bool)`

GetAccessControlListOk returns a tuple with the AccessControlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlList

`func (o *OutputLocationS3) SetAccessControlList(v Grants)`

SetAccessControlList sets AccessControlList field to given value.

### HasAccessControlList

`func (o *OutputLocationS3) HasAccessControlList() bool`

HasAccessControlList returns a boolean if a field has been set.

### GetTagging

`func (o *OutputLocationS3) GetTagging() S3LocationTagging`

GetTagging returns the Tagging field if non-nil, zero value otherwise.

### GetTaggingOk

`func (o *OutputLocationS3) GetTaggingOk() (*S3LocationTagging, bool)`

GetTaggingOk returns a tuple with the Tagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagging

`func (o *OutputLocationS3) SetTagging(v S3LocationTagging)`

SetTagging sets Tagging field to given value.

### HasTagging

`func (o *OutputLocationS3) HasTagging() bool`

HasTagging returns a boolean if a field has been set.

### GetUserMetadata

`func (o *OutputLocationS3) GetUserMetadata() Array`

GetUserMetadata returns the UserMetadata field if non-nil, zero value otherwise.

### GetUserMetadataOk

`func (o *OutputLocationS3) GetUserMetadataOk() (*Array, bool)`

GetUserMetadataOk returns a tuple with the UserMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetadata

`func (o *OutputLocationS3) SetUserMetadata(v Array)`

SetUserMetadata sets UserMetadata field to given value.

### HasUserMetadata

`func (o *OutputLocationS3) HasUserMetadata() bool`

HasUserMetadata returns a boolean if a field has been set.

### GetStorageClass

`func (o *OutputLocationS3) GetStorageClass() StorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *OutputLocationS3) GetStorageClassOk() (*StorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *OutputLocationS3) SetStorageClass(v StorageClass)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *OutputLocationS3) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


