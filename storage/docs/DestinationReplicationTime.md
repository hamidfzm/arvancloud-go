# DestinationReplicationTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ReplicationTimeStatus**](ReplicationTimeStatus.md) |  | 
**Time** | [**ReplicationTimeTime**](ReplicationTimeTime.md) |  | 

## Methods

### NewDestinationReplicationTime

`func NewDestinationReplicationTime(status ReplicationTimeStatus, time ReplicationTimeTime, ) *DestinationReplicationTime`

NewDestinationReplicationTime instantiates a new DestinationReplicationTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationReplicationTimeWithDefaults

`func NewDestinationReplicationTimeWithDefaults() *DestinationReplicationTime`

NewDestinationReplicationTimeWithDefaults instantiates a new DestinationReplicationTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DestinationReplicationTime) GetStatus() ReplicationTimeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DestinationReplicationTime) GetStatusOk() (*ReplicationTimeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DestinationReplicationTime) SetStatus(v ReplicationTimeStatus)`

SetStatus sets Status field to given value.


### GetTime

`func (o *DestinationReplicationTime) GetTime() ReplicationTimeTime`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DestinationReplicationTime) GetTimeOk() (*ReplicationTimeTime, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DestinationReplicationTime) SetTime(v ReplicationTimeTime)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


