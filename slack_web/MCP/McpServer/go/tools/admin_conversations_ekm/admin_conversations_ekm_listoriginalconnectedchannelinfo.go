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

func Admin_conversations_ekm_listoriginalconnectedchannelinfoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("token=%v", val))
		}
		if val, ok := args["channel_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("channel_ids=%v", val))
		}
		if val, ok := args["team_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("team_ids=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["cursor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cursor=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/admin.conversations.ekm.listOriginalConnectedChannelInfo%s", cfg.BaseURL, queryString)
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

func CreateAdmin_conversations_ekm_listoriginalconnectedchannelinfoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_admin.conversations.ekm.listOriginalConnectedChannelInfo",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Required(), mcp.Description("Authentication token. Requires scope: `admin.conversations:read`")),
		mcp.WithString("channel_ids", mcp.Description("A comma-separated list of channels to filter to.")),
		mcp.WithString("team_ids", mcp.Description("A comma-separated list of the workspaces to which the channels you would like returned belong.")),
		mcp.WithNumber("limit", mcp.Description("The maximum number of items to return. Must be between 1 - 1000 both inclusive.")),
		mcp.WithString("cursor", mcp.Description("Set `cursor` to `next_cursor` returned by the previous call to list items in the next page.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Admin_conversations_ekm_listoriginalconnectedchannelinfoHandler(cfg),
	}
}
