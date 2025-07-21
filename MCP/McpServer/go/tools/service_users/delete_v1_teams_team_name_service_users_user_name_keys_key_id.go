package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/okta-privileged-access/mcp-server/config"
	"github.com/okta-privileged-access/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Delete_v1_teams_team_name_service_users_user_name_keys_key_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		team_nameVal, ok := args["team_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: team_name"), nil
		}
		team_name, ok := team_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: team_name"), nil
		}
		user_nameVal, ok := args["user_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: user_name"), nil
		}
		user_name, ok := user_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: user_name"), nil
		}
		key_idVal, ok := args["key_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: key_id"), nil
		}
		key_id, ok := key_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: key_id"), nil
		}
		url := fmt.Sprintf("%s/v1/teams/%s/service_users/%s/keys/%s", cfg.BaseURL, team_name, user_name, key_id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
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

func CreateDelete_v1_teams_team_name_service_users_user_name_keys_key_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_v1_teams_team_name_service_users_user_name_keys_key_id",
		mcp.WithDescription("Delete an API key"),
		mcp.WithString("team_name", mcp.Required(), mcp.Description("(Required) The name of your Team")),
		mcp.WithString("user_name", mcp.Required(), mcp.Description("(Required) The username for an existing User")),
		mcp.WithString("key_id", mcp.Required(), mcp.Description("(Required) The UUID of a Service User key")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_v1_teams_team_name_service_users_user_name_keys_key_idHandler(cfg),
	}
}
