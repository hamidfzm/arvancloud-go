# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxImageMeta** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxPersonality** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxPersonalitySize** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxSecurityGroupRules** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxSecurityGroups** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxServerGroupMembers** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxServerGroups** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxServerMeta** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxTotalCores** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxTotalFloatingIps** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxTotalInstances** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxTotalKeypairs** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**MaxTotalRamSize** | Pointer to [**QuotaValue**](QuotaValue.md) |  | [optional] 
**TotalCoresUsed** | Pointer to **int64** |  | [optional] 
**TotalFloatingIpsUsed** | Pointer to **int64** |  | [optional] 
**TotalInstancesUsed** | Pointer to **int64** |  | [optional] 
**TotalRamUsed** | Pointer to **int64** |  | [optional] 
**TotalSecurityGroupsUsed** | Pointer to **int64** |  | [optional] 
**TotalServerGroupsUsed** | Pointer to **int64** |  | [optional] 

## Methods

### NewQuota

`func NewQuota() *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxImageMeta

`func (o *Quota) GetMaxImageMeta() QuotaValue`

GetMaxImageMeta returns the MaxImageMeta field if non-nil, zero value otherwise.

### GetMaxImageMetaOk

`func (o *Quota) GetMaxImageMetaOk() (*QuotaValue, bool)`

GetMaxImageMetaOk returns a tuple with the MaxImageMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxImageMeta

`func (o *Quota) SetMaxImageMeta(v QuotaValue)`

SetMaxImageMeta sets MaxImageMeta field to given value.

### HasMaxImageMeta

`func (o *Quota) HasMaxImageMeta() bool`

HasMaxImageMeta returns a boolean if a field has been set.

### GetMaxPersonality

`func (o *Quota) GetMaxPersonality() QuotaValue`

GetMaxPersonality returns the MaxPersonality field if non-nil, zero value otherwise.

### GetMaxPersonalityOk

`func (o *Quota) GetMaxPersonalityOk() (*QuotaValue, bool)`

GetMaxPersonalityOk returns a tuple with the MaxPersonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPersonality

`func (o *Quota) SetMaxPersonality(v QuotaValue)`

SetMaxPersonality sets MaxPersonality field to given value.

### HasMaxPersonality

`func (o *Quota) HasMaxPersonality() bool`

HasMaxPersonality returns a boolean if a field has been set.

### GetMaxPersonalitySize

`func (o *Quota) GetMaxPersonalitySize() QuotaValue`

GetMaxPersonalitySize returns the MaxPersonalitySize field if non-nil, zero value otherwise.

### GetMaxPersonalitySizeOk

`func (o *Quota) GetMaxPersonalitySizeOk() (*QuotaValue, bool)`

GetMaxPersonalitySizeOk returns a tuple with the MaxPersonalitySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPersonalitySize

`func (o *Quota) SetMaxPersonalitySize(v QuotaValue)`

SetMaxPersonalitySize sets MaxPersonalitySize field to given value.

### HasMaxPersonalitySize

`func (o *Quota) HasMaxPersonalitySize() bool`

HasMaxPersonalitySize returns a boolean if a field has been set.

### GetMaxSecurityGroupRules

`func (o *Quota) GetMaxSecurityGroupRules() QuotaValue`

GetMaxSecurityGroupRules returns the MaxSecurityGroupRules field if non-nil, zero value otherwise.

### GetMaxSecurityGroupRulesOk

`func (o *Quota) GetMaxSecurityGroupRulesOk() (*QuotaValue, bool)`

GetMaxSecurityGroupRulesOk returns a tuple with the MaxSecurityGroupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSecurityGroupRules

`func (o *Quota) SetMaxSecurityGroupRules(v QuotaValue)`

SetMaxSecurityGroupRules sets MaxSecurityGroupRules field to given value.

### HasMaxSecurityGroupRules

`func (o *Quota) HasMaxSecurityGroupRules() bool`

HasMaxSecurityGroupRules returns a boolean if a field has been set.

### GetMaxSecurityGroups

`func (o *Quota) GetMaxSecurityGroups() QuotaValue`

GetMaxSecurityGroups returns the MaxSecurityGroups field if non-nil, zero value otherwise.

### GetMaxSecurityGroupsOk

`func (o *Quota) GetMaxSecurityGroupsOk() (*QuotaValue, bool)`

GetMaxSecurityGroupsOk returns a tuple with the MaxSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSecurityGroups

`func (o *Quota) SetMaxSecurityGroups(v QuotaValue)`

SetMaxSecurityGroups sets MaxSecurityGroups field to given value.

### HasMaxSecurityGroups

`func (o *Quota) HasMaxSecurityGroups() bool`

HasMaxSecurityGroups returns a boolean if a field has been set.

### GetMaxServerGroupMembers

`func (o *Quota) GetMaxServerGroupMembers() QuotaValue`

GetMaxServerGroupMembers returns the MaxServerGroupMembers field if non-nil, zero value otherwise.

### GetMaxServerGroupMembersOk

`func (o *Quota) GetMaxServerGroupMembersOk() (*QuotaValue, bool)`

GetMaxServerGroupMembersOk returns a tuple with the MaxServerGroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerGroupMembers

`func (o *Quota) SetMaxServerGroupMembers(v QuotaValue)`

SetMaxServerGroupMembers sets MaxServerGroupMembers field to given value.

### HasMaxServerGroupMembers

`func (o *Quota) HasMaxServerGroupMembers() bool`

HasMaxServerGroupMembers returns a boolean if a field has been set.

### GetMaxServerGroups

`func (o *Quota) GetMaxServerGroups() QuotaValue`

GetMaxServerGroups returns the MaxServerGroups field if non-nil, zero value otherwise.

### GetMaxServerGroupsOk

`func (o *Quota) GetMaxServerGroupsOk() (*QuotaValue, bool)`

GetMaxServerGroupsOk returns a tuple with the MaxServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerGroups

`func (o *Quota) SetMaxServerGroups(v QuotaValue)`

SetMaxServerGroups sets MaxServerGroups field to given value.

### HasMaxServerGroups

`func (o *Quota) HasMaxServerGroups() bool`

HasMaxServerGroups returns a boolean if a field has been set.

### GetMaxServerMeta

`func (o *Quota) GetMaxServerMeta() QuotaValue`

GetMaxServerMeta returns the MaxServerMeta field if non-nil, zero value otherwise.

### GetMaxServerMetaOk

`func (o *Quota) GetMaxServerMetaOk() (*QuotaValue, bool)`

GetMaxServerMetaOk returns a tuple with the MaxServerMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerMeta

`func (o *Quota) SetMaxServerMeta(v QuotaValue)`

SetMaxServerMeta sets MaxServerMeta field to given value.

### HasMaxServerMeta

`func (o *Quota) HasMaxServerMeta() bool`

HasMaxServerMeta returns a boolean if a field has been set.

### GetMaxTotalCores

`func (o *Quota) GetMaxTotalCores() QuotaValue`

GetMaxTotalCores returns the MaxTotalCores field if non-nil, zero value otherwise.

### GetMaxTotalCoresOk

`func (o *Quota) GetMaxTotalCoresOk() (*QuotaValue, bool)`

GetMaxTotalCoresOk returns a tuple with the MaxTotalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalCores

`func (o *Quota) SetMaxTotalCores(v QuotaValue)`

SetMaxTotalCores sets MaxTotalCores field to given value.

### HasMaxTotalCores

`func (o *Quota) HasMaxTotalCores() bool`

HasMaxTotalCores returns a boolean if a field has been set.

### GetMaxTotalFloatingIps

`func (o *Quota) GetMaxTotalFloatingIps() QuotaValue`

GetMaxTotalFloatingIps returns the MaxTotalFloatingIps field if non-nil, zero value otherwise.

### GetMaxTotalFloatingIpsOk

`func (o *Quota) GetMaxTotalFloatingIpsOk() (*QuotaValue, bool)`

GetMaxTotalFloatingIpsOk returns a tuple with the MaxTotalFloatingIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalFloatingIps

`func (o *Quota) SetMaxTotalFloatingIps(v QuotaValue)`

SetMaxTotalFloatingIps sets MaxTotalFloatingIps field to given value.

### HasMaxTotalFloatingIps

`func (o *Quota) HasMaxTotalFloatingIps() bool`

HasMaxTotalFloatingIps returns a boolean if a field has been set.

### GetMaxTotalInstances

`func (o *Quota) GetMaxTotalInstances() QuotaValue`

GetMaxTotalInstances returns the MaxTotalInstances field if non-nil, zero value otherwise.

### GetMaxTotalInstancesOk

`func (o *Quota) GetMaxTotalInstancesOk() (*QuotaValue, bool)`

GetMaxTotalInstancesOk returns a tuple with the MaxTotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalInstances

`func (o *Quota) SetMaxTotalInstances(v QuotaValue)`

SetMaxTotalInstances sets MaxTotalInstances field to given value.

### HasMaxTotalInstances

`func (o *Quota) HasMaxTotalInstances() bool`

HasMaxTotalInstances returns a boolean if a field has been set.

### GetMaxTotalKeypairs

`func (o *Quota) GetMaxTotalKeypairs() QuotaValue`

GetMaxTotalKeypairs returns the MaxTotalKeypairs field if non-nil, zero value otherwise.

### GetMaxTotalKeypairsOk

`func (o *Quota) GetMaxTotalKeypairsOk() (*QuotaValue, bool)`

GetMaxTotalKeypairsOk returns a tuple with the MaxTotalKeypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalKeypairs

`func (o *Quota) SetMaxTotalKeypairs(v QuotaValue)`

SetMaxTotalKeypairs sets MaxTotalKeypairs field to given value.

### HasMaxTotalKeypairs

`func (o *Quota) HasMaxTotalKeypairs() bool`

HasMaxTotalKeypairs returns a boolean if a field has been set.

### GetMaxTotalRamSize

`func (o *Quota) GetMaxTotalRamSize() QuotaValue`

GetMaxTotalRamSize returns the MaxTotalRamSize field if non-nil, zero value otherwise.

### GetMaxTotalRamSizeOk

`func (o *Quota) GetMaxTotalRamSizeOk() (*QuotaValue, bool)`

GetMaxTotalRamSizeOk returns a tuple with the MaxTotalRamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalRamSize

`func (o *Quota) SetMaxTotalRamSize(v QuotaValue)`

SetMaxTotalRamSize sets MaxTotalRamSize field to given value.

### HasMaxTotalRamSize

`func (o *Quota) HasMaxTotalRamSize() bool`

HasMaxTotalRamSize returns a boolean if a field has been set.

### GetTotalCoresUsed

`func (o *Quota) GetTotalCoresUsed() int64`

GetTotalCoresUsed returns the TotalCoresUsed field if non-nil, zero value otherwise.

### GetTotalCoresUsedOk

`func (o *Quota) GetTotalCoresUsedOk() (*int64, bool)`

GetTotalCoresUsedOk returns a tuple with the TotalCoresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCoresUsed

`func (o *Quota) SetTotalCoresUsed(v int64)`

SetTotalCoresUsed sets TotalCoresUsed field to given value.

### HasTotalCoresUsed

`func (o *Quota) HasTotalCoresUsed() bool`

HasTotalCoresUsed returns a boolean if a field has been set.

### GetTotalFloatingIpsUsed

`func (o *Quota) GetTotalFloatingIpsUsed() int64`

GetTotalFloatingIpsUsed returns the TotalFloatingIpsUsed field if non-nil, zero value otherwise.

### GetTotalFloatingIpsUsedOk

`func (o *Quota) GetTotalFloatingIpsUsedOk() (*int64, bool)`

GetTotalFloatingIpsUsedOk returns a tuple with the TotalFloatingIpsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFloatingIpsUsed

`func (o *Quota) SetTotalFloatingIpsUsed(v int64)`

SetTotalFloatingIpsUsed sets TotalFloatingIpsUsed field to given value.

### HasTotalFloatingIpsUsed

`func (o *Quota) HasTotalFloatingIpsUsed() bool`

HasTotalFloatingIpsUsed returns a boolean if a field has been set.

### GetTotalInstancesUsed

`func (o *Quota) GetTotalInstancesUsed() int64`

GetTotalInstancesUsed returns the TotalInstancesUsed field if non-nil, zero value otherwise.

### GetTotalInstancesUsedOk

`func (o *Quota) GetTotalInstancesUsedOk() (*int64, bool)`

GetTotalInstancesUsedOk returns a tuple with the TotalInstancesUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstancesUsed

`func (o *Quota) SetTotalInstancesUsed(v int64)`

SetTotalInstancesUsed sets TotalInstancesUsed field to given value.

### HasTotalInstancesUsed

`func (o *Quota) HasTotalInstancesUsed() bool`

HasTotalInstancesUsed returns a boolean if a field has been set.

### GetTotalRamUsed

`func (o *Quota) GetTotalRamUsed() int64`

GetTotalRamUsed returns the TotalRamUsed field if non-nil, zero value otherwise.

### GetTotalRamUsedOk

`func (o *Quota) GetTotalRamUsedOk() (*int64, bool)`

GetTotalRamUsedOk returns a tuple with the TotalRamUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRamUsed

`func (o *Quota) SetTotalRamUsed(v int64)`

SetTotalRamUsed sets TotalRamUsed field to given value.

### HasTotalRamUsed

`func (o *Quota) HasTotalRamUsed() bool`

HasTotalRamUsed returns a boolean if a field has been set.

### GetTotalSecurityGroupsUsed

`func (o *Quota) GetTotalSecurityGroupsUsed() int64`

GetTotalSecurityGroupsUsed returns the TotalSecurityGroupsUsed field if non-nil, zero value otherwise.

### GetTotalSecurityGroupsUsedOk

`func (o *Quota) GetTotalSecurityGroupsUsedOk() (*int64, bool)`

GetTotalSecurityGroupsUsedOk returns a tuple with the TotalSecurityGroupsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSecurityGroupsUsed

`func (o *Quota) SetTotalSecurityGroupsUsed(v int64)`

SetTotalSecurityGroupsUsed sets TotalSecurityGroupsUsed field to given value.

### HasTotalSecurityGroupsUsed

`func (o *Quota) HasTotalSecurityGroupsUsed() bool`

HasTotalSecurityGroupsUsed returns a boolean if a field has been set.

### GetTotalServerGroupsUsed

`func (o *Quota) GetTotalServerGroupsUsed() int64`

GetTotalServerGroupsUsed returns the TotalServerGroupsUsed field if non-nil, zero value otherwise.

### GetTotalServerGroupsUsedOk

`func (o *Quota) GetTotalServerGroupsUsedOk() (*int64, bool)`

GetTotalServerGroupsUsedOk returns a tuple with the TotalServerGroupsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServerGroupsUsed

`func (o *Quota) SetTotalServerGroupsUsed(v int64)`

SetTotalServerGroupsUsed sets TotalServerGroupsUsed field to given value.

### HasTotalServerGroupsUsed

`func (o *Quota) HasTotalServerGroupsUsed() bool`

HasTotalServerGroupsUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


