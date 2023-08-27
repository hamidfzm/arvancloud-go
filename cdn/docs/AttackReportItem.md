# AttackReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttackerIp** | Pointer to **string** |  | [optional] 
**AttackerCountry** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**UriArgs** | Pointer to **string** |  | [optional] 
**Cookie** | Pointer to **[]string** |  | [optional] 
**Alerts** | Pointer to **[]string** |  | [optional] 
**UserAgent** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAttackReportItem

`func NewAttackReportItem() *AttackReportItem`

NewAttackReportItem instantiates a new AttackReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackReportItemWithDefaults

`func NewAttackReportItemWithDefaults() *AttackReportItem`

NewAttackReportItemWithDefaults instantiates a new AttackReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttackerIp

`func (o *AttackReportItem) GetAttackerIp() string`

GetAttackerIp returns the AttackerIp field if non-nil, zero value otherwise.

### GetAttackerIpOk

`func (o *AttackReportItem) GetAttackerIpOk() (*string, bool)`

GetAttackerIpOk returns a tuple with the AttackerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackerIp

`func (o *AttackReportItem) SetAttackerIp(v string)`

SetAttackerIp sets AttackerIp field to given value.

### HasAttackerIp

`func (o *AttackReportItem) HasAttackerIp() bool`

HasAttackerIp returns a boolean if a field has been set.

### GetAttackerCountry

`func (o *AttackReportItem) GetAttackerCountry() string`

GetAttackerCountry returns the AttackerCountry field if non-nil, zero value otherwise.

### GetAttackerCountryOk

`func (o *AttackReportItem) GetAttackerCountryOk() (*string, bool)`

GetAttackerCountryOk returns a tuple with the AttackerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackerCountry

`func (o *AttackReportItem) SetAttackerCountry(v string)`

SetAttackerCountry sets AttackerCountry field to given value.

### HasAttackerCountry

`func (o *AttackReportItem) HasAttackerCountry() bool`

HasAttackerCountry returns a boolean if a field has been set.

### GetMethod

`func (o *AttackReportItem) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AttackReportItem) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AttackReportItem) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AttackReportItem) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *AttackReportItem) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AttackReportItem) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AttackReportItem) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *AttackReportItem) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetHost

`func (o *AttackReportItem) GetHost() []string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AttackReportItem) GetHostOk() (*[]string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AttackReportItem) SetHost(v []string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AttackReportItem) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetTimestamp

`func (o *AttackReportItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AttackReportItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AttackReportItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AttackReportItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUriArgs

`func (o *AttackReportItem) GetUriArgs() string`

GetUriArgs returns the UriArgs field if non-nil, zero value otherwise.

### GetUriArgsOk

`func (o *AttackReportItem) GetUriArgsOk() (*string, bool)`

GetUriArgsOk returns a tuple with the UriArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriArgs

`func (o *AttackReportItem) SetUriArgs(v string)`

SetUriArgs sets UriArgs field to given value.

### HasUriArgs

`func (o *AttackReportItem) HasUriArgs() bool`

HasUriArgs returns a boolean if a field has been set.

### GetCookie

`func (o *AttackReportItem) GetCookie() []string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *AttackReportItem) GetCookieOk() (*[]string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *AttackReportItem) SetCookie(v []string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *AttackReportItem) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### GetAlerts

`func (o *AttackReportItem) GetAlerts() []string`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *AttackReportItem) GetAlertsOk() (*[]string, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *AttackReportItem) SetAlerts(v []string)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *AttackReportItem) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetUserAgent

`func (o *AttackReportItem) GetUserAgent() []string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AttackReportItem) GetUserAgentOk() (*[]string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AttackReportItem) SetUserAgent(v []string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AttackReportItem) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


