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

func Get_v1_teams_team_name_users_user_name_groupsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		user_nameVal, ok := args["user_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: user_name"), nil
		}
		user_name, ok := user_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: user_name"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["contains"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("contains=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["descending"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("descending=%v", val))
		}
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["ignore"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ignore=%v", val))
		}
		if val, ok := args["include_deleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_deleted=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["only_include_deleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("only_include_deleted=%v", val))
		}
		if val, ok := args["prev"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("prev=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/teams/%s/users/%s/groups%s", cfg.BaseURL, team_name, user_name, queryString)
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

func CreateGet_v1_teams_team_name_users_user_name_groupsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_teams_team_name_users_user_name_groups",
		mcp.WithDescription("List all Groups for a User"),
		mcp.WithString("Accept", mcp.Description("")),
		mcp.WithString("contains", mcp.Description("Only return results that include the specified value")),
		mcp.WithString("count", mcp.Description("The number of objects per page")),
		mcp.WithString("descending", mcp.Description("The object order")),
		mcp.WithString("id", mcp.Description("Only return results with the specified IDs")),
		mcp.WithString("ignore", mcp.Description("Ignore Groups with the specified names. This is case sensitive.")),
		mcp.WithString("include_deleted", mcp.Description("If `true`, include deleted Groups in the results")),
		mcp.WithString("offset", mcp.Description("The offset value for pagination. The **rel=\"next\"** and **rel=\"prev\"** `Link` headers define the offset for subsequent or previous pages.")),
		mcp.WithString("only_include_deleted", mcp.Description("If `true`, only return deleted Groups in the results")),
		mcp.WithString("prev", mcp.Description("The direction of paging")),
		mcp.WithString("team_name", mcp.Required(), mcp.Description("(Required) The name of your Team")),
		mcp.WithString("user_name", mcp.Required(), mcp.Description("(Required) The username for an existing User")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v1_teams_team_name_users_user_name_groupsHandler(cfg),
	}
}
