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

func Orgs_list_pat_grantsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["owner"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner=%v", val))
		}
		if val, ok := args["repository"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("repository=%v", val))
		}
		if val, ok := args["permission"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("permission=%v", val))
		}
		if val, ok := args["last_used_before"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_used_before=%v", val))
		}
		if val, ok := args["last_used_after"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last_used_after=%v", val))
		}
		if val, ok := args["token_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("token_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/orgs/%s/personal-access-tokens%s", cfg.BaseURL, org, queryString)
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

func CreateOrgs_list_pat_grantsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_orgs_org_personal-access-tokens",
		mcp.WithDescription("List fine-grained personal access tokens with access to organization resources"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("sort", mcp.Description("The property by which to sort the results.")),
		mcp.WithString("direction", mcp.Description("The direction to sort the results by.")),
		mcp.WithArray("owner", mcp.Description("A list of owner usernames to use to filter the results.")),
		mcp.WithString("repository", mcp.Description("The name of the repository to use to filter the results.")),
		mcp.WithString("permission", mcp.Description("The permission to use to filter the results.")),
		mcp.WithString("last_used_before", mcp.Description("Only show fine-grained personal access tokens used before the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.")),
		mcp.WithString("last_used_after", mcp.Description("Only show fine-grained personal access tokens used after the given time. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.")),
		mcp.WithArray("token_id", mcp.Description("The ID of the token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_list_pat_grantsHandler(cfg),
	}
}
