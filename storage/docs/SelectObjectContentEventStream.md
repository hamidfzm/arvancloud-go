# SelectObjectContentEventStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**SelectObjectContentEventStreamRecords**](SelectObjectContentEventStreamRecords.md) |  | [optional] 
**Stats** | Pointer to [**SelectObjectContentEventStreamStats**](SelectObjectContentEventStreamStats.md) |  | [optional] 
**Progress** | Pointer to [**SelectObjectContentEventStreamProgress**](SelectObjectContentEventStreamProgress.md) |  | [optional] 
**Cont** | Pointer to **interface{}** |  | [optional] 
**End** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSelectObjectContentEventStream

`func NewSelectObjectContentEventStream() *SelectObjectContentEventStream`

NewSelectObjectContentEventStream instantiates a new SelectObjectContentEventStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectObjectContentEventStreamWithDefaults

`func NewSelectObjectContentEventStreamWithDefaults() *SelectObjectContentEventStream`

NewSelectObjectContentEventStreamWithDefaults instantiates a new SelectObjectContentEventStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *SelectObjectContentEventStream) GetRecords() SelectObjectContentEventStreamRecords`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SelectObjectContentEventStream) GetRecordsOk() (*SelectObjectContentEventStreamRecords, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SelectObjectContentEventStream) SetRecords(v SelectObjectContentEventStreamRecords)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SelectObjectContentEventStream) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStats

`func (o *SelectObjectContentEventStream) GetStats() SelectObjectContentEventStreamStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SelectObjectContentEventStream) GetStatsOk() (*SelectObjectContentEventStreamStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SelectObjectContentEventStream) SetStats(v SelectObjectContentEventStreamStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SelectObjectContentEventStream) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetProgress

`func (o *SelectObjectContentEventStream) GetProgress() SelectObjectContentEventStreamProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SelectObjectContentEventStream) GetProgressOk() (*SelectObjectContentEventStreamProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SelectObjectContentEventStream) SetProgress(v SelectObjectContentEventStreamProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SelectObjectContentEventStream) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCont

`func (o *SelectObjectContentEventStream) GetCont() interface{}`

GetCont returns the Cont field if non-nil, zero value otherwise.

### GetContOk

`func (o *SelectObjectContentEventStream) GetContOk() (*interface{}, bool)`

GetContOk returns a tuple with the Cont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCont

`func (o *SelectObjectContentEventStream) SetCont(v interface{})`

SetCont sets Cont field to given value.

### HasCont

`func (o *SelectObjectContentEventStream) HasCont() bool`

HasCont returns a boolean if a field has been set.

### GetEnd

`func (o *SelectObjectContentEventStream) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SelectObjectContentEventStream) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SelectObjectContentEventStream) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *SelectObjectContentEventStream) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


