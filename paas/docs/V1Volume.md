# V1Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsElasticBlockStore** | Pointer to [**V1AWSElasticBlockStoreVolumeSource**](V1AWSElasticBlockStoreVolumeSource.md) |  | [optional] 
**AzureDisk** | Pointer to [**V1AzureDiskVolumeSource**](V1AzureDiskVolumeSource.md) |  | [optional] 
**AzureFile** | Pointer to [**V1AzureFileVolumeSource**](V1AzureFileVolumeSource.md) |  | [optional] 
**Cephfs** | Pointer to [**V1CephFSVolumeSource**](V1CephFSVolumeSource.md) |  | [optional] 
**Cinder** | Pointer to [**V1CinderVolumeSource**](V1CinderVolumeSource.md) |  | [optional] 
**ConfigMap** | Pointer to [**V1ConfigMapVolumeSource**](V1ConfigMapVolumeSource.md) |  | [optional] 
**DownwardAPI** | Pointer to [**V1DownwardAPIVolumeSource**](V1DownwardAPIVolumeSource.md) |  | [optional] 
**EmptyDir** | Pointer to [**V1EmptyDirVolumeSource**](V1EmptyDirVolumeSource.md) |  | [optional] 
**Fc** | Pointer to [**V1FCVolumeSource**](V1FCVolumeSource.md) |  | [optional] 
**FlexVolume** | Pointer to [**V1FlexVolumeSource**](V1FlexVolumeSource.md) |  | [optional] 
**Flocker** | Pointer to [**V1FlockerVolumeSource**](V1FlockerVolumeSource.md) |  | [optional] 
**GcePersistentDisk** | Pointer to [**V1GCEPersistentDiskVolumeSource**](V1GCEPersistentDiskVolumeSource.md) |  | [optional] 
**GitRepo** | Pointer to [**V1GitRepoVolumeSource**](V1GitRepoVolumeSource.md) |  | [optional] 
**Glusterfs** | Pointer to [**V1GlusterfsVolumeSource**](V1GlusterfsVolumeSource.md) |  | [optional] 
**HostPath** | Pointer to [**V1HostPathVolumeSource**](V1HostPathVolumeSource.md) |  | [optional] 
**Iscsi** | Pointer to [**V1ISCSIVolumeSource**](V1ISCSIVolumeSource.md) |  | [optional] 
**Name** | **string** | Volume&#39;s name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | 
**Nfs** | Pointer to [**V1NFSVolumeSource**](V1NFSVolumeSource.md) |  | [optional] 
**PersistentVolumeClaim** | Pointer to [**V1PersistentVolumeClaimVolumeSource**](V1PersistentVolumeClaimVolumeSource.md) |  | [optional] 
**PhotonPersistentDisk** | Pointer to [**V1PhotonPersistentDiskVolumeSource**](V1PhotonPersistentDiskVolumeSource.md) |  | [optional] 
**PortworxVolume** | Pointer to [**V1PortworxVolumeSource**](V1PortworxVolumeSource.md) |  | [optional] 
**Projected** | Pointer to [**V1ProjectedVolumeSource**](V1ProjectedVolumeSource.md) |  | [optional] 
**Quobyte** | Pointer to [**V1QuobyteVolumeSource**](V1QuobyteVolumeSource.md) |  | [optional] 
**Rbd** | Pointer to [**V1RBDVolumeSource**](V1RBDVolumeSource.md) |  | [optional] 
**ScaleIO** | Pointer to [**V1ScaleIOVolumeSource**](V1ScaleIOVolumeSource.md) |  | [optional] 
**Secret** | Pointer to [**V1SecretVolumeSource**](V1SecretVolumeSource.md) |  | [optional] 
**Storageos** | Pointer to [**V1StorageOSVolumeSource**](V1StorageOSVolumeSource.md) |  | [optional] 
**VsphereVolume** | Pointer to [**V1VsphereVirtualDiskVolumeSource**](V1VsphereVirtualDiskVolumeSource.md) |  | [optional] 

## Methods

### NewV1Volume

`func NewV1Volume(name string, ) *V1Volume`

NewV1Volume instantiates a new V1Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VolumeWithDefaults

`func NewV1VolumeWithDefaults() *V1Volume`

