# DomainsNameserversCheck200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NsDomain**](NsDomain.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDomainsNameserversCheck200Response

`func NewDomainsNameserversCheck200Response() *DomainsNameserversCheck200Response`

NewDomainsNameserversCheck200Response instantiates a new DomainsNameserversCheck200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainsNameserversCheck200ResponseWithDefaults

`func NewDomainsNameserversCheck200ResponseWithDefaults() *DomainsNameserversCheck200Response`

NewDomainsNameserversCheck200ResponseWithDefaults instantiates a new DomainsNameserversCheck200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DomainsNameserversCheck200Response) GetData() NsDomain`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DomainsNameserversCheck200Response) GetDataOk() (*NsDomain, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DomainsNameserversCheck200Response) SetData(v NsDomain)`

SetData sets Data field to given value.

### HasData

`func (o *DomainsNameserversCheck200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *DomainsNameserversCheck200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DomainsNameserversCheck200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DomainsNameserversCheck200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DomainsNameserversCheck200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DomainsNameserversCheck200Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DomainsNameserversCheck200Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


