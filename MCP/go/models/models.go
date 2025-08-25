package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Aws represents the Aws schema from the OpenAPI specification
type Aws struct {
	Keypair string `json:"keypair,omitempty"`
}

// Tests represents the Tests schema from the OpenAPI specification
type Tests struct {
	Tests []Teststestsinner `json:"tests,omitempty"`
}

// Projectusernameprojectsshkeypostrequest represents the Projectusernameprojectsshkeypostrequest schema from the OpenAPI specification
type Projectusernameprojectsshkeypostrequest struct {
	Hostname string `json:"hostname,omitempty"`
	Private_key string `json:"private_key,omitempty"`
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Node_index int `json:"node_index,omitempty"`
	Path string `json:"path,omitempty"`
	Pretty_path string `json:"pretty_path,omitempty"`
	Url string `json:"url,omitempty"`
}

// Teststestsinner represents the Teststestsinner schema from the OpenAPI specification
type Teststestsinner struct {
	Source string `json:"source,omitempty"`
	Classname string `json:"classname,omitempty"`
	File string `json:"file,omitempty"`
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Result string `json:"result,omitempty"`
	Run_time float64 `json:"run_time,omitempty"`
}

// PreviousBuild represents the PreviousBuild schema from the OpenAPI specification
type PreviousBuild struct {
	Build_num int `json:"build_num,omitempty"`
	Build_time_millis int `json:"build_time_millis,omitempty"`
	Status string `json:"status,omitempty"`
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Trial_end string `json:"trial_end,omitempty"`
	Parallelism int `json:"parallelism,omitempty"`
	Github_oauth_scopes []string `json:"github_oauth_scopes,omitempty"`
	In_beta_program bool `json:"in_beta_program,omitempty"`
	Sign_in_count int `json:"sign_in_count,omitempty"`
	Bitbucket int `json:"bitbucket,omitempty"`
	Containers int `json:"containers,omitempty"`
	Analytics_id string `json:"analytics_id,omitempty"`
	Enrolled_betas []string `json:"enrolled_betas,omitempty"`
	Heroku_api_key string `json:"heroku_api_key,omitempty"`
	Basic_email_prefs string `json:"basic_email_prefs,omitempty"`
	Days_left_in_trial int `json:"days_left_in_trial,omitempty"`
	Name string `json:"name,omitempty"`
	Plan string `json:"plan,omitempty"`
	Admin bool `json:"admin,omitempty"`
	Bitbucket_authorized bool `json:"bitbucket_authorized,omitempty"`
	Github_id int `json:"github_id,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Login string `json:"login,omitempty"`
	Projects map[string]interface{} `json:"projects,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	All_emails []string `json:"all_emails,omitempty"`
	Gravatar_id int `json:"gravatar_id,omitempty"`
	Organization_prefs map[string]interface{} `json:"organization_prefs,omitempty"`
	Selected_email string `json:"selected_email,omitempty"`
	Pusher_id string `json:"pusher_id,omitempty"`
	Dev_admin bool `json:"dev_admin,omitempty"`
}

// CommitDetail represents the CommitDetail schema from the OpenAPI specification
type CommitDetail struct {
	Body string `json:"body,omitempty"`
	Commit string `json:"commit,omitempty"`
	Committer_email string `json:"committer_email,omitempty"`
	Committer_login string `json:"committer_login,omitempty"`
	Author_name string `json:"author_name,omitempty"`
	Commit_url string `json:"commit_url,omitempty"`
	Subject string `json:"subject,omitempty"`
	Committer_date string `json:"committer_date,omitempty"`
	Committer_name string `json:"committer_name,omitempty"`
	Author_date string `json:"author_date,omitempty"`
	Author_email string `json:"author_email,omitempty"`
	Author_login string `json:"author_login,omitempty"`
}

