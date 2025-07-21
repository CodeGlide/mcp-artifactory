package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/cloud-memorystore-for-memcached-api/mcp-server/config"
	"github.com/cloud-memorystore-for-memcached-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Memcache_projects_locations_instances_patchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["updateMask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updateMask=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Instance
		
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
		url := fmt.Sprintf("%s/v1/%s%s", cfg.BaseURL, name, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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

func CreateMemcache_projects_locations_instances_patchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_v1_name",
		mcp.WithDescription(""),
		mcp.WithString("name", mcp.Required(), mcp.Description("Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.")),
		mcp.WithString("updateMask", mcp.Description("Required. Mask of fields to update. * `displayName`")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. The time the instance was updated.")),
		mcp.WithString("discoveryEndpoint", mcp.Description("Input parameter: Output only. Endpoint for the Discovery API.")),
		mcp.WithString("memcacheVersion", mcp.Description("Input parameter: The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.")),
		mcp.WithNumber("nodeCount", mcp.Description("Input parameter: Required. Number of nodes in the Memcached instance.")),
		mcp.WithArray("instanceMessages", mcp.Description("Input parameter: List of messages that describe the current state of the Memcached instance.")),
		mcp.WithArray("zones", mcp.Description("Input parameter: Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.")),
		mcp.WithString("authorizedNetwork", mcp.Description("Input parameter: The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.")),
		mcp.WithString("name", mcp.Description("Input parameter: Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. The time the instance was created.")),
		mcp.WithObject("maintenancePolicy", mcp.Description("Input parameter: Maintenance policy per instance.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.")),
		mcp.WithString("memcacheFullVersion", mcp.Description("Input parameter: Output only. The full version of memcached server running on this instance. System automatically determines the full memcached version for an instance based on the input MemcacheVersion. The full version format will be \"memcached-1.5.16\".")),
		mcp.WithObject("parameters", mcp.Description("")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources")),
		mcp.WithObject("maintenanceSchedule", mcp.Description("Input parameter: Upcoming maintenance schedule.")),
		mcp.WithString("state", mcp.Description("Input parameter: Output only. The state of this Memcached instance.")),
		mcp.WithArray("memcacheNodes", mcp.Description("Input parameter: Output only. List of Memcached nodes. Refer to Node message for more details.")),
		mcp.WithObject("nodeConfig", mcp.Description("Input parameter: Configuration for a Memcached Node.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Memcache_projects_locations_instances_patchHandler(cfg),
	}
}
