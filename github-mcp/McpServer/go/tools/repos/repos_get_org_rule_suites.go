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

func Repos_get_org_rule_suitesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["ref"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ref=%v", val))
		}
		if val, ok := args["repository_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("repository_name=%v", val))
		}
		if val, ok := args["time_period"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("time_period=%v", val))
		}
		if val, ok := args["actor_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("actor_name=%v", val))
		}
		if val, ok := args["rule_suite_result"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("rule_suite_result=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/orgs/%s/rulesets/rule-suites%s", cfg.BaseURL, org, queryString)
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
		var result []map[string]interface{}
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

func CreateRepos_get_org_rule_suitesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_orgs_org_rulesets_rule-suites",
		mcp.WithDescription("List organization rule suites"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("ref", mcp.Description("The name of the ref. Cannot contain wildcard characters. Optionally prefix with `refs/heads/` to limit to branches or `refs/tags/` to limit to tags. Omit the prefix to search across all refs. When specified, only rule evaluations triggered for this ref will be returned.")),
		mcp.WithString("repository_name", mcp.Description("The name of the repository to filter on.")),
		mcp.WithString("time_period", mcp.Description("The time period to filter by.\n\nFor example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for rule suites that occurred in the past 7 days (168 hours).")),
		mcp.WithString("actor_name", mcp.Description("The handle for the GitHub user account to filter on. When specified, only rule evaluations triggered by this actor will be returned.")),
		mcp.WithString("rule_suite_result", mcp.Description("The rule suite results to filter on. When specified, only suites with this result will be returned.")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repos_get_org_rule_suitesHandler(cfg),
	}
}
