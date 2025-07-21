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

func Dependabot_list_alerts_for_enterpriseHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		enterpriseVal, ok := args["enterprise"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: enterprise"), nil
		}
		enterprise, ok := enterpriseVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: enterprise"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["severity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("severity=%v", val))
		}
		if val, ok := args["ecosystem"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ecosystem=%v", val))
		}
		if val, ok := args["package"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("package=%v", val))
		}
		if val, ok := args["epss_percentage"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("epss_percentage=%v", val))
		}
		if val, ok := args["has"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("has=%v", val))
		}
		if val, ok := args["scope"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("scope=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["before"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("before=%v", val))
		}
		if val, ok := args["after"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("after=%v", val))
		}
		if val, ok := args["first"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("first=%v", val))
		}
		if val, ok := args["last"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("last=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/enterprises/%s/dependabot/alerts%s", cfg.BaseURL, enterprise, queryString)
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

func CreateDependabot_list_alerts_for_enterpriseTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_enterprises_enterprise_dependabot_alerts",
		mcp.WithDescription("List Dependabot alerts for an enterprise"),
		mcp.WithString("enterprise", mcp.Required(), mcp.Description("The slug version of the enterprise name. You can also substitute this value with the enterprise id.")),
		mcp.WithString("state", mcp.Description("A comma-separated list of states. If specified, only alerts with these states will be returned.\n\nCan be: `auto_dismissed`, `dismissed`, `fixed`, `open`")),
		mcp.WithString("severity", mcp.Description("A comma-separated list of severities. If specified, only alerts with these severities will be returned.\n\nCan be: `low`, `medium`, `high`, `critical`")),
		mcp.WithString("ecosystem", mcp.Description("A comma-separated list of ecosystems. If specified, only alerts for these ecosystems will be returned.\n\nCan be: `composer`, `go`, `maven`, `npm`, `nuget`, `pip`, `pub`, `rubygems`, `rust`")),
		mcp.WithString("package", mcp.Description("A comma-separated list of package names. If specified, only alerts for these packages will be returned.")),
		mcp.WithString("epss_percentage", mcp.Description("CVE Exploit Prediction Scoring System (EPSS) percentage. Can be specified as:\n- An exact number (`n`)\n- Comparators such as `>n`, `<n`, `>=n`, `<=n`\n- A range like `n..n`, where `n` is a number from 0.0 to 1.0\n\nFilters the list of alerts based on EPSS percentages. If specified, only alerts with the provided EPSS percentages will be returned.")),
		mcp.WithString("has", mcp.Description("Filters the list of alerts based on whether the alert has the given value. If specified, only alerts meeting this criterion will be returned.\nMultiple `has` filters can be passed to filter for alerts that have all of the values. Currently, only `patch` is supported.")),
		mcp.WithString("scope", mcp.Description("The scope of the vulnerable dependency. If specified, only alerts with this scope will be returned.")),
		mcp.WithString("sort", mcp.Description("The property by which to sort the results.\n`created` means when the alert was created.\n`updated` means when the alert's state last changed.\n`epss_percentage` sorts alerts by the Exploit Prediction Scoring System (EPSS) percentage.")),
		mcp.WithString("direction", mcp.Description("The direction to sort the results by.")),
		mcp.WithString("before", mcp.Description("A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("after", mcp.Description("A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("first", mcp.Description("**Deprecated**. The number of results per page (max 100), starting from the first matching result.\nThis parameter must not be used in combination with `last`.\nInstead, use `per_page` in combination with `after` to fetch the first page of results.")),
		mcp.WithNumber("last", mcp.Description("**Deprecated**. The number of results per page (max 100), starting from the last matching result.\nThis parameter must not be used in combination with `first`.\nInstead, use `per_page` in combination with `before` to fetch the last page of results.")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Dependabot_list_alerts_for_enterpriseHandler(cfg),
	}
}
