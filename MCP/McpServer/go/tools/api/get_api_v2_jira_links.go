package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jira-integration-api/mcp-server/config"
	"github.com/jira-integration-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_api_v2_jira_linksHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["page[after]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page[after]=%v", val))
		}
		if val, ok := args["page[size]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page[size]=%v", val))
		}
		if val, ok := args["filter[ticket_id]"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("filter[ticket_id]=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/v2/jira/links%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")
		if val, ok := args["Accept"]; ok {
			req.Header.Set("Accept", fmt.Sprintf("%v", val))
		}

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

func CreateGet_api_v2_jira_linksTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_v2_jira_links",
		mcp.WithDescription("List Links"),
		mcp.WithString("Accept", mcp.Description("")),
		mcp.WithString("page[after]", mcp.Description("When provided, the returned paginated data must have as its first item the item that is immediately after the cursor in the results list. Exception: If there are no items in the results list that fall after the cursor, the returned paginated data must be an empty array")),
		mcp.WithNumber("page[size]", mcp.Description("The number of entries that will be returned")),
		mcp.WithString("filter[ticket_id]", mcp.Description("List links for a specific Zendesk ticket or Jira issue by specifying a ticket id\nor issue id. Filtering by issue key is not currently supported\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_v2_jira_linksHandler(cfg),
	}
}
