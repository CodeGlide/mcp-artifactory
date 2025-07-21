package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Apps_remove_repo_from_installation_for_authenticated_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		installation_idVal, ok := args["installation_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: installation_id"), nil
		}
		installation_id, ok := installation_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: installation_id"), nil
		}
		repository_idVal, ok := args["repository_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repository_id"), nil
		}
		repository_id, ok := repository_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repository_id"), nil
		}
		url := fmt.Sprintf("%s/user/installations/%s/repositories/%s", cfg.BaseURL, installation_id, repository_id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
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
		var result map[string]interface{}
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

func CreateApps_remove_repo_from_installation_for_authenticated_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_user_installations_installation_id_repositories_repository_id",
		mcp.WithDescription("Remove a repository from an app installation"),
		mcp.WithNumber("installation_id", mcp.Required(), mcp.Description("The unique identifier of the installation.")),
		mcp.WithNumber("repository_id", mcp.Required(), mcp.Description("The unique identifier of the repository.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apps_remove_repo_from_installation_for_authenticated_userHandler(cfg),
	}
}
