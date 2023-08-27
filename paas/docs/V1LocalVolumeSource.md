# V1LocalVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The full path to the volume on the node. It can be either a directory or block device (disk, partition, ...). Directories can be represented only by PersistentVolume with VolumeMode&#x3D;Filesystem. Block devices can be represented only by VolumeMode&#x3D;Block, which also requires the BlockVolume alpha feature gate to be enabled. | 

## Methods

### NewV1LocalVolumeSource

`func NewV1LocalVolumeSource(path string, ) *V1LocalVolumeSource`

NewV1LocalVolumeSource instantiates a new V1LocalVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LocalVolumeSourceWithDefaults

`func NewV1LocalVolumeSourceWithDefaults() *V1LocalVolumeSource`

NewV1LocalVolumeSourceWithDefaults instantiates a new V1LocalVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *V1LocalVolumeSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1LocalVolumeSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1LocalVolumeSource) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


