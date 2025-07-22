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

func Oauth_tokenHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_id=%v", val))
		}
		if val, ok := args["client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_secret=%v", val))
		}
		if val, ok := args["code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("code=%v", val))
		}
		if val, ok := args["redirect_uri"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("redirect_uri=%v", val))
		}
		if val, ok := args["single_channel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("single_channel=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/oauth.token%s", cfg.BaseURL, queryString)
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

func CreateOauth_tokenTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_oauth.token",
		mcp.WithDescription(""),
		mcp.WithString("client_id", mcp.Description("Issued when you created your application.")),
		mcp.WithString("client_secret", mcp.Description("Issued when you created your application.")),
		mcp.WithString("code", mcp.Description("The `code` param returned via the OAuth callback.")),
		mcp.WithString("redirect_uri", mcp.Description("This must match the originally submitted URI (if one was sent).")),
		mcp.WithBoolean("single_channel", mcp.Description("Request the user to add your app only to a single channel.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Oauth_tokenHandler(cfg),
	}
}
