# TopicConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An optional unique identifier for configurations in a notification configuration. If you don&#39;t provide one, ArvanCloud S3 will assign an ID. | [optional] 
**TopicArn** | **string** |  | 
**Events** | [**Array**](array.md) |  | 
**Filter** | Pointer to [**NotificationConfigurationFilter**](NotificationConfigurationFilter.md) |  | [optional] 

## Methods

### NewTopicConfiguration

`func NewTopicConfiguration(topicArn string, events Array, ) *TopicConfiguration`

NewTopicConfiguration instantiates a new TopicConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicConfigurationWithDefaults

`func NewTopicConfigurationWithDefaults() *TopicConfiguration`

NewTopicConfigurationWithDefaults instantiates a new TopicConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopicConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopicConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopicConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TopicConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTopicArn

`func (o *TopicConfiguration) GetTopicArn() string`

GetTopicArn returns the TopicArn field if non-nil, zero value otherwise.

### GetTopicArnOk

`func (o *TopicConfiguration) GetTopicArnOk() (*string, bool)`

GetTopicArnOk returns a tuple with the TopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicArn

`func (o *TopicConfiguration) SetTopicArn(v string)`

SetTopicArn sets TopicArn field to given value.


### GetEvents

`func (o *TopicConfiguration) GetEvents() Array`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TopicConfiguration) GetEventsOk() (*Array, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TopicConfiguration) SetEvents(v Array)`

SetEvents sets Events field to given value.


### GetFilter

`func (o *TopicConfiguration) GetFilter() NotificationConfigurationFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *TopicConfiguration) GetFilterOk() (*NotificationConfigurationFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *TopicConfiguration) SetFilter(v NotificationConfigurationFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *TopicConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


