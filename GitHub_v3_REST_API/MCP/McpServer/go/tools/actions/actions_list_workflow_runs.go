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

func Actions_list_workflow_runsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		workflow_idVal, ok := args["workflow_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: workflow_id"), nil
		}
		workflow_id, ok := workflow_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: workflow_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["actor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("actor=%v", val))
		}
		if val, ok := args["branch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("branch=%v", val))
		}
		if val, ok := args["event"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("event=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["created"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created=%v", val))
		}
		if val, ok := args["exclude_pull_requests"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude_pull_requests=%v", val))
		}
		if val, ok := args["check_suite_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("check_suite_id=%v", val))
		}
		if val, ok := args["head_sha"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("head_sha=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/repos/%s/%s/actions/workflows/%s/runs%s", cfg.BaseURL, owner, repo, workflow_id, queryString)
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

func CreateActions_list_workflow_runsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repos_owner_repo_actions_workflows_workflow_id_runs",
		mcp.WithDescription("List workflow runs for a workflow"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithString("workflow_id", mcp.Required(), mcp.Description("The ID of the workflow. You can also pass the workflow file name as a string.")),
		mcp.WithString("actor", mcp.Description("Returns someone's workflow runs. Use the login for the user who created the `push` associated with the check suite or workflow run.")),
		mcp.WithString("branch", mcp.Description("Returns workflow runs associated with a branch. Use the name of the branch of the `push`.")),
		mcp.WithString("event", mcp.Description("Returns workflow run triggered by the event you specify. For example, `push`, `pull_request` or `issue`. For more information, see \"[Events that trigger workflows](https://docs.github.com/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows).\"")),
		mcp.WithString("status", mcp.Description("Returns workflow runs with the check run `status` or `conclusion` that you specify. For example, a conclusion can be `success` or a status can be `in_progress`. Only GitHub Actions can set a status of `waiting`, `pending`, or `requested`.")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("created", mcp.Description("Returns workflow runs created within the given date-time range. For more information on the syntax, see \"[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\"")),
		mcp.WithBoolean("exclude_pull_requests", mcp.Description("If `true` pull requests are omitted from the response (empty array).")),
		mcp.WithNumber("check_suite_id", mcp.Description("Returns workflow runs with the `check_suite_id` that you specify.")),
		mcp.WithString("head_sha", mcp.Description("Only returns workflow runs that are associated with the specified `head_sha`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Actions_list_workflow_runsHandler(cfg),
	}
}
