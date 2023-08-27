# V1DeploymentConfigRollbackSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**IncludeReplicationMeta** | **bool** | IncludeReplicationMeta specifies whether to include the replica count and selector. | 
**IncludeStrategy** | **bool** | IncludeStrategy specifies whether to include the deployment Strategy. | 
**IncludeTemplate** | **bool** | IncludeTemplate specifies whether to include the PodTemplateSpec. | 
**IncludeTriggers** | **bool** | IncludeTriggers specifies whether to include config Triggers. | 
**Revision** | Pointer to **int64** | Revision to rollback to. If set to 0, rollback to the last revision. | [optional] 

## Methods

### NewV1DeploymentConfigRollbackSpec

`func NewV1DeploymentConfigRollbackSpec(from V1ObjectReference, includeReplicationMeta bool, includeStrategy bool, includeTemplate bool, includeTriggers bool, ) *V1DeploymentConfigRollbackSpec`

NewV1DeploymentConfigRollbackSpec instantiates a new V1DeploymentConfigRollbackSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentConfigRollbackSpecWithDefaults

`func NewV1DeploymentConfigRollbackSpecWithDefaults() *V1DeploymentConfigRollbackSpec`

NewV1DeploymentConfigRollbackSpecWithDefaults instantiates a new V1DeploymentConfigRollbackSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1DeploymentConfigRollbackSpec) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1DeploymentConfigRollbackSpec) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1DeploymentConfigRollbackSpec) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetIncludeReplicationMeta

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeReplicationMeta() bool`

GetIncludeReplicationMeta returns the IncludeReplicationMeta field if non-nil, zero value otherwise.

### GetIncludeReplicationMetaOk

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeReplicationMetaOk() (*bool, bool)`

GetIncludeReplicationMetaOk returns a tuple with the IncludeReplicationMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReplicationMeta

`func (o *V1DeploymentConfigRollbackSpec) SetIncludeReplicationMeta(v bool)`

SetIncludeReplicationMeta sets IncludeReplicationMeta field to given value.


### GetIncludeStrategy

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeStrategy() bool`

GetIncludeStrategy returns the IncludeStrategy field if non-nil, zero value otherwise.

### GetIncludeStrategyOk

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeStrategyOk() (*bool, bool)`

GetIncludeStrategyOk returns a tuple with the IncludeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStrategy

`func (o *V1DeploymentConfigRollbackSpec) SetIncludeStrategy(v bool)`

SetIncludeStrategy sets IncludeStrategy field to given value.


### GetIncludeTemplate

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeTemplate() bool`

GetIncludeTemplate returns the IncludeTemplate field if non-nil, zero value otherwise.

### GetIncludeTemplateOk

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeTemplateOk() (*bool, bool)`

GetIncludeTemplateOk returns a tuple with the IncludeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTemplate

`func (o *V1DeploymentConfigRollbackSpec) SetIncludeTemplate(v bool)`

SetIncludeTemplate sets IncludeTemplate field to given value.


### GetIncludeTriggers

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeTriggers() bool`

GetIncludeTriggers returns the IncludeTriggers field if non-nil, zero value otherwise.

### GetIncludeTriggersOk

`func (o *V1DeploymentConfigRollbackSpec) GetIncludeTriggersOk() (*bool, bool)`

GetIncludeTriggersOk returns a tuple with the IncludeTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTriggers

`func (o *V1DeploymentConfigRollbackSpec) SetIncludeTriggers(v bool)`

SetIncludeTriggers sets IncludeTriggers field to given value.


### GetRevision

`func (o *V1DeploymentConfigRollbackSpec) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1DeploymentConfigRollbackSpec) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1DeploymentConfigRollbackSpec) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1DeploymentConfigRollbackSpec) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


