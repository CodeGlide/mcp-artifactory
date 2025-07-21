package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/okta-privileged-access/mcp-server/config"
	"github.com/okta-privileged-access/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Put_v1_teams_team_name_groups_group_name_attributes_attribute_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		group_nameVal, ok := args["group_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: group_name"), nil
		}
		group_name, ok := group_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: group_name"), nil
		}
		attribute_idVal, ok := args["attribute_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: attribute_id"), nil
		}
		attribute_id, ok := attribute_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: attribute_id"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/teams/%s/groups/%s/attributes/%s", cfg.BaseURL, team_name, group_name, attribute_id)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["Content-Type"]; ok {
			req.Header.Set("Content-Type", fmt.Sprintf("%v", val))
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

func CreatePut_v1_teams_team_name_groups_group_name_attributes_attribute_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_v1_teams_team_name_groups_group_name_attributes_attribute_id",
		mcp.WithDescription("Update a Group Attribute"),
		mcp.WithString("Content-Type", mcp.Description("")),
		mcp.WithString("team_name", mcp.Required(), mcp.Description("(Required) The name of your Team")),
		mcp.WithString("group_name", mcp.Required(), mcp.Description("(Required) The name of a Group")),
		mcp.WithString("attribute_id", mcp.Required(), mcp.Description("(Required) The UUID of an Attribute")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_v1_teams_team_name_groups_group_name_attributes_attribute_idHandler(cfg),
	}
}
