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

func Delete_v1_teams_team_name_resource_groups_resource_group_id_projects_project_id_secret_folders_secret_folder_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		resource_group_idVal, ok := args["resource_group_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resource_group_id"), nil
		}
		resource_group_id, ok := resource_group_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resource_group_id"), nil
		}
		project_idVal, ok := args["project_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: project_id"), nil
		}
		project_id, ok := project_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: project_id"), nil
		}
		secret_folder_idVal, ok := args["secret_folder_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: secret_folder_id"), nil
		}
		secret_folder_id, ok := secret_folder_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: secret_folder_id"), nil
		}
		url := fmt.Sprintf("%s/v1/teams/%s/resource_groups/%s/projects/%s/secret_folders/%s", cfg.BaseURL, team_name, resource_group_id, project_id, secret_folder_id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["Accept"]; ok {
			req.Header.Set("Accept", fmt.Sprintf("%v", val))
		}

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

func CreateDelete_v1_teams_team_name_resource_groups_resource_group_id_projects_project_id_secret_folders_secret_folder_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_v1_teams_team_name_resource_groups_resource_group_id_projects_project_id_secret_folders_secret_folder_id",
		mcp.WithDescription("Delete a Secret Folder"),
		mcp.WithString("Accept", mcp.Description("")),
		mcp.WithString("team_name", mcp.Required(), mcp.Description("(Required) The name of your Team")),
		mcp.WithString("resource_group_id", mcp.Required(), mcp.Description("(Required) The UUID of a Resource Group")),
		mcp.WithString("project_id", mcp.Required(), mcp.Description("(Required) The UUID of a Project")),
		mcp.WithString("secret_folder_id", mcp.Required(), mcp.Description("(Required) The UUID of a Secret Folder")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_v1_teams_team_name_resource_groups_resource_group_id_projects_project_id_secret_folders_secret_folder_idHandler(cfg),
	}
}
