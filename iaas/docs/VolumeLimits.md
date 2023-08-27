# VolumeLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTotalBackupGigabytes** | Pointer to **int32** |  | [optional] 
**MaxTotalBackups** | Pointer to **int32** |  | [optional] 
**MaxTotalSnapshots** | Pointer to **int32** |  | [optional] 
**MaxTotalVolumeGigabytes** | Pointer to **int32** |  | [optional] 
**MaxTotalVolumes** | Pointer to **int32** |  | [optional] 
**TotalBackupGigabytesUsed** | Pointer to **int32** |  | [optional] 
**TotalBackupsUsed** | Pointer to **int32** |  | [optional] 
**TotalGigabytesUsed** | Pointer to **int32** |  | [optional] 
**TotalSnapshotsUsed** | Pointer to **int32** |  | [optional] 
**TotalVolumesUsed** | Pointer to **int32** |  | [optional] 

## Methods

### NewVolumeLimits

`func NewVolumeLimits() *VolumeLimits`

NewVolumeLimits instantiates a new VolumeLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeLimitsWithDefaults

`func NewVolumeLimitsWithDefaults() *VolumeLimits`

NewVolumeLimitsWithDefaults instantiates a new VolumeLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTotalBackupGigabytes

`func (o *VolumeLimits) GetMaxTotalBackupGigabytes() int32`

GetMaxTotalBackupGigabytes returns the MaxTotalBackupGigabytes field if non-nil, zero value otherwise.

### GetMaxTotalBackupGigabytesOk

`func (o *VolumeLimits) GetMaxTotalBackupGigabytesOk() (*int32, bool)`

GetMaxTotalBackupGigabytesOk returns a tuple with the MaxTotalBackupGigabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalBackupGigabytes

`func (o *VolumeLimits) SetMaxTotalBackupGigabytes(v int32)`

SetMaxTotalBackupGigabytes sets MaxTotalBackupGigabytes field to given value.

### HasMaxTotalBackupGigabytes

`func (o *VolumeLimits) HasMaxTotalBackupGigabytes() bool`

HasMaxTotalBackupGigabytes returns a boolean if a field has been set.

### GetMaxTotalBackups

`func (o *VolumeLimits) GetMaxTotalBackups() int32`

GetMaxTotalBackups returns the MaxTotalBackups field if non-nil, zero value otherwise.

### GetMaxTotalBackupsOk

`func (o *VolumeLimits) GetMaxTotalBackupsOk() (*int32, bool)`

GetMaxTotalBackupsOk returns a tuple with the MaxTotalBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalBackups

`func (o *VolumeLimits) SetMaxTotalBackups(v int32)`

SetMaxTotalBackups sets MaxTotalBackups field to given value.

### HasMaxTotalBackups

`func (o *VolumeLimits) HasMaxTotalBackups() bool`

HasMaxTotalBackups returns a boolean if a field has been set.

### GetMaxTotalSnapshots

`func (o *VolumeLimits) GetMaxTotalSnapshots() int32`

GetMaxTotalSnapshots returns the MaxTotalSnapshots field if non-nil, zero value otherwise.

### GetMaxTotalSnapshotsOk

`func (o *VolumeLimits) GetMaxTotalSnapshotsOk() (*int32, bool)`

GetMaxTotalSnapshotsOk returns a tuple with the MaxTotalSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSnapshots

`func (o *VolumeLimits) SetMaxTotalSnapshots(v int32)`

SetMaxTotalSnapshots sets MaxTotalSnapshots field to given value.

### HasMaxTotalSnapshots

`func (o *VolumeLimits) HasMaxTotalSnapshots() bool`

HasMaxTotalSnapshots returns a boolean if a field has been set.

### GetMaxTotalVolumeGigabytes

`func (o *VolumeLimits) GetMaxTotalVolumeGigabytes() int32`

GetMaxTotalVolumeGigabytes returns the MaxTotalVolumeGigabytes field if non-nil, zero value otherwise.

### GetMaxTotalVolumeGigabytesOk

`func (o *VolumeLimits) GetMaxTotalVolumeGigabytesOk() (*int32, bool)`

GetMaxTotalVolumeGigabytesOk returns a tuple with the MaxTotalVolumeGigabytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalVolumeGigabytes

`func (o *VolumeLimits) SetMaxTotalVolumeGigabytes(v int32)`

SetMaxTotalVolumeGigabytes sets MaxTotalVolumeGigabytes field to given value.

### HasMaxTotalVolumeGigabytes

`func (o *VolumeLimits) HasMaxTotalVolumeGigabytes() bool`

HasMaxTotalVolumeGigabytes returns a boolean if a field has been set.

### GetMaxTotalVolumes

`func (o *VolumeLimits) GetMaxTotalVolumes() int32`

