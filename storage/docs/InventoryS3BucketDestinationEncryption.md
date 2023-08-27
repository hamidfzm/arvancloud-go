# InventoryS3BucketDestinationEncryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SSES3** | Pointer to **interface{}** |  | [optional] 
**SSEKMS** | Pointer to [**InventoryEncryptionSSEKMS**](InventoryEncryptionSSEKMS.md) |  | [optional] 

## Methods

### NewInventoryS3BucketDestinationEncryption

`func NewInventoryS3BucketDestinationEncryption() *InventoryS3BucketDestinationEncryption`

NewInventoryS3BucketDestinationEncryption instantiates a new InventoryS3BucketDestinationEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryS3BucketDestinationEncryptionWithDefaults

`func NewInventoryS3BucketDestinationEncryptionWithDefaults() *InventoryS3BucketDestinationEncryption`

NewInventoryS3BucketDestinationEncryptionWithDefaults instantiates a new InventoryS3BucketDestinationEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSSES3

`func (o *InventoryS3BucketDestinationEncryption) GetSSES3() interface{}`

GetSSES3 returns the SSES3 field if non-nil, zero value otherwise.

### GetSSES3Ok

`func (o *InventoryS3BucketDestinationEncryption) GetSSES3Ok() (*interface{}, bool)`

GetSSES3Ok returns a tuple with the SSES3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSES3

`func (o *InventoryS3BucketDestinationEncryption) SetSSES3(v interface{})`

SetSSES3 sets SSES3 field to given value.

### HasSSES3

`func (o *InventoryS3BucketDestinationEncryption) HasSSES3() bool`

HasSSES3 returns a boolean if a field has been set.

### GetSSEKMS

`func (o *InventoryS3BucketDestinationEncryption) GetSSEKMS() InventoryEncryptionSSEKMS`

GetSSEKMS returns the SSEKMS field if non-nil, zero value otherwise.

### GetSSEKMSOk

`func (o *InventoryS3BucketDestinationEncryption) GetSSEKMSOk() (*InventoryEncryptionSSEKMS, bool)`

GetSSEKMSOk returns a tuple with the SSEKMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSEKMS

`func (o *InventoryS3BucketDestinationEncryption) SetSSEKMS(v InventoryEncryptionSSEKMS)`

SetSSEKMS sets SSEKMS field to given value.

### HasSSEKMS

`func (o *InventoryS3BucketDestinationEncryption) HasSSEKMS() bool`

HasSSEKMS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


