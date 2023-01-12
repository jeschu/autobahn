# \DefaultApi

All URIs are relative to *https://verkehr.autobahn.de/o/autobahn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChargingStation**](DefaultApi.md#GetChargingStation) | **Get** /details/electric_charging_station/{stationId} | Details zu einer Ladestation
[**GetClosure**](DefaultApi.md#GetClosure) | **Get** /details/closure/{closureId} | Details zu einer Sperrung
[**GetParkingLorry**](DefaultApi.md#GetParkingLorry) | **Get** /details/parking_lorry/{lorryId} | Details eines Rastplatzes
[**GetRoadwork**](DefaultApi.md#GetRoadwork) | **Get** /details/roadworks/{roadworkId} | Details einer Baustelle
[**GetWarning**](DefaultApi.md#GetWarning) | **Get** /details/warning/{warningId} | Details zu einer Verkehrsmeldung
[**GetWebcam**](DefaultApi.md#GetWebcam) | **Get** /details/webcam/{webcamId} | Details einer Webcam
[**ListAutobahnen**](DefaultApi.md#ListAutobahnen) | **Get** / | Liste verfügbarer Autobahnen
[**ListChargingStations**](DefaultApi.md#ListChargingStations) | **Get** /{roadId}/services/electric_charging_station | Liste aktueller Ladestationen
[**ListClosures**](DefaultApi.md#ListClosures) | **Get** /{roadId}/services/closure | Liste aktueller Sperrungen
[**ListParkingLorries**](DefaultApi.md#ListParkingLorries) | **Get** /{roadId}/services/parking_lorry | Liste verfügbarer Rastplätze
[**ListRoadworks**](DefaultApi.md#ListRoadworks) | **Get** /{roadId}/services/roadworks | Liste aktueller Baustellen
[**ListWarnings**](DefaultApi.md#ListWarnings) | **Get** /{roadId}/services/warning | Liste aktueller Verkehrsmeldungen
[**ListWebcams**](DefaultApi.md#ListWebcams) | **Get** /{roadId}/services/webcam | Liste verfügbarer Webcams



## GetChargingStation

> ElectricChargingStation GetChargingStation(ctx, stationId).Execute()

Details zu einer Ladestation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationId := string([B@9a2ec9b) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetChargingStation(context.Background(), stationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetChargingStation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChargingStation`: ElectricChargingStation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetChargingStation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChargingStationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ElectricChargingStation**](ElectricChargingStation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClosure

> Closure GetClosure(ctx, closureId).Execute()

Details zu einer Sperrung



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    closureId := string([B@719843e5) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetClosure(context.Background(), closureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClosure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClosure`: Closure
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClosure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**closureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClosureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Closure**](Closure.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParkingLorry

> ParkingLorry GetParkingLorry(ctx, lorryId).Execute()

Details eines Rastplatzes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    lorryId := string([B@73044cdf) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetParkingLorry(context.Background(), lorryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetParkingLorry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParkingLorry`: ParkingLorry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetParkingLorry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lorryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParkingLorryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParkingLorry**](ParkingLorry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoadwork

> Roadwork GetRoadwork(ctx, roadworkId).Execute()

Details einer Baustelle



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadworkId := string([B@7fc645e4) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRoadwork(context.Background(), roadworkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRoadwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoadwork`: Roadwork
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRoadwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadworkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoadworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Roadwork**](Roadwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarning

> Warning GetWarning(ctx, warningId).Execute()

Details zu einer Verkehrsmeldung



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    warningId := string([B@3cbcd8f3) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWarning(context.Background(), warningId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWarning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWarning`: Warning
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWarning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warningId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Warning**](Warning.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebcam

> Webcam GetWebcam(ctx, webcamId).Execute()

Details einer Webcam



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webcamId := string([B@205b132e) // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWebcam(context.Background(), webcamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWebcam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebcam`: Webcam
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWebcam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webcamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebcamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webcam**](Webcam.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutobahnen

> Roads ListAutobahnen(ctx).Execute()

Liste verfügbarer Autobahnen



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAutobahnen(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAutobahnen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutobahnen`: Roads
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAutobahnen`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAutobahnenRequest struct via the builder pattern


### Return type

[**Roads**](Roads.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChargingStations

> ElectricChargingStations ListChargingStations(ctx, roadId).Execute()

Liste aktueller Ladestationen



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadId := "roadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListChargingStations(context.Background(), roadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListChargingStations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChargingStations`: ElectricChargingStations
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListChargingStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListChargingStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ElectricChargingStations**](ElectricChargingStations.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClosures

> Closures ListClosures(ctx, roadId).Execute()

Liste aktueller Sperrungen



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadId := "roadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListClosures(context.Background(), roadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListClosures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClosures`: Closures
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListClosures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClosuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Closures**](Closures.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParkingLorries

> ParkingLorries ListParkingLorries(ctx, roadId).Execute()

Liste verfügbarer Rastplätze



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadId := "roadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListParkingLorries(context.Background(), roadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListParkingLorries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListParkingLorries`: ParkingLorries
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListParkingLorries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListParkingLorriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParkingLorries**](ParkingLorries.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoadworks

> Roadworks ListRoadworks(ctx, roadId).Execute()

Liste aktueller Baustellen



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadId := "roadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRoadworks(context.Background(), roadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRoadworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoadworks`: Roadworks
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRoadworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRoadworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Roadworks**](Roadworks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWarnings

> Warnings ListWarnings(ctx, roadId).Execute()

Liste aktueller Verkehrsmeldungen



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadId := "roadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListWarnings(context.Background(), roadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWarnings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWarnings`: Warnings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWarnings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWarningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Warnings**](Warnings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebcams

> Webcams ListWebcams(ctx, roadId).Execute()

Liste verfügbarer Webcams



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    roadId := "roadId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListWebcams(context.Background(), roadId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebcams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebcams`: Webcams
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebcams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebcamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webcams**](Webcams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

