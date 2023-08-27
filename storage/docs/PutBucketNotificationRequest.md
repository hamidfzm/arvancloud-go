# PutBucketNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationConfiguration** | [**NotificationConfigurationDeprecated**](NotificationConfigurationDeprecated.md) |  | 

## Methods

### NewPutBucketNotificationRequest

`func NewPutBucketNotificationRequest(notificationConfiguration NotificationConfigurationDeprecated, ) *PutBucketNotificationRequest`

NewPutBucketNotificationRequest instantiates a new PutBucketNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketNotificationRequestWithDefaults

`func NewPutBucketNotificationRequestWithDefaults() *PutBucketNotificationRequest`

NewPutBucketNotificationRequestWithDefaults instantiates a new PutBucketNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationConfiguration

`func (o *PutBucketNotificationRequest) GetNotificationConfiguration() NotificationConfigurationDeprecated`

GetNotificationConfiguration returns the NotificationConfiguration field if non-nil, zero value otherwise.

### GetNotificationConfigurationOk

`func (o *PutBucketNotificationRequest) GetNotificationConfigurationOk() (*NotificationConfigurationDeprecated, bool)`

GetNotificationConfigurationOk returns a tuple with the NotificationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationConfiguration

`func (o *PutBucketNotificationRequest) SetNotificationConfiguration(v NotificationConfigurationDeprecated)`

SetNotificationConfiguration sets NotificationConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


