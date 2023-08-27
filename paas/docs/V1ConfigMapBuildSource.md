# V1ConfigMapBuildSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMap** | [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | 
**DestinationDir** | Pointer to **string** | destinationDir is the directory where the files from the configmap should be available for the build time. For the Source build strategy, these will be injected into a container where the assemble script runs. For the Docker build strategy, these will be copied into the build directory, where the Dockerfile is located, so users can ADD or COPY them during docker build. | [optional] 

## Methods

### NewV1ConfigMapBuildSource

`func NewV1ConfigMapBuildSource(configMap V1LocalObjectReference, ) *V1ConfigMapBuildSource`

NewV1ConfigMapBuildSource instantiates a new V1ConfigMapBuildSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigMapBuildSourceWithDefaults

`func NewV1ConfigMapBuildSourceWithDefaults() *V1ConfigMapBuildSource`

NewV1ConfigMapBuildSourceWithDefaults instantiates a new V1ConfigMapBuildSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMap

`func (o *V1ConfigMapBuildSource) GetConfigMap() V1LocalObjectReference`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *V1ConfigMapBuildSource) GetConfigMapOk() (*V1LocalObjectReference, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *V1ConfigMapBuildSource) SetConfigMap(v V1LocalObjectReference)`

SetConfigMap sets ConfigMap field to given value.


### GetDestinationDir

`func (o *V1ConfigMapBuildSource) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *V1ConfigMapBuildSource) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *V1ConfigMapBuildSource) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *V1ConfigMapBuildSource) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


