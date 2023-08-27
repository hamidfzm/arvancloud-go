# V1ISCSIPersistentVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChapAuthDiscovery** | Pointer to **bool** | whether support iSCSI Discovery CHAP authentication | [optional] 
**ChapAuthSession** | Pointer to **bool** | whether support iSCSI Session CHAP authentication | [optional] 
**FsType** | Pointer to **string** | Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi | [optional] 
**InitiatorName** | Pointer to **string** | Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface &lt;target portal&gt;:&lt;volume name&gt; will be created for the connection. | [optional] 
**Iqn** | **string** | Target iSCSI Qualified Name. | 
**IscsiInterface** | Pointer to **string** | iSCSI Interface Name that uses an iSCSI transport. Defaults to &#39;default&#39; (tcp). | [optional] 
**Lun** | **int32** | iSCSI Target Lun number. | 
**Portals** | Pointer to **[]string** | iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | [optional] 
**ReadOnly** | Pointer to **bool** | ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. | [optional] 
**SecretRef** | Pointer to [**V1SecretReference**](V1SecretReference.md) |  | [optional] 
**TargetPortal** | **string** | iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). | 

## Methods

### NewV1ISCSIPersistentVolumeSource

`func NewV1ISCSIPersistentVolumeSource(iqn string, lun int32, targetPortal string, ) *V1ISCSIPersistentVolumeSource`

NewV1ISCSIPersistentVolumeSource instantiates a new V1ISCSIPersistentVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ISCSIPersistentVolumeSourceWithDefaults

`func NewV1ISCSIPersistentVolumeSourceWithDefaults() *V1ISCSIPersistentVolumeSource`

NewV1ISCSIPersistentVolumeSourceWithDefaults instantiates a new V1ISCSIPersistentVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChapAuthDiscovery

`func (o *V1ISCSIPersistentVolumeSource) GetChapAuthDiscovery() bool`

GetChapAuthDiscovery returns the ChapAuthDiscovery field if non-nil, zero value otherwise.

### GetChapAuthDiscoveryOk

`func (o *V1ISCSIPersistentVolumeSource) GetChapAuthDiscoveryOk() (*bool, bool)`

GetChapAuthDiscoveryOk returns a tuple with the ChapAuthDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapAuthDiscovery

`func (o *V1ISCSIPersistentVolumeSource) SetChapAuthDiscovery(v bool)`

SetChapAuthDiscovery sets ChapAuthDiscovery field to given value.

### HasChapAuthDiscovery

`func (o *V1ISCSIPersistentVolumeSource) HasChapAuthDiscovery() bool`

HasChapAuthDiscovery returns a boolean if a field has been set.

### GetChapAuthSession

`func (o *V1ISCSIPersistentVolumeSource) GetChapAuthSession() bool`

GetChapAuthSession returns the ChapAuthSession field if non-nil, zero value otherwise.

### GetChapAuthSessionOk

`func (o *V1ISCSIPersistentVolumeSource) GetChapAuthSessionOk() (*bool, bool)`

GetChapAuthSessionOk returns a tuple with the ChapAuthSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapAuthSession

`func (o *V1ISCSIPersistentVolumeSource) SetChapAuthSession(v bool)`

SetChapAuthSession sets ChapAuthSession field to given value.

### HasChapAuthSession

`func (o *V1ISCSIPersistentVolumeSource) HasChapAuthSession() bool`

HasChapAuthSession returns a boolean if a field has been set.

### GetFsType

`func (o *V1ISCSIPersistentVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *V1ISCSIPersistentVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *V1ISCSIPersistentVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *V1ISCSIPersistentVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetInitiatorName

`func (o *V1ISCSIPersistentVolumeSource) GetInitiatorName() string`

GetInitiatorName returns the InitiatorName field if non-nil, zero value otherwise.

### GetInitiatorNameOk

`func (o *V1ISCSIPersistentVolumeSource) GetInitiatorNameOk() (*string, bool)`

GetInitiatorNameOk returns a tuple with the InitiatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorName

`func (o *V1ISCSIPersistentVolumeSource) SetInitiatorName(v string)`

SetInitiatorName sets InitiatorName field to given value.

### HasInitiatorName

`func (o *V1ISCSIPersistentVolumeSource) HasInitiatorName() bool`

HasInitiatorName returns a boolean if a field has been set.

### GetIqn

`func (o *V1ISCSIPersistentVolumeSource) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *V1ISCSIPersistentVolumeSource) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *V1ISCSIPersistentVolumeSource) SetIqn(v string)`

SetIqn sets Iqn field to given value.


### GetIscsiInterface

`func (o *V1ISCSIPersistentVolumeSource) GetIscsiInterface() string`

GetIscsiInterface returns the IscsiInterface field if non-nil, zero value otherwise.

### GetIscsiInterfaceOk

`func (o *V1ISCSIPersistentVolumeSource) GetIscsiInterfaceOk() (*string, bool)`

GetIscsiInterfaceOk returns a tuple with the IscsiInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiInterface

`func (o *V1ISCSIPersistentVolumeSource) SetIscsiInterface(v string)`

SetIscsiInterface sets IscsiInterface field to given value.

### HasIscsiInterface

`func (o *V1ISCSIPersistentVolumeSource) HasIscsiInterface() bool`

HasIscsiInterface returns a boolean if a field has been set.

### GetLun

`func (o *V1ISCSIPersistentVolumeSource) GetLun() int32`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *V1ISCSIPersistentVolumeSource) GetLunOk() (*int32, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *V1ISCSIPersistentVolumeSource) SetLun(v int32)`

SetLun sets Lun field to given value.


### GetPortals

`func (o *V1ISCSIPersistentVolumeSource) GetPortals() []string`

GetPortals returns the Portals field if non-nil, zero value otherwise.

### GetPortalsOk

`func (o *V1ISCSIPersistentVolumeSource) GetPortalsOk() (*[]string, bool)`

GetPortalsOk returns a tuple with the Portals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortals

`func (o *V1ISCSIPersistentVolumeSource) SetPortals(v []string)`

SetPortals sets Portals field to given value.

### HasPortals

`func (o *V1ISCSIPersistentVolumeSource) HasPortals() bool`

HasPortals returns a boolean if a field has been set.

### GetReadOnly

`func (o *V1ISCSIPersistentVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *V1ISCSIPersistentVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *V1ISCSIPersistentVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *V1ISCSIPersistentVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSecretRef

`func (o *V1ISCSIPersistentVolumeSource) GetSecretRef() V1SecretReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *V1ISCSIPersistentVolumeSource) GetSecretRefOk() (*V1SecretReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *V1ISCSIPersistentVolumeSource) SetSecretRef(v V1SecretReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *V1ISCSIPersistentVolumeSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetTargetPortal

`func (o *V1ISCSIPersistentVolumeSource) GetTargetPortal() string`

GetTargetPortal returns the TargetPortal field if non-nil, zero value otherwise.

### GetTargetPortalOk

`func (o *V1ISCSIPersistentVolumeSource) GetTargetPortalOk() (*string, bool)`

GetTargetPortalOk returns a tuple with the TargetPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPortal

`func (o *V1ISCSIPersistentVolumeSource) SetTargetPortal(v string)`

SetTargetPortal sets TargetPortal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


