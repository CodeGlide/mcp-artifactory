package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/cloud-memorystore-for-memcached-api/mcp-server/config"
	"github.com/cloud-memorystore-for-memcached-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Memcache_projects_locations_instances_reschedulemaintenanceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		instanceVal, ok := args["instance"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: instance"), nil
		}
		instance, ok := instanceVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: instance"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.RescheduleMaintenanceRequest
		
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
		url := fmt.Sprintf("%s/v1/%s:rescheduleMaintenance", cfg.BaseURL, instance)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Operation
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

func CreateMemcache_projects_locations_instances_reschedulemaintenanceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_instance:rescheduleMaintenance",
		mcp.WithDescription(""),
		mcp.WithString("instance", mcp.Required(), mcp.Description("Required. Memcache instance resource name using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` where `location_id` refers to a GCP region.")),
		mcp.WithString("scheduleTime", mcp.Description("Input parameter: Timestamp when the maintenance shall be rescheduled to if reschedule_type=SPECIFIC_TIME, in RFC 3339 format, for example `2012-11-15T16:19:00.094Z`.")),
		mcp.WithString("rescheduleType", mcp.Description("Input parameter: Required. If reschedule type is SPECIFIC_TIME, must set up schedule_time as well.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Memcache_projects_locations_instances_reschedulemaintenanceHandler(cfg),
	}
}
