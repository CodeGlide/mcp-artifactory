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

func Usergroups_listHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["include_users"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_users=%v", val))
		}
		if val, ok := args["token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("token=%v", val))
		}
		if val, ok := args["include_count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_count=%v", val))
		}
		if val, ok := args["include_disabled"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_disabled=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/usergroups.list%s", cfg.BaseURL, queryString)
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

func CreateUsergroups_listTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_usergroups.list",
		mcp.WithDescription(""),
		mcp.WithBoolean("include_users", mcp.Description("Include the list of users for each User Group.")),
		mcp.WithString("token", mcp.Required(), mcp.Description("Authentication token. Requires scope: `usergroups:read`")),
		mcp.WithBoolean("include_count", mcp.Description("Include the number of users in each User Group.")),
		mcp.WithBoolean("include_disabled", mcp.Description("Include disabled User Groups.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Usergroups_listHandler(cfg),
	}
}
