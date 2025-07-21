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

func Patch_payrollinputs_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody models.PayrollInputViewdc697bd5bc1b100005f3109b925b00f6
		
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
		url := fmt.Sprintf("%s/payrollInputs/%s", cfg.BaseURL, ID)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.PayrollInputViewdc697bd5bc1b100005f3109b925b00f6
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

func CreatePatch_payrollinputs_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_payrollInputs_ID",
		mcp.WithDescription("Partially updates an existing payroll input instance."),
		mcp.WithString("ID", mcp.Required(), mcp.Description("The Workday ID of the resource.")),
		mcp.WithArray("inputDetails", mcp.Description("Input parameter: The details for this payroll input.")),
		mcp.WithString("startDate", mcp.Description("Input parameter: The start date before which this input does not apply.")),
		mcp.WithString("fieldEditability", mcp.Description("Input parameter: The editability status indicating the fields that can be updated in the payroll input request. Possible values: all, none, endDateOnly")),
		mcp.WithObject("worker", mcp.Description("Input parameter: The worker for this payroll input.")),
		mcp.WithObject("position", mcp.Description("Input parameter: The worker's position the payroll input applies to if Multi Job Payroll is used.")),
		mcp.WithObject("currency", mcp.Description("Input parameter: The currency for the payroll input. If no currency exists, the system assumes the Pay Group currency. The Pay Group currency is derived from the default currency for the Pay Group country.")),
		mcp.WithString("comment", mcp.Description("Input parameter: The text comment for this input.")),
		mcp.WithArray("runCategories", mcp.Description("Input parameter: The run category for the payroll input.")),
		mcp.WithArray("worktags", mcp.Description("Input parameter: The worktags associated with the payroll input.")),
		mcp.WithBoolean("adjustment", mcp.Description("Input parameter: If true, the input is for an adjustment as opposed to an override.")),
		mcp.WithObject("payComponent", mcp.Description("Input parameter: The pay component for this payroll input.")),
		mcp.WithString("id", mcp.Description("Input parameter: Id of the instance")),
		mcp.WithString("descriptor", mcp.Description("Input parameter: A preview of the instance")),
		mcp.WithString("endDate", mcp.Description("Input parameter: The end date after which this input does not apply.")),
		mcp.WithBoolean("ongoing", mcp.Description("Input parameter: If true, the payroll input is ongoing.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_payrollinputs_idHandler(cfg),
	}
}
