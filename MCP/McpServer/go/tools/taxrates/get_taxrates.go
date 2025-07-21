package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/payroll/mcp-server/config"
	"github.com/payroll/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_taxratesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["company"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("company=%v", val))
		}
		if val, ok := args["effective"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("effective=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["payrollStateAuthorityTaxCode"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("payrollStateAuthorityTaxCode=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/taxRates%s", cfg.BaseURL, queryString)
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

func CreateGet_taxratesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_taxRates",
		mcp.WithDescription("Retrieves a single or a collection of company SUI rates."),
		mcp.WithString("company", mcp.Description("The company reference ID or WID that represents 1 or more companies. Example: company=comp1&company=comp2&company=cb550da820584750aae8f807882fa79a")),
		mcp.WithString("effective", mcp.Description("The effective date for the SUI rate, using the yyyy-mm-dd format.")),
		mcp.WithNumber("limit", mcp.Description("The maximum number of objects in a single response. The default is 20. The maximum is 100.")),
		mcp.WithNumber("offset", mcp.Description("The zero-based index of the first object in a response collection. The default is 0. Use offset with the limit parameter to control paging of a response collection. Example: If limit is 5 and offset is 9, the response returns a collection of 5 objects starting with the 10th object.")),
		mcp.WithString("payrollStateAuthorityTaxCode", mcp.Description("The FIPS code or WID that represents 1 or more states. Example: payrollStateAuthorityTaxCode=06&payrollStateAuthorityTaxCode=3b3d378d5f4a48b8b3ac46fee0703226&payrollStateAuthorityTaxCode=48")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_taxratesHandler(cfg),
	}
}