GetMaxTotalVolumes returns the MaxTotalVolumes field if non-nil, zero value otherwise.

### GetMaxTotalVolumesOk

`func (o *VolumeLimits) GetMaxTotalVolumesOk() (*int32, bool)`

GetMaxTotalVolumesOk returns a tuple with the MaxTotalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalVolumes

`func (o *VolumeLimits) SetMaxTotalVolumes(v int32)`

SetMaxTotalVolumes sets MaxTotalVolumes field to given value.

### HasMaxTotalVolumes

`func (o *VolumeLimits) HasMaxTotalVolumes() bool`

HasMaxTotalVolumes returns a boolean if a field has been set.

### GetTotalBackupGigabytesUsed

`func (o *VolumeLimits) GetTotalBackupGigabytesUsed() int32`

GetTotalBackupGigabytesUsed returns the TotalBackupGigabytesUsed field if non-nil, zero value otherwise.

### GetTotalBackupGigabytesUsedOk

`func (o *VolumeLimits) GetTotalBackupGigabytesUsedOk() (*int32, bool)`

GetTotalBackupGigabytesUsedOk returns a tuple with the TotalBackupGigabytesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupGigabytesUsed

`func (o *VolumeLimits) SetTotalBackupGigabytesUsed(v int32)`

SetTotalBackupGigabytesUsed sets TotalBackupGigabytesUsed field to given value.

### HasTotalBackupGigabytesUsed

`func (o *VolumeLimits) HasTotalBackupGigabytesUsed() bool`

HasTotalBackupGigabytesUsed returns a boolean if a field has been set.

### GetTotalBackupsUsed

`func (o *VolumeLimits) GetTotalBackupsUsed() int32`

GetTotalBackupsUsed returns the TotalBackupsUsed field if non-nil, zero value otherwise.

### GetTotalBackupsUsedOk

`func (o *VolumeLimits) GetTotalBackupsUsedOk() (*int32, bool)`

GetTotalBackupsUsedOk returns a tuple with the TotalBackupsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackupsUsed

`func (o *VolumeLimits) SetTotalBackupsUsed(v int32)`

SetTotalBackupsUsed sets TotalBackupsUsed field to given value.

### HasTotalBackupsUsed

`func (o *VolumeLimits) HasTotalBackupsUsed() bool`

HasTotalBackupsUsed returns a boolean if a field has been set.

### GetTotalGigabytesUsed

`func (o *VolumeLimits) GetTotalGigabytesUsed() int32`

GetTotalGigabytesUsed returns the TotalGigabytesUsed field if non-nil, zero value otherwise.

### GetTotalGigabytesUsedOk

`func (o *VolumeLimits) GetTotalGigabytesUsedOk() (*int32, bool)`

GetTotalGigabytesUsedOk returns a tuple with the TotalGigabytesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGigabytesUsed

`func (o *VolumeLimits) SetTotalGigabytesUsed(v int32)`

SetTotalGigabytesUsed sets TotalGigabytesUsed field to given value.

### HasTotalGigabytesUsed

`func (o *VolumeLimits) HasTotalGigabytesUsed() bool`

HasTotalGigabytesUsed returns a boolean if a field has been set.

### GetTotalSnapshotsUsed

`func (o *VolumeLimits) GetTotalSnapshotsUsed() int32`

GetTotalSnapshotsUsed returns the TotalSnapshotsUsed field if non-nil, zero value otherwise.

### GetTotalSnapshotsUsedOk

`func (o *VolumeLimits) GetTotalSnapshotsUsedOk() (*int32, bool)`

GetTotalSnapshotsUsedOk returns a tuple with the TotalSnapshotsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSnapshotsUsed

`func (o *VolumeLimits) SetTotalSnapshotsUsed(v int32)`

SetTotalSnapshotsUsed sets TotalSnapshotsUsed field to given value.

### HasTotalSnapshotsUsed

`func (o *VolumeLimits) HasTotalSnapshotsUsed() bool`

HasTotalSnapshotsUsed returns a boolean if a field has been set.

### GetTotalVolumesUsed

`func (o *VolumeLimits) GetTotalVolumesUsed() int32`

GetTotalVolumesUsed returns the TotalVolumesUsed field if non-nil, zero value otherwise.

### GetTotalVolumesUsedOk

`func (o *VolumeLimits) GetTotalVolumesUsedOk() (*int32, bool)`

GetTotalVolumesUsedOk returns a tuple with the TotalVolumesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumesUsed

`func (o *VolumeLimits) SetTotalVolumesUsed(v int32)`

SetTotalVolumesUsed sets TotalVolumesUsed field to given value.

### HasTotalVolumesUsed

`func (o *VolumeLimits) HasTotalVolumesUsed() bool`

HasTotalVolumesUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