NewV1VolumeWithDefaults instantiates a new V1Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsElasticBlockStore

`func (o *V1Volume) GetAwsElasticBlockStore() V1AWSElasticBlockStoreVolumeSource`

GetAwsElasticBlockStore returns the AwsElasticBlockStore field if non-nil, zero value otherwise.

### GetAwsElasticBlockStoreOk

`func (o *V1Volume) GetAwsElasticBlockStoreOk() (*V1AWSElasticBlockStoreVolumeSource, bool)`

GetAwsElasticBlockStoreOk returns a tuple with the AwsElasticBlockStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsElasticBlockStore

`func (o *V1Volume) SetAwsElasticBlockStore(v V1AWSElasticBlockStoreVolumeSource)`

SetAwsElasticBlockStore sets AwsElasticBlockStore field to given value.

### HasAwsElasticBlockStore

`func (o *V1Volume) HasAwsElasticBlockStore() bool`

HasAwsElasticBlockStore returns a boolean if a field has been set.

### GetAzureDisk

`func (o *V1Volume) GetAzureDisk() V1AzureDiskVolumeSource`

GetAzureDisk returns the AzureDisk field if non-nil, zero value otherwise.

### GetAzureDiskOk

`func (o *V1Volume) GetAzureDiskOk() (*V1AzureDiskVolumeSource, bool)`

GetAzureDiskOk returns a tuple with the AzureDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisk

`func (o *V1Volume) SetAzureDisk(v V1AzureDiskVolumeSource)`

SetAzureDisk sets AzureDisk field to given value.

### HasAzureDisk

`func (o *V1Volume) HasAzureDisk() bool`

HasAzureDisk returns a boolean if a field has been set.

### GetAzureFile

`func (o *V1Volume) GetAzureFile() V1AzureFileVolumeSource`

GetAzureFile returns the AzureFile field if non-nil, zero value otherwise.

### GetAzureFileOk

`func (o *V1Volume) GetAzureFileOk() (*V1AzureFileVolumeSource, bool)`

GetAzureFileOk returns a tuple with the AzureFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFile

`func (o *V1Volume) SetAzureFile(v V1AzureFileVolumeSource)`

SetAzureFile sets AzureFile field to given value.

### HasAzureFile

`func (o *V1Volume) HasAzureFile() bool`

HasAzureFile returns a boolean if a field has been set.

### GetCephfs

`func (o *V1Volume) GetCephfs() V1CephFSVolumeSource`

GetCephfs returns the Cephfs field if non-nil, zero value otherwise.

### GetCephfsOk

`func (o *V1Volume) GetCephfsOk() (*V1CephFSVolumeSource, bool)`

GetCephfsOk returns a tuple with the Cephfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephfs

`func (o *V1Volume) SetCephfs(v V1CephFSVolumeSource)`

SetCephfs sets Cephfs field to given value.

### HasCephfs

`func (o *V1Volume) HasCephfs() bool`

HasCephfs returns a boolean if a field has been set.

### GetCinder

`func (o *V1Volume) GetCinder() V1CinderVolumeSource`

GetCinder returns the Cinder field if non-nil, zero value otherwise.

### GetCinderOk

`func (o *V1Volume) GetCinderOk() (*V1CinderVolumeSource, bool)`

GetCinderOk returns a tuple with the Cinder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCinder

`func (o *V1Volume) SetCinder(v V1CinderVolumeSource)`

SetCinder sets Cinder field to given value.

### HasCinder

`func (o *V1Volume) HasCinder() bool`

HasCinder returns a boolean if a field has been set.

### GetConfigMap

`func (o *V1Volume) GetConfigMap() V1ConfigMapVolumeSource`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *V1Volume) GetConfigMapOk() (*V1ConfigMapVolumeSource, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *V1Volume) SetConfigMap(v V1ConfigMapVolumeSource)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *V1Volume) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetDownwardAPI

`func (o *V1Volume) GetDownwardAPI() V1DownwardAPIVolumeSource`

GetDownwardAPI returns the DownwardAPI field if non-nil, zero value otherwise.

### GetDownwardAPIOk

`func (o *V1Volume) GetDownwardAPIOk() (*V1DownwardAPIVolumeSource, bool)`

