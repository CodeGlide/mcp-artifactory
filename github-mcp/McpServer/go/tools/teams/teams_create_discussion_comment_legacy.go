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

func Teams_create_discussion_comment_legacyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		team_idVal, ok := args["team_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: team_id"), nil
		}
		team_id, ok := team_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: team_id"), nil
		}
		discussion_numberVal, ok := args["discussion_number"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: discussion_number"), nil
		}
		discussion_number, ok := discussion_numberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: discussion_number"), nil
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
		url := fmt.Sprintf("%s/teams/%s/discussions/%s/comments", cfg.BaseURL, team_id, discussion_number)
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
		var result models.GeneratedType
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

func CreateTeams_create_discussion_comment_legacyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_teams_team_id_discussions_discussion_number_comments",
		mcp.WithDescription("Create a discussion comment (Legacy)"),
		mcp.WithNumber("team_id", mcp.Required(), mcp.Description("The unique identifier of the team.")),
		mcp.WithNumber("discussion_number", mcp.Required(), mcp.Description("The number that identifies the discussion.")),
		mcp.WithString("body", mcp.Required(), mcp.Description("Input parameter: The discussion comment's body text.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Teams_create_discussion_comment_legacyHandler(cfg),
	}
}
