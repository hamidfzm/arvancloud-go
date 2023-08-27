# DomainsStore422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] [default to false]
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** | List of parameters and related errors | [optional] 

## Methods

### NewDomainsStore422Response

`func NewDomainsStore422Response() *DomainsStore422Response`

NewDomainsStore422Response instantiates a new DomainsStore422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainsStore422ResponseWithDefaults

`func NewDomainsStore422ResponseWithDefaults() *DomainsStore422Response`

NewDomainsStore422ResponseWithDefaults instantiates a new DomainsStore422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DomainsStore422Response) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainsStore422Response) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainsStore422Response) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainsStore422Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DomainsStore422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DomainsStore422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DomainsStore422Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DomainsStore422Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *DomainsStore422Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DomainsStore422Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DomainsStore422Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DomainsStore422Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *DomainsStore422Response) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *DomainsStore422Response) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


