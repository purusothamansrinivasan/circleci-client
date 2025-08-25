package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/circleci-rest-api/mcp-server/config"
	"github.com/circleci-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_project_username_projectHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter=%v", val))
		}
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("circle-token=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/project/%s/%s%s", cfg.BaseURL, username, project, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result []Build
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

func CreateGet_project_username_projectTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_project_username_project",
		mcp.WithDescription("Build summary for each of the last 30 builds for a single git repo.
"),
		mcp.WithString("username", mcp.Required(), mcp.Description("XXXXXXXXX\n")),
		mcp.WithString("project", mcp.Required(), mcp.Description("XXXXXXXXX\n")),
		mcp.WithNumber("limit", mcp.Description("The number of builds to return. Maximum 100, defaults to 30.\n")),
		mcp.WithNumber("offset", mcp.Description("The API returns builds starting from this offset, defaults to 0.\n")),
		mcp.WithString("filter", mcp.Description("Restricts which builds are returned.\nSet to \"completed\", \"successful\", \"failed\", \"running\", or defaults to no filter.\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_project_username_projectHandler(cfg),
	}
}