GetDownwardAPIOk returns a tuple with the DownwardAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardAPI

`func (o *V1Volume) SetDownwardAPI(v V1DownwardAPIVolumeSource)`

SetDownwardAPI sets DownwardAPI field to given value.

### HasDownwardAPI

`func (o *V1Volume) HasDownwardAPI() bool`

HasDownwardAPI returns a boolean if a field has been set.

### GetEmptyDir

`func (o *V1Volume) GetEmptyDir() V1EmptyDirVolumeSource`

GetEmptyDir returns the EmptyDir field if non-nil, zero value otherwise.

### GetEmptyDirOk

`func (o *V1Volume) GetEmptyDirOk() (*V1EmptyDirVolumeSource, bool)`

GetEmptyDirOk returns a tuple with the EmptyDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyDir

`func (o *V1Volume) SetEmptyDir(v V1EmptyDirVolumeSource)`

SetEmptyDir sets EmptyDir field to given value.

### HasEmptyDir

`func (o *V1Volume) HasEmptyDir() bool`

HasEmptyDir returns a boolean if a field has been set.

### GetFc

`func (o *V1Volume) GetFc() V1FCVolumeSource`

GetFc returns the Fc field if non-nil, zero value otherwise.

### GetFcOk

`func (o *V1Volume) GetFcOk() (*V1FCVolumeSource, bool)`

GetFcOk returns a tuple with the Fc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFc

`func (o *V1Volume) SetFc(v V1FCVolumeSource)`

SetFc sets Fc field to given value.

### HasFc

`func (o *V1Volume) HasFc() bool`

HasFc returns a boolean if a field has been set.

### GetFlexVolume

`func (o *V1Volume) GetFlexVolume() V1FlexVolumeSource`

GetFlexVolume returns the FlexVolume field if non-nil, zero value otherwise.

### GetFlexVolumeOk

`func (o *V1Volume) GetFlexVolumeOk() (*V1FlexVolumeSource, bool)`

GetFlexVolumeOk returns a tuple with the FlexVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexVolume

`func (o *V1Volume) SetFlexVolume(v V1FlexVolumeSource)`

SetFlexVolume sets FlexVolume field to given value.

### HasFlexVolume

`func (o *V1Volume) HasFlexVolume() bool`

HasFlexVolume returns a boolean if a field has been set.

### GetFlocker

`func (o *V1Volume) GetFlocker() V1FlockerVolumeSource`

GetFlocker returns the Flocker field if non-nil, zero value otherwise.

### GetFlockerOk

`func (o *V1Volume) GetFlockerOk() (*V1FlockerVolumeSource, bool)`

GetFlockerOk returns a tuple with the Flocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlocker

`func (o *V1Volume) SetFlocker(v V1FlockerVolumeSource)`

SetFlocker sets Flocker field to given value.

### HasFlocker

`func (o *V1Volume) HasFlocker() bool`

HasFlocker returns a boolean if a field has been set.

### GetGcePersistentDisk

`func (o *V1Volume) GetGcePersistentDisk() V1GCEPersistentDiskVolumeSource`

GetGcePersistentDisk returns the GcePersistentDisk field if non-nil, zero value otherwise.

### GetGcePersistentDiskOk

`func (o *V1Volume) GetGcePersistentDiskOk() (*V1GCEPersistentDiskVolumeSource, bool)`

GetGcePersistentDiskOk returns a tuple with the GcePersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcePersistentDisk

`func (o *V1Volume) SetGcePersistentDisk(v V1GCEPersistentDiskVolumeSource)`

SetGcePersistentDisk sets GcePersistentDisk field to given value.

### HasGcePersistentDisk

`func (o *V1Volume) HasGcePersistentDisk() bool`

HasGcePersistentDisk returns a boolean if a field has been set.

### GetGitRepo

