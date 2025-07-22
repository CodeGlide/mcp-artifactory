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

func Views_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["view_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("view_id=%v", val))
		}
		if val, ok := args["external_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("external_id=%v", val))
		}
		if val, ok := args["view"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("view=%v", val))
		}
		if val, ok := args["hash"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hash=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/views.update%s", cfg.BaseURL, queryString)
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

func CreateViews_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_views.update",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Required(), mcp.Description("Authentication token. Requires scope: `none`")),
		mcp.WithString("view_id", mcp.Description("A unique identifier of the view to be updated. Either `view_id` or `external_id` is required.")),
		mcp.WithString("external_id", mcp.Description("A unique identifier of the view set by the developer. Must be unique for all views on a team. Max length of 255 characters. Either `view_id` or `external_id` is required.")),
		mcp.WithString("view", mcp.Description("A [view object](/reference/surfaces/views). This must be a JSON-encoded string.")),
		mcp.WithString("hash", mcp.Description("A string that represents view state to protect against possible race conditions.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Views_updateHandler(cfg),
	}
}
