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

func Orgs_get_org_ruleset_versionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		ruleset_idVal, ok := args["ruleset_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ruleset_id"), nil
		}
		ruleset_id, ok := ruleset_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ruleset_id"), nil
		}
		version_idVal, ok := args["version_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: version_id"), nil
		}
		version_id, ok := version_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: version_id"), nil
		}
		url := fmt.Sprintf("%s/orgs/%s/rulesets/%s/history/%s", cfg.BaseURL, org, ruleset_id, version_id)
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

func CreateOrgs_get_org_ruleset_versionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_orgs_org_rulesets_ruleset_id_history_version_id",
		mcp.WithDescription("Get organization ruleset version"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithNumber("ruleset_id", mcp.Required(), mcp.Description("The ID of the ruleset.")),
		mcp.WithNumber("version_id", mcp.Required(), mcp.Description("The ID of the version")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_get_org_ruleset_versionHandler(cfg),
	}
}
