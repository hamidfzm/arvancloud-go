# BulkVisitorReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** | Domain&#39;s ID | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **NullableString** | The error message | [optional] 
**Data** | Pointer to [**BulkVisitorReportData**](BulkVisitorReportData.md) |  | [optional] 

## Methods

### NewBulkVisitorReport

`func NewBulkVisitorReport() *BulkVisitorReport`

NewBulkVisitorReport instantiates a new BulkVisitorReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkVisitorReportWithDefaults

`func NewBulkVisitorReportWithDefaults() *BulkVisitorReport`

NewBulkVisitorReportWithDefaults instantiates a new BulkVisitorReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *BulkVisitorReport) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BulkVisitorReport) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BulkVisitorReport) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *BulkVisitorReport) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSuccess

`func (o *BulkVisitorReport) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkVisitorReport) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkVisitorReport) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BulkVisitorReport) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *BulkVisitorReport) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulkVisitorReport) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulkVisitorReport) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulkVisitorReport) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *BulkVisitorReport) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *BulkVisitorReport) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetData

`func (o *BulkVisitorReport) GetData() BulkVisitorReportData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkVisitorReport) GetDataOk() (*BulkVisitorReportData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkVisitorReport) SetData(v BulkVisitorReportData)`

SetData sets Data field to given value.

### HasData

`func (o *BulkVisitorReport) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


