

# ProjectUsernameProjectTreeBranchPostRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**buildParameters** | **Object** | Additional environment variables to inject into the build environment. A map of names to values.  |  [optional] |
|**parallel** | **String** | The number of containers to use to run the build. Default is null and the project default is used.  |  [optional] |
|**revision** | **String** | The specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.  |  [optional] |



