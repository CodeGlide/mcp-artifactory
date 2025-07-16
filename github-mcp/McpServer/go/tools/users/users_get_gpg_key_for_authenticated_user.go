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

func Users_get_gpg_key_for_authenticated_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		gpg_key_idVal, ok := args["gpg_key_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: gpg_key_id"), nil
		}
		gpg_key_id, ok := gpg_key_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: gpg_key_id"), nil
		}
		url := fmt.Sprintf("%s/user/gpg_keys/%s", cfg.BaseURL, gpg_key_id)
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

func CreateUsers_get_gpg_key_for_authenticated_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_user_gpg_keys_gpg_key_id",
		mcp.WithDescription("Get a GPG key for the authenticated user"),
		mcp.WithNumber("gpg_key_id", mcp.Required(), mcp.Description("The unique identifier of the GPG key.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Users_get_gpg_key_for_authenticated_userHandler(cfg),
	}
}
