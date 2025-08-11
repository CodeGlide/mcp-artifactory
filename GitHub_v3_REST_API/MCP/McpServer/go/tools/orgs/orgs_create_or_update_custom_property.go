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

func Orgs_create_or_update_custom_propertyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		custom_property_nameVal, ok := args["custom_property_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: custom_property_name"), nil
		}
		custom_property_name, ok := custom_property_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: custom_property_name"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.GeneratedType
		
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
		url := fmt.Sprintf("%s/orgs/%s/properties/schema/%s", cfg.BaseURL, org, custom_property_name)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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

func CreateOrgs_create_or_update_custom_propertyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_orgs_org_properties_schema_custom_property_name",
		mcp.WithDescription("Create or update a custom property for an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithString("custom_property_name", mcp.Required(), mcp.Description("The custom property name")),
		mcp.WithString("values_editable_by", mcp.Description("Input parameter: Who can edit the values of the property")),
		mcp.WithArray("allowed_values", mcp.Description("Input parameter: An ordered list of the allowed values of the property.\nThe property can have up to 200 allowed values.")),
		mcp.WithString("default_value", mcp.Description("Input parameter: Default value of the property")),
		mcp.WithString("description", mcp.Description("Input parameter: Short description of the property")),
		mcp.WithBoolean("required", mcp.Description("Input parameter: Whether the property is required.")),
		mcp.WithString("value_type", mcp.Required(), mcp.Description("Input parameter: The type of the value for the property")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_create_or_update_custom_propertyHandler(cfg),
	}
}
