# QueueConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An optional unique identifier for configurations in a notification configuration. If you don&#39;t provide one, ArvanCloud S3 will assign an ID. | [optional] 
**QueueArn** | **string** |  | 
**Events** | [**Array**](array.md) |  | 
**Filter** | Pointer to [**NotificationConfigurationFilter**](NotificationConfigurationFilter.md) |  | [optional] 

## Methods

### NewQueueConfiguration

`func NewQueueConfiguration(queueArn string, events Array, ) *QueueConfiguration`

NewQueueConfiguration instantiates a new QueueConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueConfigurationWithDefaults

`func NewQueueConfigurationWithDefaults() *QueueConfiguration`

NewQueueConfigurationWithDefaults instantiates a new QueueConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueueConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueueConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQueueArn

`func (o *QueueConfiguration) GetQueueArn() string`

GetQueueArn returns the QueueArn field if non-nil, zero value otherwise.

### GetQueueArnOk

`func (o *QueueConfiguration) GetQueueArnOk() (*string, bool)`

GetQueueArnOk returns a tuple with the QueueArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueArn

`func (o *QueueConfiguration) SetQueueArn(v string)`

SetQueueArn sets QueueArn field to given value.


### GetEvents

`func (o *QueueConfiguration) GetEvents() Array`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *QueueConfiguration) GetEventsOk() (*Array, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *QueueConfiguration) SetEvents(v Array)`

SetEvents sets Events field to given value.


### GetFilter

`func (o *QueueConfiguration) GetFilter() NotificationConfigurationFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *QueueConfiguration) GetFilterOk() (*NotificationConfigurationFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *QueueConfiguration) SetFilter(v NotificationConfigurationFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *QueueConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


