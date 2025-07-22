package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/meraki-dashboard-api-v1-9-0/mcp-server/config"
	"github.com/meraki-dashboard-api-v1-9-0/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_organizations_organizationid_insight_monitoredmediaservers_monitoredmediaserveridHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		organizationIdVal, ok := args["organizationId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: organizationId"), nil
		}
		organizationId, ok := organizationIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: organizationId"), nil
		}
		monitoredMediaServerIdVal, ok := args["monitoredMediaServerId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: monitoredMediaServerId"), nil
		}
		monitoredMediaServerId, ok := monitoredMediaServerIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: monitoredMediaServerId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody string
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/organizations/%s/insight/monitoredMediaServers/%s", cfg.BaseURL, organizationId, monitoredMediaServerId)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Cisco-Meraki-API-Key"]; ok {
			req.Header.Set("X-Cisco-Meraki-API-Key", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Content-Type"]; ok {
			req.Header.Set("Content-Type", fmt.Sprintf("%v", val))
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

func CreatePut_organizations_organizationid_insight_monitoredmediaservers_monitoredmediaserveridTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_organizations_organizationId_insight_monitoredMediaServers_monitoredMediaServerId",
		mcp.WithDescription("Update A Monitored Media Server For This Organization"),
		mcp.WithString("X-Cisco-Meraki-API-Key", mcp.Description("")),
		mcp.WithString("Content-Type", mcp.Description("")),
		mcp.WithString("organizationId", mcp.Required(), mcp.Description("(Required) ")),
		mcp.WithString("monitoredMediaServerId", mcp.Required(), mcp.Description("(Required) ")),
		mcp.WithString("value", mcp.Required(), mcp.Description("String value")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_organizations_organizationid_insight_monitoredmediaservers_monitoredmediaserveridHandler(cfg),
	}
}
