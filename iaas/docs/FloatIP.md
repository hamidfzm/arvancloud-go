# FloatIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FixedIpAddress** | Pointer to **string** |  | [optional] 
**FloatingIpAddress** | Pointer to **string** |  | [optional] 
**FloatingNetworkId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PortId** | Pointer to **string** |  | [optional] 
**RevisionNumber** | Pointer to **string** |  | [optional] 
**RouterId** | Pointer to **string** |  | [optional] 
**Server** | Pointer to [**ServerDetail**](ServerDetail.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewFloatIP

`func NewFloatIP() *FloatIP`

NewFloatIP instantiates a new FloatIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloatIPWithDefaults

`func NewFloatIPWithDefaults() *FloatIP`

NewFloatIPWithDefaults instantiates a new FloatIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FloatIP) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FloatIP) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FloatIP) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FloatIP) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FloatIP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FloatIP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FloatIP) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FloatIP) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedIpAddress

`func (o *FloatIP) GetFixedIpAddress() string`

GetFixedIpAddress returns the FixedIpAddress field if non-nil, zero value otherwise.

### GetFixedIpAddressOk

`func (o *FloatIP) GetFixedIpAddressOk() (*string, bool)`

GetFixedIpAddressOk returns a tuple with the FixedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddress

`func (o *FloatIP) SetFixedIpAddress(v string)`

SetFixedIpAddress sets FixedIpAddress field to given value.

### HasFixedIpAddress

`func (o *FloatIP) HasFixedIpAddress() bool`

HasFixedIpAddress returns a boolean if a field has been set.

### GetFloatingIpAddress

`func (o *FloatIP) GetFloatingIpAddress() string`

GetFloatingIpAddress returns the FloatingIpAddress field if non-nil, zero value otherwise.

### GetFloatingIpAddressOk

`func (o *FloatIP) GetFloatingIpAddressOk() (*string, bool)`

GetFloatingIpAddressOk returns a tuple with the FloatingIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpAddress

`func (o *FloatIP) SetFloatingIpAddress(v string)`

SetFloatingIpAddress sets FloatingIpAddress field to given value.

### HasFloatingIpAddress

`func (o *FloatIP) HasFloatingIpAddress() bool`

HasFloatingIpAddress returns a boolean if a field has been set.

### GetFloatingNetworkId

`func (o *FloatIP) GetFloatingNetworkId() string`

GetFloatingNetworkId returns the FloatingNetworkId field if non-nil, zero value otherwise.

### GetFloatingNetworkIdOk

`func (o *FloatIP) GetFloatingNetworkIdOk() (*string, bool)`

GetFloatingNetworkIdOk returns a tuple with the FloatingNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingNetworkId

`func (o *FloatIP) SetFloatingNetworkId(v string)`

SetFloatingNetworkId sets FloatingNetworkId field to given value.

### HasFloatingNetworkId

`func (o *FloatIP) HasFloatingNetworkId() bool`

HasFloatingNetworkId returns a boolean if a field has been set.

### GetId

`func (o *FloatIP) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloatIP) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloatIP) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FloatIP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortId

`func (o *FloatIP) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FloatIP) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FloatIP) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FloatIP) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *FloatIP) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *FloatIP) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *FloatIP) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *FloatIP) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetRouterId

`func (o *FloatIP) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *FloatIP) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *FloatIP) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *FloatIP) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetServer

`func (o *FloatIP) GetServer() ServerDetail`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FloatIP) GetServerOk() (*ServerDetail, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FloatIP) SetServer(v ServerDetail)`

SetServer sets Server field to given value.

### HasServer

`func (o *FloatIP) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetStatus

`func (o *FloatIP) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FloatIP) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FloatIP) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FloatIP) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *FloatIP) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FloatIP) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FloatIP) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FloatIP) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FloatIP) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FloatIP) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FloatIP) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FloatIP) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


