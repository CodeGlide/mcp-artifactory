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

func Code_scanning_get_variant_analysis_repo_taskHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ownerVal, ok := args["owner"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: owner"), nil
		}
		owner, ok := ownerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: owner"), nil
		}
		repoVal, ok := args["repo"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo"), nil
		}
		repo, ok := repoVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo"), nil
		}
		codeql_variant_analysis_idVal, ok := args["codeql_variant_analysis_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: codeql_variant_analysis_id"), nil
		}
		codeql_variant_analysis_id, ok := codeql_variant_analysis_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: codeql_variant_analysis_id"), nil
		}
		repo_ownerVal, ok := args["repo_owner"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo_owner"), nil
		}
		repo_owner, ok := repo_ownerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo_owner"), nil
		}
		repo_nameVal, ok := args["repo_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repo_name"), nil
		}
		repo_name, ok := repo_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repo_name"), nil
		}
		url := fmt.Sprintf("%s/repos/%s/%s/code-scanning/codeql/variant-analyses/%s/repos/%s/%s", cfg.BaseURL, owner, repo, codeql_variant_analysis_id, repo_owner, repo_name)
		req, err := http.NewRequest("GET", url, nil)
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

func CreateCode_scanning_get_variant_analysis_repo_taskTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_repos_owner_repo_code-scanning_codeql_variant-analyses_codeql_variant_analysis_id_repos_repo_owner_repo_name",
		mcp.WithDescription("Get the analysis status of a repository in a CodeQL variant analysis"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the controller repository.")),
		mcp.WithNumber("codeql_variant_analysis_id", mcp.Required(), mcp.Description("The ID of the variant analysis.")),
		mcp.WithString("repo_owner", mcp.Required(), mcp.Description("The account owner of the variant analysis repository. The name is not case sensitive.")),
		mcp.WithString("repo_name", mcp.Required(), mcp.Description("The name of the variant analysis repository.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Code_scanning_get_variant_analysis_repo_taskHandler(cfg),
	}
}
