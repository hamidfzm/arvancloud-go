# DnsRecordGeneric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDnsRecordGeneric

`func NewDnsRecordGeneric() *DnsRecordGeneric`

NewDnsRecordGeneric instantiates a new DnsRecordGeneric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordGenericWithDefaults

`func NewDnsRecordGenericWithDefaults() *DnsRecordGeneric`

NewDnsRecordGenericWithDefaults instantiates a new DnsRecordGeneric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnsRecordGeneric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnsRecordGeneric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnsRecordGeneric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnsRecordGeneric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DnsRecordGeneric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsRecordGeneric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsRecordGeneric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DnsRecordGeneric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *DnsRecordGeneric) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsRecordGeneric) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsRecordGeneric) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsRecordGeneric) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetCloud

`func (o *DnsRecordGeneric) GetCloud() bool`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *DnsRecordGeneric) GetCloudOk() (*bool, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *DnsRecordGeneric) SetCloud(v bool)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *DnsRecordGeneric) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUpstreamHttps

`func (o *DnsRecordGeneric) GetUpstreamHttps() string`

GetUpstreamHttps returns the UpstreamHttps field if non-nil, zero value otherwise.

### GetUpstreamHttpsOk

`func (o *DnsRecordGeneric) GetUpstreamHttpsOk() (*string, bool)`

GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamHttps

`func (o *DnsRecordGeneric) SetUpstreamHttps(v string)`

SetUpstreamHttps sets UpstreamHttps field to given value.

### HasUpstreamHttps

`func (o *DnsRecordGeneric) HasUpstreamHttps() bool`

HasUpstreamHttps returns a boolean if a field has been set.

### GetIpFilterMode

`func (o *DnsRecordGeneric) GetIpFilterMode() DnsRecordIpFilterMode`

GetIpFilterMode returns the IpFilterMode field if non-nil, zero value otherwise.

### GetIpFilterModeOk

`func (o *DnsRecordGeneric) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool)`

GetIpFilterModeOk returns a tuple with the IpFilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilterMode

`func (o *DnsRecordGeneric) SetIpFilterMode(v DnsRecordIpFilterMode)`

SetIpFilterMode sets IpFilterMode field to given value.

### HasIpFilterMode

`func (o *DnsRecordGeneric) HasIpFilterMode() bool`

HasIpFilterMode returns a boolean if a field has been set.

### GetIsProtected

`func (o *DnsRecordGeneric) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *DnsRecordGeneric) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *DnsRecordGeneric) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *DnsRecordGeneric) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUsage

`func (o *DnsRecordGeneric) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DnsRecordGeneric) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DnsRecordGeneric) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DnsRecordGeneric) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DnsRecordGeneric) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DnsRecordGeneric) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DnsRecordGeneric) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DnsRecordGeneric) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DnsRecordGeneric) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DnsRecordGeneric) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DnsRecordGeneric) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DnsRecordGeneric) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


