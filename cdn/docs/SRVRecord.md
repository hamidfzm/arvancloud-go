# SRVRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**SRVRecordValue**](SRVRecordValue.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "srv"]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Cloud** | Pointer to **bool** |  | [optional] [default to false]
**UpstreamHttps** | Pointer to **string** |  | [optional] 
**IpFilterMode** | Pointer to [**DnsRecordIpFilterMode**](DnsRecordIpFilterMode.md) |  | [optional] 
**IsProtected** | Pointer to **bool** | Protected records cannot be modified or deleted by user. | [optional] [readonly] [default to false]
**Usage** | Pointer to **[]string** |  | [optional] [readonly] [default to []]
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSRVRecord

`func NewSRVRecord() *SRVRecord`

NewSRVRecord instantiates a new SRVRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRVRecordWithDefaults

`func NewSRVRecordWithDefaults() *SRVRecord`

NewSRVRecordWithDefaults instantiates a new SRVRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SRVRecord) GetValue() SRVRecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SRVRecord) GetValueOk() (*SRVRecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SRVRecord) SetValue(v SRVRecordValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *SRVRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *SRVRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SRVRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SRVRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SRVRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SRVRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SRVRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SRVRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SRVRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SRVRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SRVRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SRVRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SRVRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *SRVRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SRVRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SRVRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SRVRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetCloud

`func (o *SRVRecord) GetCloud() bool`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *SRVRecord) GetCloudOk() (*bool, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *SRVRecord) SetCloud(v bool)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *SRVRecord) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUpstreamHttps

`func (o *SRVRecord) GetUpstreamHttps() string`

GetUpstreamHttps returns the UpstreamHttps field if non-nil, zero value otherwise.

### GetUpstreamHttpsOk

`func (o *SRVRecord) GetUpstreamHttpsOk() (*string, bool)`

GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamHttps

`func (o *SRVRecord) SetUpstreamHttps(v string)`

SetUpstreamHttps sets UpstreamHttps field to given value.

### HasUpstreamHttps

`func (o *SRVRecord) HasUpstreamHttps() bool`

HasUpstreamHttps returns a boolean if a field has been set.

### GetIpFilterMode

`func (o *SRVRecord) GetIpFilterMode() DnsRecordIpFilterMode`

GetIpFilterMode returns the IpFilterMode field if non-nil, zero value otherwise.

### GetIpFilterModeOk

`func (o *SRVRecord) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool)`

GetIpFilterModeOk returns a tuple with the IpFilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilterMode

`func (o *SRVRecord) SetIpFilterMode(v DnsRecordIpFilterMode)`

SetIpFilterMode sets IpFilterMode field to given value.

### HasIpFilterMode

`func (o *SRVRecord) HasIpFilterMode() bool`

HasIpFilterMode returns a boolean if a field has been set.

### GetIsProtected

`func (o *SRVRecord) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *SRVRecord) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *SRVRecord) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *SRVRecord) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUsage

`func (o *SRVRecord) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SRVRecord) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SRVRecord) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SRVRecord) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SRVRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SRVRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SRVRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SRVRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SRVRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SRVRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SRVRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SRVRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


