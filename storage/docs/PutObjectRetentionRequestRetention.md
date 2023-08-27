# PutObjectRetentionRequestRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**ObjectLockRetentionMode**](ObjectLockRetentionMode.md) |  | [optional] 
**RetainUntilDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewPutObjectRetentionRequestRetention

`func NewPutObjectRetentionRequestRetention() *PutObjectRetentionRequestRetention`

NewPutObjectRetentionRequestRetention instantiates a new PutObjectRetentionRequestRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectRetentionRequestRetentionWithDefaults

`func NewPutObjectRetentionRequestRetentionWithDefaults() *PutObjectRetentionRequestRetention`

NewPutObjectRetentionRequestRetentionWithDefaults instantiates a new PutObjectRetentionRequestRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PutObjectRetentionRequestRetention) GetMode() ObjectLockRetentionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PutObjectRetentionRequestRetention) GetModeOk() (*ObjectLockRetentionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PutObjectRetentionRequestRetention) SetMode(v ObjectLockRetentionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PutObjectRetentionRequestRetention) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRetainUntilDate

`func (o *PutObjectRetentionRequestRetention) GetRetainUntilDate() time.Time`

GetRetainUntilDate returns the RetainUntilDate field if non-nil, zero value otherwise.

### GetRetainUntilDateOk

`func (o *PutObjectRetentionRequestRetention) GetRetainUntilDateOk() (*time.Time, bool)`

GetRetainUntilDateOk returns a tuple with the RetainUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainUntilDate

`func (o *PutObjectRetentionRequestRetention) SetRetainUntilDate(v time.Time)`

SetRetainUntilDate sets RetainUntilDate field to given value.

### HasRetainUntilDate

`func (o *PutObjectRetentionRequestRetention) HasRetainUntilDate() bool`

HasRetainUntilDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


