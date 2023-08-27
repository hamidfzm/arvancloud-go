# InventoryS3BucketDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Bucket** | **string** |  | 
**Format** | [**InventoryFormat**](InventoryFormat.md) |  | 
**Prefix** | Pointer to **string** |  | [optional] 
**Encryption** | Pointer to [**InventoryS3BucketDestinationEncryption**](InventoryS3BucketDestinationEncryption.md) |  | [optional] 

## Methods

### NewInventoryS3BucketDestination

`func NewInventoryS3BucketDestination(bucket string, format InventoryFormat, ) *InventoryS3BucketDestination`

NewInventoryS3BucketDestination instantiates a new InventoryS3BucketDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryS3BucketDestinationWithDefaults

`func NewInventoryS3BucketDestinationWithDefaults() *InventoryS3BucketDestination`

NewInventoryS3BucketDestinationWithDefaults instantiates a new InventoryS3BucketDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *InventoryS3BucketDestination) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InventoryS3BucketDestination) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InventoryS3BucketDestination) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InventoryS3BucketDestination) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBucket

`func (o *InventoryS3BucketDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *InventoryS3BucketDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *InventoryS3BucketDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetFormat

`func (o *InventoryS3BucketDestination) GetFormat() InventoryFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *InventoryS3BucketDestination) GetFormatOk() (*InventoryFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *InventoryS3BucketDestination) SetFormat(v InventoryFormat)`

SetFormat sets Format field to given value.


### GetPrefix

`func (o *InventoryS3BucketDestination) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *InventoryS3BucketDestination) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *InventoryS3BucketDestination) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *InventoryS3BucketDestination) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetEncryption

`func (o *InventoryS3BucketDestination) GetEncryption() InventoryS3BucketDestinationEncryption`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *InventoryS3BucketDestination) GetEncryptionOk() (*InventoryS3BucketDestinationEncryption, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *InventoryS3BucketDestination) SetEncryption(v InventoryS3BucketDestinationEncryption)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *InventoryS3BucketDestination) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


