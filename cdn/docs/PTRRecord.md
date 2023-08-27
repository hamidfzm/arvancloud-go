# PTRRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PTRRecordValue**](PTRRecordValue.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "ptr"]
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

### NewPTRRecord

`func NewPTRRecord() *PTRRecord`

NewPTRRecord instantiates a new PTRRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPTRRecordWithDefaults

`func NewPTRRecordWithDefaults() *PTRRecord`

NewPTRRecordWithDefaults instantiates a new PTRRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PTRRecord) GetValue() PTRRecordValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PTRRecord) GetValueOk() (*PTRRecordValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PTRRecord) SetValue(v PTRRecordValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *PTRRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *PTRRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PTRRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PTRRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PTRRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PTRRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PTRRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PTRRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PTRRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PTRRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PTRRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PTRRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PTRRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *PTRRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PTRRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PTRRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PTRRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetCloud

`func (o *PTRRecord) GetCloud() bool`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *PTRRecord) GetCloudOk() (*bool, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *PTRRecord) SetCloud(v bool)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *PTRRecord) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUpstreamHttps

`func (o *PTRRecord) GetUpstreamHttps() string`

GetUpstreamHttps returns the UpstreamHttps field if non-nil, zero value otherwise.

### GetUpstreamHttpsOk

`func (o *PTRRecord) GetUpstreamHttpsOk() (*string, bool)`

GetUpstreamHttpsOk returns a tuple with the UpstreamHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamHttps

`func (o *PTRRecord) SetUpstreamHttps(v string)`

SetUpstreamHttps sets UpstreamHttps field to given value.

### HasUpstreamHttps

`func (o *PTRRecord) HasUpstreamHttps() bool`

HasUpstreamHttps returns a boolean if a field has been set.

### GetIpFilterMode

`func (o *PTRRecord) GetIpFilterMode() DnsRecordIpFilterMode`

GetIpFilterMode returns the IpFilterMode field if non-nil, zero value otherwise.

### GetIpFilterModeOk

`func (o *PTRRecord) GetIpFilterModeOk() (*DnsRecordIpFilterMode, bool)`

GetIpFilterModeOk returns a tuple with the IpFilterMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFilterMode

`func (o *PTRRecord) SetIpFilterMode(v DnsRecordIpFilterMode)`

SetIpFilterMode sets IpFilterMode field to given value.

### HasIpFilterMode

`func (o *PTRRecord) HasIpFilterMode() bool`

HasIpFilterMode returns a boolean if a field has been set.

### GetIsProtected

`func (o *PTRRecord) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *PTRRecord) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *PTRRecord) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *PTRRecord) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUsage

`func (o *PTRRecord) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PTRRecord) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PTRRecord) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PTRRecord) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PTRRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PTRRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PTRRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PTRRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PTRRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PTRRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PTRRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PTRRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


