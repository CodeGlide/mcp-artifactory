package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Actions_create_self_hosted_runner_group_for_orgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
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
		url := fmt.Sprintf("%s/orgs/%s/actions/runner-groups", cfg.BaseURL, org)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateActions_create_self_hosted_runner_group_for_orgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_orgs_org_actions_runner-groups",
		mcp.WithDescription("Create a self-hosted runner group for an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithArray("selected_repository_ids", mcp.Description("Input parameter: List of repository IDs that can access the runner group.")),
		mcp.WithArray("selected_workflows", mcp.Description("Input parameter: List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.")),
		mcp.WithString("visibility", mcp.Description("Input parameter: Visibility of a runner group. You can select all repositories, select individual repositories, or limit access to private repositories.")),
		mcp.WithBoolean("allows_public_repositories", mcp.Description("Input parameter: Whether the runner group can be used by `public` repositories.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: Name of the runner group.")),
		mcp.WithString("network_configuration_id", mcp.Description("Input parameter: The identifier of a hosted compute network configuration.")),
		mcp.WithBoolean("restricted_to_workflows", mcp.Description("Input parameter: If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.")),
		mcp.WithArray("runners", mcp.Description("Input parameter: List of runner IDs to add to the runner group.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Actions_create_self_hosted_runner_group_for_orgHandler(cfg),
	}
}
