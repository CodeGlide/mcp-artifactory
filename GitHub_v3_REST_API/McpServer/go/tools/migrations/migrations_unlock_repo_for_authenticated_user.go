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

func Migrations_unlock_repo_for_authenticated_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		migration_idVal, ok := args["migration_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: migration_id"), nil
		}
		migration_id, ok := migration_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: migration_id"), nil
		}
		repo_nameVal, ok := args["repo_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo_name"), nil
		}
		repo_name, ok := repo_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo_name"), nil
		}
		url := fmt.Sprintf("%s/user/migrations/%s/repos/%s/lock", cfg.BaseURL, migration_id, repo_name)
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

func CreateMigrations_unlock_repo_for_authenticated_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_user_migrations_migration_id_repos_repo_name_lock",
		mcp.WithDescription("Unlock a user repository"),
		mcp.WithNumber("migration_id", mcp.Required(), mcp.Description("The unique identifier of the migration.")),
		mcp.WithString("repo_name", mcp.Required(), mcp.Description("repo_name parameter")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Migrations_unlock_repo_for_authenticated_userHandler(cfg),
	}
}
