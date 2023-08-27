# V1PersistentVolumeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to [**[]V1PersistentVolumeAccessMode**](V1PersistentVolumeAccessMode.md) | AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes | [optional] 
**AwsElasticBlockStore** | Pointer to [**V1AWSElasticBlockStoreVolumeSource**](V1AWSElasticBlockStoreVolumeSource.md) |  | [optional] 
**AzureDisk** | Pointer to [**V1AzureDiskVolumeSource**](V1AzureDiskVolumeSource.md) |  | [optional] 
**AzureFile** | Pointer to [**V1AzureFilePersistentVolumeSource**](V1AzureFilePersistentVolumeSource.md) |  | [optional] 
**Capacity** | Pointer to **map[string]interface{}** | A description of the persistent volume&#39;s resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] 
**Cephfs** | Pointer to [**V1CephFSPersistentVolumeSource**](V1CephFSPersistentVolumeSource.md) |  | [optional] 
**Cinder** | Pointer to [**V1CinderPersistentVolumeSource**](V1CinderPersistentVolumeSource.md) |  | [optional] 
**ClaimRef** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**Csi** | Pointer to [**V1CSIPersistentVolumeSource**](V1CSIPersistentVolumeSource.md) |  | [optional] 
**Fc** | Pointer to [**V1FCVolumeSource**](V1FCVolumeSource.md) |  | [optional] 
**FlexVolume** | Pointer to [**V1FlexPersistentVolumeSource**](V1FlexPersistentVolumeSource.md) |  | [optional] 
**Flocker** | Pointer to [**V1FlockerVolumeSource**](V1FlockerVolumeSource.md) |  | [optional] 
**GcePersistentDisk** | Pointer to [**V1GCEPersistentDiskVolumeSource**](V1GCEPersistentDiskVolumeSource.md) |  | [optional] 
**Glusterfs** | Pointer to [**V1GlusterfsVolumeSource**](V1GlusterfsVolumeSource.md) |  | [optional] 
**HostPath** | Pointer to [**V1HostPathVolumeSource**](V1HostPathVolumeSource.md) |  | [optional] 
**Iscsi** | Pointer to [**V1ISCSIPersistentVolumeSource**](V1ISCSIPersistentVolumeSource.md) |  | [optional] 
**Local** | Pointer to [**V1LocalVolumeSource**](V1LocalVolumeSource.md) |  | [optional] 
**MountOptions** | Pointer to **[]string** | A list of mount options, e.g. [\&quot;ro\&quot;, \&quot;soft\&quot;]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options | [optional] 
**Nfs** | Pointer to [**V1NFSVolumeSource**](V1NFSVolumeSource.md) |  | [optional] 
**NodeAffinity** | Pointer to [**V1VolumeNodeAffinity**](V1VolumeNodeAffinity.md) |  | [optional] 
**PersistentVolumeReclaimPolicy** | Pointer to **string** | What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming | [optional] 
**PhotonPersistentDisk** | Pointer to [**V1PhotonPersistentDiskVolumeSource**](V1PhotonPersistentDiskVolumeSource.md) |  | [optional] 
**PortworxVolume** | Pointer to [**V1PortworxVolumeSource**](V1PortworxVolumeSource.md) |  | [optional] 
**Quobyte** | Pointer to [**V1QuobyteVolumeSource**](V1QuobyteVolumeSource.md) |  | [optional] 
**Rbd** | Pointer to [**V1RBDPersistentVolumeSource**](V1RBDPersistentVolumeSource.md) |  | [optional] 
**ScaleIO** | Pointer to [**V1ScaleIOPersistentVolumeSource**](V1ScaleIOPersistentVolumeSource.md) |  | [optional] 
**StorageClassName** | Pointer to **string** | Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass. | [optional] 
**Storageos** | Pointer to [**V1StorageOSPersistentVolumeSource**](V1StorageOSPersistentVolumeSource.md) |  | [optional] 
**VolumeMode** | Pointer to [**V1PersistentVolumeMode**](V1PersistentVolumeMode.md) |  | [optional] 
**VsphereVolume** | Pointer to [**V1VsphereVirtualDiskVolumeSource**](V1VsphereVirtualDiskVolumeSource.md) |  | [optional] 

