# Bitget\MarginIsolatedFinflowApi

All URIs are relative to https://api.bitget.com, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**isolatedFinList()**](MarginIsolatedFinflowApi.md#isolatedFinList) | **GET** /api/margin/v1/isolated/fin/list | list |


## `isolatedFinList()`

```php
isolatedFinList($symbol, $start_time, $coin, $margin_type, $end_time, $loan_id, $page_size, $page_id): \Bitget\Model\ApiResponseResultOfMarginIsolatedFinFlowResult
```

list

Get finance flow List

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: ACCESS_KEY
$config = Bitget\Configuration::getDefaultConfiguration()->setApiKey('ACCESS-KEY', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Bitget\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ACCESS-KEY', 'Bearer');

// Configure API key authorization: ACCESS_PASSPHRASE
$config = Bitget\Configuration::getDefaultConfiguration()->setApiKey('ACCESS-PASSPHRASE', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Bitget\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ACCESS-PASSPHRASE', 'Bearer');

// Configure API key authorization: ACCESS_SIGN
$config = Bitget\Configuration::getDefaultConfiguration()->setApiKey('ACCESS-SIGN', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Bitget\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ACCESS-SIGN', 'Bearer');

// Configure API key authorization: ACCESS_TIMESTAMP
$config = Bitget\Configuration::getDefaultConfiguration()->setApiKey('ACCESS-TIMESTAMP', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Bitget\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ACCESS-TIMESTAMP', 'Bearer');

// Configure API key authorization: SECRET_KEY
$config = Bitget\Configuration::getDefaultConfiguration()->setApiKey('SECRET-KEY', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = Bitget\Configuration::getDefaultConfiguration()->setApiKeyPrefix('SECRET-KEY', 'Bearer');


$apiInstance = new Bitget\Api\MarginIsolatedFinflowApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$symbol = BTCUSDT; // string | symbol
$start_time = 1678193338000; // string | startTime
$coin = USDT; // string | coin
$margin_type = transfer_in; // string | marginType
$end_time = 1678193338000; // string | endTime
$loan_id = 'loan_id_example'; // string | loanId
$page_size = 10; // string | pageSize
$page_id = 'page_id_example'; // string | pageId

try {
    $result = $apiInstance->isolatedFinList($symbol, $start_time, $coin, $margin_type, $end_time, $loan_id, $page_size, $page_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MarginIsolatedFinflowApi->isolatedFinList: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **symbol** | **string**| symbol | |
| **start_time** | **string**| startTime | |
| **coin** | **string**| coin | [optional] |
| **margin_type** | **string**| marginType | [optional] |
| **end_time** | **string**| endTime | [optional] |
| **loan_id** | **string**| loanId | [optional] |
| **page_size** | **string**| pageSize | [optional] |
| **page_id** | **string**| pageId | [optional] |

### Return type

[**\Bitget\Model\ApiResponseResultOfMarginIsolatedFinFlowResult**](../Model/ApiResponseResultOfMarginIsolatedFinFlowResult.md)

### Authorization

[ACCESS_KEY](../../README.md#ACCESS_KEY), [ACCESS_PASSPHRASE](../../README.md#ACCESS_PASSPHRASE), [ACCESS_SIGN](../../README.md#ACCESS_SIGN), [ACCESS_TIMESTAMP](../../README.md#ACCESS_TIMESTAMP), [SECRET_KEY](../../README.md#SECRET_KEY)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)