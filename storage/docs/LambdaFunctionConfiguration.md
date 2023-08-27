# LambdaFunctionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An optional unique identifier for configurations in a notification configuration. If you don&#39;t provide one, ArvanCloud S3 will assign an ID. | [optional] 
**LambdaFunctionArn** | **string** |  | 
**Events** | [**Array**](array.md) |  | 
**Filter** | Pointer to [**NotificationConfigurationFilter**](NotificationConfigurationFilter.md) |  | [optional] 

## Methods

### NewLambdaFunctionConfiguration

`func NewLambdaFunctionConfiguration(lambdaFunctionArn string, events Array, ) *LambdaFunctionConfiguration`

NewLambdaFunctionConfiguration instantiates a new LambdaFunctionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLambdaFunctionConfigurationWithDefaults

`func NewLambdaFunctionConfigurationWithDefaults() *LambdaFunctionConfiguration`

NewLambdaFunctionConfigurationWithDefaults instantiates a new LambdaFunctionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LambdaFunctionConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LambdaFunctionConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LambdaFunctionConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LambdaFunctionConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLambdaFunctionArn

`func (o *LambdaFunctionConfiguration) GetLambdaFunctionArn() string`

GetLambdaFunctionArn returns the LambdaFunctionArn field if non-nil, zero value otherwise.

### GetLambdaFunctionArnOk

`func (o *LambdaFunctionConfiguration) GetLambdaFunctionArnOk() (*string, bool)`

GetLambdaFunctionArnOk returns a tuple with the LambdaFunctionArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionArn

`func (o *LambdaFunctionConfiguration) SetLambdaFunctionArn(v string)`

SetLambdaFunctionArn sets LambdaFunctionArn field to given value.


### GetEvents

`func (o *LambdaFunctionConfiguration) GetEvents() Array`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *LambdaFunctionConfiguration) GetEventsOk() (*Array, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *LambdaFunctionConfiguration) SetEvents(v Array)`

SetEvents sets Events field to given value.


### GetFilter

`func (o *LambdaFunctionConfiguration) GetFilter() NotificationConfigurationFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LambdaFunctionConfiguration) GetFilterOk() (*NotificationConfigurationFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LambdaFunctionConfiguration) SetFilter(v NotificationConfigurationFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LambdaFunctionConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


