package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/slack-web-api/mcp-server/config"
	"github.com/slack-web-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Workflows_updatestepHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["workflow_step_edit_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("workflow_step_edit_id=%v", val))
		}
		if val, ok := args["inputs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("inputs=%v", val))
		}
		if val, ok := args["outputs"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("outputs=%v", val))
		}
		if val, ok := args["step_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("step_name=%v", val))
		}
		if val, ok := args["step_image_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("step_image_url=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/workflows.updateStep%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["token"]; ok {
			req.Header.Set("token", fmt.Sprintf("%v", val))
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

func CreateWorkflows_updatestepTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_workflows.updateStep",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Required(), mcp.Description("Authentication token. Requires scope: `workflow.steps:execute`")),
		mcp.WithString("workflow_step_edit_id", mcp.Required(), mcp.Description("A context identifier provided with `view_submission` payloads used to call back to `workflows.updateStep`.")),
		mcp.WithString("inputs", mcp.Description("A JSON key-value map of inputs required from a user during configuration. This is the data your app expects to receive when the workflow step starts. **Please note**: the embedded variable format is set and replaced by the workflow system. You cannot create custom variables that will be replaced at runtime. [Read more about variables in workflow steps here](/workflows/steps#variables).")),
		mcp.WithString("outputs", mcp.Description("An JSON array of output objects used during step execution. This is the data your app agrees to provide when your workflow step was executed.")),
		mcp.WithString("step_name", mcp.Description("An optional field that can be used to override the step name that is shown in the Workflow Builder.")),
		mcp.WithString("step_image_url", mcp.Description("An optional field that can be used to override app image that is shown in the Workflow Builder.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Workflows_updatestepHandler(cfg),
	}
}