`func (o *V1Volume) GetGitRepo() V1GitRepoVolumeSource`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *V1Volume) GetGitRepoOk() (*V1GitRepoVolumeSource, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *V1Volume) SetGitRepo(v V1GitRepoVolumeSource)`

SetGitRepo sets GitRepo field to given value.

### HasGitRepo

`func (o *V1Volume) HasGitRepo() bool`

HasGitRepo returns a boolean if a field has been set.

### GetGlusterfs

`func (o *V1Volume) GetGlusterfs() V1GlusterfsVolumeSource`

GetGlusterfs returns the Glusterfs field if non-nil, zero value otherwise.

### GetGlusterfsOk

`func (o *V1Volume) GetGlusterfsOk() (*V1GlusterfsVolumeSource, bool)`

GetGlusterfsOk returns a tuple with the Glusterfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlusterfs

`func (o *V1Volume) SetGlusterfs(v V1GlusterfsVolumeSource)`

SetGlusterfs sets Glusterfs field to given value.

### HasGlusterfs

`func (o *V1Volume) HasGlusterfs() bool`

HasGlusterfs returns a boolean if a field has been set.

### GetHostPath

`func (o *V1Volume) GetHostPath() V1HostPathVolumeSource`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *V1Volume) GetHostPathOk() (*V1HostPathVolumeSource, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *V1Volume) SetHostPath(v V1HostPathVolumeSource)`

SetHostPath sets HostPath field to given value.

### HasHostPath

`func (o *V1Volume) HasHostPath() bool`

HasHostPath returns a boolean if a field has been set.

### GetIscsi

`func (o *V1Volume) GetIscsi() V1ISCSIVolumeSource`

GetIscsi returns the Iscsi field if non-nil, zero value otherwise.

### GetIscsiOk

`func (o *V1Volume) GetIscsiOk() (*V1ISCSIVolumeSource, bool)`

GetIscsiOk returns a tuple with the Iscsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsi

`func (o *V1Volume) SetIscsi(v V1ISCSIVolumeSource)`

SetIscsi sets Iscsi field to given value.

### HasIscsi

`func (o *V1Volume) HasIscsi() bool`

HasIscsi returns a boolean if a field has been set.

### GetName

`func (o *V1Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Volume) SetName(v string)`

SetName sets Name field to given value.


### GetNfs

`func (o *V1Volume) GetNfs() V1NFSVolumeSource`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *V1Volume) GetNfsOk() (*V1NFSVolumeSource, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *V1Volume) SetNfs(v V1NFSVolumeSource)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *V1Volume) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetPersistentVolumeClaim

`func (o *V1Volume) GetPersistentVolumeClaim() V1PersistentVolumeClaimVolumeSource`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *V1Volume) GetPersistentVolumeClaimOk() (*V1PersistentVolumeClaimVolumeSource, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *V1Volume) SetPersistentVolumeClaim(v V1PersistentVolumeClaimVolumeSource)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *V1Volume) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.

### GetPhotonPersistentDisk

`func (o *V1Volume) GetPhotonPersistentDisk() V1PhotonPersistentDiskVolumeSource`

GetPhotonPersistentDisk returns the PhotonPersistentDisk field if non-nil, zero value otherwise.

### GetPhotonPersistentDiskOk

`func (o *V1Volume) GetPhotonPersistentDiskOk() (*V1PhotonPersistentDiskVolumeSource, bool)`

GetPhotonPersistentDiskOk returns a tuple with the PhotonPersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonPersistentDisk

`func (o *V1Volume) SetPhotonPersistentDisk(v V1PhotonPersistentDiskVolumeSource)`

SetPhotonPersistentDisk sets PhotonPersistentDisk field to given value.

### HasPhotonPersistentDisk

`func (o *V1Volume) HasPhotonPersistentDisk() bool`

HasPhotonPersistentDisk returns a boolean if a field has been set.

### GetPortworxVolume

`func (o *V1Volume) GetPortworxVolume() V1PortworxVolumeSource`

GetPortworxVolume returns the PortworxVolume field if non-nil, zero value otherwise.

### GetPortworxVolumeOk

`func (o *V1Volume) GetPortworxVolumeOk() (*V1PortworxVolumeSource, bool)`

GetPortworxVolumeOk returns a tuple with the PortworxVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortworxVolume

`func (o *V1Volume) SetPortworxVolume(v V1PortworxVolumeSource)`

SetPortworxVolume sets PortworxVolume field to given value.

### HasPortworxVolume

`func (o *V1Volume) HasPortworxVolume() bool`

