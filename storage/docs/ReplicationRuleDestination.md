# ReplicationRuleDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Account** | Pointer to **string** |  | [optional] 
**StorageClass** | Pointer to [**StorageClass**](StorageClass.md) |  | [optional] 
**AccessControlTranslation** | Pointer to [**DestinationAccessControlTranslation**](DestinationAccessControlTranslation.md) |  | [optional] 
**EncryptionConfiguration** | Pointer to [**DestinationEncryptionConfiguration**](DestinationEncryptionConfiguration.md) |  | [optional] 
**ReplicationTime** | Pointer to [**DestinationReplicationTime**](DestinationReplicationTime.md) |  | [optional] 
**Metrics** | Pointer to [**DestinationMetrics**](DestinationMetrics.md) |  | [optional] 

## Methods

### NewReplicationRuleDestination

`func NewReplicationRuleDestination(bucket string, ) *ReplicationRuleDestination`

NewReplicationRuleDestination instantiates a new ReplicationRuleDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRuleDestinationWithDefaults

`func NewReplicationRuleDestinationWithDefaults() *ReplicationRuleDestination`

NewReplicationRuleDestinationWithDefaults instantiates a new ReplicationRuleDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *ReplicationRuleDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ReplicationRuleDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ReplicationRuleDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetAccount

`func (o *ReplicationRuleDestination) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ReplicationRuleDestination) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ReplicationRuleDestination) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ReplicationRuleDestination) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetStorageClass

`func (o *ReplicationRuleDestination) GetStorageClass() StorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ReplicationRuleDestination) GetStorageClassOk() (*StorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ReplicationRuleDestination) SetStorageClass(v StorageClass)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *ReplicationRuleDestination) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetAccessControlTranslation

`func (o *ReplicationRuleDestination) GetAccessControlTranslation() DestinationAccessControlTranslation`

GetAccessControlTranslation returns the AccessControlTranslation field if non-nil, zero value otherwise.

### GetAccessControlTranslationOk

`func (o *ReplicationRuleDestination) GetAccessControlTranslationOk() (*DestinationAccessControlTranslation, bool)`

GetAccessControlTranslationOk returns a tuple with the AccessControlTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlTranslation

`func (o *ReplicationRuleDestination) SetAccessControlTranslation(v DestinationAccessControlTranslation)`

SetAccessControlTranslation sets AccessControlTranslation field to given value.

### HasAccessControlTranslation

`func (o *ReplicationRuleDestination) HasAccessControlTranslation() bool`

HasAccessControlTranslation returns a boolean if a field has been set.

### GetEncryptionConfiguration

`func (o *ReplicationRuleDestination) GetEncryptionConfiguration() DestinationEncryptionConfiguration`

GetEncryptionConfiguration returns the EncryptionConfiguration field if non-nil, zero value otherwise.

### GetEncryptionConfigurationOk

`func (o *ReplicationRuleDestination) GetEncryptionConfigurationOk() (*DestinationEncryptionConfiguration, bool)`

GetEncryptionConfigurationOk returns a tuple with the EncryptionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfiguration

`func (o *ReplicationRuleDestination) SetEncryptionConfiguration(v DestinationEncryptionConfiguration)`

SetEncryptionConfiguration sets EncryptionConfiguration field to given value.

### HasEncryptionConfiguration

`func (o *ReplicationRuleDestination) HasEncryptionConfiguration() bool`

HasEncryptionConfiguration returns a boolean if a field has been set.

### GetReplicationTime

`func (o *ReplicationRuleDestination) GetReplicationTime() DestinationReplicationTime`

GetReplicationTime returns the ReplicationTime field if non-nil, zero value otherwise.

### GetReplicationTimeOk

`func (o *ReplicationRuleDestination) GetReplicationTimeOk() (*DestinationReplicationTime, bool)`

GetReplicationTimeOk returns a tuple with the ReplicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationTime

`func (o *ReplicationRuleDestination) SetReplicationTime(v DestinationReplicationTime)`

SetReplicationTime sets ReplicationTime field to given value.

### HasReplicationTime

`func (o *ReplicationRuleDestination) HasReplicationTime() bool`

HasReplicationTime returns a boolean if a field has been set.

### GetMetrics

`func (o *ReplicationRuleDestination) GetMetrics() DestinationMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ReplicationRuleDestination) GetMetricsOk() (*DestinationMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ReplicationRuleDestination) SetMetrics(v DestinationMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ReplicationRuleDestination) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


