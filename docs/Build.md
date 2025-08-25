

# Build


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**body** | **String** | commit message body |  [optional] |
|**branch** | **String** |  |  [optional] |
|**buildTimeMillis** | **Integer** |  |  [optional] |
|**buildUrl** | **URI** |  |  [optional] |
|**committerEmail** | **String** |  |  [optional] |
|**committerName** | **String** |  |  [optional] |
|**dontBuild** | **String** | reason why we didn&#39;t build, if we didn&#39;t build |  [optional] |
|**lifecycle** | **Lifecycle** |  |  [optional] |
|**previous** | [**PreviousBuild**](PreviousBuild.md) |  |  [optional] |
|**queuedAt** | **OffsetDateTime** | time build was queued |  [optional] |
|**reponame** | **String** |  |  [optional] |
|**retryOf** | **Integer** | build_num of the build this is a retry of |  [optional] |
|**startTime** | **OffsetDateTime** | time build started |  [optional] |
|**stopTime** | **OffsetDateTime** | time build finished |  [optional] |
|**subject** | **String** |  |  [optional] |
|**username** | **String** |  |  [optional] |
|**vcsUrl** | **URI** |  |  [optional] |
|**why** | **String** | short string explaining the reason we built |  [optional] |



