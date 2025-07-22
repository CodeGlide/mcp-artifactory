package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/slack-web-api/mcp-server/config"
	"github.com/slack-web-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Chat_unfurlHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody interface{}
		
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
		url := fmt.Sprintf("%s/chat.unfurl", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["token"]; ok {
			req.Header.Set("token", fmt.Sprintf("%v", val))
		}

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

func CreateChat_unfurlTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_chat.unfurl",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Required(), mcp.Description("Authentication token. Requires scope: `links:write`")),
		mcp.WithString("channel", mcp.Required(), mcp.Description("Input parameter: Channel ID of the message")),
		mcp.WithString("ts", mcp.Required(), mcp.Description("Input parameter: Timestamp of the message to add unfurl behavior to.")),
		mcp.WithString("unfurls", mcp.Description("Input parameter: URL-encoded JSON map with keys set to URLs featured in the the message, pointing to their unfurl blocks or message attachments.")),
		mcp.WithString("user_auth_message", mcp.Description("Input parameter: Provide a simply-formatted string to send as an ephemeral message to the user as invitation to authenticate further and enable full unfurling behavior")),
		mcp.WithBoolean("user_auth_required", mcp.Description("Input parameter: Set to `true` or `1` to indicate the user must install your Slack app to trigger unfurls for this domain")),
		mcp.WithString("user_auth_url", mcp.Description("Input parameter: Send users to this custom URL where they will complete authentication in your app to fully trigger unfurling. Value should be properly URL-encoded.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Chat_unfurlHandler(cfg),
	}
}