## Methods

### NewV1PersistentVolumeSpec

`func NewV1PersistentVolumeSpec() *V1PersistentVolumeSpec`

NewV1PersistentVolumeSpec instantiates a new V1PersistentVolumeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PersistentVolumeSpecWithDefaults

`func NewV1PersistentVolumeSpecWithDefaults() *V1PersistentVolumeSpec`

NewV1PersistentVolumeSpecWithDefaults instantiates a new V1PersistentVolumeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *V1PersistentVolumeSpec) GetAccessModes() []V1PersistentVolumeAccessMode`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *V1PersistentVolumeSpec) GetAccessModesOk() (*[]V1PersistentVolumeAccessMode, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *V1PersistentVolumeSpec) SetAccessModes(v []V1PersistentVolumeAccessMode)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *V1PersistentVolumeSpec) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetAwsElasticBlockStore

`func (o *V1PersistentVolumeSpec) GetAwsElasticBlockStore() V1AWSElasticBlockStoreVolumeSource`

GetAwsElasticBlockStore returns the AwsElasticBlockStore field if non-nil, zero value otherwise.

### GetAwsElasticBlockStoreOk

`func (o *V1PersistentVolumeSpec) GetAwsElasticBlockStoreOk() (*V1AWSElasticBlockStoreVolumeSource, bool)`

GetAwsElasticBlockStoreOk returns a tuple with the AwsElasticBlockStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsElasticBlockStore

`func (o *V1PersistentVolumeSpec) SetAwsElasticBlockStore(v V1AWSElasticBlockStoreVolumeSource)`

SetAwsElasticBlockStore sets AwsElasticBlockStore field to given value.

### HasAwsElasticBlockStore

`func (o *V1PersistentVolumeSpec) HasAwsElasticBlockStore() bool`

HasAwsElasticBlockStore returns a boolean if a field has been set.

### GetAzureDisk

`func (o *V1PersistentVolumeSpec) GetAzureDisk() V1AzureDiskVolumeSource`

GetAzureDisk returns the AzureDisk field if non-nil, zero value otherwise.

### GetAzureDiskOk

`func (o *V1PersistentVolumeSpec) GetAzureDiskOk() (*V1AzureDiskVolumeSource, bool)`

GetAzureDiskOk returns a tuple with the AzureDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisk

`func (o *V1PersistentVolumeSpec) SetAzureDisk(v V1AzureDiskVolumeSource)`

SetAzureDisk sets AzureDisk field to given value.

### HasAzureDisk

`func (o *V1PersistentVolumeSpec) HasAzureDisk() bool`

HasAzureDisk returns a boolean if a field has been set.

### GetAzureFile

`func (o *V1PersistentVolumeSpec) GetAzureFile() V1AzureFilePersistentVolumeSource`

GetAzureFile returns the AzureFile field if non-nil, zero value otherwise.

### GetAzureFileOk

`func (o *V1PersistentVolumeSpec) GetAzureFileOk() (*V1AzureFilePersistentVolumeSource, bool)`

GetAzureFileOk returns a tuple with the AzureFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFile

`func (o *V1PersistentVolumeSpec) SetAzureFile(v V1AzureFilePersistentVolumeSource)`

SetAzureFile sets AzureFile field to given value.

### HasAzureFile

`func (o *V1PersistentVolumeSpec) HasAzureFile() bool`

HasAzureFile returns a boolean if a field has been set.

### GetCapacity

`func (o *V1PersistentVolumeSpec) GetCapacity() map[string]interface{}`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *V1PersistentVolumeSpec) GetCapacityOk() (*map[string]interface{}, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *V1PersistentVolumeSpec) SetCapacity(v map[string]interface{})`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *V1PersistentVolumeSpec) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCephfs

`func (o *V1PersistentVolumeSpec) GetCephfs() V1CephFSPersistentVolumeSource`

GetCephfs returns the Cephfs field if non-nil, zero value otherwise.

### GetCephfsOk

`func (o *V1PersistentVolumeSpec) GetCephfsOk() (*V1CephFSPersistentVolumeSource, bool)`

