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

func Repos_create_using_templateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		template_ownerVal, ok := args["template_owner"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: template_owner"), nil
		}
		template_owner, ok := template_ownerVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: template_owner"), nil
		}
		template_repoVal, ok := args["template_repo"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: template_repo"), nil
		}
		template_repo, ok := template_repoVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: template_repo"), nil
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
		url := fmt.Sprintf("%s/repos/%s/%s/generate", cfg.BaseURL, template_owner, template_repo)
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

func CreateRepos_create_using_templateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_repos_template_owner_template_repo_generate",
		mcp.WithDescription("Create a repository using a template"),
		mcp.WithString("template_owner", mcp.Required(), mcp.Description("The account owner of the template repository. The name is not case sensitive.")),
		mcp.WithString("template_repo", mcp.Required(), mcp.Description("The name of the template repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithBoolean("private", mcp.Description("Input parameter: Either `true` to create a new private repository or `false` to create a new public one.")),
		mcp.WithString("description", mcp.Description("Input parameter: A short description of the new repository.")),
		mcp.WithBoolean("include_all_branches", mcp.Description("Input parameter: Set to `true` to include the directory structure and files from all branches in the template repository, and not just the default branch. Default: `false`.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the new repository.")),
		mcp.WithString("owner", mcp.Description("Input parameter: The organization or person who will own the new repository. To create a new repository in an organization, the authenticated user must be a member of the specified organization.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repos_create_using_templateHandler(cfg),
	}
}
