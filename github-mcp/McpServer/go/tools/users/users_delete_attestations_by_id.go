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

func Users_delete_attestations_by_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		usernameVal, ok := args["username"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: username"), nil
		}
		username, ok := usernameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: username"), nil
		}
		attestation_idVal, ok := args["attestation_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: attestation_id"), nil
		}
		attestation_id, ok := attestation_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: attestation_id"), nil
		}
		url := fmt.Sprintf("%s/users/%s/attestations/%s", cfg.BaseURL, username, attestation_id)
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

func CreateUsers_delete_attestations_by_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_users_username_attestations_attestation_id",
		mcp.WithDescription("Delete attestations by ID"),
		mcp.WithString("username", mcp.Required(), mcp.Description("The handle for the GitHub user account.")),
		mcp.WithNumber("attestation_id", mcp.Required(), mcp.Description("Attestation ID")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Users_delete_attestations_by_idHandler(cfg),
	}
}