GetCephfsOk returns a tuple with the Cephfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephfs

`func (o *V1PersistentVolumeSpec) SetCephfs(v V1CephFSPersistentVolumeSource)`

SetCephfs sets Cephfs field to given value.

### HasCephfs

`func (o *V1PersistentVolumeSpec) HasCephfs() bool`

HasCephfs returns a boolean if a field has been set.

### GetCinder

`func (o *V1PersistentVolumeSpec) GetCinder() V1CinderPersistentVolumeSource`

GetCinder returns the Cinder field if non-nil, zero value otherwise.

### GetCinderOk

`func (o *V1PersistentVolumeSpec) GetCinderOk() (*V1CinderPersistentVolumeSource, bool)`

GetCinderOk returns a tuple with the Cinder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCinder

`func (o *V1PersistentVolumeSpec) SetCinder(v V1CinderPersistentVolumeSource)`

SetCinder sets Cinder field to given value.

### HasCinder

`func (o *V1PersistentVolumeSpec) HasCinder() bool`

HasCinder returns a boolean if a field has been set.

### GetClaimRef

`func (o *V1PersistentVolumeSpec) GetClaimRef() V1ObjectReference`

GetClaimRef returns the ClaimRef field if non-nil, zero value otherwise.

### GetClaimRefOk

`func (o *V1PersistentVolumeSpec) GetClaimRefOk() (*V1ObjectReference, bool)`

GetClaimRefOk returns a tuple with the ClaimRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRef

`func (o *V1PersistentVolumeSpec) SetClaimRef(v V1ObjectReference)`

SetClaimRef sets ClaimRef field to given value.

### HasClaimRef

`func (o *V1PersistentVolumeSpec) HasClaimRef() bool`

HasClaimRef returns a boolean if a field has been set.

### GetCsi

`func (o *V1PersistentVolumeSpec) GetCsi() V1CSIPersistentVolumeSource`

GetCsi returns the Csi field if non-nil, zero value otherwise.

### GetCsiOk

`func (o *V1PersistentVolumeSpec) GetCsiOk() (*V1CSIPersistentVolumeSource, bool)`

GetCsiOk returns a tuple with the Csi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsi

`func (o *V1PersistentVolumeSpec) SetCsi(v V1CSIPersistentVolumeSource)`

SetCsi sets Csi field to given value.

### HasCsi

`func (o *V1PersistentVolumeSpec) HasCsi() bool`

HasCsi returns a boolean if a field has been set.

### GetFc

`func (o *V1PersistentVolumeSpec) GetFc() V1FCVolumeSource`

GetFc returns the Fc field if non-nil, zero value otherwise.

### GetFcOk

`func (o *V1PersistentVolumeSpec) GetFcOk() (*V1FCVolumeSource, bool)`

GetFcOk returns a tuple with the Fc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFc

`func (o *V1PersistentVolumeSpec) SetFc(v V1FCVolumeSource)`

SetFc sets Fc field to given value.

### HasFc

`func (o *V1PersistentVolumeSpec) HasFc() bool`

HasFc returns a boolean if a field has been set.

### GetFlexVolume

`func (o *V1PersistentVolumeSpec) GetFlexVolume() V1FlexPersistentVolumeSource`

GetFlexVolume returns the FlexVolume field if non-nil, zero value otherwise.

### GetFlexVolumeOk

`func (o *V1PersistentVolumeSpec) GetFlexVolumeOk() (*V1FlexPersistentVolumeSource, bool)`

GetFlexVolumeOk returns a tuple with the FlexVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexVolume

`func (o *V1PersistentVolumeSpec) SetFlexVolume(v V1FlexPersistentVolumeSource)`

SetFlexVolume sets FlexVolume field to given value.

### HasFlexVolume

`func (o *V1PersistentVolumeSpec) HasFlexVolume() bool`

HasFlexVolume returns a boolean if a field has been set.

### GetFlocker

`func (o *V1PersistentVolumeSpec) GetFlocker() V1FlockerVolumeSource`

GetFlocker returns the Flocker field if non-nil, zero value otherwise.

### GetFlockerOk

