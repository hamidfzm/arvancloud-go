# ReplicationTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ReplicationTimeStatus**](ReplicationTimeStatus.md) |  | 
**Time** | [**ReplicationTimeTime**](ReplicationTimeTime.md) |  | 

## Methods

### NewReplicationTime

`func NewReplicationTime(status ReplicationTimeStatus, time ReplicationTimeTime, ) *ReplicationTime`

NewReplicationTime instantiates a new ReplicationTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationTimeWithDefaults

`func NewReplicationTimeWithDefaults() *ReplicationTime`

NewReplicationTimeWithDefaults instantiates a new ReplicationTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ReplicationTime) GetStatus() ReplicationTimeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReplicationTime) GetStatusOk() (*ReplicationTimeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReplicationTime) SetStatus(v ReplicationTimeStatus)`

SetStatus sets Status field to given value.


### GetTime

`func (o *ReplicationTime) GetTime() ReplicationTimeTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ReplicationTime) GetTimeOk() (*ReplicationTimeTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ReplicationTime) SetTime(v ReplicationTimeTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


