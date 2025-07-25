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

func Reactions_delete_for_commit_commentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ownerVal, ok := args["owner"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: owner"), nil
		}
		owner, ok := ownerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: owner"), nil
		}
		repoVal, ok := args["repo"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo"), nil
		}
		repo, ok := repoVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo"), nil
		}
		comment_idVal, ok := args["comment_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: comment_id"), nil
		}
		comment_id, ok := comment_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: comment_id"), nil
		}
		reaction_idVal, ok := args["reaction_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: reaction_id"), nil
		}
		reaction_id, ok := reaction_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: reaction_id"), nil
		}
		url := fmt.Sprintf("%s/repos/%s/%s/comments/%s/reactions/%s", cfg.BaseURL, owner, repo, comment_id, reaction_id)
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

func CreateReactions_delete_for_commit_commentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_repos_owner_repo_comments_comment_id_reactions_reaction_id",
		mcp.WithDescription("Delete a commit comment reaction"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithNumber("comment_id", mcp.Required(), mcp.Description("The unique identifier of the comment.")),
		mcp.WithNumber("reaction_id", mcp.Required(), mcp.Description("The unique identifier of the reaction.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Reactions_delete_for_commit_commentHandler(cfg),
	}
}
