# DefaultApi

All URIs are relative to *https://circleci.com/api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**meGet**](DefaultApi.md#meGet) | **GET** /me |  |
| [**projectUsernameProjectBuildCacheDelete**](DefaultApi.md#projectUsernameProjectBuildCacheDelete) | **DELETE** /project/{username}/{project}/build-cache |  |
| [**projectUsernameProjectBuildNumArtifactsGet**](DefaultApi.md#projectUsernameProjectBuildNumArtifactsGet) | **GET** /project/{username}/{project}/{build_num}/artifacts |  |
| [**projectUsernameProjectBuildNumCancelPost**](DefaultApi.md#projectUsernameProjectBuildNumCancelPost) | **POST** /project/{username}/{project}/{build_num}/cancel |  |
| [**projectUsernameProjectBuildNumGet**](DefaultApi.md#projectUsernameProjectBuildNumGet) | **GET** /project/{username}/{project}/{build_num} |  |
| [**projectUsernameProjectBuildNumRetryPost**](DefaultApi.md#projectUsernameProjectBuildNumRetryPost) | **POST** /project/{username}/{project}/{build_num}/retry |  |
| [**projectUsernameProjectBuildNumTestsGet**](DefaultApi.md#projectUsernameProjectBuildNumTestsGet) | **GET** /project/{username}/{project}/{build_num}/tests |  |
| [**projectUsernameProjectCheckoutKeyFingerprintDelete**](DefaultApi.md#projectUsernameProjectCheckoutKeyFingerprintDelete) | **DELETE** /project/{username}/{project}/checkout-key/{fingerprint} |  |
| [**projectUsernameProjectCheckoutKeyFingerprintGet**](DefaultApi.md#projectUsernameProjectCheckoutKeyFingerprintGet) | **GET** /project/{username}/{project}/checkout-key/{fingerprint} |  |
| [**projectUsernameProjectCheckoutKeyGet**](DefaultApi.md#projectUsernameProjectCheckoutKeyGet) | **GET** /project/{username}/{project}/checkout-key |  |
| [**projectUsernameProjectCheckoutKeyPost**](DefaultApi.md#projectUsernameProjectCheckoutKeyPost) | **POST** /project/{username}/{project}/checkout-key |  |
| [**projectUsernameProjectEnvvarGet**](DefaultApi.md#projectUsernameProjectEnvvarGet) | **GET** /project/{username}/{project}/envvar |  |
| [**projectUsernameProjectEnvvarNameDelete**](DefaultApi.md#projectUsernameProjectEnvvarNameDelete) | **DELETE** /project/{username}/{project}/envvar/{name} |  |
| [**projectUsernameProjectEnvvarNameGet**](DefaultApi.md#projectUsernameProjectEnvvarNameGet) | **GET** /project/{username}/{project}/envvar/{name} |  |
| [**projectUsernameProjectEnvvarPost**](DefaultApi.md#projectUsernameProjectEnvvarPost) | **POST** /project/{username}/{project}/envvar |  |
| [**projectUsernameProjectGet**](DefaultApi.md#projectUsernameProjectGet) | **GET** /project/{username}/{project} |  |
| [**projectUsernameProjectPost**](DefaultApi.md#projectUsernameProjectPost) | **POST** /project/{username}/{project} |  |
| [**projectUsernameProjectSshKeyPost**](DefaultApi.md#projectUsernameProjectSshKeyPost) | **POST** /project/{username}/{project}/ssh-key |  |
| [**projectUsernameProjectTreeBranchPost**](DefaultApi.md#projectUsernameProjectTreeBranchPost) | **POST** /project/{username}/{project}/tree/{branch} |  |
| [**projectsGet**](DefaultApi.md#projectsGet) | **GET** /projects |  |
| [**recentBuildsGet**](DefaultApi.md#recentBuildsGet) | **GET** /recent-builds |  |
| [**userHerokuKeyPost**](DefaultApi.md#userHerokuKeyPost) | **POST** /user/heroku-key |  |


<a id="meGet"></a>
# **meGet**
> User meGet()



Provides information about the signed in user. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      User result = apiInstance.meGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#meGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | signed in user |  -  |

<a id="projectUsernameProjectBuildCacheDelete"></a>
# **projectUsernameProjectBuildCacheDelete**
> ProjectUsernameProjectBuildCacheDelete200Response projectUsernameProjectBuildCacheDelete(username, project)



Clears the cache for a project. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    try {
      ProjectUsernameProjectBuildCacheDelete200Response result = apiInstance.projectUsernameProjectBuildCacheDelete(username, project);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectBuildCacheDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |

### Return type

[**ProjectUsernameProjectBuildCacheDelete200Response**](ProjectUsernameProjectBuildCacheDelete200Response.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | status message |  -  |

<a id="projectUsernameProjectBuildNumArtifactsGet"></a>
# **projectUsernameProjectBuildNumArtifactsGet**
> List&lt;Artifact&gt; projectUsernameProjectBuildNumArtifactsGet(username, project, buildNum)



List the artifacts produced by a given build. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    Integer buildNum = 56; // Integer | XXXXXXXXXX 
    try {
      List<Artifact> result = apiInstance.projectUsernameProjectBuildNumArtifactsGet(username, project, buildNum);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectBuildNumArtifactsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **buildNum** | **Integer**| XXXXXXXXXX  | |

### Return type

[**List&lt;Artifact&gt;**](Artifact.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List the artifacts produced by a given build |  -  |

<a id="projectUsernameProjectBuildNumCancelPost"></a>
# **projectUsernameProjectBuildNumCancelPost**
> Build projectUsernameProjectBuildNumCancelPost(username, project, buildNum)



Cancels the build, returns a summary of the build. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    Integer buildNum = 56; // Integer | XXXXXXXXXX 
    try {
      Build result = apiInstance.projectUsernameProjectBuildNumCancelPost(username, project, buildNum);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectBuildNumCancelPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **buildNum** | **Integer**| XXXXXXXXXX  | |

### Return type

[**Build**](Build.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | returns a summary of the build |  -  |

<a id="projectUsernameProjectBuildNumGet"></a>
# **projectUsernameProjectBuildNumGet**
> BuildDetail projectUsernameProjectBuildNumGet(username, project, buildNum)



Full details for a single build. The response includes all of the fields from the build summary. This is also the payload for the [notification webhooks](/docs/configuration/#notify), in which case this object is the value to a key named &#39;payload&#39;. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    Integer buildNum = 56; // Integer | XXXXXXXXXX 
    try {
      BuildDetail result = apiInstance.projectUsernameProjectBuildNumGet(username, project, buildNum);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectBuildNumGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **buildNum** | **Integer**| XXXXXXXXXX  | |

### Return type

[**BuildDetail**](BuildDetail.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Full details for a single build |  -  |

<a id="projectUsernameProjectBuildNumRetryPost"></a>
# **projectUsernameProjectBuildNumRetryPost**
> Build projectUsernameProjectBuildNumRetryPost(username, project, buildNum)



Retries the build, returns a summary of the new build. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    Integer buildNum = 56; // Integer | XXXXXXXXXX 
    try {
      Build result = apiInstance.projectUsernameProjectBuildNumRetryPost(username, project, buildNum);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectBuildNumRetryPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **buildNum** | **Integer**| XXXXXXXXXX  | |

### Return type

[**Build**](Build.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | returns a summary of the new build |  -  |

<a id="projectUsernameProjectBuildNumTestsGet"></a>
# **projectUsernameProjectBuildNumTestsGet**
> Tests projectUsernameProjectBuildNumTestsGet(username, project, buildNum)



Provides test metadata for a build Note: [Learn how to set up your builds to collect test metadata](https://circleci.com/docs/test-metadata/) 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    Integer buildNum = 56; // Integer | XXXXXXXXXX 
    try {
      Tests result = apiInstance.projectUsernameProjectBuildNumTestsGet(username, project, buildNum);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectBuildNumTestsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **buildNum** | **Integer**| XXXXXXXXXX  | |

### Return type

[**Tests**](Tests.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | test metadata for a build  |  -  |

<a id="projectUsernameProjectCheckoutKeyFingerprintDelete"></a>
# **projectUsernameProjectCheckoutKeyFingerprintDelete**
> ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response projectUsernameProjectCheckoutKeyFingerprintDelete(username, project, fingerprint)



Delete a checkout key. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String fingerprint = "fingerprint_example"; // String | XXXXXXXXXX 
    try {
      ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response result = apiInstance.projectUsernameProjectCheckoutKeyFingerprintDelete(username, project, fingerprint);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectCheckoutKeyFingerprintDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **fingerprint** | **String**| XXXXXXXXXX  | |

### Return type

[**ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response**](ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | status message |  -  |

<a id="projectUsernameProjectCheckoutKeyFingerprintGet"></a>
# **projectUsernameProjectCheckoutKeyFingerprintGet**
> Key projectUsernameProjectCheckoutKeyFingerprintGet(username, project, fingerprint)



Get a checkout key. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String fingerprint = "fingerprint_example"; // String | XXXXXXXXXX 
    try {
      Key result = apiInstance.projectUsernameProjectCheckoutKeyFingerprintGet(username, project, fingerprint);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectCheckoutKeyFingerprintGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **fingerprint** | **String**| XXXXXXXXXX  | |

### Return type

[**Key**](Key.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | checkout key |  -  |

<a id="projectUsernameProjectCheckoutKeyGet"></a>
# **projectUsernameProjectCheckoutKeyGet**
> List&lt;Key&gt; projectUsernameProjectCheckoutKeyGet(username, project)



Lists checkout keys. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    try {
      List<Key> result = apiInstance.projectUsernameProjectCheckoutKeyGet(username, project);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectCheckoutKeyGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |

### Return type

[**List&lt;Key&gt;**](Key.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | checkout keys |  -  |

<a id="projectUsernameProjectCheckoutKeyPost"></a>
# **projectUsernameProjectCheckoutKeyPost**
> Key projectUsernameProjectCheckoutKeyPost(username, project, body)



Creates a new checkout key. Only usable with a user API token. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String body = "body_example"; // String | The type of key to create. Can be 'deploy-key' or 'github-user-key'. 
    try {
      Key result = apiInstance.projectUsernameProjectCheckoutKeyPost(username, project, body);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectCheckoutKeyPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **body** | **String**| The type of key to create. Can be &#39;deploy-key&#39; or &#39;github-user-key&#39;.  | [optional] |

### Return type

[**Key**](Key.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | checkout key |  -  |

<a id="projectUsernameProjectEnvvarGet"></a>
# **projectUsernameProjectEnvvarGet**
> List&lt;Envvar&gt; projectUsernameProjectEnvvarGet(username, project)



Lists the environment variables for :project 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    try {
      List<Envvar> result = apiInstance.projectUsernameProjectEnvvarGet(username, project);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectEnvvarGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |

### Return type

[**List&lt;Envvar&gt;**](Envvar.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | XXX |  -  |

<a id="projectUsernameProjectEnvvarNameDelete"></a>
# **projectUsernameProjectEnvvarNameDelete**
> ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response projectUsernameProjectEnvvarNameDelete(username, project, name)



Deletes the environment variable named &#39;:name&#39; 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String name = "name_example"; // String | XXXXXXXXXX 
    try {
      ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response result = apiInstance.projectUsernameProjectEnvvarNameDelete(username, project, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectEnvvarNameDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **name** | **String**| XXXXXXXXXX  | |

### Return type

[**ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response**](ProjectUsernameProjectCheckoutKeyFingerprintDelete200Response.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Deletes the environment variable named &#39;:name&#39;  |  -  |

<a id="projectUsernameProjectEnvvarNameGet"></a>
# **projectUsernameProjectEnvvarNameGet**
> Envvar projectUsernameProjectEnvvarNameGet(username, project, name)



Gets the hidden value of environment variable :name 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String name = "name_example"; // String | XXXXXXXXXX 
    try {
      Envvar result = apiInstance.projectUsernameProjectEnvvarNameGet(username, project, name);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectEnvvarNameGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **name** | **String**| XXXXXXXXXX  | |

### Return type

[**Envvar**](Envvar.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | XXX |  -  |

<a id="projectUsernameProjectEnvvarPost"></a>
# **projectUsernameProjectEnvvarPost**
> Envvar projectUsernameProjectEnvvarPost(username, project)



Creates a new environment variable 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    try {
      Envvar result = apiInstance.projectUsernameProjectEnvvarPost(username, project);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectEnvvarPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |

### Return type

[**Envvar**](Envvar.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | XXX |  -  |

<a id="projectUsernameProjectGet"></a>
# **projectUsernameProjectGet**
> List&lt;Build&gt; projectUsernameProjectGet(username, project, limit, offset, filter)



Build summary for each of the last 30 builds for a single git repo. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    Integer limit = 30; // Integer | The number of builds to return. Maximum 100, defaults to 30. 
    Integer offset = 0; // Integer | The API returns builds starting from this offset, defaults to 0. 
    String filter = "completed"; // String | Restricts which builds are returned. Set to \"completed\", \"successful\", \"failed\", \"running\", or defaults to no filter. 
    try {
      List<Build> result = apiInstance.projectUsernameProjectGet(username, project, limit, offset, filter);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **limit** | **Integer**| The number of builds to return. Maximum 100, defaults to 30.  | [optional] [default to 30] |
| **offset** | **Integer**| The API returns builds starting from this offset, defaults to 0.  | [optional] [default to 0] |
| **filter** | **String**| Restricts which builds are returned. Set to \&quot;completed\&quot;, \&quot;successful\&quot;, \&quot;failed\&quot;, \&quot;running\&quot;, or defaults to no filter.  | [optional] [enum: completed, successful, failed, running] |

### Return type

[**List&lt;Build&gt;**](Build.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Build summary for each of the last 30 builds |  -  |

<a id="projectUsernameProjectPost"></a>
# **projectUsernameProjectPost**
> BuildSummary projectUsernameProjectPost(username, project, projectUsernameProjectPostRequest)



Triggers a new build, returns a summary of the build. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    ProjectUsernameProjectPostRequest projectUsernameProjectPostRequest = new ProjectUsernameProjectPostRequest(); // ProjectUsernameProjectPostRequest | 
    try {
      BuildSummary result = apiInstance.projectUsernameProjectPost(username, project, projectUsernameProjectPostRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **projectUsernameProjectPostRequest** | [**ProjectUsernameProjectPostRequest**](ProjectUsernameProjectPostRequest.md)|  | [optional] |

### Return type

[**BuildSummary**](BuildSummary.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | returns a summary of the build |  -  |

<a id="projectUsernameProjectSshKeyPost"></a>
# **projectUsernameProjectSshKeyPost**
> ProjectUsernameProjectSshKeyPostDefaultResponse projectUsernameProjectSshKeyPost(username, project, contentType, projectUsernameProjectSshKeyPostRequest)



Create an ssh key used to access external systems that require SSH key-based authentication 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String contentType = "application/json"; // String | 
    ProjectUsernameProjectSshKeyPostRequest projectUsernameProjectSshKeyPostRequest = new ProjectUsernameProjectSshKeyPostRequest(); // ProjectUsernameProjectSshKeyPostRequest | 
    try {
      ProjectUsernameProjectSshKeyPostDefaultResponse result = apiInstance.projectUsernameProjectSshKeyPost(username, project, contentType, projectUsernameProjectSshKeyPostRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectSshKeyPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **contentType** | **String**|  | [enum: application/json] |
| **projectUsernameProjectSshKeyPostRequest** | [**ProjectUsernameProjectSshKeyPostRequest**](ProjectUsernameProjectSshKeyPostRequest.md)|  | |

### Return type

[**ProjectUsernameProjectSshKeyPostDefaultResponse**](ProjectUsernameProjectSshKeyPostDefaultResponse.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **0** | no response expected |  -  |

<a id="projectUsernameProjectTreeBranchPost"></a>
# **projectUsernameProjectTreeBranchPost**
> Build projectUsernameProjectTreeBranchPost(username, project, branch, projectUsernameProjectTreeBranchPostRequest)



Triggers a new build, returns a summary of the build. Optional build parameters can be set using an experimental API.  Note: For more about build parameters, read about [using parameterized builds](https://circleci.com/docs/parameterized-builds/) 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String username = "username_example"; // String | XXXXXXXXX 
    String project = "project_example"; // String | XXXXXXXXX 
    String branch = "branch_example"; // String | The branch name should be url-encoded. 
    ProjectUsernameProjectTreeBranchPostRequest projectUsernameProjectTreeBranchPostRequest = new ProjectUsernameProjectTreeBranchPostRequest(); // ProjectUsernameProjectTreeBranchPostRequest | 
    try {
      Build result = apiInstance.projectUsernameProjectTreeBranchPost(username, project, branch, projectUsernameProjectTreeBranchPostRequest);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectUsernameProjectTreeBranchPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| XXXXXXXXX  | |
| **project** | **String**| XXXXXXXXX  | |
| **branch** | **String**| The branch name should be url-encoded.  | |
| **projectUsernameProjectTreeBranchPostRequest** | [**ProjectUsernameProjectTreeBranchPostRequest**](ProjectUsernameProjectTreeBranchPostRequest.md)|  | [optional] |

### Return type

[**Build**](Build.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | returns a summary of the build |  * Location -  <br>  |

<a id="projectsGet"></a>
# **projectsGet**
> List&lt;Project&gt; projectsGet()



List of all the projects you&#39;re following on CircleCI, with build information organized by branch. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      List<Project> result = apiInstance.projectsGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#projectsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List&lt;Project&gt;**](Project.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List of all the projects you&#39;re following on CircleCI  |  -  |

<a id="recentBuildsGet"></a>
# **recentBuildsGet**
> List&lt;Build&gt; recentBuildsGet(limit, offset)



Build summary for each of the last 30 recent builds, ordered by build_num. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    Integer limit = 30; // Integer | The number of builds to return. Maximum 100, defaults to 30. 
    Integer offset = 0; // Integer | The API returns builds starting from this offset, defaults to 0. 
    try {
      List<Build> result = apiInstance.recentBuildsGet(limit, offset);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#recentBuildsGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **limit** | **Integer**| The number of builds to return. Maximum 100, defaults to 30.  | [optional] [default to 30] |
| **offset** | **Integer**| The API returns builds starting from this offset, defaults to 0.  | [optional] [default to 0] |

### Return type

[**List&lt;Build&gt;**](Build.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Build summary for each of the last 30 recent builds |  -  |

<a id="userHerokuKeyPost"></a>
# **userHerokuKeyPost**
> userHerokuKeyPost()



Adds your Heroku API key to CircleCI, takes apikey as form param name. 

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://circleci.com/api/v1");
    
    // Configure API key authorization: apikey
    ApiKeyAuth apikey = (ApiKeyAuth) defaultClient.getAuthentication("apikey");
    apikey.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      apiInstance.userHerokuKeyPost();
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#userHerokuKeyPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **403** | Your Heroku API key is invalid.  |  -  |

