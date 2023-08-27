# ARecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]ARecordValue**](ARecordValue.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "a"]
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

### NewARecord

`func NewARecord() *ARecord`

NewARecord instantiates a new ARecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewARecordWithDefaults

`func NewARecordWithDefaults() *ARecord`

NewARecordWithDefaults instantiates a new ARecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ARecord) GetValue() []ARecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ARecord) GetValueOk() (*[]ARecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ARecord) SetValue(v []ARecordValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ARecord) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *ARecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ARecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ARecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ARecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ARecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ARecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ARecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ARecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ARecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ARecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ARecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ARecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *ARecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ARecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ARecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ARecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetCloud

`func (o *ARecord) GetCloud() bool`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ARecord) GetCloudOk() (*bool, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ARecord) SetCloud(v bool)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ARecord) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUpstreamHttps

`func (o *ARecord) GetUpstreamHttps() string`

GetUpstreamHttps returns the UpstreamHttps field if non-nil, zero value otherwise.

### GetUpstreamHttpsOk

`func (o *ARecord) GetUpstreamHttpsOk() (*string, bool)`

GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamHttps

`func (o *ARecord) SetUpstreamHttps(v string)`

SetUpstreamHttps sets UpstreamHttps field to given value.

### HasUpstreamHttps

`func (o *ARecord) HasUpstreamHttps() bool`

HasUpstreamHttps returns a boolean if a field has been set.

### GetIpFilterMode

`func (o *ARecord) GetIpFilterMode() DnsRecordIpFilterMode`

GetIpFilterMode returns the IpFilterMode field if non-nil, zero value otherwise.

### GetIpFilterModeOk

`func (o *ARecord) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool)`

GetIpFilterModeOk returns a tuple with the IpFilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilterMode

`func (o *ARecord) SetIpFilterMode(v DnsRecordIpFilterMode)`

SetIpFilterMode sets IpFilterMode field to given value.

### HasIpFilterMode

`func (o *ARecord) HasIpFilterMode() bool`

HasIpFilterMode returns a boolean if a field has been set.

### GetIsProtected

`func (o *ARecord) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *ARecord) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *ARecord) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *ARecord) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUsage

`func (o *ARecord) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ARecord) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ARecord) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ARecord) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ARecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ARecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ARecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ARecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ARecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ARecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ARecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ARecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


