# V1SecretSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPath** | **string** | mountPath is the path at which to mount the secret | 
**SecretSource** | [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | 

## Methods

### NewV1SecretSpec

`func NewV1SecretSpec(mountPath string, secretSource V1LocalObjectReference, ) *V1SecretSpec`

NewV1SecretSpec instantiates a new V1SecretSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SecretSpecWithDefaults

`func NewV1SecretSpecWithDefaults() *V1SecretSpec`

NewV1SecretSpecWithDefaults instantiates a new V1SecretSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPath

`func (o *V1SecretSpec) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *V1SecretSpec) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *V1SecretSpec) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetSecretSource

`func (o *V1SecretSpec) GetSecretSource() V1LocalObjectReference`

GetSecretSource returns the SecretSource field if non-nil, zero value otherwise.

### GetSecretSourceOk

`func (o *V1SecretSpec) GetSecretSourceOk() (*V1LocalObjectReference, bool)`

GetSecretSourceOk returns a tuple with the SecretSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretSource

`func (o *V1SecretSpec) SetSecretSource(v V1LocalObjectReference)`

SetSecretSource sets SecretSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


