package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jira-integration-api/mcp-server/config"
	"github.com/jira-integration-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_api_services_jira_links_link_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		link_idVal, ok := args["link_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: link_id"), nil
		}
		link_id, ok := link_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: link_id"), nil
		}
		url := fmt.Sprintf("%s/api/services/jira/links/%s", cfg.BaseURL, link_id)
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

func CreateGet_api_services_jira_links_link_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_services_jira_links_link_id",
		mcp.WithDescription("Show Link"),
		mcp.WithString("Accept", mcp.Description("")),
		mcp.WithString("link_id", mcp.Required(), mcp.Description("(Required) The id of the link")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_services_jira_links_link_idHandler(cfg),
	}
}
