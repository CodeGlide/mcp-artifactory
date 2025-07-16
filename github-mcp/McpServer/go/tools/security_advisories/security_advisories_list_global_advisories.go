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

func Security_advisories_list_global_advisoriesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["ghsa_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ghsa_id=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["cve_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cve_id=%v", val))
		}
		if val, ok := args["ecosystem"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ecosystem=%v", val))
		}
		if val, ok := args["severity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("severity=%v", val))
		}
		if val, ok := args["cwes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cwes=%v", val))
		}
		if val, ok := args["is_withdrawn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_withdrawn=%v", val))
		}
		if val, ok := args["affects"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("affects=%v", val))
		}
		if val, ok := args["published"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("published=%v", val))
		}
		if val, ok := args["updated"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("updated=%v", val))
		}
		if val, ok := args["modified"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified=%v", val))
		}
		if val, ok := args["epss_percentage"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("epss_percentage=%v", val))
		}
		if val, ok := args["epss_percentile"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("epss_percentile=%v", val))
		}
		if val, ok := args["before"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("before=%v", val))
		}
		if val, ok := args["after"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("after=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/advisories%s", cfg.BaseURL, queryString)
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

func CreateSecurity_advisories_list_global_advisoriesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_advisories",
		mcp.WithDescription("List global security advisories"),
		mcp.WithString("ghsa_id", mcp.Description("If specified, only advisories with this GHSA (GitHub Security Advisory) identifier will be returned.")),
		mcp.WithString("type", mcp.Description("If specified, only advisories of this type will be returned. By default, a request with no other parameters defined will only return reviewed advisories that are not malware.")),
		mcp.WithString("cve_id", mcp.Description("If specified, only advisories with this CVE (Common Vulnerabilities and Exposures) identifier will be returned.")),
		mcp.WithString("ecosystem", mcp.Description("If specified, only advisories for these ecosystems will be returned.")),
		mcp.WithString("severity", mcp.Description("If specified, only advisories with these severities will be returned.")),
		mcp.WithString("cwes", mcp.Description("If specified, only advisories with these Common Weakness Enumerations (CWEs) will be returned.\n\nExample: `cwes=79,284,22` or `cwes[]=79&cwes[]=284&cwes[]=22`")),
		mcp.WithBoolean("is_withdrawn", mcp.Description("Whether to only return advisories that have been withdrawn.")),
		mcp.WithString("affects", mcp.Description("If specified, only return advisories that affect any of `package` or `package@version`. A maximum of 1000 packages can be specified.\nIf the query parameter causes the URL to exceed the maximum URL length supported by your client, you must specify fewer packages.\n\nExample: `affects=package1,package2@1.0.0,package3@^2.0.0` or `affects[]=package1&affects[]=package2@1.0.0`")),
		mcp.WithString("published", mcp.Description("If specified, only return advisories that were published on a date or date range.\n\nFor more information on the syntax of the date range, see \"[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\"")),
		mcp.WithString("updated", mcp.Description("If specified, only return advisories that were updated on a date or date range.\n\nFor more information on the syntax of the date range, see \"[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\"")),
		mcp.WithString("modified", mcp.Description("If specified, only show advisories that were updated or published on a date or date range.\n\nFor more information on the syntax of the date range, see \"[Understanding the search syntax](https://docs.github.com/search-github/getting-started-with-searching-on-github/understanding-the-search-syntax#query-for-dates).\"")),
		mcp.WithString("epss_percentage", mcp.Description("If specified, only return advisories that have an EPSS percentage score that matches the provided value.\nThe EPSS percentage represents the likelihood of a CVE being exploited.")),
		mcp.WithString("epss_percentile", mcp.Description("If specified, only return advisories that have an EPSS percentile score that matches the provided value.\nThe EPSS percentile represents the relative rank of the CVE's likelihood of being exploited compared to other CVEs.")),
		mcp.WithString("before", mcp.Description("A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("after", mcp.Description("A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("direction", mcp.Description("The direction to sort the results by.")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("sort", mcp.Description("The property to sort the results by.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Security_advisories_list_global_advisoriesHandler(cfg),
	}
}
