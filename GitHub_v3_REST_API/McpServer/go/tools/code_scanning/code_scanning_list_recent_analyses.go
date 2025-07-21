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

func Code_scanning_list_recent_analysesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["tool_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tool_name=%v", val))
		}
		if val, ok := args["tool_guid"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tool_guid=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["pr"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pr=%v", val))
		}
		if val, ok := args["ref"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ref=%v", val))
		}
		if val, ok := args["sarif_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sarif_id=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/repos/%s/%s/code-scanning/analyses%s", cfg.BaseURL, owner, repo, queryString)
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
		var result []GeneratedType
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

func CreateCode_scanning_list_recent_analysesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repos_owner_repo_code-scanning_analyses",
		mcp.WithDescription("List code scanning analyses for a repository"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithString("tool_name", mcp.Description("The name of a code scanning tool. Only results by this tool will be listed. You can specify the tool by using either `tool_name` or `tool_guid`, but not both.")),
		mcp.WithString("tool_guid", mcp.Description("The GUID of a code scanning tool. Only results by this tool will be listed. Note that some code scanning tools may not include a GUID in their analysis data. You can specify the tool by using either `tool_guid` or `tool_name`, but not both.")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("pr", mcp.Description("The number of the pull request for the results you want to list.")),
		mcp.WithString("ref", mcp.Description("The Git reference for the analyses you want to list. The `ref` for a branch can be formatted either as `refs/heads/<branch name>` or simply `<branch name>`. To reference a pull request use `refs/pull/<number>/merge`.")),
		mcp.WithString("sarif_id", mcp.Description("Filter analyses belonging to the same SARIF upload.")),
		mcp.WithString("direction", mcp.Description("The direction to sort the results by.")),
		mcp.WithString("sort", mcp.Description("The property by which to sort the results.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Code_scanning_list_recent_analysesHandler(cfg),
	}
}
