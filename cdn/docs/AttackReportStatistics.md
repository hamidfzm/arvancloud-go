# AttackReportStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attacks** | Pointer to [**DnsRequestReportStatistics**](DnsRequestReportStatistics.md) |  | [optional] 

## Methods

### NewAttackReportStatistics

`func NewAttackReportStatistics() *AttackReportStatistics`

NewAttackReportStatistics instantiates a new AttackReportStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackReportStatisticsWithDefaults

`func NewAttackReportStatisticsWithDefaults() *AttackReportStatistics`

NewAttackReportStatisticsWithDefaults instantiates a new AttackReportStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttacks

`func (o *AttackReportStatistics) GetAttacks() DnsRequestReportStatistics`

GetAttacks returns the Attacks field if non-nil, zero value otherwise.

### GetAttacksOk

`func (o *AttackReportStatistics) GetAttacksOk() (*DnsRequestReportStatistics, bool)`

GetAttacksOk returns a tuple with the Attacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttacks

`func (o *AttackReportStatistics) SetAttacks(v DnsRequestReportStatistics)`

SetAttacks sets Attacks field to given value.

### HasAttacks

`func (o *AttackReportStatistics) HasAttacks() bool`

HasAttacks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


