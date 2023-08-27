# V1AzureFileVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnly** | Pointer to **bool** | Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. | [optional] 
**SecretName** | **string** | the name of secret that contains Azure Storage Account Name and Key | 
**ShareName** | **string** | Share Name | 

## Methods

### NewV1AzureFileVolumeSource

`func NewV1AzureFileVolumeSource(secretName string, shareName string, ) *V1AzureFileVolumeSource`

NewV1AzureFileVolumeSource instantiates a new V1AzureFileVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AzureFileVolumeSourceWithDefaults

`func NewV1AzureFileVolumeSourceWithDefaults() *V1AzureFileVolumeSource`

NewV1AzureFileVolumeSourceWithDefaults instantiates a new V1AzureFileVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadOnly

`func (o *V1AzureFileVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *V1AzureFileVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *V1AzureFileVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *V1AzureFileVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretName

`func (o *V1AzureFileVolumeSource) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *V1AzureFileVolumeSource) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *V1AzureFileVolumeSource) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetShareName

`func (o *V1AzureFileVolumeSource) GetShareName() string`

GetShareName returns the ShareName field if non-nil, zero value otherwise.

### GetShareNameOk

`func (o *V1AzureFileVolumeSource) GetShareNameOk() (*string, bool)`

GetShareNameOk returns a tuple with the ShareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareName

`func (o *V1AzureFileVolumeSource) SetShareName(v string)`

SetShareName sets ShareName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


