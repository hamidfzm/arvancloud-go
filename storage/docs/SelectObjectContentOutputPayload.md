# SelectObjectContentOutputPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**SelectObjectContentEventStreamRecords**](SelectObjectContentEventStreamRecords.md) |  | [optional] 
**Stats** | Pointer to [**SelectObjectContentEventStreamStats**](SelectObjectContentEventStreamStats.md) |  | [optional] 
**Progress** | Pointer to [**SelectObjectContentEventStreamProgress**](SelectObjectContentEventStreamProgress.md) |  | [optional] 
**Cont** | Pointer to **interface{}** |  | [optional] 
**End** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSelectObjectContentOutputPayload

`func NewSelectObjectContentOutputPayload() *SelectObjectContentOutputPayload`

NewSelectObjectContentOutputPayload instantiates a new SelectObjectContentOutputPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectObjectContentOutputPayloadWithDefaults

`func NewSelectObjectContentOutputPayloadWithDefaults() *SelectObjectContentOutputPayload`

NewSelectObjectContentOutputPayloadWithDefaults instantiates a new SelectObjectContentOutputPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *SelectObjectContentOutputPayload) GetRecords() SelectObjectContentEventStreamRecords`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SelectObjectContentOutputPayload) GetRecordsOk() (*SelectObjectContentEventStreamRecords, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SelectObjectContentOutputPayload) SetRecords(v SelectObjectContentEventStreamRecords)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SelectObjectContentOutputPayload) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStats

`func (o *SelectObjectContentOutputPayload) GetStats() SelectObjectContentEventStreamStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SelectObjectContentOutputPayload) GetStatsOk() (*SelectObjectContentEventStreamStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SelectObjectContentOutputPayload) SetStats(v SelectObjectContentEventStreamStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SelectObjectContentOutputPayload) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetProgress

`func (o *SelectObjectContentOutputPayload) GetProgress() SelectObjectContentEventStreamProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SelectObjectContentOutputPayload) GetProgressOk() (*SelectObjectContentEventStreamProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SelectObjectContentOutputPayload) SetProgress(v SelectObjectContentEventStreamProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SelectObjectContentOutputPayload) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetCont

`func (o *SelectObjectContentOutputPayload) GetCont() interface{}`

GetCont returns the Cont field if non-nil, zero value otherwise.

### GetContOk

`func (o *SelectObjectContentOutputPayload) GetContOk() (*interface{}, bool)`

GetContOk returns a tuple with the Cont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCont

`func (o *SelectObjectContentOutputPayload) SetCont(v interface{})`

SetCont sets Cont field to given value.

### HasCont

`func (o *SelectObjectContentOutputPayload) HasCont() bool`

HasCont returns a boolean if a field has been set.

### GetEnd

`func (o *SelectObjectContentOutputPayload) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SelectObjectContentOutputPayload) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SelectObjectContentOutputPayload) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *SelectObjectContentOutputPayload) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


