package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/circleci-rest-api/mcp-server/config"
	"github.com/circleci-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_project_username_project_tree_branchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		usernameVal, ok := args["username"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: username"), nil
		}
		username, ok := usernameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: username"), nil
		}
		projectVal, ok := args["project"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: project"), nil
		}
		project, ok := projectVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: project"), nil
		}
		branchVal, ok := args["branch"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: branch"), nil
		}
		branch, ok := branchVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: branch"), nil
		}
		queryParams := make([]string, 0)
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("circle-token=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Projectusernameprojecttreebranchpostrequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/project/%s/%s/tree/%s%s", cfg.BaseURL, username, project, branch, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			// API key already added to query string
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Build
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePost_project_username_project_tree_branchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_project_username_project_tree_branch",
		mcp.WithDescription("Triggers a new build, returns a summary of the build.
Optional build parameters can be set using an experimental API.

Note: For more about build parameters, read about [using parameterized builds](https://circleci.com/docs/parameterized-builds/)
"),
		mcp.WithString("username", mcp.Required(), mcp.Description("XXXXXXXXX\n")),
		mcp.WithString("project", mcp.Required(), mcp.Description("XXXXXXXXX\n")),
		mcp.WithString("branch", mcp.Required(), mcp.Description("The branch name should be url-encoded.\n")),
		mcp.WithString("parallel", mcp.Description("Input parameter: The number of containers to use to run the build. Default is null and the project default is used.\n")),
		mcp.WithString("revision", mcp.Description("Input parameter: The specific revision to build.\nDefault is null and the head of the branch is used. Cannot be used with tag parameter.\n")),
		mcp.WithObject("build_parameters", mcp.Description("Input parameter: Additional environment variables to inject into the build environment. A map of names to values.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_project_username_project_tree_branchHandler(cfg),
	}
}
