# Metrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**MetricsStatus**](MetricsStatus.md) |  | 
**EventThreshold** | Pointer to [**MetricsEventThreshold**](MetricsEventThreshold.md) |  | [optional] 

## Methods

### NewMetrics

`func NewMetrics(status MetricsStatus, ) *Metrics`

NewMetrics instantiates a new Metrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsWithDefaults

`func NewMetricsWithDefaults() *Metrics`

NewMetricsWithDefaults instantiates a new Metrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Metrics) GetStatus() MetricsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Metrics) GetStatusOk() (*MetricsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Metrics) SetStatus(v MetricsStatus)`

SetStatus sets Status field to given value.


### GetEventThreshold

`func (o *Metrics) GetEventThreshold() MetricsEventThreshold`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *Metrics) GetEventThresholdOk() (*MetricsEventThreshold, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *Metrics) SetEventThreshold(v MetricsEventThreshold)`

SetEventThreshold sets EventThreshold field to given value.

### HasEventThreshold

`func (o *Metrics) HasEventThreshold() bool`

HasEventThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


