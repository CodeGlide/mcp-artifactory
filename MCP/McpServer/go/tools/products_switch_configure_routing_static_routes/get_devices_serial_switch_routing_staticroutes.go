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

func Get_devices_serial_switch_routing_staticroutesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		serialVal, ok := args["serial"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serial"), nil
		}
		serial, ok := serialVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serial"), nil
		}
		url := fmt.Sprintf("%s/devices/%s/switch/routing/staticRoutes", cfg.BaseURL, serial)
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

func CreateGet_devices_serial_switch_routing_staticroutesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_devices_serial_switch_routing_staticRoutes",
		mcp.WithDescription("List Layer 3 Static Routes For A Switch"),
		mcp.WithString("X-Cisco-Meraki-API-Key", mcp.Description("")),
		mcp.WithString("serial", mcp.Required(), mcp.Description("(Required) ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_devices_serial_switch_routing_staticroutesHandler(cfg),
	}
}