HasPortworxVolume returns a boolean if a field has been set.

### GetProjected

`func (o *V1Volume) GetProjected() V1ProjectedVolumeSource`

GetProjected returns the Projected field if non-nil, zero value otherwise.

### GetProjectedOk

`func (o *V1Volume) GetProjectedOk() (*V1ProjectedVolumeSource, bool)`

GetProjectedOk returns a tuple with the Projected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjected

`func (o *V1Volume) SetProjected(v V1ProjectedVolumeSource)`

SetProjected sets Projected field to given value.

### HasProjected

`func (o *V1Volume) HasProjected() bool`

HasProjected returns a boolean if a field has been set.

### GetQuobyte

`func (o *V1Volume) GetQuobyte() V1QuobyteVolumeSource`

GetQuobyte returns the Quobyte field if non-nil, zero value otherwise.

### GetQuobyteOk

`func (o *V1Volume) GetQuobyteOk() (*V1QuobyteVolumeSource, bool)`

GetQuobyteOk returns a tuple with the Quobyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuobyte

`func (o *V1Volume) SetQuobyte(v V1QuobyteVolumeSource)`

SetQuobyte sets Quobyte field to given value.

### HasQuobyte

`func (o *V1Volume) HasQuobyte() bool`

HasQuobyte returns a boolean if a field has been set.

### GetRbd

`func (o *V1Volume) GetRbd() V1RBDVolumeSource`

GetRbd returns the Rbd field if non-nil, zero value otherwise.

### GetRbdOk

`func (o *V1Volume) GetRbdOk() (*V1RBDVolumeSource, bool)`

GetRbdOk returns a tuple with the Rbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbd

`func (o *V1Volume) SetRbd(v V1RBDVolumeSource)`

SetRbd sets Rbd field to given value.

### HasRbd

`func (o *V1Volume) HasRbd() bool`

HasRbd returns a boolean if a field has been set.

### GetScaleIO

`func (o *V1Volume) GetScaleIO() V1ScaleIOVolumeSource`

GetScaleIO returns the ScaleIO field if non-nil, zero value otherwise.

### GetScaleIOOk

`func (o *V1Volume) GetScaleIOOk() (*V1ScaleIOVolumeSource, bool)`

GetScaleIOOk returns a tuple with the ScaleIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIO

`func (o *V1Volume) SetScaleIO(v V1ScaleIOVolumeSource)`

SetScaleIO sets ScaleIO field to given value.

### HasScaleIO

`func (o *V1Volume) HasScaleIO() bool`

HasScaleIO returns a boolean if a field has been set.

### GetSecret

`func (o *V1Volume) GetSecret() V1SecretVolumeSource`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1Volume) GetSecretOk() (*V1SecretVolumeSource, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1Volume) SetSecret(v V1SecretVolumeSource)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V1Volume) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStorageos

`func (o *V1Volume) GetStorageos() V1StorageOSVolumeSource`

GetStorageos returns the Storageos field if non-nil, zero value otherwise.

### GetStorageosOk

`func (o *V1Volume) GetStorageosOk() (*V1StorageOSVolumeSource, bool)`

GetStorageosOk returns a tuple with the Storageos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageos

`func (o *V1Volume) SetStorageos(v V1StorageOSVolumeSource)`

SetStorageos sets Storageos field to given value.

### HasStorageos

`func (o *V1Volume) HasStorageos() bool`

HasStorageos returns a boolean if a field has been set.

### GetVsphereVolume

`func (o *V1Volume) GetVsphereVolume() V1VsphereVirtualDiskVolumeSource`

GetVsphereVolume returns the VsphereVolume field if non-nil, zero value otherwise.

### GetVsphereVolumeOk

`func (o *V1Volume) GetVsphereVolumeOk() (*V1VsphereVirtualDiskVolumeSource, bool)`

GetVsphereVolumeOk returns a tuple with the VsphereVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereVolume

`func (o *V1Volume) SetVsphereVolume(v V1VsphereVirtualDiskVolumeSource)`

SetVsphereVolume sets VsphereVolume field to given value.

### HasVsphereVolume

`func (o *V1Volume) HasVsphereVolume() bool`

HasVsphereVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


