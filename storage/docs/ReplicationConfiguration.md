# ReplicationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**Role**](Role.md) |  | 
**Rules** | [**ReplicationRules**](ReplicationRules.md) |  | 

## Methods

### NewReplicationConfiguration

`func NewReplicationConfiguration(role Role, rules ReplicationRules, ) *ReplicationConfiguration`

NewReplicationConfiguration instantiates a new ReplicationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationConfigurationWithDefaults

`func NewReplicationConfigurationWithDefaults() *ReplicationConfiguration`

NewReplicationConfigurationWithDefaults instantiates a new ReplicationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ReplicationConfiguration) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ReplicationConfiguration) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ReplicationConfiguration) SetRole(v Role)`

SetRole sets Role field to given value.


### GetRules

`func (o *ReplicationConfiguration) GetRules() ReplicationRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ReplicationConfiguration) GetRulesOk() (*ReplicationRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ReplicationConfiguration) SetRules(v ReplicationRules)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


