# V1VolumeProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMap** | Pointer to [**V1ConfigMapProjection**](V1ConfigMapProjection.md) |  | [optional] 
**DownwardAPI** | Pointer to [**V1DownwardAPIProjection**](V1DownwardAPIProjection.md) |  | [optional] 
**Secret** | Pointer to [**V1SecretProjection**](V1SecretProjection.md) |  | [optional] 
**ServiceAccountToken** | Pointer to [**V1ServiceAccountTokenProjection**](V1ServiceAccountTokenProjection.md) |  | [optional] 

## Methods

### NewV1VolumeProjection

`func NewV1VolumeProjection() *V1VolumeProjection`

NewV1VolumeProjection instantiates a new V1VolumeProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VolumeProjectionWithDefaults

`func NewV1VolumeProjectionWithDefaults() *V1VolumeProjection`

NewV1VolumeProjectionWithDefaults instantiates a new V1VolumeProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMap

`func (o *V1VolumeProjection) GetConfigMap() V1ConfigMapProjection`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *V1VolumeProjection) GetConfigMapOk() (*V1ConfigMapProjection, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *V1VolumeProjection) SetConfigMap(v V1ConfigMapProjection)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *V1VolumeProjection) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetDownwardAPI

`func (o *V1VolumeProjection) GetDownwardAPI() V1DownwardAPIProjection`

GetDownwardAPI returns the DownwardAPI field if non-nil, zero value otherwise.

### GetDownwardAPIOk

`func (o *V1VolumeProjection) GetDownwardAPIOk() (*V1DownwardAPIProjection, bool)`

GetDownwardAPIOk returns a tuple with the DownwardAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardAPI

`func (o *V1VolumeProjection) SetDownwardAPI(v V1DownwardAPIProjection)`

SetDownwardAPI sets DownwardAPI field to given value.

### HasDownwardAPI

`func (o *V1VolumeProjection) HasDownwardAPI() bool`

HasDownwardAPI returns a boolean if a field has been set.

### GetSecret

`func (o *V1VolumeProjection) GetSecret() V1SecretProjection`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1VolumeProjection) GetSecretOk() (*V1SecretProjection, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1VolumeProjection) SetSecret(v V1SecretProjection)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V1VolumeProjection) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServiceAccountToken

`func (o *V1VolumeProjection) GetServiceAccountToken() V1ServiceAccountTokenProjection`

GetServiceAccountToken returns the ServiceAccountToken field if non-nil, zero value otherwise.

### GetServiceAccountTokenOk

`func (o *V1VolumeProjection) GetServiceAccountTokenOk() (*V1ServiceAccountTokenProjection, bool)`

GetServiceAccountTokenOk returns a tuple with the ServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountToken

`func (o *V1VolumeProjection) SetServiceAccountToken(v V1ServiceAccountTokenProjection)`

SetServiceAccountToken sets ServiceAccountToken field to given value.

### HasServiceAccountToken

`func (o *V1VolumeProjection) HasServiceAccountToken() bool`

HasServiceAccountToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


