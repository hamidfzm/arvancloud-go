# ServerIPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **string** |  | [optional] 
**HasPublicIp** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpData** | Pointer to [**[]IPInfo**](IPInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewServerIPInfo

`func NewServerIPInfo() *ServerIPInfo`

NewServerIPInfo instantiates a new ServerIPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerIPInfoWithDefaults

`func NewServerIPInfoWithDefaults() *ServerIPInfo`

NewServerIPInfoWithDefaults instantiates a new ServerIPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *ServerIPInfo) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ServerIPInfo) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ServerIPInfo) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ServerIPInfo) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetHasPublicIp

`func (o *ServerIPInfo) GetHasPublicIp() bool`

GetHasPublicIp returns the HasPublicIp field if non-nil, zero value otherwise.

### GetHasPublicIpOk

`func (o *ServerIPInfo) GetHasPublicIpOk() (*bool, bool)`

GetHasPublicIpOk returns a tuple with the HasPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPublicIp

`func (o *ServerIPInfo) SetHasPublicIp(v bool)`

SetHasPublicIp sets HasPublicIp field to given value.

### HasHasPublicIp

`func (o *ServerIPInfo) HasHasPublicIp() bool`

HasHasPublicIp returns a boolean if a field has been set.

### GetId

`func (o *ServerIPInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerIPInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerIPInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerIPInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpData

`func (o *ServerIPInfo) GetIpData() []IPInfo`

GetIpData returns the IpData field if non-nil, zero value otherwise.

### GetIpDataOk

`func (o *ServerIPInfo) GetIpDataOk() (*[]IPInfo, bool)`

GetIpDataOk returns a tuple with the IpData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpData

`func (o *ServerIPInfo) SetIpData(v []IPInfo)`

SetIpData sets IpData field to given value.

### HasIpData

`func (o *ServerIPInfo) HasIpData() bool`

HasIpData returns a boolean if a field has been set.

### GetName

`func (o *ServerIPInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerIPInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerIPInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerIPInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ServerIPInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerIPInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerIPInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerIPInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


