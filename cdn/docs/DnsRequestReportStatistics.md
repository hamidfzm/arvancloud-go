# DnsRequestReportStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Top** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDnsRequestReportStatistics

`func NewDnsRequestReportStatistics() *DnsRequestReportStatistics`

NewDnsRequestReportStatistics instantiates a new DnsRequestReportStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRequestReportStatisticsWithDefaults

`func NewDnsRequestReportStatisticsWithDefaults() *DnsRequestReportStatistics`

NewDnsRequestReportStatisticsWithDefaults instantiates a new DnsRequestReportStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *DnsRequestReportStatistics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DnsRequestReportStatistics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DnsRequestReportStatistics) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DnsRequestReportStatistics) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTop

`func (o *DnsRequestReportStatistics) GetTop() time.Time`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *DnsRequestReportStatistics) GetTopOk() (*time.Time, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *DnsRequestReportStatistics) SetTop(v time.Time)`

SetTop sets Top field to given value.

### HasTop

`func (o *DnsRequestReportStatistics) HasTop() bool`

HasTop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