`func (o *V1PersistentVolumeSpec) GetFlockerOk() (*V1FlockerVolumeSource, bool)`

GetFlockerOk returns a tuple with the Flocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlocker

`func (o *V1PersistentVolumeSpec) SetFlocker(v V1FlockerVolumeSource)`

SetFlocker sets Flocker field to given value.

### HasFlocker

`func (o *V1PersistentVolumeSpec) HasFlocker() bool`

HasFlocker returns a boolean if a field has been set.

### GetGcePersistentDisk

`func (o *V1PersistentVolumeSpec) GetGcePersistentDisk() V1GCEPersistentDiskVolumeSource`

GetGcePersistentDisk returns the GcePersistentDisk field if non-nil, zero value otherwise.

### GetGcePersistentDiskOk

`func (o *V1PersistentVolumeSpec) GetGcePersistentDiskOk() (*V1GCEPersistentDiskVolumeSource, bool)`

GetGcePersistentDiskOk returns a tuple with the GcePersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcePersistentDisk

`func (o *V1PersistentVolumeSpec) SetGcePersistentDisk(v V1GCEPersistentDiskVolumeSource)`

SetGcePersistentDisk sets GcePersistentDisk field to given value.

### HasGcePersistentDisk

`func (o *V1PersistentVolumeSpec) HasGcePersistentDisk() bool`

HasGcePersistentDisk returns a boolean if a field has been set.

### GetGlusterfs

`func (o *V1PersistentVolumeSpec) GetGlusterfs() V1GlusterfsVolumeSource`

GetGlusterfs returns the Glusterfs field if non-nil, zero value otherwise.

### GetGlusterfsOk

`func (o *V1PersistentVolumeSpec) GetGlusterfsOk() (*V1GlusterfsVolumeSource, bool)`

GetGlusterfsOk returns a tuple with the Glusterfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlusterfs

`func (o *V1PersistentVolumeSpec) SetGlusterfs(v V1GlusterfsVolumeSource)`

SetGlusterfs sets Glusterfs field to given value.

### HasGlusterfs

`func (o *V1PersistentVolumeSpec) HasGlusterfs() bool`

HasGlusterfs returns a boolean if a field has been set.

### GetHostPath

`func (o *V1PersistentVolumeSpec) GetHostPath() V1HostPathVolumeSource`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *V1PersistentVolumeSpec) GetHostPathOk() (*V1HostPathVolumeSource, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *V1PersistentVolumeSpec) SetHostPath(v V1HostPathVolumeSource)`

SetHostPath sets HostPath field to given value.

### HasHostPath

`func (o *V1PersistentVolumeSpec) HasHostPath() bool`

HasHostPath returns a boolean if a field has been set.

### GetIscsi

`func (o *V1PersistentVolumeSpec) GetIscsi() V1ISCSIPersistentVolumeSource`

GetIscsi returns the Iscsi field if non-nil, zero value otherwise.

### GetIscsiOk

`func (o *V1PersistentVolumeSpec) GetIscsiOk() (*V1ISCSIPersistentVolumeSource, bool)`

GetIscsiOk returns a tuple with the Iscsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsi

`func (o *V1PersistentVolumeSpec) SetIscsi(v V1ISCSIPersistentVolumeSource)`

SetIscsi sets Iscsi field to given value.

### HasIscsi

`func (o *V1PersistentVolumeSpec) HasIscsi() bool`

HasIscsi returns a boolean if a field has been set.

### GetLocal

`func (o *V1PersistentVolumeSpec) GetLocal() V1LocalVolumeSource`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *V1PersistentVolumeSpec) GetLocalOk() (*V1LocalVolumeSource, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *V1PersistentVolumeSpec) SetLocal(v V1LocalVolumeSource)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *V1PersistentVolumeSpec) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetMountOptions

`func (o *V1PersistentVolumeSpec) GetMountOptions() []string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *V1PersistentVolumeSpec) GetMountOptionsOk() (*[]string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *V1PersistentVolumeSpec) SetMountOptions(v []string)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *V1PersistentVolumeSpec) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetNfs

