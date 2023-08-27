# DnsRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DnsRecordGeneric**](DnsRecordGeneric.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDnsRecordResponse

`func NewDnsRecordResponse() *DnsRecordResponse`

NewDnsRecordResponse instantiates a new DnsRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordResponseWithDefaults

`func NewDnsRecordResponseWithDefaults() *DnsRecordResponse`

NewDnsRecordResponseWithDefaults instantiates a new DnsRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DnsRecordResponse) GetData() DnsRecordGeneric`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsRecordResponse) GetDataOk() (*DnsRecordGeneric, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsRecordResponse) SetData(v DnsRecordGeneric)`

SetData sets Data field to given value.

### HasData

`func (o *DnsRecordResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *DnsRecordResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DnsRecordResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DnsRecordResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DnsRecordResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DnsRecordResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DnsRecordResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


