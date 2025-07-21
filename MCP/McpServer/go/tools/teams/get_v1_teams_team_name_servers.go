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

func Get_v1_teams_team_name_serversHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["AdServers"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("AdServers=%v", val))
		}
		if val, ok := args["alt_names_contains"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("alt_names_contains=%v", val))
		}
		if val, ok := args["bastion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bastion=%v", val))
		}
		if val, ok := args["canonical_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("canonical_name=%v", val))
		}
		if val, ok := args["cloud_account"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cloud_account=%v", val))
		}
		if val, ok := args["cloud_provider"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cloud_provider=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["credentialed"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("credentialed=%v", val))
		}
		if val, ok := args["descending"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("descending=%v", val))
		}
		if val, ok := args["has_account_under_management"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("has_account_under_management=%v", val))
		}
		if val, ok := args["hostname"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hostname=%v", val))
		}
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["instance_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("instance_id=%v", val))
		}
		if val, ok := args["managed"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("managed=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["prev"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("prev=%v", val))
		}
		if val, ok := args["project_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("project_name=%v", val))
		}
		if val, ok := args["state"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("state=%v", val))
		}
		if val, ok := args["selector"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("selector=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/teams/%s/servers%s", cfg.BaseURL, team_name, queryString)
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

func CreateGet_v1_teams_team_name_serversTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_teams_team_name_servers",
		mcp.WithDescription("List all Servers for a Team"),
		mcp.WithString("Accept", mcp.Description("")),
		mcp.WithString("AdServers", mcp.Description("If `true`, only return AD servers. If `false`, only return Non-AD servers")),
		mcp.WithString("alt_names_contains", mcp.Description("Only return Servers that contain the specified alternate name")),
		mcp.WithString("bastion", mcp.Description("Only return Servers associated with the specified bastion")),
		mcp.WithString("canonical_name", mcp.Description("A canonical name")),
		mcp.WithString("cloud_account", mcp.Description("Only return Servers associated with the specified cloud account")),
		mcp.WithString("cloud_provider", mcp.Description("Only return Servers associated with the specified cloud provider. Possible values: `aws` or `gce`")),
		mcp.WithString("count", mcp.Description("The number of objects per page")),
		mcp.WithString("credentialed", mcp.Description("If `true`, only return unmanaged Servers with credential issuance enabled. If `false`, only return unmanaged Servers with credential issuance disabled.")),
		mcp.WithString("descending", mcp.Description("The object order")),
		mcp.WithString("has_account_under_management", mcp.Description("If `true`, only return Servers that currently have at least one account's password under management'.  If `false`, only return servers that do not currently have an account whose password is under management.")),
		mcp.WithString("hostname", mcp.Description("Only return Servers that contain the specified hostname")),
		mcp.WithString("id", mcp.Description("Only return Servers with the specified IDs. Only usable for PAM administrative views of servers, not end-user Server views.")),
		mcp.WithString("instance_id", mcp.Description("Only return Servers that contain the specified instance ID")),
		mcp.WithString("managed", mcp.Description("If `true`, only return managed servers. If `false`, only return unmanaged servers.")),
		mcp.WithString("offset", mcp.Description("The offset value for pagination. The **rel=\"next\"** and **rel=\"prev\"** `Link` headers define the offset for subsequent or previous pages.")),
		mcp.WithString("prev", mcp.Description("The direction of paging")),
		mcp.WithString("project_name", mcp.Description("Only return Servers that belong to the specified Project")),
		mcp.WithString("state", mcp.Description("Include Servers with the specified state. Valid statuses: `ACTIVE` or `INACTIVE`.")),
		mcp.WithString("selector", mcp.Description("Only return Servers that contain the specified Server selectors.")),
		mcp.WithString("team_name", mcp.Required(), mcp.Description("(Required) The name of your Team")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v1_teams_team_name_serversHandler(cfg),
	}
}
