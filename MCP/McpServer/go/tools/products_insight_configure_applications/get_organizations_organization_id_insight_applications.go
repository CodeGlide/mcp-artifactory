package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/meraki-dashboard-api-v1-9-0/mcp-server/config"
	"github.com/meraki-dashboard-api-v1-9-0/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_organizations_organization_id_insight_applicationsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		organization_idVal, ok := args["organization_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: organization_id"), nil
		}
		organization_id, ok := organization_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: organization_id"), nil
		}
		url := fmt.Sprintf("%s/organizations/%s/insight/applications", cfg.BaseURL, organization_id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Cisco-Meraki-API-Key"]; ok {
			req.Header.Set("X-Cisco-Meraki-API-Key", fmt.Sprintf("%v", val))
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

func CreateGet_organizations_organization_id_insight_applicationsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_organizations_organization_id_insight_applications",
		mcp.WithDescription("List all Insight tracked applications"),
		mcp.WithString("X-Cisco-Meraki-API-Key", mcp.Description("")),
		mcp.WithString("organization_id", mcp.Required(), mcp.Description("(Required) ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_organizations_organization_id_insight_applicationsHandler(cfg),
	}
}
