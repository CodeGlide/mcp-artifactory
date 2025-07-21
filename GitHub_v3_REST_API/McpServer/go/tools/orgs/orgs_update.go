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

func Orgs_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/orgs/%s", cfg.BaseURL, org)
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

func CreateOrgs_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_orgs_org",
		mcp.WithDescription("Update an organization"),
		mcp.WithString("org", mcp.Required(), mcp.Description("The organization name. The name is not case sensitive.")),
		mcp.WithBoolean("members_can_create_public_pages", mcp.Description("Input parameter: Whether organization members can create public GitHub Pages sites. Existing published sites will not be impacted.")),
		mcp.WithString("billing_email", mcp.Description("Input parameter: Billing email address. This address is not publicized.")),
		mcp.WithBoolean("advanced_security_enabled_for_new_repositories", mcp.Description("Input parameter: **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead.\n\nWhether GitHub Advanced Security is automatically enabled for new repositories and repositories transferred to this organization.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nYou can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.")),
		mcp.WithBoolean("members_can_create_private_pages", mcp.Description("Input parameter: Whether organization members can create private GitHub Pages sites. Existing published sites will not be impacted.")),
		mcp.WithBoolean("dependabot_alerts_enabled_for_new_repositories", mcp.Description("Input parameter: **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead.\n\nWhether Dependabot alerts are automatically enabled for new repositories and repositories transferred to this organization.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nYou can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.")),
		mcp.WithBoolean("has_repository_projects", mcp.Description("Input parameter: Whether repositories that belong to the organization can use repository projects.")),
		mcp.WithString("twitter_username", mcp.Description("Input parameter: The Twitter username of the company.")),
		mcp.WithString("blog", mcp.Description("")),
		mcp.WithString("name", mcp.Description("Input parameter: The shorthand name of the company.")),
		mcp.WithBoolean("members_can_create_private_repositories", mcp.Description("Input parameter: Whether organization members can create private repositories, which are visible to organization members with permission. For more information, see \"[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)\" in the GitHub Help documentation.")),
		mcp.WithString("email", mcp.Description("Input parameter: The publicly visible email address.")),
		mcp.WithString("default_repository_permission", mcp.Description("Input parameter: Default permission level members have for organization repositories.")),
		mcp.WithBoolean("members_can_create_internal_repositories", mcp.Description("Input parameter: Whether organization members can create internal repositories, which are visible to all enterprise members. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see \"[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)\" in the GitHub Help documentation.")),
		mcp.WithBoolean("members_can_create_repositories", mcp.Description("Input parameter: Whether of non-admin organization members can create repositories. **Note:** A parameter can override this parameter. See `members_allowed_repository_creation_type` in this table for details.")),
		mcp.WithBoolean("dependabot_security_updates_enabled_for_new_repositories", mcp.Description("Input parameter: **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead.\n\nWhether Dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nYou can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.")),
		mcp.WithBoolean("dependency_graph_enabled_for_new_repositories", mcp.Description("Input parameter: **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead.\n\nWhether dependency graph is automatically enabled for new repositories and repositories transferred to this organization.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nYou can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.")),
		mcp.WithBoolean("deploy_keys_enabled_for_repositories", mcp.Description("Input parameter: Controls whether or not deploy keys may be added and used for repositories in the organization.")),
		mcp.WithBoolean("members_can_create_public_repositories", mcp.Description("Input parameter: Whether organization members can create public repositories, which are visible to anyone. For more information, see \"[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)\" in the GitHub Help documentation.")),
		mcp.WithBoolean("secret_scanning_push_protection_custom_link_enabled", mcp.Description("Input parameter: Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.")),
		mcp.WithBoolean("secret_scanning_push_protection_enabled_for_new_repositories", mcp.Description("Input parameter: **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead.\n\nWhether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nYou can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.")),
		mcp.WithString("description", mcp.Description("Input parameter: The description of the company. The maximum size is 160 characters.")),
		mcp.WithString("location", mcp.Description("Input parameter: The location.")),
		mcp.WithBoolean("secret_scanning_enabled_for_new_repositories", mcp.Description("Input parameter: **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead.\n\nWhether secret scanning is automatically enabled for new repositories and repositories transferred to this organization.\n\nTo use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see \"[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization).\"\n\nYou can check which security and analysis features are currently enabled by using a `GET /orgs/{org}` request.")),
		mcp.WithString("company", mcp.Description("Input parameter: The company name.")),
		mcp.WithBoolean("members_can_create_pages", mcp.Description("Input parameter: Whether organization members can create GitHub Pages sites. Existing published sites will not be impacted.")),
		mcp.WithBoolean("has_organization_projects", mcp.Description("Input parameter: Whether an organization can use organization projects.")),
		mcp.WithBoolean("web_commit_signoff_required", mcp.Description("Input parameter: Whether contributors to organization repositories are required to sign off on commits they make through GitHub's web interface.")),
		mcp.WithString("members_allowed_repository_creation_type", mcp.Description("Input parameter: Specifies which types of repositories non-admin organization members can create. `private` is only available to repositories that are part of an organization on GitHub Enterprise Cloud. \n**Note:** This parameter is closing down and will be removed in the future. Its return value ignores internal repositories. Using this parameter overrides values set in `members_can_create_repositories`. See the parameter deprecation notice in the operation description for details.")),
		mcp.WithBoolean("members_can_fork_private_repositories", mcp.Description("Input parameter: Whether organization members can fork private organization repositories.")),
		mcp.WithString("secret_scanning_push_protection_custom_link", mcp.Description("Input parameter: If `secret_scanning_push_protection_custom_link_enabled` is true, the URL that will be displayed to contributors who are blocked from pushing a secret.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Orgs_updateHandler(cfg),
	}
}
