# PlanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MonthlyCost** | Pointer to **float32** |  | [optional] 
**Discount** | Pointer to **float32** | between 0 to 100 is the percentage of the discount | [optional] 
**NeededBalance** | Pointer to **float32** | How much balance the account needs for selected plan | [optional] 

## Methods

### NewPlanInfo

`func NewPlanInfo() *PlanInfo`

NewPlanInfo instantiates a new PlanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanInfoWithDefaults

`func NewPlanInfoWithDefaults() *PlanInfo`

NewPlanInfoWithDefaults instantiates a new PlanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlanInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlanInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlanInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PlanInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *PlanInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlanInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMonthlyCost

`func (o *PlanInfo) GetMonthlyCost() float32`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *PlanInfo) GetMonthlyCostOk() (*float32, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *PlanInfo) SetMonthlyCost(v float32)`

SetMonthlyCost sets MonthlyCost field to given value.

### HasMonthlyCost

`func (o *PlanInfo) HasMonthlyCost() bool`

HasMonthlyCost returns a boolean if a field has been set.

### GetDiscount

`func (o *PlanInfo) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *PlanInfo) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *PlanInfo) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *PlanInfo) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetNeededBalance

`func (o *PlanInfo) GetNeededBalance() float32`

GetNeededBalance returns the NeededBalance field if non-nil, zero value otherwise.

### GetNeededBalanceOk

`func (o *PlanInfo) GetNeededBalanceOk() (*float32, bool)`

GetNeededBalanceOk returns a tuple with the NeededBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeededBalance

`func (o *PlanInfo) SetNeededBalance(v float32)`

SetNeededBalance sets NeededBalance field to given value.

### HasNeededBalance

`func (o *PlanInfo) HasNeededBalance() bool`

HasNeededBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


