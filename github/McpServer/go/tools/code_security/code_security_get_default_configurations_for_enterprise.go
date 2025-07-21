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

func Code_security_get_default_configurations_for_enterpriseHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		enterpriseVal, ok := args["enterprise"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: enterprise"), nil
		}
		enterprise, ok := enterpriseVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: enterprise"), nil
		}
		url := fmt.Sprintf("%s/enterprises/%s/code-security/configurations/defaults", cfg.BaseURL, enterprise)
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
		var result []map[string]interface{}
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

func CreateCode_security_get_default_configurations_for_enterpriseTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_enterprises_enterprise_code-security_configurations_defaults",
		mcp.WithDescription("Get default code security configurations for an enterprise"),
		mcp.WithString("enterprise", mcp.Required(), mcp.Description("The slug version of the enterprise name. You can also substitute this value with the enterprise id.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Code_security_get_default_configurations_for_enterpriseHandler(cfg),
	}
}
