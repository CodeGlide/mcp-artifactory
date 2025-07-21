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

func Gists_delete_commentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		gist_idVal, ok := args["gist_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: gist_id"), nil
		}
		gist_id, ok := gist_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: gist_id"), nil
		}
		comment_idVal, ok := args["comment_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: comment_id"), nil
		}
		comment_id, ok := comment_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: comment_id"), nil
		}
		url := fmt.Sprintf("%s/gists/%s/comments/%s", cfg.BaseURL, gist_id, comment_id)
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

func CreateGists_delete_commentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_gists_gist_id_comments_comment_id",
		mcp.WithDescription("Delete a gist comment"),
		mcp.WithString("gist_id", mcp.Required(), mcp.Description("The unique identifier of the gist.")),
		mcp.WithNumber("comment_id", mcp.Required(), mcp.Description("The unique identifier of the comment.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Gists_delete_commentHandler(cfg),
	}
}
