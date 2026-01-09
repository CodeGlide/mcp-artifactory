package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/slack-web-api/mcp-server/config"
	"github.com/slack-web-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Conversations_infoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("token=%v", val))
		}
		if val, ok := args["channel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("channel=%v", val))
		}
		if val, ok := args["include_locale"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_locale=%v", val))
		}
		if val, ok := args["include_num_members"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_num_members=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/conversations.info%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
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

func CreateConversations_infoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_conversations.info",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Description("Authentication token. Requires scope: `conversations:read`")),
		mcp.WithString("channel", mcp.Description("Conversation ID to learn more about")),
		mcp.WithBoolean("include_locale", mcp.Description("Set this to `true` to receive the locale for this conversation. Defaults to `false`")),
		mcp.WithBoolean("include_num_members", mcp.Description("Set to `true` to include the member count for the specified conversation. Defaults to `false`")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Conversations_infoHandler(cfg),
	}
}
