package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/payroll/mcp-server/config"
	"github.com/payroll/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_paygroupdetails_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		IDVal, ok := args["ID"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ID"), nil
		}
		ID, ok := IDVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ID"), nil
		}
		url := fmt.Sprintf("%s/payGroupDetails/%s", cfg.BaseURL, ID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
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
		var result models.PayGroupDetailViewe5811059356910001433d1f0fa2e2256
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

func CreateGet_paygroupdetails_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_payGroupDetails_ID",
		mcp.WithDescription("Retrieves a single pay group detail instance."),
		mcp.WithString("ID", mcp.Required(), mcp.Description("The Workday ID of the resource.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_paygroupdetails_idHandler(cfg),
	}
}
