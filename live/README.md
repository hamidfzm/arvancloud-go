# Go API client for live

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 2.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import live "github.com/hamidfzm/arvancloud-go/live"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), live.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), live.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), live.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), live.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://napi.arvancloud.ir/live/2.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnalyticsAPI* | [**AnalyticsPlayCountGet**](docs/AnalyticsAPI.md#analyticsplaycountget) | **Get** /analytics/play-count | Return appropriate play count
*AnalyticsAPI* | [**AnalyticsTrafficGet**](docs/AnalyticsAPI.md#analyticstrafficget) | **Get** /analytics/traffic | Return appropriate traffic
*AnalyticsAPI* | [**AnalyticsWatchTimeGet**](docs/AnalyticsAPI.md#analyticswatchtimeget) | **Get** /analytics/watch-time | Return appropriate watch time
*DomainAPI* | [**DomainGet**](docs/DomainAPI.md#domainget) | **Get** /domain | Return User Domain.
*DomainAPI* | [**DomainPost**](docs/DomainAPI.md#domainpost) | **Post** /domain | Set subdomain for LIVE service.
*GeneralReportAPI* | [**ReportGeoGet**](docs/GeneralReportAPI.md#reportgeoget) | **Get** /report/geo | Return Domain Geo Report.
*GeneralReportAPI* | [**ReportStatisticsGet**](docs/GeneralReportAPI.md#reportstatisticsget) | **Get** /report/statistics | Return Domain statistics report.
*GeneralReportAPI* | [**ReportTrafficsGet**](docs/GeneralReportAPI.md#reporttrafficsget) | **Get** /report/traffics | Return Domain Traffic.
*GeneralReportAPI* | [**ReportUserAgentGet**](docs/GeneralReportAPI.md#reportuseragentget) | **Get** /report/user-agent | Return User Agent.
*GeneralReportAPI* | [**ReportVisitorsGet**](docs/GeneralReportAPI.md#reportvisitorsget) | **Get** /report/visitors | Return Domain Visitors.
*StreamAPI* | [**StreamsGet**](docs/StreamAPI.md#streamsget) | **Get** /streams | Return all streams.
*StreamAPI* | [**StreamsPost**](docs/StreamAPI.md#streamspost) | **Post** /streams | Store a newly created stream.
*StreamAPI* | [**StreamsStreamDelete**](docs/StreamAPI.md#streamsstreamdelete) | **Delete** /streams/{stream} | Remove the specified stream.
*StreamAPI* | [**StreamsStreamGet**](docs/StreamAPI.md#streamsstreamget) | **Get** /streams/{stream} | Return the specified stream.
*StreamAPI* | [**StreamsStreamPatch**](docs/StreamAPI.md#streamsstreampatch) | **Patch** /streams/{stream} | Update the specified stream.
*StreamAPI* | [**StreamsStreamStartRecordGet**](docs/StreamAPI.md#streamsstreamstartrecordget) | **Get** /streams/{stream}/start-record | Start record live stream.
*StreamAPI* | [**StreamsStreamStopRecordGet**](docs/StreamAPI.md#streamsstreamstoprecordget) | **Get** /streams/{stream}/stop-record | Stop record live stream.
*WatermarkAPI* | [**WatermarksGet**](docs/WatermarkAPI.md#watermarksget) | **Get** /watermarks | Return all watermarks.
*WatermarkAPI* | [**WatermarksPost**](docs/WatermarkAPI.md#watermarkspost) | **Post** /watermarks | Store a newly created Watermark.
*WatermarkAPI* | [**WatermarksWatermarkDelete**](docs/WatermarkAPI.md#watermarkswatermarkdelete) | **Delete** /watermarks/{watermark} | Remove the specified watermark.
*WatermarkAPI* | [**WatermarksWatermarkGet**](docs/WatermarkAPI.md#watermarkswatermarkget) | **Get** /watermarks/{watermark} | Return the specified watermark.
*WatermarkAPI* | [**WatermarksWatermarkPatch**](docs/WatermarkAPI.md#watermarkswatermarkpatch) | **Patch** /watermarks/{watermark} | Update the specified watermark.


## Documentation For Models

 - [DomainPostRequest](docs/DomainPostRequest.md)
 - [StreamsPostRequest](docs/StreamsPostRequest.md)
 - [StreamsPostRequestConvertInfoInner](docs/StreamsPostRequestConvertInfoInner.md)
 - [StreamsStreamPatchRequest](docs/StreamsStreamPatchRequest.md)
 - [WatermarksWatermarkPatchRequest](docs/WatermarksWatermarkPatchRequest.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api_key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


