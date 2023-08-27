# LogForwarderEventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **bool** |  | [optional] 
**Method** | Pointer to **bool** |  | [optional] 
**Scheme** | Pointer to **bool** |  | [optional] 
**Ip** | Pointer to **bool** |  | [optional] 
**Country** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**ServerIp** | Pointer to **bool** |  | [optional] 
**ServerPort** | Pointer to **bool** |  | [optional] 
**Uri** | Pointer to **bool** |  | [optional] 
**QueryString** | Pointer to **bool** |  | [optional] 
**Firewall** | Pointer to **bool** |  | [optional] 
**Proxy** | Pointer to **bool** |  | [optional] 
**DnsResolver** | Pointer to **bool** |  | [optional] 
**Ddos** | Pointer to **bool** |  | [optional] 
**Ratelimit** | Pointer to **bool** |  | [optional] 
**Waf** | Pointer to **bool** |  | [optional] 

## Methods

### NewLogForwarderEventType

`func NewLogForwarderEventType() *LogForwarderEventType`

NewLogForwarderEventType instantiates a new LogForwarderEventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogForwarderEventTypeWithDefaults

`func NewLogForwarderEventTypeWithDefaults() *LogForwarderEventType`

NewLogForwarderEventTypeWithDefaults instantiates a new LogForwarderEventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *LogForwarderEventType) GetDomain() bool`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LogForwarderEventType) GetDomainOk() (*bool, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LogForwarderEventType) SetDomain(v bool)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LogForwarderEventType) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUuid

`func (o *LogForwarderEventType) GetUuid() bool`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LogForwarderEventType) GetUuidOk() (*bool, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LogForwarderEventType) SetUuid(v bool)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LogForwarderEventType) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTimestamp

`func (o *LogForwarderEventType) GetTimestamp() bool`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LogForwarderEventType) GetTimestampOk() (*bool, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LogForwarderEventType) SetTimestamp(v bool)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LogForwarderEventType) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMethod

`func (o *LogForwarderEventType) GetMethod() bool`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LogForwarderEventType) GetMethodOk() (*bool, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LogForwarderEventType) SetMethod(v bool)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LogForwarderEventType) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetScheme

`func (o *LogForwarderEventType) GetScheme() bool`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *LogForwarderEventType) GetSchemeOk() (*bool, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *LogForwarderEventType) SetScheme(v bool)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *LogForwarderEventType) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetIp

`func (o *LogForwarderEventType) GetIp() bool`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LogForwarderEventType) GetIpOk() (*bool, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LogForwarderEventType) SetIp(v bool)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LogForwarderEventType) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCountry

`func (o *LogForwarderEventType) GetCountry() bool`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LogForwarderEventType) GetCountryOk() (*bool, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LogForwarderEventType) SetCountry(v bool)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LogForwarderEventType) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStatus

`func (o *LogForwarderEventType) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogForwarderEventType) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogForwarderEventType) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogForwarderEventType) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetServerIp

`func (o *LogForwarderEventType) GetServerIp() bool`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *LogForwarderEventType) GetServerIpOk() (*bool, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *LogForwarderEventType) SetServerIp(v bool)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *LogForwarderEventType) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *LogForwarderEventType) GetServerPort() bool`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *LogForwarderEventType) GetServerPortOk() (*bool, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *LogForwarderEventType) SetServerPort(v bool)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *LogForwarderEventType) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetUri

`func (o *LogForwarderEventType) GetUri() bool`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *LogForwarderEventType) GetUriOk() (*bool, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *LogForwarderEventType) SetUri(v bool)`

SetUri sets Uri field to given value.

### HasUri

`func (o *LogForwarderEventType) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetQueryString

`func (o *LogForwarderEventType) GetQueryString() bool`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogForwarderEventType) GetQueryStringOk() (*bool, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogForwarderEventType) SetQueryString(v bool)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *LogForwarderEventType) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetFirewall

`func (o *LogForwarderEventType) GetFirewall() bool`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *LogForwarderEventType) GetFirewallOk() (*bool, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *LogForwarderEventType) SetFirewall(v bool)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *LogForwarderEventType) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetProxy

`func (o *LogForwarderEventType) GetProxy() bool`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *LogForwarderEventType) GetProxyOk() (*bool, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *LogForwarderEventType) SetProxy(v bool)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *LogForwarderEventType) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetDnsResolver

`func (o *LogForwarderEventType) GetDnsResolver() bool`

GetDnsResolver returns the DnsResolver field if non-nil, zero value otherwise.

### GetDnsResolverOk

`func (o *LogForwarderEventType) GetDnsResolverOk() (*bool, bool)`

GetDnsResolverOk returns a tuple with the DnsResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolver

`func (o *LogForwarderEventType) SetDnsResolver(v bool)`

SetDnsResolver sets DnsResolver field to given value.

### HasDnsResolver

`func (o *LogForwarderEventType) HasDnsResolver() bool`

HasDnsResolver returns a boolean if a field has been set.

### GetDdos

`func (o *LogForwarderEventType) GetDdos() bool`

GetDdos returns the Ddos field if non-nil, zero value otherwise.

### GetDdosOk

`func (o *LogForwarderEventType) GetDdosOk() (*bool, bool)`

GetDdosOk returns a tuple with the Ddos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdos

`func (o *LogForwarderEventType) SetDdos(v bool)`

SetDdos sets Ddos field to given value.

### HasDdos

`func (o *LogForwarderEventType) HasDdos() bool`

HasDdos returns a boolean if a field has been set.

### GetRatelimit

`func (o *LogForwarderEventType) GetRatelimit() bool`

GetRatelimit returns the Ratelimit field if non-nil, zero value otherwise.

### GetRatelimitOk

`func (o *LogForwarderEventType) GetRatelimitOk() (*bool, bool)`

GetRatelimitOk returns a tuple with the Ratelimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimit

`func (o *LogForwarderEventType) SetRatelimit(v bool)`

SetRatelimit sets Ratelimit field to given value.

### HasRatelimit

`func (o *LogForwarderEventType) HasRatelimit() bool`

HasRatelimit returns a boolean if a field has been set.

### GetWaf

`func (o *LogForwarderEventType) GetWaf() bool`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *LogForwarderEventType) GetWafOk() (*bool, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *LogForwarderEventType) SetWaf(v bool)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *LogForwarderEventType) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


