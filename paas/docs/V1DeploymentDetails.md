# V1DeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Causes** | [**[]V1DeploymentCause**](V1DeploymentCause.md) | Causes are extended data associated with all the causes for creating a new deployment | 
**Message** | Pointer to **string** | Message is the user specified change message, if this deployment was triggered manually by the user | [optional] 

## Methods

### NewV1DeploymentDetails

`func NewV1DeploymentDetails(causes []V1DeploymentCause, ) *V1DeploymentDetails`

NewV1DeploymentDetails instantiates a new V1DeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentDetailsWithDefaults

`func NewV1DeploymentDetailsWithDefaults() *V1DeploymentDetails`

NewV1DeploymentDetailsWithDefaults instantiates a new V1DeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCauses

`func (o *V1DeploymentDetails) GetCauses() []V1DeploymentCause`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *V1DeploymentDetails) GetCausesOk() (*[]V1DeploymentCause, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *V1DeploymentDetails) SetCauses(v []V1DeploymentCause)`

SetCauses sets Causes field to given value.


### GetMessage

`func (o *V1DeploymentDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1DeploymentDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1DeploymentDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1DeploymentDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


