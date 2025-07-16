package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Migrations_start_for_authenticated_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
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
		url := fmt.Sprintf("%s/user/migrations", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Migration
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

func CreateMigrations_start_for_authenticated_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_user_migrations",
		mcp.WithDescription("Start a user migration"),
		mcp.WithArray("exclude", mcp.Description("Input parameter: Exclude attributes from the API response to improve performance")),
		mcp.WithBoolean("exclude_attachments", mcp.Description("Input parameter: Do not include attachments in the migration")),
		mcp.WithBoolean("exclude_owner_projects", mcp.Description("Input parameter: Indicates whether projects owned by the organization or users should be excluded.")),
		mcp.WithArray("repositories", mcp.Required(), mcp.Description("")),
		mcp.WithBoolean("exclude_git_data", mcp.Description("Input parameter: Indicates whether the repository git data should be excluded from the migration.")),
		mcp.WithBoolean("lock_repositories", mcp.Description("Input parameter: Lock the repositories being migrated at the start of the migration")),
		mcp.WithBoolean("exclude_releases", mcp.Description("Input parameter: Do not include releases in the migration")),
		mcp.WithBoolean("exclude_metadata", mcp.Description("Input parameter: Indicates whether metadata should be excluded and only git source should be included for the migration.")),
		mcp.WithBoolean("org_metadata_only", mcp.Description("Input parameter: Indicates whether this should only include organization metadata (repositories array should be empty and will ignore other flags).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Migrations_start_for_authenticated_userHandler(cfg),
	}
}
