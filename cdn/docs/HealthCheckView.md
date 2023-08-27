# HealthCheckView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestConfig** | Pointer to **map[string]interface{}** |  | [optional] 
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

### NewHealthCheckView

`func NewHealthCheckView() *HealthCheckView`

NewHealthCheckView instantiates a new HealthCheckView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckViewWithDefaults

`func NewHealthCheckViewWithDefaults() *HealthCheckView`

NewHealthCheckViewWithDefaults instantiates a new HealthCheckView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestConfig

`func (o *HealthCheckView) GetRequestConfig() map[string]interface{}`

GetRequestConfig returns the RequestConfig field if non-nil, zero value otherwise.

### GetRequestConfigOk

`func (o *HealthCheckView) GetRequestConfigOk() (*map[string]interface{}, bool)`

GetRequestConfigOk returns a tuple with the RequestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestConfig

`func (o *HealthCheckView) SetRequestConfig(v map[string]interface{})`

SetRequestConfig sets RequestConfig field to given value.

### HasRequestConfig

`func (o *HealthCheckView) HasRequestConfig() bool`

HasRequestConfig returns a boolean if a field has been set.

### GetId

`func (o *HealthCheckView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthCheckView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthCheckView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HealthCheckView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HealthCheckView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheckView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheckView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthCheckView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *HealthCheckView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HealthCheckView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HealthCheckView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HealthCheckView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrigin

`func (o *HealthCheckView) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *HealthCheckView) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *HealthCheckView) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *HealthCheckView) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOriginType

`func (o *HealthCheckView) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *HealthCheckView) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *HealthCheckView) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *HealthCheckView) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetUpstreams

`func (o *HealthCheckView) GetUpstreams() []string`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *HealthCheckView) GetUpstreamsOk() (*[]string, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *HealthCheckView) SetUpstreams(v []string)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *HealthCheckView) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetInterval

`func (o *HealthCheckView) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HealthCheckView) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HealthCheckView) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HealthCheckView) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetThreshold

`func (o *HealthCheckView) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *HealthCheckView) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *HealthCheckView) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *HealthCheckView) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetType

`func (o *HealthCheckView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HealthCheckView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HealthCheckView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HealthCheckView) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckView) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckView) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckView) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheckView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRetries

`func (o *HealthCheckView) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *HealthCheckView) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *HealthCheckView) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *HealthCheckView) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetZones

`func (o *HealthCheckView) GetZones() []HealthCheckZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *HealthCheckView) GetZonesOk() (*[]HealthCheckZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *HealthCheckView) SetZones(v []HealthCheckZone)`

SetZones sets Zones field to given value.

### HasZones

`func (o *HealthCheckView) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetMonitoringUpdatedAt

`func (o *HealthCheckView) GetMonitoringUpdatedAt() time.Time`

GetMonitoringUpdatedAt returns the MonitoringUpdatedAt field if non-nil, zero value otherwise.

### GetMonitoringUpdatedAtOk

`func (o *HealthCheckView) GetMonitoringUpdatedAtOk() (*time.Time, bool)`

GetMonitoringUpdatedAtOk returns a tuple with the MonitoringUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUpdatedAt

`func (o *HealthCheckView) SetMonitoringUpdatedAt(v time.Time)`

SetMonitoringUpdatedAt sets MonitoringUpdatedAt field to given value.

### HasMonitoringUpdatedAt

`func (o *HealthCheckView) HasMonitoringUpdatedAt() bool`

HasMonitoringUpdatedAt returns a boolean if a field has been set.

### SetMonitoringUpdatedAtNil

`func (o *HealthCheckView) SetMonitoringUpdatedAtNil(b bool)`

 SetMonitoringUpdatedAtNil sets the value for MonitoringUpdatedAt to be an explicit nil

### UnsetMonitoringUpdatedAt
`func (o *HealthCheckView) UnsetMonitoringUpdatedAt()`

UnsetMonitoringUpdatedAt ensures that no value is present for MonitoringUpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


