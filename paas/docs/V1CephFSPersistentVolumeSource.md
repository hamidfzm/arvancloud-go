# V1CephFSPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitors** | **[]string** | Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | 
**Path** | Pointer to **string** | Optional: Used as the mounted root, rather than the full Ceph tree, default is / | [optional] 
**ReadOnly** | Pointer to **bool** | Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] 
**SecretFile** | Pointer to **string** | Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] 
**SecretRef** | Pointer to [**V1SecretReference**](V1SecretReference.md) |  | [optional] 
**User** | Pointer to **string** | Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it | [optional] 

## Methods

### NewV1CephFSPersistentVolumeSource

`func NewV1CephFSPersistentVolumeSource(monitors []string, ) *V1CephFSPersistentVolumeSource`

NewV1CephFSPersistentVolumeSource instantiates a new V1CephFSPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CephFSPersistentVolumeSourceWithDefaults

`func NewV1CephFSPersistentVolumeSourceWithDefaults() *V1CephFSPersistentVolumeSource`

NewV1CephFSPersistentVolumeSourceWithDefaults instantiates a new V1CephFSPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitors

`func (o *V1CephFSPersistentVolumeSource) GetMonitors() []string`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *V1CephFSPersistentVolumeSource) GetMonitorsOk() (*[]string, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *V1CephFSPersistentVolumeSource) SetMonitors(v []string)`

SetMonitors sets Monitors field to given value.


### GetPath

`func (o *V1CephFSPersistentVolumeSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1CephFSPersistentVolumeSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1CephFSPersistentVolumeSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *V1CephFSPersistentVolumeSource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReadOnly

`func (o *V1CephFSPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *V1CephFSPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *V1CephFSPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *V1CephFSPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretFile

`func (o *V1CephFSPersistentVolumeSource) GetSecretFile() string`

GetSecretFile returns the SecretFile field if non-nil, zero value otherwise.

### GetSecretFileOk

`func (o *V1CephFSPersistentVolumeSource) GetSecretFileOk() (*string, bool)`

GetSecretFileOk returns a tuple with the SecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretFile

`func (o *V1CephFSPersistentVolumeSource) SetSecretFile(v string)`

SetSecretFile sets SecretFile field to given value.

### HasSecretFile

`func (o *V1CephFSPersistentVolumeSource) HasSecretFile() bool`

HasSecretFile returns a boolean if a field has been set.

### GetSecretRef

`func (o *V1CephFSPersistentVolumeSource) GetSecretRef() V1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *V1CephFSPersistentVolumeSource) GetSecretRefOk() (*V1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *V1CephFSPersistentVolumeSource) SetSecretRef(v V1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *V1CephFSPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUser

`func (o *V1CephFSPersistentVolumeSource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V1CephFSPersistentVolumeSource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V1CephFSPersistentVolumeSource) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V1CephFSPersistentVolumeSource) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


