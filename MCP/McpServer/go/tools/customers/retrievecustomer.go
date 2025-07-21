package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/square-connect-api/mcp-server/config"
	"github.com/square-connect-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func RetrievecustomerHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		customer_idVal, ok := args["customer_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: customer_id"), nil
		}
		customer_id, ok := customer_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: customer_id"), nil
		}
		url := fmt.Sprintf("%s/customers/%s", cfg.BaseURL, customer_id)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.APIKey != "" {
			req.Header.Set("_connectv2-docs_session", cfg.APIKey)
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
		var result models.Customer
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

func CreateRetrievecustomerTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_customers_customer_id",
		mcp.WithDescription("RetrieveCustomer"),
		mcp.WithString("customer_id", mcp.Required(), mcp.Description("The customer's ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    RetrievecustomerHandler(cfg),
	}
}
