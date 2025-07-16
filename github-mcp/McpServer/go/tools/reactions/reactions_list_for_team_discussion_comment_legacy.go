package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Reactions_list_for_team_discussion_comment_legacyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		comment_numberVal, ok := args["comment_number"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: comment_number"), nil
		}
		comment_number, ok := comment_numberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: comment_number"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["content"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("content=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/teams/%s/discussions/%s/comments/%s/reactions%s", cfg.BaseURL, team_id, discussion_number, comment_number, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result []Reaction
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

func CreateReactions_list_for_team_discussion_comment_legacyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_teams_team_id_discussions_discussion_number_comments_comment_number_reactions",
		mcp.WithDescription("List reactions for a team discussion comment (Legacy)"),
		mcp.WithNumber("team_id", mcp.Required(), mcp.Description("The unique identifier of the team.")),
		mcp.WithNumber("discussion_number", mcp.Required(), mcp.Description("The number that identifies the discussion.")),
		mcp.WithNumber("comment_number", mcp.Required(), mcp.Description("The number that identifies the comment.")),
		mcp.WithString("content", mcp.Description("Returns a single [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions). Omit this parameter to list all reactions to a team discussion comment.")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Reactions_list_for_team_discussion_comment_legacyHandler(cfg),
	}
}
