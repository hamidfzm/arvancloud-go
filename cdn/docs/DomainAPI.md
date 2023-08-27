# \DomainAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsClone**](DomainAPI.md#DomainsClone) | **Post** /domains/{domain}/clone | Clone a domain config from another one
[**DomainsCnameSetupCheck**](DomainAPI.md#DomainsCnameSetupCheck) | **Get** /domains/{domain}/cname-setup/check | Check Cname Setup to find whether domain is activated
[**DomainsCnameSetupConvert**](DomainAPI.md#DomainsCnameSetupConvert) | **Post** /domains/{domain}/cname-setup/convert | Convert domain setup to cname
[**DomainsCnameSetupReset**](DomainAPI.md#DomainsCnameSetupReset) | **Delete** /domains/{domain}/cname-setup/custom | Reset the custom record of CNAME Setup to the default value
[**DomainsCnameSetupSet**](DomainAPI.md#DomainsCnameSetupSet) | **Put** /domains/{domain}/cname-setup/custom | Set a custom record for using CNAME Setup
[**DomainsDestroy**](DomainAPI.md#DomainsDestroy) | **Delete** /domains/{domain} | Remove the domain
[**DomainsIndex**](DomainAPI.md#DomainsIndex) | **Get** /domains | Get the list of domains
[**DomainsNameserversCheck**](DomainAPI.md#DomainsNameserversCheck) | **Get** /domains/{domain}/ns-keys/check | Check NS to find whether domain is activated
[**DomainsNameserversDeprecatedCheck**](DomainAPI.md#DomainsNameserversDeprecatedCheck) | **Put** /domains/{domain}/dns-service/check-ns | Deprecated in favor /ns-keys/check
[**DomainsNameserversOptional**](DomainAPI.md#DomainsNameserversOptional) | **Post** /domains/{domain}/ns-keys/use-optional-keys | Use optional NS keys
[**DomainsNameserversReset**](DomainAPI.md#DomainsNameserversReset) | **Delete** /domains/{domain}/ns-keys | Reset custom Nameserver keys to the default values for the domain
[**DomainsNameserversSet**](DomainAPI.md#DomainsNameserversSet) | **Put** /domains/{domain}/ns-keys | Set custom NS records for the domain
[**DomainsRegenerate**](DomainAPI.md#DomainsRegenerate) | **Post** /domains/{domain}/regenerate | Regenerate domain settings for edge servers
[**DomainsShow**](DomainAPI.md#DomainsShow) | **Get** /domains/{domain} | Get information of the domain
[**DomainsStore**](DomainAPI.md#DomainsStore) | **Post** /domains/dns-service | Create new domain



## DomainsClone

> MessageResponse DomainsClone(ctx, domain).CloneDomain(cloneDomain).Execute()

Clone a domain config from another one

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
    cloneDomain := *openapiclient.NewCloneDomain("From_example") // CloneDomain |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainsClone(context.Background(), domain).CloneDomain(cloneDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsClone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsClone`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneDomain** | [**CloneDomain**](CloneDomain.md) |  | 

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


## DomainsCnameSetupCheck

> DomainResponse DomainsCnameSetupCheck(ctx, domain).Execute()

Check Cname Setup to find whether domain is activated

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
    resp, r, err := apiClient.DomainAPI.DomainsCnameSetupCheck(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsCnameSetupCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCnameSetupCheck`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsCnameSetupCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCnameSetupCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsCnameSetupConvert

> DomainResponse DomainsCnameSetupConvert(ctx, domain).Execute()

Convert domain setup to cname

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
    resp, r, err := apiClient.DomainAPI.DomainsCnameSetupConvert(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsCnameSetupConvert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCnameSetupConvert`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsCnameSetupConvert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCnameSetupConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsCnameSetupReset

> DomainResponse DomainsCnameSetupReset(ctx, domain).Execute()

Reset the custom record of CNAME Setup to the default value

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
    resp, r, err := apiClient.DomainAPI.DomainsCnameSetupReset(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsCnameSetupReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCnameSetupReset`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsCnameSetupReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCnameSetupResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsCnameSetupSet

> DomainResponse DomainsCnameSetupSet(ctx, domain).CustomCname(customCname).Execute()

Set a custom record for using CNAME Setup

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
    customCname := *openapiclient.NewCustomCname("Address_example") // CustomCname |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainsCnameSetupSet(context.Background(), domain).CustomCname(customCname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsCnameSetupSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsCnameSetupSet`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsCnameSetupSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsCnameSetupSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customCname** | [**CustomCname**](CustomCname.md) |  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDestroy

> MessageResponse DomainsDestroy(ctx, domain).Id(id).Execute()

Remove the domain

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainsDestroy(context.Background(), domain).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** |  | 

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


## DomainsIndex

> DomainsIndex200Response DomainsIndex(ctx).Search(search).PerPage(perPage).Page(page).Execute()

Get the list of domains



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
    search := "search_example" // string | Search term (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainsIndex(context.Background()).Search(search).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsIndex`: DomainsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Search term | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**DomainsIndex200Response**](DomainsIndex200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsNameserversCheck

> DomainsNameserversCheck200Response DomainsNameserversCheck(ctx, domain).Execute()

Check NS to find whether domain is activated

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
    resp, r, err := apiClient.DomainAPI.DomainsNameserversCheck(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsNameserversCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsNameserversCheck`: DomainsNameserversCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsNameserversCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsNameserversCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainsNameserversCheck200Response**](DomainsNameserversCheck200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsNameserversDeprecatedCheck

> DomainsNameserversDeprecatedCheck200Response DomainsNameserversDeprecatedCheck(ctx, domain).Execute()

Deprecated in favor /ns-keys/check

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
    resp, r, err := apiClient.DomainAPI.DomainsNameserversDeprecatedCheck(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsNameserversDeprecatedCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsNameserversDeprecatedCheck`: DomainsNameserversDeprecatedCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsNameserversDeprecatedCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsNameserversDeprecatedCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainsNameserversDeprecatedCheck200Response**](DomainsNameserversDeprecatedCheck200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsNameserversOptional

> NsKeysResponse DomainsNameserversOptional(ctx, domain).Execute()

Use optional NS keys

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
    resp, r, err := apiClient.DomainAPI.DomainsNameserversOptional(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsNameserversOptional``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsNameserversOptional`: NsKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsNameserversOptional`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsNameserversOptionalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NsKeysResponse**](NsKeysResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsNameserversReset

> NsKeysResponse DomainsNameserversReset(ctx, domain).Execute()

Reset custom Nameserver keys to the default values for the domain

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
    resp, r, err := apiClient.DomainAPI.DomainsNameserversReset(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsNameserversReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsNameserversReset`: NsKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsNameserversReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsNameserversResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NsKeysResponse**](NsKeysResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsNameserversSet

> NsKeysResponse DomainsNameserversSet(ctx, domain).NsKeys(nsKeys).Execute()

Set custom NS records for the domain

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
    nsKeys := *openapiclient.NewNsKeys() // NsKeys |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainsNameserversSet(context.Background(), domain).NsKeys(nsKeys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsNameserversSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsNameserversSet`: NsKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsNameserversSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsNameserversSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nsKeys** | [**NsKeys**](NsKeys.md) |  | 

### Return type

[**NsKeysResponse**](NsKeysResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsRegenerate

> MessageResponse DomainsRegenerate(ctx, domain).Execute()

Regenerate domain settings for edge servers

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
    resp, r, err := apiClient.DomainAPI.DomainsRegenerate(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsRegenerate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsRegenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsRegenerateRequest struct via the builder pattern


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


## DomainsShow

> DomainResponse DomainsShow(ctx, domain).Execute()

Get information of the domain

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
    resp, r, err := apiClient.DomainAPI.DomainsShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsShow`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsStore

> DomainResponse DomainsStore(ctx).DomainStore(domainStore).Execute()

Create new domain

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
    domainStore := *openapiclient.NewDomainStore("Domain_example") // DomainStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainAPI.DomainsStore(context.Background()).DomainStore(domainStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.DomainsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsStore`: DomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DomainAPI.DomainsStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainStore** | [**DomainStore**](DomainStore.md) |  | 

### Return type

[**DomainResponse**](DomainResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

