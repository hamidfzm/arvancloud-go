# PlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**FeatureSets**](FeatureSets.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPlanResponse

`func NewPlanResponse() *PlanResponse`

NewPlanResponse instantiates a new PlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanResponseWithDefaults

`func NewPlanResponseWithDefaults() *PlanResponse`

NewPlanResponseWithDefaults instantiates a new PlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PlanResponse) GetData() FeatureSets`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlanResponse) GetDataOk() (*FeatureSets, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlanResponse) SetData(v FeatureSets)`

SetData sets Data field to given value.

### HasData

`func (o *PlanResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *PlanResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlanResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlanResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlanResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PlanResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PlanResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


