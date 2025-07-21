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

func Issues_check_user_can_be_assigned_to_issueHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		issue_numberVal, ok := args["issue_number"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: issue_number"), nil
		}
		issue_number, ok := issue_numberVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: issue_number"), nil
		}
		assigneeVal, ok := args["assignee"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: assignee"), nil
		}
		assignee, ok := assigneeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: assignee"), nil
		}
		url := fmt.Sprintf("%s/repos/%s/%s/issues/%s/assignees/%s", cfg.BaseURL, owner, repo, issue_number, assignee)
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

func CreateIssues_check_user_can_be_assigned_to_issueTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repos_owner_repo_issues_issue_number_assignees_assignee",
		mcp.WithDescription("Check if a user can be assigned to a issue"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithNumber("issue_number", mcp.Required(), mcp.Description("The number that identifies the issue.")),
		mcp.WithString("assignee", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Issues_check_user_can_be_assigned_to_issueHandler(cfg),
	}
}
