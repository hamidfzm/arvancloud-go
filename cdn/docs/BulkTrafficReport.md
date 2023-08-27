# BulkTrafficReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** | Domain&#39;s ID | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** | The error message | [optional] 
**Data** | Pointer to [**BulkTrafficReportData**](BulkTrafficReportData.md) |  | [optional] 

## Methods

### NewBulkTrafficReport

`func NewBulkTrafficReport() *BulkTrafficReport`

NewBulkTrafficReport instantiates a new BulkTrafficReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTrafficReportWithDefaults

`func NewBulkTrafficReportWithDefaults() *BulkTrafficReport`

NewBulkTrafficReportWithDefaults instantiates a new BulkTrafficReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *BulkTrafficReport) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BulkTrafficReport) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BulkTrafficReport) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *BulkTrafficReport) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSuccess

`func (o *BulkTrafficReport) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkTrafficReport) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkTrafficReport) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BulkTrafficReport) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *BulkTrafficReport) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulkTrafficReport) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulkTrafficReport) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulkTrafficReport) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BulkTrafficReport) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BulkTrafficReport) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetData

`func (o *BulkTrafficReport) GetData() BulkTrafficReportData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkTrafficReport) GetDataOk() (*BulkTrafficReportData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkTrafficReport) SetData(v BulkTrafficReportData)`

SetData sets Data field to given value.

### HasData

`func (o *BulkTrafficReport) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


