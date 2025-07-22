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

func Conversations_repliesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["ts"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ts=%v", val))
		}
		if val, ok := args["latest"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("latest=%v", val))
		}
		if val, ok := args["oldest"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("oldest=%v", val))
		}
		if val, ok := args["inclusive"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("inclusive=%v", val))
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
		url := fmt.Sprintf("%s/conversations.replies%s", cfg.BaseURL, queryString)
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

func CreateConversations_repliesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_conversations.replies",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Description("Authentication token. Requires scope: `conversations:history`")),
		mcp.WithString("channel", mcp.Description("Conversation ID to fetch thread from.")),
		mcp.WithString("ts", mcp.Description("Unique identifier of a thread's parent message. `ts` must be the timestamp of an existing message with 0 or more replies. If there are no replies then just the single message referenced by `ts` will return - it is just an ordinary, unthreaded message.")),
		mcp.WithString("latest", mcp.Description("End of time range of messages to include in results.")),
		mcp.WithString("oldest", mcp.Description("Start of time range of messages to include in results.")),
		mcp.WithBoolean("inclusive", mcp.Description("Include messages with latest or oldest timestamp in results only when either timestamp is specified.")),
		mcp.WithNumber("limit", mcp.Description("The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn't been reached.")),
		mcp.WithString("cursor", mcp.Description("Paginate through collections of data by setting the `cursor` parameter to a `next_cursor` attribute returned by a previous request's `response_metadata`. Default value fetches the first \"page\" of the collection. See [pagination](/docs/pagination) for more detail.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Conversations_repliesHandler(cfg),
	}
}