// Projectusernameprojectpostrequest represents the Projectusernameprojectpostrequest schema from the OpenAPI specification
type Projectusernameprojectpostrequest struct {
	Build_parameters map[string]interface{} `json:"build_parameters,omitempty"` // Additional environment variables to inject into the build environment. A map of names to values.
	Parallel string `json:"parallel,omitempty"` // The number of containers to use to run the build. Default is null and the project default is used.
	Revision string `json:"revision,omitempty"` // The specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.
	Tag string `json:"tag,omitempty"` // The tag to build. Default is null. Cannot be used with revision parameter.
}

// Projectusernameprojectsshkeypostdefaultresponse represents the Projectusernameprojectsshkeypostdefaultresponse schema from the OpenAPI specification
type Projectusernameprojectsshkeypostdefaultresponse struct {
	Message string `json:"message,omitempty"`
}

// Projectfeatureflags represents the Projectfeatureflags schema from the OpenAPI specification
type Projectfeatureflags struct {
	Junit bool `json:"junit,omitempty"`
	Oss bool `json:"oss,omitempty"`
	Osx bool `json:"osx,omitempty"`
	Set_github_status bool `json:"set-github-status,omitempty"`
	Trusty_beta bool `json:"trusty-beta,omitempty"`
	Build_fork_prs bool `json:"build-fork-prs,omitempty"`
	Fleet bool `json:"fleet,omitempty"`
}

