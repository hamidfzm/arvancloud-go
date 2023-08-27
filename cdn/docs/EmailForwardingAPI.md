# \EmailForwardingAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailForwardingsActivate**](EmailForwardingAPI.md#EmailForwardingsActivate) | **Post** /domains/{domain}/email-forwardings/activate | Activate email forwarding
[**EmailForwardingsAliasesDestroy**](EmailForwardingAPI.md#EmailForwardingsAliasesDestroy) | **Delete** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId} | Delete an alias
[**EmailForwardingsAliasesIndex**](EmailForwardingAPI.md#EmailForwardingsAliasesIndex) | **Get** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases | List of email forwarding aliases for given domain
[**EmailForwardingsAliasesStore**](EmailForwardingAPI.md#EmailForwardingsAliasesStore) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases | Create new alias
[**EmailForwardingsAliasesToggleActivation**](EmailForwardingAPI.md#EmailForwardingsAliasesToggleActivation) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId}/toggle-activation | Toggle alias activation
[**EmailForwardingsAliasesUpdateRecipients**](EmailForwardingAPI.md#EmailForwardingsAliasesUpdateRecipients) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId}/recipients | Update alias recipients
[**EmailForwardingsCatchAllActivate**](EmailForwardingAPI.md#EmailForwardingsCatchAllActivate) | **Post** /domains/{domain}/email-forwardings/catch-all/activate | Activate email forwarding catch all
[**EmailForwardingsCatchAllDeactivate**](EmailForwardingAPI.md#EmailForwardingsCatchAllDeactivate) | **Post** /domains/{domain}/email-forwardings/catch-all/deactivate | Deactivate email forwarding catch all
[**EmailForwardingsDeactivate**](EmailForwardingAPI.md#EmailForwardingsDeactivate) | **Post** /domains/{domain}/email-forwardings/deactivate | Deactivate email forwarding
[**EmailForwardingsRecipientsDestroy**](EmailForwardingAPI.md#EmailForwardingsRecipientsDestroy) | **Delete** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId} | Delete a recipient
[**EmailForwardingsRecipientsIndex**](EmailForwardingAPI.md#EmailForwardingsRecipientsIndex) | **Get** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients | List recipients of an email forwarding
[**EmailForwardingsRecipientsResendVerification**](EmailForwardingAPI.md#EmailForwardingsRecipientsResendVerification) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/resend-verification | Resend Verification
[**EmailForwardingsRecipientsSetDefault**](EmailForwardingAPI.md#EmailForwardingsRecipientsSetDefault) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/set-default | Set default recipient
[**EmailForwardingsRecipientsStore**](EmailForwardingAPI.md#EmailForwardingsRecipientsStore) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients | Create new recipient
[**EmailForwardingsRecipientsVerify**](EmailForwardingAPI.md#EmailForwardingsRecipientsVerify) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/verify | Verify recipient
[**EmailForwardingsStats**](EmailForwardingAPI.md#EmailForwardingsStats) | **Get** /domains/{domain}/email-forwardings/stats | Show stats of domain&#39;s email forwarding



## EmailForwardingsActivate

> MessageResponse EmailForwardingsActivate(ctx, domain).Execute()

Activate email forwarding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsActivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsActivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsAliasesDestroy

> MessageResponse EmailForwardingsAliasesDestroy(ctx, domain, emailForwardingId, emailForwardingAliasId).Execute()

Delete an alias

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding alias id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesDestroy(context.Background(), domain, emailForwardingId, emailForwardingAliasId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsAliasesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsAliasesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingAliasId** | **string** | Email forwarding alias id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsAliasesIndex

> EmailForwardingAliasesListData EmailForwardingsAliasesIndex(ctx, domain, emailForwardingId).PerPage(perPage).Page(page).Execute()

List of email forwarding aliases for given domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesIndex(context.Background(), domain, emailForwardingId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsAliasesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesIndex`: EmailForwardingAliasesListData
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsAliasesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**EmailForwardingAliasesListData**](EmailForwardingAliasesListData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsAliasesStore

> EmailForwardingsAliasesStore201Response EmailForwardingsAliasesStore(ctx, domain, emailForwardingId).EmailForwardingAliasesStore(emailForwardingAliasesStore).Execute()

Create new alias

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasesStore := *openapiclient.NewEmailForwardingAliasesStore("LocalPart_example", []string{"Recipients_example"}) // EmailForwardingAliasesStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesStore(context.Background(), domain, emailForwardingId).EmailForwardingAliasesStore(emailForwardingAliasesStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsAliasesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesStore`: EmailForwardingsAliasesStore201Response
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsAliasesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailForwardingAliasesStore** | [**EmailForwardingAliasesStore**](EmailForwardingAliasesStore.md) |  | 

### Return type

[**EmailForwardingsAliasesStore201Response**](EmailForwardingsAliasesStore201Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsAliasesToggleActivation

> MessageResponse EmailForwardingsAliasesToggleActivation(ctx, domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesToggleActivation(emailForwardingAliasesToggleActivation).Execute()

Toggle alias activation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding alias id
    emailForwardingAliasesToggleActivation := *openapiclient.NewEmailForwardingAliasesToggleActivation(false) // EmailForwardingAliasesToggleActivation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesToggleActivation(context.Background(), domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesToggleActivation(emailForwardingAliasesToggleActivation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsAliasesToggleActivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesToggleActivation`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsAliasesToggleActivation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingAliasId** | **string** | Email forwarding alias id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesToggleActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **emailForwardingAliasesToggleActivation** | [**EmailForwardingAliasesToggleActivation**](EmailForwardingAliasesToggleActivation.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsAliasesUpdateRecipients

> MessageResponse EmailForwardingsAliasesUpdateRecipients(ctx, domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesRecipients(emailForwardingAliasesRecipients).Execute()

Update alias recipients

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingAliasId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding alias id
    emailForwardingAliasesRecipients := *openapiclient.NewEmailForwardingAliasesRecipients([]string{"Recipients_example"}) // EmailForwardingAliasesRecipients |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesUpdateRecipients(context.Background(), domain, emailForwardingId, emailForwardingAliasId).EmailForwardingAliasesRecipients(emailForwardingAliasesRecipients).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsAliasesUpdateRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsAliasesUpdateRecipients`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsAliasesUpdateRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingAliasId** | **string** | Email forwarding alias id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsAliasesUpdateRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **emailForwardingAliasesRecipients** | [**EmailForwardingAliasesRecipients**](EmailForwardingAliasesRecipients.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsCatchAllActivate

> MessageResponse EmailForwardingsCatchAllActivate(ctx, domain).Execute()

Activate email forwarding catch all

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsCatchAllActivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsCatchAllActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsCatchAllActivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsCatchAllActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsCatchAllActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsCatchAllDeactivate

> MessageResponse EmailForwardingsCatchAllDeactivate(ctx, domain).Execute()

Deactivate email forwarding catch all

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsCatchAllDeactivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsCatchAllDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsCatchAllDeactivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsCatchAllDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsCatchAllDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsDeactivate

> MessageResponse EmailForwardingsDeactivate(ctx, domain).Execute()

Deactivate email forwarding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsDeactivate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsDeactivate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsRecipientsDestroy

> MessageResponse EmailForwardingsRecipientsDestroy(ctx, domain, emailForwardingId, emailForwardingRecipientId).Execute()

Delete a recipient

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsDestroy(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsRecipientsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsRecipientsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsRecipientsIndex

> EmailForwardingRecipientsListData EmailForwardingsRecipientsIndex(ctx, domain, emailForwardingId).PerPage(perPage).Page(page).Execute()

List recipients of an email forwarding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsIndex(context.Background(), domain, emailForwardingId).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsRecipientsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsIndex`: EmailForwardingRecipientsListData
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsRecipientsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**EmailForwardingRecipientsListData**](EmailForwardingRecipientsListData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsRecipientsResendVerification

> MessageResponse EmailForwardingsRecipientsResendVerification(ctx, domain, emailForwardingId, emailForwardingRecipientId).Execute()

Resend Verification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsResendVerification(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsRecipientsResendVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsResendVerification`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsRecipientsResendVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsResendVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsRecipientsSetDefault

> MessageResponse EmailForwardingsRecipientsSetDefault(ctx, domain, emailForwardingId, emailForwardingRecipientId).Execute()

Set default recipient

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsSetDefault(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsRecipientsSetDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsSetDefault`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsRecipientsSetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsSetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsRecipientsStore

> EmailForwardingsRecipientsStore201Response EmailForwardingsRecipientsStore(ctx, domain, emailForwardingId).EmailForwardingRecipientsStore(emailForwardingRecipientsStore).Execute()

Create new recipient

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientsStore := *openapiclient.NewEmailForwardingRecipientsStore("Email_example") // EmailForwardingRecipientsStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsStore(context.Background(), domain, emailForwardingId).EmailForwardingRecipientsStore(emailForwardingRecipientsStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsRecipientsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsStore`: EmailForwardingsRecipientsStore201Response
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsRecipientsStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailForwardingRecipientsStore** | [**EmailForwardingRecipientsStore**](EmailForwardingRecipientsStore.md) |  | 

### Return type

[**EmailForwardingsRecipientsStore201Response**](EmailForwardingsRecipientsStore201Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsRecipientsVerify

> MessageResponse EmailForwardingsRecipientsVerify(ctx, domain, emailForwardingId, emailForwardingRecipientId).EmailForwardingRecipientsVerify(emailForwardingRecipientsVerify).Execute()

Verify recipient

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    emailForwardingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding id
    emailForwardingRecipientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Email forwarding recipient id
    emailForwardingRecipientsVerify := *openapiclient.NewEmailForwardingRecipientsVerify("Code_example") // EmailForwardingRecipientsVerify |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsVerify(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).EmailForwardingRecipientsVerify(emailForwardingRecipientsVerify).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsRecipientsVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsRecipientsVerify`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsRecipientsVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**emailForwardingId** | **string** | Email forwarding id | 
**emailForwardingRecipientId** | **string** | Email forwarding recipient id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsRecipientsVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **emailForwardingRecipientsVerify** | [**EmailForwardingRecipientsVerify**](EmailForwardingRecipientsVerify.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailForwardingsStats

> EmailForwardingStatsData EmailForwardingsStats(ctx, domain).Execute()

Show stats of domain's email forwarding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailForwardingAPI.EmailForwardingsStats(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailForwardingAPI.EmailForwardingsStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmailForwardingsStats`: EmailForwardingStatsData
    fmt.Fprintf(os.Stdout, "Response from `EmailForwardingAPI.EmailForwardingsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmailForwardingsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailForwardingStatsData**](EmailForwardingStatsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

