package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Packages_delete_package_version_for_orgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		package_typeVal, ok := args["package_type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: package_type"), nil
		}
		package_type, ok := package_typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: package_type"), nil
		}
		package_nameVal, ok := args["package_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: package_name"), nil
		}
		package_name, ok := package_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: package_name"), nil
		}
		orgVal, ok := args["org"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: org"), nil
		}
		org, ok := orgVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: org"), nil
		}
		package_version_idVal, ok := args["package_version_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: package_version_id"), nil
		}
		package_version_id, ok := package_version_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: package_version_id"), nil
		}
		url := fmt.Sprintf("%s/orgs/%s/packages/%s/%s/versions/%s", cfg.BaseURL, package_type, package_name, org, package_version_id)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreatePackages_delete_package_version_for_orgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_orgs_org_packages_package_type_package_name_versions_package_version_id",
		mcp.WithDescription("Delete package version for an organization"),
		mcp.WithString("package_type", mcp.Required(), mcp.Description("The type of supported package. Packages in GitHub's Gradle registry have the type `maven`. Docker images pushed to GitHub's Container registry (`ghcr.io`) have the type `container`. You can use the type `docker` to find images that were pushed to GitHub's Docker registry (`docker.pkg.github.com`), even if these have now been migrated to the Container registry.")),
		mcp.WithString("package_name", mcp.Required(), mcp.Description("The name of the package.")),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithNumber("package_version_id", mcp.Required(), mcp.Description("Unique identifier of the package version.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Packages_delete_package_version_for_orgHandler(cfg),
	}
}
