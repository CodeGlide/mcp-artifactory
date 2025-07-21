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

func Repos_disable_deployment_protection_ruleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		environment_nameVal, ok := args["environment_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: environment_name"), nil
		}
		environment_name, ok := environment_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: environment_name"), nil
		}
		repoVal, ok := args["repo"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo"), nil
		}
		repo, ok := repoVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo"), nil
		}
		ownerVal, ok := args["owner"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: owner"), nil
		}
		owner, ok := ownerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: owner"), nil
		}
		protection_rule_idVal, ok := args["protection_rule_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: protection_rule_id"), nil
		}
		protection_rule_id, ok := protection_rule_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: protection_rule_id"), nil
		}
		url := fmt.Sprintf("%s/repos/%s/%s/environments/%s/deployment_protection_rules/%s", cfg.BaseURL, environment_name, repo, owner, protection_rule_id)
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

func CreateRepos_disable_deployment_protection_ruleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_repos_owner_repo_environments_environment_name_deployment_protection_rules_protection_rule_id",
		mcp.WithDescription("Disable a custom protection rule for an environment"),
		mcp.WithString("environment_name", mcp.Required(), mcp.Description("The name of the environment. The name must be URL encoded. For example, any slashes in the name must be replaced with `%2F`.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithNumber("protection_rule_id", mcp.Required(), mcp.Description("The unique identifier of the protection rule.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repos_disable_deployment_protection_ruleHandler(cfg),
	}
}
