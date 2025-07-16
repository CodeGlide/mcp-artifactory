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

func Repos_create_for_authenticated_userHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		url := fmt.Sprintf("%s/user/repos", cfg.BaseURL)
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

func CreateRepos_create_for_authenticated_userTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_user_repos",
		mcp.WithDescription("Create a repository for the authenticated user"),
		mcp.WithString("license_template", mcp.Description("Input parameter: The license keyword of the open source license for this repository.")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name of the repository.")),
		mcp.WithString("gitignore_template", mcp.Description("Input parameter: The desired language or platform to apply to the .gitignore.")),
		mcp.WithString("homepage", mcp.Description("Input parameter: A URL with more information about the repository.")),
		mcp.WithBoolean("is_template", mcp.Description("Input parameter: Whether this repository acts as a template that can be used to generate new repositories.")),
		mcp.WithBoolean("private", mcp.Description("Input parameter: Whether the repository is private.")),
		mcp.WithString("description", mcp.Description("Input parameter: A short description of the repository.")),
		mcp.WithString("merge_commit_title", mcp.Description("Input parameter: Required when using `merge_commit_message`.\n\nThe default value for a merge commit title.\n\n- `PR_TITLE` - default to the pull request's title.\n- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).")),
		mcp.WithBoolean("has_projects", mcp.Description("Input parameter: Whether projects are enabled.")),
		mcp.WithNumber("team_id", mcp.Description("Input parameter: The id of the team that will be granted access to this repository. This is only valid when creating a repository in an organization.")),
		mcp.WithString("squash_merge_commit_title", mcp.Description("Input parameter: Required when using `squash_merge_commit_message`.\n\nThe default value for a squash merge commit title:\n\n- `PR_TITLE` - default to the pull request's title.\n- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).")),
		mcp.WithBoolean("auto_init", mcp.Description("Input parameter: Whether the repository is initialized with a minimal README.")),
		mcp.WithBoolean("allow_merge_commit", mcp.Description("Input parameter: Whether to allow merge commits for pull requests.")),
		mcp.WithBoolean("allow_rebase_merge", mcp.Description("Input parameter: Whether to allow rebase merges for pull requests.")),
		mcp.WithString("merge_commit_message", mcp.Description("Input parameter: The default value for a merge commit message.\n\n- `PR_TITLE` - default to the pull request's title.\n- `PR_BODY` - default to the pull request's body.\n- `BLANK` - default to a blank commit message.")),
		mcp.WithBoolean("allow_squash_merge", mcp.Description("Input parameter: Whether to allow squash merges for pull requests.")),
		mcp.WithBoolean("has_discussions", mcp.Description("Input parameter: Whether discussions are enabled.")),
		mcp.WithBoolean("allow_auto_merge", mcp.Description("Input parameter: Whether to allow Auto-merge to be used on pull requests.")),
		mcp.WithBoolean("has_wiki", mcp.Description("Input parameter: Whether the wiki is enabled.")),
		mcp.WithString("squash_merge_commit_message", mcp.Description("Input parameter: The default value for a squash merge commit message:\n\n- `PR_BODY` - default to the pull request's body.\n- `COMMIT_MESSAGES` - default to the branch's commit messages.\n- `BLANK` - default to a blank commit message.")),
		mcp.WithBoolean("delete_branch_on_merge", mcp.Description("Input parameter: Whether to delete head branches when pull requests are merged")),
		mcp.WithBoolean("has_downloads", mcp.Description("Input parameter: Whether downloads are enabled.")),
		mcp.WithBoolean("has_issues", mcp.Description("Input parameter: Whether issues are enabled.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Repos_create_for_authenticated_userHandler(cfg),
	}
}
