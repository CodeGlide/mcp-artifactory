package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Apps_scope_tokenHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		client_idVal, ok := args["client_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: client_id"), nil
		}
		client_id, ok := client_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: client_id"), nil
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
		url := fmt.Sprintf("%s/applications/%s/token/scoped", cfg.BaseURL, client_id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Authorization
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

func CreateApps_scope_tokenTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_applications_client_id_token_scoped",
		mcp.WithDescription("Create a scoped access token"),
		mcp.WithString("client_id", mcp.Required(), mcp.Description("The client ID of the GitHub app.")),
		mcp.WithNumber("target_id", mcp.Description("Input parameter: The ID of the user or organization to scope the user access token to. **Required** unless `target` is specified.")),
		mcp.WithString("access_token", mcp.Required(), mcp.Description("Input parameter: The access token used to authenticate to the GitHub API.")),
		mcp.WithObject("permissions", mcp.Description("Input parameter: The permissions granted to the user access token.")),
		mcp.WithArray("repositories", mcp.Description("Input parameter: The list of repository names to scope the user access token to. `repositories` may not be specified if `repository_ids` is specified.")),
		mcp.WithArray("repository_ids", mcp.Description("Input parameter: The list of repository IDs to scope the user access token to. `repository_ids` may not be specified if `repositories` is specified.")),
		mcp.WithString("target", mcp.Description("Input parameter: The name of the user or organization to scope the user access token to. **Required** unless `target_id` is specified.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apps_scope_tokenHandler(cfg),
	}
}
