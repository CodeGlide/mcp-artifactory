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

func Files_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("token=%v", val))
		}
		if val, ok := args["user"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("user=%v", val))
		}
		if val, ok := args["channel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("channel=%v", val))
		}
		if val, ok := args["ts_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ts_from=%v", val))
		}
		if val, ok := args["ts_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ts_to=%v", val))
		}
		if val, ok := args["types"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("types=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["show_files_hidden_by_limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("show_files_hidden_by_limit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/files.list%s", cfg.BaseURL, queryString)
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

func CreateFiles_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_files.list",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Description("Authentication token. Requires scope: `files:read`")),
		mcp.WithString("user", mcp.Description("Filter files created by a single user.")),
		mcp.WithString("channel", mcp.Description("Filter files appearing in a specific channel, indicated by its ID.")),
		mcp.WithString("ts_from", mcp.Description("Filter files created after this timestamp (inclusive).")),
		mcp.WithString("ts_to", mcp.Description("Filter files created before this timestamp (inclusive).")),
		mcp.WithString("types", mcp.Description("Filter files by type ([see below](#file_types)). You can pass multiple values in the types argument, like `types=spaces,snippets`.The default value is `all`, which does not filter the list.")),
		mcp.WithString("count", mcp.Description("")),
		mcp.WithString("page", mcp.Description("")),
		mcp.WithBoolean("show_files_hidden_by_limit", mcp.Description("Show truncated file info for files hidden due to being too old, and the team who owns the file being over the file limit.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Files_listHandler(cfg),
	}
}
