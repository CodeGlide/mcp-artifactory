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

func Orgs_redeliver_webhook_deliveryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		orgVal, ok := args["org"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org"), nil
		}
		org, ok := orgVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org"), nil
		}
		hook_idVal, ok := args["hook_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: hook_id"), nil
		}
		hook_id, ok := hook_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: hook_id"), nil
		}
		delivery_idVal, ok := args["delivery_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: delivery_id"), nil
		}
		delivery_id, ok := delivery_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: delivery_id"), nil
		}
		url := fmt.Sprintf("%s/orgs/%s/hooks/%s/deliveries/%s/attempts", cfg.BaseURL, org, hook_id, delivery_id)
		req, err := http.NewRequest("POST", url, nil)
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

func CreateOrgs_redeliver_webhook_deliveryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_orgs_org_hooks_hook_id_deliveries_delivery_id_attempts",
		mcp.WithDescription("Redeliver a delivery for an organization webhook"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithNumber("hook_id", mcp.Required(), mcp.Description("The unique identifier of the hook. You can find this value in the `X-GitHub-Hook-ID` header of a webhook delivery.")),
		mcp.WithNumber("delivery_id", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_redeliver_webhook_deliveryHandler(cfg),
	}
}
