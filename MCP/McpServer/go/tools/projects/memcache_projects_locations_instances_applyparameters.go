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

func Memcache_projects_locations_instances_applyparametersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.ApplyParametersRequest
		
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
		url := fmt.Sprintf("%s/v1/%s:applyParameters", cfg.BaseURL, name)
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

func CreateMemcache_projects_locations_instances_applyparametersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_name:applyParameters",
		mcp.WithDescription(""),
		mcp.WithString("name", mcp.Required(), mcp.Description("Required. Resource name of the Memcached instance for which parameter group updates should be applied.")),
		mcp.WithBoolean("applyAll", mcp.Description("Input parameter: Whether to apply instance-level parameter group to all nodes. If set to true, users are restricted from specifying individual nodes, and `ApplyParameters` updates all nodes within the instance.")),
		mcp.WithArray("nodeIds", mcp.Description("Input parameter: Nodes to which the instance-level parameter group is applied.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Memcache_projects_locations_instances_applyparametersHandler(cfg),
	}
}
