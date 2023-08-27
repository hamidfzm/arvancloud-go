# Encryption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionType** | [**ServerSideEncryption**](ServerSideEncryption.md) |  | 
**KMSKeyId** | Pointer to **string** |  | [optional] 
**KMSContext** | Pointer to **string** |  | [optional] 

## Methods

### NewEncryption

`func NewEncryption(encryptionType ServerSideEncryption, ) *Encryption`

NewEncryption instantiates a new Encryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionWithDefaults

`func NewEncryptionWithDefaults() *Encryption`

NewEncryptionWithDefaults instantiates a new Encryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionType

`func (o *Encryption) GetEncryptionType() ServerSideEncryption`

GetEncryptionType returns the EncryptionType field if non-nil, zero value otherwise.

### GetEncryptionTypeOk

`func (o *Encryption) GetEncryptionTypeOk() (*ServerSideEncryption, bool)`

GetEncryptionTypeOk returns a tuple with the EncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionType

`func (o *Encryption) SetEncryptionType(v ServerSideEncryption)`

SetEncryptionType sets EncryptionType field to given value.


### GetKMSKeyId

`func (o *Encryption) GetKMSKeyId() string`

GetKMSKeyId returns the KMSKeyId field if non-nil, zero value otherwise.

### GetKMSKeyIdOk

`func (o *Encryption) GetKMSKeyIdOk() (*string, bool)`

GetKMSKeyIdOk returns a tuple with the KMSKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKMSKeyId

`func (o *Encryption) SetKMSKeyId(v string)`

SetKMSKeyId sets KMSKeyId field to given value.

### HasKMSKeyId

`func (o *Encryption) HasKMSKeyId() bool`

HasKMSKeyId returns a boolean if a field has been set.

### GetKMSContext

`func (o *Encryption) GetKMSContext() string`

GetKMSContext returns the KMSContext field if non-nil, zero value otherwise.

### GetKMSContextOk

`func (o *Encryption) GetKMSContextOk() (*string, bool)`

GetKMSContextOk returns a tuple with the KMSContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKMSContext

`func (o *Encryption) SetKMSContext(v string)`

SetKMSContext sets KMSContext field to given value.

### HasKMSContext

`func (o *Encryption) HasKMSContext() bool`

HasKMSContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


