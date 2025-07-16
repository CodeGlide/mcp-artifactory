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

func Repos_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/repos/%s/%s", cfg.BaseURL, owner, repo)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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

func CreateRepos_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_repos_owner_repo",
		mcp.WithDescription("Update a repository"),
		mcp.WithString("owner", mcp.Required(), mcp.Description("The account owner of the repository. The name is not case sensitive.")),
		mcp.WithString("repo", mcp.Required(), mcp.Description("The name of the repository without the `.git` extension. The name is not case sensitive.")),
		mcp.WithBoolean("web_commit_signoff_required", mcp.Description("Input parameter: Either `true` to require contributors to sign off on web-based commits, or `false` to not require contributors to sign off on web-based commits.")),
		mcp.WithBoolean("is_template", mcp.Description("Input parameter: Either `true` to make this repo available as a template repository or `false` to prevent it.")),
		mcp.WithString("merge_commit_message", mcp.Description("Input parameter: The default value for a merge commit message.\n\n- `PR_TITLE` - default to the pull request's title.\n- `PR_BODY` - default to the pull request's body.\n- `BLANK` - default to a blank commit message.")),
		mcp.WithString("merge_commit_title", mcp.Description("Input parameter: Required when using `merge_commit_message`.\n\nThe default value for a merge commit title.\n\n- `PR_TITLE` - default to the pull request's title.\n- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).")),
		mcp.WithBoolean("use_squash_pr_title_as_default", mcp.Description("Input parameter: Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property is closing down. Please use `squash_merge_commit_title` instead.")),
		mcp.WithBoolean("private", mcp.Description("Input parameter: Either `true` to make the repository private or `false` to make it public. Default: `false`.  \n**Note**: You will get a `422` error if the organization restricts [changing repository visibility](https://docs.github.com/articles/repository-permission-levels-for-an-organization#changing-the-visibility-of-repositories) to organization owners and a non-owner tries to change the value of private.")),
		mcp.WithObject("security_and_analysis", mcp.Description("Input parameter: Specify which security and analysis features to enable or disable for the repository.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nFor example, to enable GitHub Advanced Security, use this data in the body of the `PATCH` request:\n`{ \"security_and_analysis\": {\"advanced_security\": { \"status\": \"enabled\" } } }`.\n\nYou can check which security and analysis features are currently enabled by using a `GET /repos/{owner}/{repo}` request.")),
		mcp.WithBoolean("allow_squash_merge", mcp.Description("Input parameter: Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.")),
		mcp.WithString("description", mcp.Description("Input parameter: A short description of the repository.")),
		mcp.WithString("name", mcp.Description("Input parameter: The name of the repository.")),
		mcp.WithString("squash_merge_commit_title", mcp.Description("Input parameter: Required when using `squash_merge_commit_message`.\n\nThe default value for a squash merge commit title:\n\n- `PR_TITLE` - default to the pull request's title.\n- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).")),
		mcp.WithString("visibility", mcp.Description("Input parameter: The visibility of the repository.")),
		mcp.WithBoolean("allow_merge_commit", mcp.Description("Input parameter: Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.")),
		mcp.WithBoolean("allow_rebase_merge", mcp.Description("Input parameter: Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.")),
		mcp.WithBoolean("has_wiki", mcp.Description("Input parameter: Either `true` to enable the wiki for this repository or `false` to disable it.")),
		mcp.WithBoolean("has_issues", mcp.Description("Input parameter: Either `true` to enable issues for this repository or `false` to disable them.")),
		mcp.WithBoolean("has_projects", mcp.Description("Input parameter: Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.")),
		mcp.WithString("default_branch", mcp.Description("Input parameter: Updates the default branch for this repository.")),
		mcp.WithString("homepage", mcp.Description("Input parameter: A URL with more information about the repository.")),
		mcp.WithBoolean("delete_branch_on_merge", mcp.Description("Input parameter: Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion.")),
		mcp.WithBoolean("allow_update_branch", mcp.Description("Input parameter: Either `true` to always allow a pull request head branch that is behind its base branch to be updated even if it is not required to be up to date before merging, or false otherwise.")),
		mcp.WithBoolean("archived", mcp.Description("Input parameter: Whether to archive this repository. `false` will unarchive a previously archived repository.")),
		mcp.WithBoolean("allow_forking", mcp.Description("Input parameter: Either `true` to allow private forks, or `false` to prevent private forks.")),
		mcp.WithString("squash_merge_commit_message", mcp.Description("Input parameter: The default value for a squash merge commit message:\n\n- `PR_BODY` - default to the pull request's body.\n- `COMMIT_MESSAGES` - default to the branch's commit messages.\n- `BLANK` - default to a blank commit message.")),
		mcp.WithBoolean("allow_auto_merge", mcp.Description("Input parameter: Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repos_updateHandler(cfg),
	}
}
