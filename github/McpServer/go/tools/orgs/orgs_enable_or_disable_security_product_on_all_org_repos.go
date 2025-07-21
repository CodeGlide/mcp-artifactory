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

func Orgs_enable_or_disable_security_product_on_all_org_reposHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		security_productVal, ok := args["security_product"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: security_product"), nil
		}
		security_product, ok := security_productVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: security_product"), nil
		}
		enablementVal, ok := args["enablement"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: enablement"), nil
		}
		enablement, ok := enablementVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: enablement"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.interface{}
		
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
		url := fmt.Sprintf("%s/orgs/%s/%s/%s", cfg.BaseURL, org, security_product, enablement)
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

func CreateOrgs_enable_or_disable_security_product_on_all_org_reposTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_orgs_org_security_product_enablement",
		mcp.WithDescription("Enable or disable a security feature for an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("security_product", mcp.Required(), mcp.Description("The security feature to enable or disable.")),
		mcp.WithString("enablement", mcp.Required(), mcp.Description("The action to take.\n\n`enable_all` means to enable the specified security feature for all repositories in the organization.\n`disable_all` means to disable the specified security feature for all repositories in the organization.")),
		mcp.WithString("query_suite", mcp.Description("Input parameter: CodeQL query suite to be used. If you specify the `query_suite` parameter, the default setup will be configured with this query suite only on all repositories that didn't have default setup already configured. It will not change the query suite on repositories that already have default setup configured.\nIf you don't specify any `query_suite` in your request, the preferred query suite of the organization will be applied.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_enable_or_disable_security_product_on_all_org_reposHandler(cfg),
	}
}
