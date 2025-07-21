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

func Repos_create_in_orgHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/orgs/%s/repos", cfg.BaseURL, org)
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

func CreateRepos_create_in_orgTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_orgs_org_repos",
		mcp.WithDescription("Create an organization repository"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithBoolean("allow_auto_merge", mcp.Description("Input parameter: Either `true` to allow auto-merge on pull requests, or `false` to disallow auto-merge.")),
		mcp.WithBoolean("is_template", mcp.Description("Input parameter: Either `true` to make this repo available as a template repository or `false` to prevent it.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the repository.")),
		mcp.WithBoolean("use_squash_pr_title_as_default", mcp.Description("Input parameter: Either `true` to allow squash-merge commits to use pull request title, or `false` to use commit message. **This property is closing down. Please use `squash_merge_commit_title` instead.")),
		mcp.WithBoolean("auto_init", mcp.Description("Input parameter: Pass `true` to create an initial commit with empty README.")),
		mcp.WithString("squash_merge_commit_title", mcp.Description("Input parameter: Required when using `squash_merge_commit_message`.\n\nThe default value for a squash merge commit title:\n\n- `PR_TITLE` - default to the pull request's title.\n- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).")),
		mcp.WithObject("custom_properties", mcp.Description("Input parameter: The custom properties for the new repository. The keys are the custom property names, and the values are the corresponding custom property values.")),
		mcp.WithString("description", mcp.Description("Input parameter: A short description of the repository.")),
		mcp.WithBoolean("private", mcp.Description("Input parameter: Whether the repository is private.")),
		mcp.WithString("homepage", mcp.Description("Input parameter: A URL with more information about the repository.")),
		mcp.WithNumber("team_id", mcp.Description("Input parameter: The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.")),
		mcp.WithBoolean("has_issues", mcp.Description("Input parameter: Either `true` to enable issues for this repository or `false` to disable them.")),
		mcp.WithBoolean("allow_merge_commit", mcp.Description("Input parameter: Either `true` to allow merging pull requests with a merge commit, or `false` to prevent merging pull requests with merge commits.")),
		mcp.WithBoolean("allow_squash_merge", mcp.Description("Input parameter: Either `true` to allow squash-merging pull requests, or `false` to prevent squash-merging.")),
		mcp.WithString("gitignore_template", mcp.Description("Input parameter: Desired language or platform [.gitignore template](https://github.com/github/gitignore) to apply. Use the name of the template without the extension. For example, \"Haskell\".")),
		mcp.WithBoolean("has_downloads", mcp.Description("Input parameter: Whether downloads are enabled.")),
		mcp.WithString("visibility", mcp.Description("Input parameter: The visibility of the repository.")),
		mcp.WithBoolean("delete_branch_on_merge", mcp.Description("Input parameter: Either `true` to allow automatically deleting head branches when pull requests are merged, or `false` to prevent automatic deletion. **The authenticated user must be an organization owner to set this property to `true`.**")),
		mcp.WithBoolean("has_projects", mcp.Description("Input parameter: Either `true` to enable projects for this repository or `false` to disable them. **Note:** If you're creating a repository in an organization that has disabled repository projects, the default is `false`, and if you pass `true`, the API returns an error.")),
		mcp.WithString("squash_merge_commit_message", mcp.Description("Input parameter: The default value for a squash merge commit message:\n\n- `PR_BODY` - default to the pull request's body.\n- `COMMIT_MESSAGES` - default to the branch's commit messages.\n- `BLANK` - default to a blank commit message.")),
		mcp.WithBoolean("allow_rebase_merge", mcp.Description("Input parameter: Either `true` to allow rebase-merging pull requests, or `false` to prevent rebase-merging.")),
		mcp.WithBoolean("has_wiki", mcp.Description("Input parameter: Either `true` to enable the wiki for this repository or `false` to disable it.")),
		mcp.WithString("license_template", mcp.Description("Input parameter: Choose an [open source license template](https://choosealicense.com/) that best suits your needs, and then use the [license keyword](https://docs.github.com/articles/licensing-a-repository/#searching-github-by-license-type) as the `license_template` string. For example, \"mit\" or \"mpl-2.0\".")),
		mcp.WithString("merge_commit_message", mcp.Description("Input parameter: The default value for a merge commit message.\n\n- `PR_TITLE` - default to the pull request's title.\n- `PR_BODY` - default to the pull request's body.\n- `BLANK` - default to a blank commit message.")),
		mcp.WithString("merge_commit_title", mcp.Description("Input parameter: Required when using `merge_commit_message`.\n\nThe default value for a merge commit title.\n\n- `PR_TITLE` - default to the pull request's title.\n- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repos_create_in_orgHandler(cfg),
	}
}