`func (o *V1PersistentVolumeSpec) GetNfs() V1NFSVolumeSource`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *V1PersistentVolumeSpec) GetNfsOk() (*V1NFSVolumeSource, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *V1PersistentVolumeSpec) SetNfs(v V1NFSVolumeSource)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *V1PersistentVolumeSpec) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetNodeAffinity

`func (o *V1PersistentVolumeSpec) GetNodeAffinity() V1VolumeNodeAffinity`

GetNodeAffinity returns the NodeAffinity field if non-nil, zero value otherwise.

### GetNodeAffinityOk

`func (o *V1PersistentVolumeSpec) GetNodeAffinityOk() (*V1VolumeNodeAffinity, bool)`

GetNodeAffinityOk returns a tuple with the NodeAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinity

`func (o *V1PersistentVolumeSpec) SetNodeAffinity(v V1VolumeNodeAffinity)`

SetNodeAffinity sets NodeAffinity field to given value.

### HasNodeAffinity

`func (o *V1PersistentVolumeSpec) HasNodeAffinity() bool`

HasNodeAffinity returns a boolean if a field has been set.

### GetPersistentVolumeReclaimPolicy

`func (o *V1PersistentVolumeSpec) GetPersistentVolumeReclaimPolicy() string`

GetPersistentVolumeReclaimPolicy returns the PersistentVolumeReclaimPolicy field if non-nil, zero value otherwise.

### GetPersistentVolumeReclaimPolicyOk

`func (o *V1PersistentVolumeSpec) GetPersistentVolumeReclaimPolicyOk() (*string, bool)`

GetPersistentVolumeReclaimPolicyOk returns a tuple with the PersistentVolumeReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeReclaimPolicy

`func (o *V1PersistentVolumeSpec) SetPersistentVolumeReclaimPolicy(v string)`

SetPersistentVolumeReclaimPolicy sets PersistentVolumeReclaimPolicy field to given value.

### HasPersistentVolumeReclaimPolicy

`func (o *V1PersistentVolumeSpec) HasPersistentVolumeReclaimPolicy() bool`

HasPersistentVolumeReclaimPolicy returns a boolean if a field has been set.

### GetPhotonPersistentDisk

`func (o *V1PersistentVolumeSpec) GetPhotonPersistentDisk() V1PhotonPersistentDiskVolumeSource`

GetPhotonPersistentDisk returns the PhotonPersistentDisk field if non-nil, zero value otherwise.

### GetPhotonPersistentDiskOk

`func (o *V1PersistentVolumeSpec) GetPhotonPersistentDiskOk() (*V1PhotonPersistentDiskVolumeSource, bool)`

GetPhotonPersistentDiskOk returns a tuple with the PhotonPersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonPersistentDisk

`func (o *V1PersistentVolumeSpec) SetPhotonPersistentDisk(v V1PhotonPersistentDiskVolumeSource)`

SetPhotonPersistentDisk sets PhotonPersistentDisk field to given value.

### HasPhotonPersistentDisk

`func (o *V1PersistentVolumeSpec) HasPhotonPersistentDisk() bool`

HasPhotonPersistentDisk returns a boolean if a field has been set.

### GetPortworxVolume

`func (o *V1PersistentVolumeSpec) GetPortworxVolume() V1PortworxVolumeSource`

GetPortworxVolume returns the PortworxVolume field if non-nil, zero value otherwise.

### GetPortworxVolumeOk

`func (o *V1PersistentVolumeSpec) GetPortworxVolumeOk() (*V1PortworxVolumeSource, bool)`

GetPortworxVolumeOk returns a tuple with the PortworxVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortworxVolume

`func (o *V1PersistentVolumeSpec) SetPortworxVolume(v V1PortworxVolumeSource)`

SetPortworxVolume sets PortworxVolume field to given value.

### HasPortworxVolume

`func (o *V1PersistentVolumeSpec) HasPortworxVolume() bool`

HasPortworxVolume returns a boolean if a field has been set.

### GetQuobyte

`func (o *V1PersistentVolumeSpec) GetQuobyte() V1QuobyteVolumeSource`

GetQuobyte returns the Quobyte field if non-nil, zero value otherwise.