// BuildSummary represents the BuildSummary schema from the OpenAPI specification
type BuildSummary struct {
	Vcs_revision string `json:"vcs_revision,omitempty"`
	Added_at string `json:"added_at,omitempty"`
	Build_num int `json:"build_num,omitempty"`
	Outcome string `json:"outcome,omitempty"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Status string `json:"status,omitempty"`
}

// Envvar represents the Envvar schema from the OpenAPI specification
type Envvar struct {
	Value string `json:"value,omitempty"`
	Name string `json:"name,omitempty"`
}

// Projectusernameprojectcheckoutkeyfingerprintdelete200response represents the Projectusernameprojectcheckoutkeyfingerprintdelete200response schema from the OpenAPI specification
type Projectusernameprojectcheckoutkeyfingerprintdelete200response struct {
	Message string `json:"message,omitempty"`
}

// Build represents the Build schema from the OpenAPI specification
type Build struct {
	Branch string `json:"branch,omitempty"`
	Previous PreviousBuild `json:"previous,omitempty"` // previous build
	Queued_at string `json:"queued_at,omitempty"` // time build was queued
	Reponame string `json:"reponame,omitempty"`
	Subject string `json:"subject,omitempty"`
	Lifecycle string `json:"lifecycle,omitempty"`
	Username string `json:"username,omitempty"`
	Vcs_url string `json:"vcs_url,omitempty"`
	Start_time string `json:"start_time,omitempty"` // time build started
	Body string `json:"body,omitempty"` // commit message body
	Dont_build string `json:"dont_build,omitempty"` // reason why we didn't build, if we didn't build
	Retry_of int `json:"retry_of,omitempty"` // build_num of the build this is a retry of
	Stop_time string `json:"stop_time,omitempty"` // time build finished
	Committer_email string `json:"committer_email,omitempty"`
	Why string `json:"why,omitempty"` // short string explaining the reason we built
	Build_time_millis int `json:"build_time_millis,omitempty"`
	Build_url string `json:"build_url,omitempty"`
	Committer_name string `json:"committer_name,omitempty"`
}

// Projectusernameprojecttreebranchpostrequest represents the Projectusernameprojecttreebranchpostrequest schema from the OpenAPI specification
type Projectusernameprojecttreebranchpostrequest struct {
	Parallel string `json:"parallel,omitempty"` // The number of containers to use to run the build. Default is null and the project default is used.
	Revision string `json:"revision,omitempty"` // The specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.
	Build_parameters map[string]interface{} `json:"build_parameters,omitempty"` // Additional environment variables to inject into the build environment. A map of names to values.
}

// BuildParameters represents the BuildParameters schema from the OpenAPI specification
type BuildParameters struct {
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Slack_webhook_url string `json:"slack_webhook_url,omitempty"`
	Branches map[string]interface{} `json:"branches,omitempty"`
	Oss bool `json:"oss,omitempty"`
	Reponame string `json:"reponame,omitempty"`
	Slack_channel string `json:"slack_channel,omitempty"`
	Slack_subdomain string `json:"slack_subdomain,omitempty"`
	Campfire_token string `json:"campfire_token,omitempty"`
	Compile string `json:"compile,omitempty"`
	Irc_keyword string `json:"irc_keyword,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Hipchat_notify_prefs string `json:"hipchat_notify_prefs,omitempty"`
	Irc_password string `json:"irc_password,omitempty"`
	Dependencies string `json:"dependencies,omitempty"`
	Hipchat_api_token string `json:"hipchat_api_token,omitempty"`
	Heroku_deploy_user string `json:"heroku_deploy_user,omitempty"`
	Campfire_notify_prefs string `json:"campfire_notify_prefs,omitempty"`
	Feature_flags Projectfeatureflags `json:"feature_flags,omitempty"`
	Irc_server string `json:"irc_server,omitempty"`
	Irc_username string `json:"irc_username,omitempty"`
	Language string `json:"language,omitempty"`
	Parallel int `json:"parallel,omitempty"`
	Username string `json:"username,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Vcs_type string `json:"vcs_type,omitempty"`
	Campfire_subdomain string `json:"campfire_subdomain,omitempty"`
	Setup string `json:"setup,omitempty"`
	Vcs_url string `json:"vcs_url,omitempty"`
	Irc_channel string `json:"irc_channel,omitempty"`
	Campfire_room string `json:"campfire_room,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Hipchat_notify string `json:"hipchat_notify,omitempty"`
	Hipchat_room string `json:"hipchat_room,omitempty"`
	Irc_notify_prefs string `json:"irc_notify_prefs,omitempty"`
	Slack_api_token string `json:"slack_api_token,omitempty"`
	Has_usable_key bool `json:"has_usable_key,omitempty"`
	Slack_channel_override string `json:"slack_channel_override,omitempty"`
	Aws Aws `json:"aws,omitempty"`
	Extra string `json:"extra,omitempty"`
	Slack_notify_prefs string `json:"slack_notify_prefs,omitempty"`
	Test string `json:"test,omitempty"`
	Followed bool `json:"followed,omitempty"`
	Flowdock_api_token string `json:"flowdock_api_token,omitempty"`
}

// BuildDetail represents the BuildDetail schema from the OpenAPI specification
type BuildDetail struct {
	Timedout bool `json:"timedout,omitempty"`
	Node interface{} `json:"node,omitempty"`
	Usage_queued_at string `json:"usage_queued_at,omitempty"`
	User User `json:"user,omitempty"`
	All_commit_details []CommitDetail `json:"all_commit_details,omitempty"`
	Compare string `json:"compare,omitempty"`
	Job_name string `json:"job_name,omitempty"`
	Previous_successful_build PreviousBuild `json:"previous_successful_build,omitempty"` // previous build
	Retries bool `json:"retries,omitempty"`
	Ssh_enabled bool `json:"ssh_enabled,omitempty"`
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Preferred bool `json:"preferred,omitempty"`
	Public_key string `json:"public_key,omitempty"`
	Time string `json:"time,omitempty"` // when the key was issued
	TypeField string `json:"type,omitempty"` // can be "deploy-key" or "github-user-key"
	Fingerprint string `json:"fingerprint,omitempty"`
}

// Projectusernameprojectbuildcachedelete200response represents the Projectusernameprojectbuildcachedelete200response schema from the OpenAPI specification
type Projectusernameprojectbuildcachedelete200response struct {
	Status string `json:"status,omitempty"`
}
