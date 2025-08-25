

# Key


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**fingerprint** | **String** |  |  [optional] |
|**preferred** | **Boolean** |  |  [optional] |
|**publicKey** | **String** |  |  [optional] |
|**time** | **OffsetDateTime** | when the key was issued |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | can be \&quot;deploy-key\&quot; or \&quot;github-user-key\&quot;  |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| DEPLOY_KEY | &quot;deploy-key&quot; |
| GITHUB_USER_KEY | &quot;github-user-key&quot; |



