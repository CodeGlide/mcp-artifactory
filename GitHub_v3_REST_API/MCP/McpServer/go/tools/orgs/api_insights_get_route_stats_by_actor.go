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

func Api_insights_get_route_stats_by_actorHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		actor_typeVal, ok := args["actor_type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: actor_type"), nil
		}
		actor_type, ok := actor_typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: actor_type"), nil
		}
		actor_idVal, ok := args["actor_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: actor_id"), nil
		}
		actor_id, ok := actor_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: actor_id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["min_timestamp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("min_timestamp=%v", val))
		}
		if val, ok := args["max_timestamp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("max_timestamp=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["api_route_substring"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api_route_substring=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/orgs/%s/insights/api/route-stats/%s/%s%s", cfg.BaseURL, org, actor_type, actor_id, queryString)
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

func CreateApi_insights_get_route_stats_by_actorTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_orgs_org_insights_api_route-stats_actor_type_actor_id",
		mcp.WithDescription("Get route stats by actor"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("actor_type", mcp.Required(), mcp.Description("The type of the actor")),
		mcp.WithNumber("actor_id", mcp.Required(), mcp.Description("The ID of the actor")),
		mcp.WithString("min_timestamp", mcp.Required(), mcp.Description("The minimum timestamp to query for stats. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.")),
		mcp.WithString("max_timestamp", mcp.Description("The maximum timestamp to query for stats. Defaults to the time 30 days ago. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("direction", mcp.Description("The direction to sort the results by.")),
		mcp.WithArray("sort", mcp.Description("The property to sort the results by.")),
		mcp.WithString("api_route_substring", mcp.Description("Providing a substring will filter results where the API route contains the substring. This is a case-insensitive search.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Api_insights_get_route_stats_by_actorHandler(cfg),
	}
}