### GetQuobyteOk

`func (o *V1PersistentVolumeSpec) GetQuobyteOk() (*V1QuobyteVolumeSource, bool)`

GetQuobyteOk returns a tuple with the Quobyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuobyte

`func (o *V1PersistentVolumeSpec) SetQuobyte(v V1QuobyteVolumeSource)`

SetQuobyte sets Quobyte field to given value.

### HasQuobyte

`func (o *V1PersistentVolumeSpec) HasQuobyte() bool`

HasQuobyte returns a boolean if a field has been set.

### GetRbd

`func (o *V1PersistentVolumeSpec) GetRbd() V1RBDPersistentVolumeSource`

GetRbd returns the Rbd field if non-nil, zero value otherwise.

### GetRbdOk

`func (o *V1PersistentVolumeSpec) GetRbdOk() (*V1RBDPersistentVolumeSource, bool)`

GetRbdOk returns a tuple with the Rbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbd

`func (o *V1PersistentVolumeSpec) SetRbd(v V1RBDPersistentVolumeSource)`

SetRbd sets Rbd field to given value.

### HasRbd

`func (o *V1PersistentVolumeSpec) HasRbd() bool`

HasRbd returns a boolean if a field has been set.

### GetScaleIO

`func (o *V1PersistentVolumeSpec) GetScaleIO() V1ScaleIOPersistentVolumeSource`

GetScaleIO returns the ScaleIO field if non-nil, zero value otherwise.

### GetScaleIOOk

`func (o *V1PersistentVolumeSpec) GetScaleIOOk() (*V1ScaleIOPersistentVolumeSource, bool)`

GetScaleIOOk returns a tuple with the ScaleIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIO

`func (o *V1PersistentVolumeSpec) SetScaleIO(v V1ScaleIOPersistentVolumeSource)`

SetScaleIO sets ScaleIO field to given value.

### HasScaleIO

`func (o *V1PersistentVolumeSpec) HasScaleIO() bool`

HasScaleIO returns a boolean if a field has been set.

### GetStorageClassName

`func (o *V1PersistentVolumeSpec) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *V1PersistentVolumeSpec) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *V1PersistentVolumeSpec) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *V1PersistentVolumeSpec) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetStorageos

`func (o *V1PersistentVolumeSpec) GetStorageos() V1StorageOSPersistentVolumeSource`

GetStorageos returns the Storageos field if non-nil, zero value otherwise.

### GetStorageosOk

`func (o *V1PersistentVolumeSpec) GetStorageosOk() (*V1StorageOSPersistentVolumeSource, bool)`

GetStorageosOk returns a tuple with the Storageos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageos

`func (o *V1PersistentVolumeSpec) SetStorageos(v V1StorageOSPersistentVolumeSource)`

SetStorageos sets Storageos field to given value.

### HasStorageos

`func (o *V1PersistentVolumeSpec) HasStorageos() bool`

HasStorageos returns a boolean if a field has been set.

### GetVolumeMode

`func (o *V1PersistentVolumeSpec) GetVolumeMode() V1PersistentVolumeMode`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *V1PersistentVolumeSpec) GetVolumeModeOk() (*V1PersistentVolumeMode, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *V1PersistentVolumeSpec) SetVolumeMode(v V1PersistentVolumeMode)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *V1PersistentVolumeSpec) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVsphereVolume

`func (o *V1PersistentVolumeSpec) GetVsphereVolume() V1VsphereVirtualDiskVolumeSource`

GetVsphereVolume returns the VsphereVolume field if non-nil, zero value otherwise.

### GetVsphereVolumeOk

`func (o *V1PersistentVolumeSpec) GetVsphereVolumeOk() (*V1VsphereVirtualDiskVolumeSource, bool)`

GetVsphereVolumeOk returns a tuple with the VsphereVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereVolume

`func (o *V1PersistentVolumeSpec) SetVsphereVolume(v V1VsphereVirtualDiskVolumeSource)`

SetVsphereVolume sets VsphereVolume field to given value.

### HasVsphereVolume

`func (o *V1PersistentVolumeSpec) HasVsphereVolume() bool`

HasVsphereVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


