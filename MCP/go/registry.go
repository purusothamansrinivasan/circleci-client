package main

import (
	"github.com/circleci-rest-api/mcp-server/config"
	"github.com/circleci-rest-api/mcp-server/models"
	tools_project "github.com/circleci-rest-api/mcp-server/tools/project"
	tools_projects "github.com/circleci-rest-api/mcp-server/tools/projects"
	tools_recent_builds "github.com/circleci-rest-api/mcp-server/tools/recent_builds"
	tools_user "github.com/circleci-rest-api/mcp-server/tools/user"
	tools_me "github.com/circleci-rest-api/mcp-server/tools/me"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_project.CreateDelete_project_username_project_build_cacheTool(cfg),
		tools_project.CreateDelete_project_username_project_envvar_nameTool(cfg),
		tools_project.CreateGet_project_username_project_envvar_nameTool(cfg),
		tools_projects.CreateGet_projectsTool(cfg),
		tools_recent_builds.CreateGet_recent_buildsTool(cfg),
		tools_project.CreateGet_project_username_project_build_numTool(cfg),
		tools_user.CreatePost_user_heroku_keyTool(cfg),
		tools_project.CreateGet_project_username_project_envvarTool(cfg),
		tools_project.CreatePost_project_username_project_envvarTool(cfg),
		tools_project.CreatePost_project_username_project_ssh_keyTool(cfg),
		tools_project.CreateGet_project_username_project_build_num_testsTool(cfg),
		tools_project.CreateGet_project_username_project_checkout_keyTool(cfg),
		tools_project.CreatePost_project_username_project_checkout_keyTool(cfg),
		tools_project.CreateGet_project_username_project_build_num_artifactsTool(cfg),
		tools_me.CreateGet_meTool(cfg),
		tools_project.CreateGet_project_username_projectTool(cfg),
		tools_project.CreatePost_project_username_projectTool(cfg),
		tools_project.CreatePost_project_username_project_tree_branchTool(cfg),
		tools_project.CreatePost_project_username_project_build_num_retryTool(cfg),
		tools_project.CreateGet_project_username_project_checkout_key_fingerprintTool(cfg),
		tools_project.CreateDelete_project_username_project_checkout_key_fingerprintTool(cfg),
		tools_project.CreatePost_project_username_project_build_num_cancelTool(cfg),
	}
}
