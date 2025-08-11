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

func Secret_scanning_list_alerts_for_orgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["secret_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("secret_type=%v", val))
		}
		if val, ok := args["resolution"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resolution=%v", val))
		}
		if val, ok := args["sort"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort=%v", val))
		}
		if val, ok := args["direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("direction=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["per_page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("per_page=%v", val))
		}
		if val, ok := args["before"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("before=%v", val))
		}
		if val, ok := args["after"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("after=%v", val))
		}
		if val, ok := args["validity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("validity=%v", val))
		}
		if val, ok := args["is_publicly_leaked"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_publicly_leaked=%v", val))
		}
		if val, ok := args["is_multi_repo"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_multi_repo=%v", val))
		}
		if val, ok := args["hide_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hide_secret=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/orgs/%s/secret-scanning/alerts%s", cfg.BaseURL, org, queryString)
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

func CreateSecret_scanning_list_alerts_for_orgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_orgs_org_secret-scanning_alerts",
		mcp.WithDescription("List secret scanning alerts for an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("state", mcp.Description("Set to `open` or `resolved` to only list secret scanning alerts in a specific state.")),
		mcp.WithString("secret_type", mcp.Description("A comma-separated list of secret types to return. All default secret patterns are returned. To return generic patterns, pass the token name(s) in the parameter. See \"[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)\" for a complete list of secret types.")),
		mcp.WithString("resolution", mcp.Description("A comma-separated list of resolutions. Only secret scanning alerts with one of these resolutions are listed. Valid resolutions are `false_positive`, `wont_fix`, `revoked`, `pattern_edited`, `pattern_deleted` or `used_in_tests`.")),
		mcp.WithString("sort", mcp.Description("The property to sort the results by. `created` means when the alert was created. `updated` means when the alert was updated or resolved.")),
		mcp.WithString("direction", mcp.Description("The direction to sort the results by.")),
		mcp.WithNumber("page", mcp.Description("The page number of the results to fetch. For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithNumber("per_page", mcp.Description("The number of results per page (max 100). For more information, see \"[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api).\"")),
		mcp.WithString("before", mcp.Description("A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events before this cursor. To receive an initial cursor on your first request, include an empty \"before\" query string.")),
		mcp.WithString("after", mcp.Description("A cursor, as given in the [Link header](https://docs.github.com/rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for events after this cursor.  To receive an initial cursor on your first request, include an empty \"after\" query string.")),
		mcp.WithString("validity", mcp.Description("A comma-separated list of validities that, when present, will return alerts that match the validities in this list. Valid options are `active`, `inactive`, and `unknown`.")),
		mcp.WithBoolean("is_publicly_leaked", mcp.Description("A boolean value representing whether or not to filter alerts by the publicly-leaked tag being present.")),
		mcp.WithBoolean("is_multi_repo", mcp.Description("A boolean value representing whether or not to filter alerts by the multi-repo tag being present.")),
		mcp.WithBoolean("hide_secret", mcp.Description("A boolean value representing whether or not to hide literal secrets in the results.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Secret_scanning_list_alerts_for_orgHandler(cfg),
	}
}
