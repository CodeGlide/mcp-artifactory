package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Codespaces_get_export_details_for_authenticated_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		codespace_nameVal, ok := args["codespace_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: codespace_name"), nil
		}
		codespace_name, ok := codespace_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: codespace_name"), nil
		}
		export_idVal, ok := args["export_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: export_id"), nil
		}
		export_id, ok := export_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: export_id"), nil
		}
		url := fmt.Sprintf("%s/user/codespaces/%s/exports/%s", cfg.BaseURL, codespace_name, export_id)
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
		var result models.GeneratedType
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

func CreateCodespaces_get_export_details_for_authenticated_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_user_codespaces_codespace_name_exports_export_id",
		mcp.WithDescription("Get details about a codespace export"),
		mcp.WithString("codespace_name", mcp.Required(), mcp.Description("The name of the codespace.")),
		mcp.WithString("export_id", mcp.Required(), mcp.Description("The ID of the export operation, or `latest`. Currently only `latest` is currently supported.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Codespaces_get_export_details_for_authenticated_userHandler(cfg),
	}
}
