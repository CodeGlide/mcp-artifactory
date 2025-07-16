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

func Teams_delete_discussion_comment_in_orgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		discussion_numberVal, ok := args["discussion_number"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: discussion_number"), nil
		}
		discussion_number, ok := discussion_numberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: discussion_number"), nil
		}
		comment_numberVal, ok := args["comment_number"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: comment_number"), nil
		}
		comment_number, ok := comment_numberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: comment_number"), nil
		}
		url := fmt.Sprintf("%s/orgs/%s/teams/%s/discussions/%s/comments/%s", cfg.BaseURL, org, team_slug, discussion_number, comment_number)
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

func CreateTeams_delete_discussion_comment_in_orgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_orgs_org_teams_team_slug_discussions_discussion_number_comments_comment_number",
		mcp.WithDescription("Delete a discussion comment"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("team_slug", mcp.Required(), mcp.Description("The slug of the team name.")),
		mcp.WithNumber("discussion_number", mcp.Required(), mcp.Description("The number that identifies the discussion.")),
		mcp.WithNumber("comment_number", mcp.Required(), mcp.Description("The number that identifies the comment.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Teams_delete_discussion_comment_in_orgHandler(cfg),
	}
}
