# CreateServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**CreateType** | Pointer to **string** |  | [optional] 
**DiskSize** | Pointer to **int64** |  | [optional] 
**FlavorId** | Pointer to **string** |  | [optional] 
**HaEnabled** | Pointer to **bool** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**InitScript** | Pointer to **string** |  | [optional] 
**IsSandbox** | Pointer to **bool** |  | [optional] 
**KeyName** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkId** | Pointer to **string** |  | [optional] 
**NetworkIds** | Pointer to **[]string** |  | [optional] 
**OsVolumeId** | Pointer to **string** |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupName**](SecurityGroupName.md) |  | [optional] 
**ServerVolumes** | Pointer to [**[]ServerVolume**](ServerVolume.md) |  | [optional] 
**SshKey** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateServerRequest

`func NewCreateServerRequest() *CreateServerRequest`

NewCreateServerRequest instantiates a new CreateServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServerRequestWithDefaults

`func NewCreateServerRequestWithDefaults() *CreateServerRequest`

NewCreateServerRequestWithDefaults instantiates a new CreateServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *CreateServerRequest) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *CreateServerRequest) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *CreateServerRequest) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *CreateServerRequest) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetCount

`func (o *CreateServerRequest) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateServerRequest) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateServerRequest) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateServerRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateType

`func (o *CreateServerRequest) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *CreateServerRequest) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *CreateServerRequest) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *CreateServerRequest) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetDiskSize

`func (o *CreateServerRequest) GetDiskSize() int64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *CreateServerRequest) GetDiskSizeOk() (*int64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *CreateServerRequest) SetDiskSize(v int64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *CreateServerRequest) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetFlavorId

`func (o *CreateServerRequest) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreateServerRequest) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreateServerRequest) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *CreateServerRequest) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### GetHaEnabled

`func (o *CreateServerRequest) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *CreateServerRequest) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *CreateServerRequest) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *CreateServerRequest) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetImageId

`func (o *CreateServerRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateServerRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateServerRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateServerRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetInitScript

`func (o *CreateServerRequest) GetInitScript() string`

GetInitScript returns the InitScript field if non-nil, zero value otherwise.

### GetInitScriptOk

`func (o *CreateServerRequest) GetInitScriptOk() (*string, bool)`

GetInitScriptOk returns a tuple with the InitScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitScript

`func (o *CreateServerRequest) SetInitScript(v string)`

SetInitScript sets InitScript field to given value.

### HasInitScript

`func (o *CreateServerRequest) HasInitScript() bool`

HasInitScript returns a boolean if a field has been set.

### GetIsSandbox

`func (o *CreateServerRequest) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *CreateServerRequest) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *CreateServerRequest) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *CreateServerRequest) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.

### GetKeyName

`func (o *CreateServerRequest) GetKeyName() map[string]interface{}`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateServerRequest) GetKeyNameOk() (*map[string]interface{}, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateServerRequest) SetKeyName(v map[string]interface{})`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *CreateServerRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetName

`func (o *CreateServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateServerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *CreateServerRequest) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreateServerRequest) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreateServerRequest) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CreateServerRequest) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkIds

`func (o *CreateServerRequest) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *CreateServerRequest) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *CreateServerRequest) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *CreateServerRequest) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.

### GetOsVolumeId

`func (o *CreateServerRequest) GetOsVolumeId() string`

GetOsVolumeId returns the OsVolumeId field if non-nil, zero value otherwise.

### GetOsVolumeIdOk

`func (o *CreateServerRequest) GetOsVolumeIdOk() (*string, bool)`

GetOsVolumeIdOk returns a tuple with the OsVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVolumeId

`func (o *CreateServerRequest) SetOsVolumeId(v string)`

SetOsVolumeId sets OsVolumeId field to given value.

### HasOsVolumeId

`func (o *CreateServerRequest) HasOsVolumeId() bool`

HasOsVolumeId returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CreateServerRequest) GetSecurityGroups() []SecurityGroupName`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateServerRequest) GetSecurityGroupsOk() (*[]SecurityGroupName, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateServerRequest) SetSecurityGroups(v []SecurityGroupName)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateServerRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetServerVolumes

`func (o *CreateServerRequest) GetServerVolumes() []ServerVolume`

GetServerVolumes returns the ServerVolumes field if non-nil, zero value otherwise.

### GetServerVolumesOk

`func (o *CreateServerRequest) GetServerVolumesOk() (*[]ServerVolume, bool)`

GetServerVolumesOk returns a tuple with the ServerVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVolumes

`func (o *CreateServerRequest) SetServerVolumes(v []ServerVolume)`

SetServerVolumes sets ServerVolumes field to given value.

### HasServerVolumes

`func (o *CreateServerRequest) HasServerVolumes() bool`

HasServerVolumes returns a boolean if a field has been set.

### GetSshKey

`func (o *CreateServerRequest) GetSshKey() bool`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *CreateServerRequest) GetSshKeyOk() (*bool, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *CreateServerRequest) SetSshKey(v bool)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *CreateServerRequest) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


