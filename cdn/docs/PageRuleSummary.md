# PageRuleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**DomainId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** | Order of the page-rule | [optional] 
**UrlType** | Pointer to **string** | This flag is deprecated in favor of is_protected flag | [optional] [default to "default"]
**IsProtected** | Pointer to **bool** | Protected records cannot be modified or deleted by user. | [optional] [readonly] [default to false]
**Url** | Pointer to **string** | URL pattern of target pages | [optional] 
**CacheLevel** | Pointer to **string** |  | [optional] [default to "query_string"]
**WafStatus** | Pointer to **bool** |  | [optional] [default to true]
**FwStatus** | Pointer to **bool** | Shows whether firewall is enabled or not | [optional] [default to true]
**Acceleration** | Pointer to [**Acceleration**](Acceleration.md) |  | [optional] 
**SlinkStatus** | Pointer to **bool** | Secure link is enabled or not | [optional] [default to false]
**Status** | Pointer to **bool** | Is the page-rule enabled? | [optional] [default to true]
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPageRuleSummary

`func NewPageRuleSummary() *PageRuleSummary`

NewPageRuleSummary instantiates a new PageRuleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRuleSummaryWithDefaults

`func NewPageRuleSummaryWithDefaults() *PageRuleSummary`

NewPageRuleSummaryWithDefaults instantiates a new PageRuleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageRuleSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageRuleSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageRuleSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageRuleSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomainId

`func (o *PageRuleSummary) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PageRuleSummary) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PageRuleSummary) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *PageRuleSummary) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetSeq

`func (o *PageRuleSummary) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PageRuleSummary) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PageRuleSummary) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PageRuleSummary) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetUrlType

`func (o *PageRuleSummary) GetUrlType() string`

GetUrlType returns the UrlType field if non-nil, zero value otherwise.

### GetUrlTypeOk

`func (o *PageRuleSummary) GetUrlTypeOk() (*string, bool)`

GetUrlTypeOk returns a tuple with the UrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlType

`func (o *PageRuleSummary) SetUrlType(v string)`

SetUrlType sets UrlType field to given value.

### HasUrlType

`func (o *PageRuleSummary) HasUrlType() bool`

HasUrlType returns a boolean if a field has been set.

### GetIsProtected

`func (o *PageRuleSummary) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *PageRuleSummary) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *PageRuleSummary) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *PageRuleSummary) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUrl

`func (o *PageRuleSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PageRuleSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PageRuleSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PageRuleSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCacheLevel

`func (o *PageRuleSummary) GetCacheLevel() string`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *PageRuleSummary) GetCacheLevelOk() (*string, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *PageRuleSummary) SetCacheLevel(v string)`

SetCacheLevel sets CacheLevel field to given value.

### HasCacheLevel

`func (o *PageRuleSummary) HasCacheLevel() bool`

HasCacheLevel returns a boolean if a field has been set.

### GetWafStatus

`func (o *PageRuleSummary) GetWafStatus() bool`

GetWafStatus returns the WafStatus field if non-nil, zero value otherwise.

### GetWafStatusOk

`func (o *PageRuleSummary) GetWafStatusOk() (*bool, bool)`

GetWafStatusOk returns a tuple with the WafStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus

`func (o *PageRuleSummary) SetWafStatus(v bool)`

SetWafStatus sets WafStatus field to given value.

### HasWafStatus

`func (o *PageRuleSummary) HasWafStatus() bool`

HasWafStatus returns a boolean if a field has been set.

### GetFwStatus

`func (o *PageRuleSummary) GetFwStatus() bool`

GetFwStatus returns the FwStatus field if non-nil, zero value otherwise.

### GetFwStatusOk

`func (o *PageRuleSummary) GetFwStatusOk() (*bool, bool)`

GetFwStatusOk returns a tuple with the FwStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwStatus

`func (o *PageRuleSummary) SetFwStatus(v bool)`

SetFwStatus sets FwStatus field to given value.

### HasFwStatus

`func (o *PageRuleSummary) HasFwStatus() bool`

HasFwStatus returns a boolean if a field has been set.

### GetAcceleration

`func (o *PageRuleSummary) GetAcceleration() Acceleration`

GetAcceleration returns the Acceleration field if non-nil, zero value otherwise.

### GetAccelerationOk

`func (o *PageRuleSummary) GetAccelerationOk() (*Acceleration, bool)`

GetAccelerationOk returns a tuple with the Acceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleration

`func (o *PageRuleSummary) SetAcceleration(v Acceleration)`

SetAcceleration sets Acceleration field to given value.

### HasAcceleration

`func (o *PageRuleSummary) HasAcceleration() bool`

HasAcceleration returns a boolean if a field has been set.

### GetSlinkStatus

`func (o *PageRuleSummary) GetSlinkStatus() bool`

GetSlinkStatus returns the SlinkStatus field if non-nil, zero value otherwise.

### GetSlinkStatusOk

`func (o *PageRuleSummary) GetSlinkStatusOk() (*bool, bool)`

GetSlinkStatusOk returns a tuple with the SlinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkStatus

`func (o *PageRuleSummary) SetSlinkStatus(v bool)`

SetSlinkStatus sets SlinkStatus field to given value.

### HasSlinkStatus

`func (o *PageRuleSummary) HasSlinkStatus() bool`

HasSlinkStatus returns a boolean if a field has been set.

### GetStatus

`func (o *PageRuleSummary) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageRuleSummary) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageRuleSummary) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageRuleSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PageRuleSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageRuleSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageRuleSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageRuleSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PageRuleSummary) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageRuleSummary) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageRuleSummary) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PageRuleSummary) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


