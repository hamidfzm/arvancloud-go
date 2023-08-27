# V1CSIPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerPublishSecretRef** | Pointer to [**V1SecretReference**](V1SecretReference.md) |  | [optional] 
**Driver** | **string** | Driver is the name of the driver to use for this volume. Required. | 
**FsType** | Pointer to **string** | Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. | [optional] 
**NodePublishSecretRef** | Pointer to [**V1SecretReference**](V1SecretReference.md) |  | [optional] 
**NodeStageSecretRef** | Pointer to [**V1SecretReference**](V1SecretReference.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** | Optional: The value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write). | [optional] 
**VolumeAttributes** | Pointer to **map[string]interface{}** | Attributes of the volume to publish. | [optional] 
**VolumeHandle** | **string** | VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required. | 

## Methods

### NewV1CSIPersistentVolumeSource

`func NewV1CSIPersistentVolumeSource(driver string, volumeHandle string, ) *V1CSIPersistentVolumeSource`

NewV1CSIPersistentVolumeSource instantiates a new V1CSIPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CSIPersistentVolumeSourceWithDefaults

`func NewV1CSIPersistentVolumeSourceWithDefaults() *V1CSIPersistentVolumeSource`

NewV1CSIPersistentVolumeSourceWithDefaults instantiates a new V1CSIPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerPublishSecretRef

`func (o *V1CSIPersistentVolumeSource) GetControllerPublishSecretRef() V1SecretReference`

GetControllerPublishSecretRef returns the ControllerPublishSecretRef field if non-nil, zero value otherwise.

### GetControllerPublishSecretRefOk

`func (o *V1CSIPersistentVolumeSource) GetControllerPublishSecretRefOk() (*V1SecretReference, bool)`

GetControllerPublishSecretRefOk returns a tuple with the ControllerPublishSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerPublishSecretRef

`func (o *V1CSIPersistentVolumeSource) SetControllerPublishSecretRef(v V1SecretReference)`

SetControllerPublishSecretRef sets ControllerPublishSecretRef field to given value.

### HasControllerPublishSecretRef

`func (o *V1CSIPersistentVolumeSource) HasControllerPublishSecretRef() bool`

HasControllerPublishSecretRef returns a boolean if a field has been set.

### GetDriver

`func (o *V1CSIPersistentVolumeSource) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *V1CSIPersistentVolumeSource) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *V1CSIPersistentVolumeSource) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetFsType

`func (o *V1CSIPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *V1CSIPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *V1CSIPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *V1CSIPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetNodePublishSecretRef

`func (o *V1CSIPersistentVolumeSource) GetNodePublishSecretRef() V1SecretReference`

GetNodePublishSecretRef returns the NodePublishSecretRef field if non-nil, zero value otherwise.

### GetNodePublishSecretRefOk

`func (o *V1CSIPersistentVolumeSource) GetNodePublishSecretRefOk() (*V1SecretReference, bool)`

GetNodePublishSecretRefOk returns a tuple with the NodePublishSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePublishSecretRef

`func (o *V1CSIPersistentVolumeSource) SetNodePublishSecretRef(v V1SecretReference)`

SetNodePublishSecretRef sets NodePublishSecretRef field to given value.

### HasNodePublishSecretRef

`func (o *V1CSIPersistentVolumeSource) HasNodePublishSecretRef() bool`

HasNodePublishSecretRef returns a boolean if a field has been set.

### GetNodeStageSecretRef

`func (o *V1CSIPersistentVolumeSource) GetNodeStageSecretRef() V1SecretReference`

GetNodeStageSecretRef returns the NodeStageSecretRef field if non-nil, zero value otherwise.

### GetNodeStageSecretRefOk

`func (o *V1CSIPersistentVolumeSource) GetNodeStageSecretRefOk() (*V1SecretReference, bool)`

GetNodeStageSecretRefOk returns a tuple with the NodeStageSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStageSecretRef

`func (o *V1CSIPersistentVolumeSource) SetNodeStageSecretRef(v V1SecretReference)`

SetNodeStageSecretRef sets NodeStageSecretRef field to given value.

### HasNodeStageSecretRef

`func (o *V1CSIPersistentVolumeSource) HasNodeStageSecretRef() bool`

HasNodeStageSecretRef returns a boolean if a field has been set.

### GetReadOnly

`func (o *V1CSIPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *V1CSIPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *V1CSIPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *V1CSIPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetVolumeAttributes

`func (o *V1CSIPersistentVolumeSource) GetVolumeAttributes() map[string]interface{}`

GetVolumeAttributes returns the VolumeAttributes field if non-nil, zero value otherwise.

### GetVolumeAttributesOk

`func (o *V1CSIPersistentVolumeSource) GetVolumeAttributesOk() (*map[string]interface{}, bool)`

GetVolumeAttributesOk returns a tuple with the VolumeAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttributes

`func (o *V1CSIPersistentVolumeSource) SetVolumeAttributes(v map[string]interface{})`

SetVolumeAttributes sets VolumeAttributes field to given value.

### HasVolumeAttributes

`func (o *V1CSIPersistentVolumeSource) HasVolumeAttributes() bool`

HasVolumeAttributes returns a boolean if a field has been set.

### GetVolumeHandle

`func (o *V1CSIPersistentVolumeSource) GetVolumeHandle() string`

GetVolumeHandle returns the VolumeHandle field if non-nil, zero value otherwise.

### GetVolumeHandleOk

`func (o *V1CSIPersistentVolumeSource) GetVolumeHandleOk() (*string, bool)`

GetVolumeHandleOk returns a tuple with the VolumeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeHandle

`func (o *V1CSIPersistentVolumeSource) SetVolumeHandle(v string)`

SetVolumeHandle sets VolumeHandle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


