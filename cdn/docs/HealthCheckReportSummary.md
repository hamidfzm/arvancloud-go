# HealthCheckReportSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to [**[]HealthCheckReportSummaryDetail**](HealthCheckReportSummaryDetail.md) |  | [optional] 

## Methods

### NewHealthCheckReportSummary

`func NewHealthCheckReportSummary() *HealthCheckReportSummary`

NewHealthCheckReportSummary instantiates a new HealthCheckReportSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckReportSummaryWithDefaults

`func NewHealthCheckReportSummaryWithDefaults() *HealthCheckReportSummary`

NewHealthCheckReportSummaryWithDefaults instantiates a new HealthCheckReportSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *HealthCheckReportSummary) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *HealthCheckReportSummary) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *HealthCheckReportSummary) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *HealthCheckReportSummary) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckReportSummary) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckReportSummary) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckReportSummary) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheckReportSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *HealthCheckReportSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *HealthCheckReportSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *HealthCheckReportSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *HealthCheckReportSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetFailed

`func (o *HealthCheckReportSummary) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *HealthCheckReportSummary) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *HealthCheckReportSummary) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *HealthCheckReportSummary) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetDetails

`func (o *HealthCheckReportSummary) GetDetails() []HealthCheckReportSummaryDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HealthCheckReportSummary) GetDetailsOk() (*[]HealthCheckReportSummaryDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HealthCheckReportSummary) SetDetails(v []HealthCheckReportSummaryDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HealthCheckReportSummary) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


