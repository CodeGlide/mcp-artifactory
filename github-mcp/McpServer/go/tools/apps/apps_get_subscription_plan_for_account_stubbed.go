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

func Apps_get_subscription_plan_for_account_stubbedHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		account_idVal, ok := args["account_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: account_id"), nil
		}
		account_id, ok := account_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: account_id"), nil
		}
		url := fmt.Sprintf("%s/marketplace_listing/stubbed/accounts/%s", cfg.BaseURL, account_id)
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

func CreateApps_get_subscription_plan_for_account_stubbedTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_marketplace_listing_stubbed_accounts_account_id",
		mcp.WithDescription("Get a subscription plan for an account (stubbed)"),
		mcp.WithNumber("account_id", mcp.Required(), mcp.Description("account_id parameter")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apps_get_subscription_plan_for_account_stubbedHandler(cfg),
	}
}
