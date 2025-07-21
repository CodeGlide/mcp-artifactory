package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/nowpayments-api/mcp-server/config"
	"github.com/nowpayments-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetlistofpaymentsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["sortBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortBy=%v", val))
		}
		if val, ok := args["orderBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orderBy=%v", val))
		}
		if val, ok := args["dateFrom"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dateFrom=%v", val))
		}
		if val, ok := args["dateTo"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dateTo=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/payment/%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.APIKey != "" {
			req.Header.Set("x-api-key", cfg.APIKey)
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
		var result models.PaymentListResponse
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

func CreateGetlistofpaymentsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_payment",
		mcp.WithDescription("7. Get list of payments"),
		mcp.WithString("limit", mcp.Description("")),
		mcp.WithString("page", mcp.Description("")),
		mcp.WithString("sortBy", mcp.Description("")),
		mcp.WithString("orderBy", mcp.Description("")),
		mcp.WithString("dateFrom", mcp.Description("")),
		mcp.WithString("dateTo", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetlistofpaymentsHandler(cfg),
	}
}
