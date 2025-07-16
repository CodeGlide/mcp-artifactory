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

func Actions_delete_hosted_runner_for_orgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		hosted_runner_idVal, ok := args["hosted_runner_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: hosted_runner_id"), nil
		}
		hosted_runner_id, ok := hosted_runner_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: hosted_runner_id"), nil
		}
		url := fmt.Sprintf("%s/orgs/%s/actions/hosted-runners/%s", cfg.BaseURL, org, hosted_runner_id)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateActions_delete_hosted_runner_for_orgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_orgs_org_actions_hosted-runners_hosted_runner_id",
		mcp.WithDescription("Delete a GitHub-hosted runner for an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithNumber("hosted_runner_id", mcp.Required(), mcp.Description("Unique identifier of the GitHub-hosted runner.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Actions_delete_hosted_runner_for_orgHandler(cfg),
	}
}
