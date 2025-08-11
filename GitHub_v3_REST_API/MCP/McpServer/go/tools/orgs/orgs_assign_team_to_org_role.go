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

func Orgs_assign_team_to_org_roleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		orgVal, ok := args["org"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org"), nil
		}
		org, ok := orgVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org"), nil
		}
		team_slugVal, ok := args["team_slug"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: team_slug"), nil
		}
		team_slug, ok := team_slugVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: team_slug"), nil
		}
		role_idVal, ok := args["role_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: role_id"), nil
		}
		role_id, ok := role_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: role_id"), nil
		}
		url := fmt.Sprintf("%s/orgs/%s/organization-roles/teams/%s/%s", cfg.BaseURL, org, team_slug, role_id)
		req, err := http.NewRequest("PUT", url, nil)
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

func CreateOrgs_assign_team_to_org_roleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_orgs_org_organization-roles_teams_team_slug_role_id",
		mcp.WithDescription("Assign an organization role to a team"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("team_slug", mcp.Required(), mcp.Description("The slug of the team name.")),
		mcp.WithNumber("role_id", mcp.Required(), mcp.Description("The unique identifier of the role.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_assign_team_to_org_roleHandler(cfg),
	}
}
