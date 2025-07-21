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

func Code_security_create_configurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		orgVal, ok := args["org"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org"), nil
		}
		org, ok := orgVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org"), nil
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
		url := fmt.Sprintf("%s/orgs/%s/code-security/configurations", cfg.BaseURL, org)
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

func CreateCode_security_create_configurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_orgs_org_code-security_configurations",
		mcp.WithDescription("Create a code security configuration"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("secret_scanning_validity_checks", mcp.Description("Input parameter: The enablement status of secret scanning validity checks")),
		mcp.WithString("code_scanning_delegated_alert_dismissal", mcp.Description("Input parameter: The enablement status of code scanning delegated alert dismissal")),
		mcp.WithString("enforcement", mcp.Description("Input parameter: The enforcement status for a security configuration")),
		mcp.WithObject("dependency_graph_autosubmit_action_options", mcp.Description("Input parameter: Feature options for Automatic dependency submission")),
		mcp.WithString("secret_scanning", mcp.Description("Input parameter: The enablement status of secret scanning")),
		mcp.WithString("code_scanning_default_setup", mcp.Description("Input parameter: The enablement status of code scanning default setup")),
		mcp.WithObject("code_scanning_default_setup_options", mcp.Description("Input parameter: Feature options for code scanning default setup")),
		mcp.WithString("secret_scanning_delegated_alert_dismissal", mcp.Description("Input parameter: The enablement status of secret scanning delegated alert dismissal")),
		mcp.WithString("secret_scanning_delegated_bypass", mcp.Description("Input parameter: The enablement status of secret scanning delegated bypass")),
		mcp.WithString("secret_scanning_push_protection", mcp.Description("Input parameter: The enablement status of secret scanning push protection")),
		mcp.WithString("dependabot_alerts", mcp.Description("Input parameter: The enablement status of Dependabot alerts")),
		mcp.WithString("private_vulnerability_reporting", mcp.Description("Input parameter: The enablement status of private vulnerability reporting")),
		mcp.WithString("advanced_security", mcp.Description("Input parameter: The enablement status of GitHub Advanced Security features. `enabled` will enable both Code Security and Secret Protection features.")),
		mcp.WithObject("secret_scanning_delegated_bypass_options", mcp.Description("Input parameter: Feature options for secret scanning delegated bypass")),
		mcp.WithString("secret_scanning_non_provider_patterns", mcp.Description("Input parameter: The enablement status of secret scanning non provider patterns")),
		mcp.WithString("dependency_graph_autosubmit_action", mcp.Description("Input parameter: The enablement status of Automatic dependency submission")),
		mcp.WithString("secret_scanning_generic_secrets", mcp.Description("Input parameter: The enablement status of Copilot secret scanning")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Input parameter: A description of the code security configuration")),
		mcp.WithString("dependency_graph", mcp.Description("Input parameter: The enablement status of Dependency Graph")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the code security configuration. Must be unique within the organization.")),
		mcp.WithString("dependabot_security_updates", mcp.Description("Input parameter: The enablement status of Dependabot security updates")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Code_security_create_configurationHandler(cfg),
	}
}
