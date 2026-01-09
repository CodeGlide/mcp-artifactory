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

func Chat_scheduledmessages_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["channel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("channel=%v", val))
		}
		if val, ok := args["latest"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("latest=%v", val))
		}
		if val, ok := args["oldest"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("oldest=%v", val))
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
		url := fmt.Sprintf("%s/chat.scheduledMessages.list%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["token"]; ok {
			req.Header.Set("token", fmt.Sprintf("%v", val))
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

func CreateChat_scheduledmessages_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_chat.scheduledMessages.list",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Description("Authentication token. Requires scope: `none`")),
		mcp.WithString("channel", mcp.Description("The channel of the scheduled messages")),
		mcp.WithString("latest", mcp.Description("A UNIX timestamp of the latest value in the time range")),
		mcp.WithString("oldest", mcp.Description("A UNIX timestamp of the oldest value in the time range")),
		mcp.WithNumber("limit", mcp.Description("Maximum number of original entries to return.")),
		mcp.WithString("cursor", mcp.Description("For pagination purposes, this is the `cursor` value returned from a previous call to `chat.scheduledmessages.list` indicating where you want to start this call from.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Chat_scheduledmessages_listHandler(cfg),
	}
}
