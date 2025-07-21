package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/okta-privileged-access/mcp-server/config"
	"github.com/okta-privileged-access/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_v1_teams_team_name_clientsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		team_nameVal, ok := args["team_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: team_name"), nil
		}
		team_name, ok := team_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: team_name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["all"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("all=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["descending"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("descending=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["prev"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("prev=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["username"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("username=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/teams/%s/clients%s", cfg.BaseURL, team_name, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
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

func CreateGet_v1_teams_team_name_clientsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_teams_team_name_clients",
		mcp.WithDescription("List all Clients"),
		mcp.WithString("Accept", mcp.Description("")),
		mcp.WithString("all", mcp.Description("When `true`, returns all Clients for the Team")),
		mcp.WithString("count", mcp.Description("The number of objects per page")),
		mcp.WithString("descending", mcp.Description("The object order")),
		mcp.WithString("offset", mcp.Description("The offset value for pagination. The **rel=\"next\"** and **rel=\"prev\"** `Link` headers define the offset for subsequent or previous pages.")),
		mcp.WithString("prev", mcp.Description("The direction of paging")),
		mcp.WithString("state", mcp.Description("Only return Clients with the specified state. Valid statuses: `ACTIVE`, `PENDING`, or `DELETED`.")),
		mcp.WithString("username", mcp.Description("Only return Clients assigned to the specified User. An empty string returns unassigned Clients.")),
		mcp.WithString("team_name", mcp.Required(), mcp.Description("(Required) The name of your Team")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v1_teams_team_name_clientsHandler(cfg),
	}
}
