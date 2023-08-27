# V1NodeConfigSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMap** | Pointer to [**V1ConfigMapNodeConfigSource**](V1ConfigMapNodeConfigSource.md) |  | [optional] 

## Methods

### NewV1NodeConfigSource

`func NewV1NodeConfigSource() *V1NodeConfigSource`

NewV1NodeConfigSource instantiates a new V1NodeConfigSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeConfigSourceWithDefaults

`func NewV1NodeConfigSourceWithDefaults() *V1NodeConfigSource`

NewV1NodeConfigSourceWithDefaults instantiates a new V1NodeConfigSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMap

`func (o *V1NodeConfigSource) GetConfigMap() V1ConfigMapNodeConfigSource`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *V1NodeConfigSource) GetConfigMapOk() (*V1ConfigMapNodeConfigSource, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *V1NodeConfigSource) SetConfigMap(v V1ConfigMapNodeConfigSource)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *V1NodeConfigSource) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


