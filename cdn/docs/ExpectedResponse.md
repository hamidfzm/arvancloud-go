# ExpectedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | Pointer to **[]int32** | List of expected http status code | [optional] 
**Headers** | Pointer to **map[string][]string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 

## Methods

### NewExpectedResponse

`func NewExpectedResponse() *ExpectedResponse`

NewExpectedResponse instantiates a new ExpectedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectedResponseWithDefaults

`func NewExpectedResponseWithDefaults() *ExpectedResponse`

NewExpectedResponseWithDefaults instantiates a new ExpectedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *ExpectedResponse) GetCodes() []int32`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *ExpectedResponse) GetCodesOk() (*[]int32, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *ExpectedResponse) SetCodes(v []int32)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *ExpectedResponse) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetHeaders

`func (o *ExpectedResponse) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ExpectedResponse) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ExpectedResponse) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ExpectedResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *ExpectedResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ExpectedResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ExpectedResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ExpectedResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


