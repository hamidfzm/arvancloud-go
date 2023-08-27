# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestConfig** | Pointer to [**HealthCheckRequestConfig**](HealthCheckRequestConfig.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** | can be IP/Host when type is &#x60;upstream&#x60;, otherwise it must be a valid record ID | [optional] 
**OriginType** | Pointer to **string** |  | [optional] 
**Upstreams** | Pointer to **[]string** |  | [optional] 
**Interval** | Pointer to **int32** | In milliseconds | [optional] 
**Threshold** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** | The health-check is off or on | [optional] [default to true]
**Retries** | Pointer to **int32** | Number of immediate retries in case of a timeout | [optional] 
**Zones** | Pointer to [**[]HealthCheckZone**](HealthCheckZone.md) |  | [optional] 
**MonitoringUpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewHealthCheck

`func NewHealthCheck() *HealthCheck`

NewHealthCheck instantiates a new HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckWithDefaults

`func NewHealthCheckWithDefaults() *HealthCheck`

NewHealthCheckWithDefaults instantiates a new HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestConfig

`func (o *HealthCheck) GetRequestConfig() HealthCheckRequestConfig`

GetRequestConfig returns the RequestConfig field if non-nil, zero value otherwise.

### GetRequestConfigOk

`func (o *HealthCheck) GetRequestConfigOk() (*HealthCheckRequestConfig, bool)`

GetRequestConfigOk returns a tuple with the RequestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfig

`func (o *HealthCheck) SetRequestConfig(v HealthCheckRequestConfig)`

SetRequestConfig sets RequestConfig field to given value.

### HasRequestConfig

`func (o *HealthCheck) HasRequestConfig() bool`

HasRequestConfig returns a boolean if a field has been set.

### GetId

`func (o *HealthCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HealthCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HealthCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *HealthCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrigin

`func (o *HealthCheck) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *HealthCheck) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *HealthCheck) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *HealthCheck) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOriginType

`func (o *HealthCheck) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *HealthCheck) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *HealthCheck) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *HealthCheck) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetUpstreams

`func (o *HealthCheck) GetUpstreams() []string`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *HealthCheck) GetUpstreamsOk() (*[]string, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *HealthCheck) SetUpstreams(v []string)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *HealthCheck) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetInterval

`func (o *HealthCheck) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HealthCheck) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HealthCheck) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HealthCheck) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetThreshold

`func (o *HealthCheck) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *HealthCheck) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *HealthCheck) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *HealthCheck) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetType

`func (o *HealthCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HealthCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheck) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheck) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheck) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRetries

`func (o *HealthCheck) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *HealthCheck) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *HealthCheck) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *HealthCheck) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetZones

`func (o *HealthCheck) GetZones() []HealthCheckZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *HealthCheck) GetZonesOk() (*[]HealthCheckZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *HealthCheck) SetZones(v []HealthCheckZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *HealthCheck) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetMonitoringUpdatedAt

`func (o *HealthCheck) GetMonitoringUpdatedAt() time.Time`

GetMonitoringUpdatedAt returns the MonitoringUpdatedAt field if non-nil, zero value otherwise.

### GetMonitoringUpdatedAtOk

`func (o *HealthCheck) GetMonitoringUpdatedAtOk() (*time.Time, bool)`

GetMonitoringUpdatedAtOk returns a tuple with the MonitoringUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUpdatedAt

`func (o *HealthCheck) SetMonitoringUpdatedAt(v time.Time)`

SetMonitoringUpdatedAt sets MonitoringUpdatedAt field to given value.

### HasMonitoringUpdatedAt

`func (o *HealthCheck) HasMonitoringUpdatedAt() bool`

HasMonitoringUpdatedAt returns a boolean if a field has been set.

### SetMonitoringUpdatedAtNil

`func (o *HealthCheck) SetMonitoringUpdatedAtNil(b bool)`

 SetMonitoringUpdatedAtNil sets the value for MonitoringUpdatedAt to be an explicit nil

### UnsetMonitoringUpdatedAt
`func (o *HealthCheck) UnsetMonitoringUpdatedAt()`

UnsetMonitoringUpdatedAt ensures that no value is present for MonitoringUpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


