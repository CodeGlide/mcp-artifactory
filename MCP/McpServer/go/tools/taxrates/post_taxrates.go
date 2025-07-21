package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/payroll/mcp-server/config"
	"github.com/payroll/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_taxratesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CompanySuiRateDetails985b8306e95d100002daf466c84d0005
		
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
		url := fmt.Sprintf("%s/taxRates", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CompanySuiRateDetails985b8306e95d100002daf466c84d0005
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

func CreatePost_taxratesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_taxRates",
		mcp.WithDescription("Creates a single or a collection of Company SUI Rates."),
		mcp.WithString("id", mcp.Description("Input parameter: Id of the instance")),
		mcp.WithString("startDate", mcp.Description("Input parameter: The start date for company tax reporting.")),
		mcp.WithString("endDate", mcp.Description("Input parameter: The end date for company tax reporting.")),
		mcp.WithString("taxCode", mcp.Description("Input parameter: The deduction for company tax reporting.")),
		mcp.WithString("ein", mcp.Description("Input parameter: The payroll tax authority EIN field for company tax reporting.")),
		mcp.WithString("applicableRate", mcp.Description("Input parameter: The tax override rate for company tax reporting.")),
		mcp.WithObject("companyInstance", mcp.Description("Input parameter: The company object for company tax reporting.")),
		mcp.WithBoolean("exempt", mcp.Description("Input parameter: If true, the SUI rate is exempt.")),
		mcp.WithObject("stateInstance", mcp.Description("Input parameter: The payroll tax authority object for company tax reporting.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_taxratesHandler(cfg),
	}
}
