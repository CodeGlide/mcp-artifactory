package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Actions_get_workflow_run_attemptHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ownerVal, ok := args["owner"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: owner"), nil
		}
		owner, ok := ownerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: owner"), nil
		}
		repoVal, ok := args["repo"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo"), nil
		}
		repo, ok := repoVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo"), nil
		}
		run_idVal, ok := args["run_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: run_id"), nil
		}
		run_id, ok := run_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: run_id"), nil
		}
		attempt_numberVal, ok := args["attempt_number"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: attempt_number"), nil
		}
		attempt_number, ok := attempt_numberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: attempt_number"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["exclude_pull_requests"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude_pull_requests=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/repos/%s/%s/actions/runs/%s/attempts/%s%s", cfg.BaseURL, owner, repo, run_id, attempt_number, queryString)
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

func CreateActions_get_workflow_run_attemptTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repos_owner_repo_actions_runs_run_id_attempts_attempt_number",
		mcp.WithDescription("Get a workflow run attempt"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithNumber("run_id", mcp.Required(), mcp.Description("The unique identifier of the workflow run.")),
		mcp.WithNumber("attempt_number", mcp.Required(), mcp.Description("The attempt number of the workflow run.")),
		mcp.WithBoolean("exclude_pull_requests", mcp.Description("If `true` pull requests are omitted from the response (empty array).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Actions_get_workflow_run_attemptHandler(cfg),
	}
}
