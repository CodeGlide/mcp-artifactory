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

func Chat_postmessageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/chat.postMessage", cfg.BaseURL)
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

func CreateChat_postmessageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_chat.postMessage",
		mcp.WithDescription(""),
		mcp.WithString("token", mcp.Required(), mcp.Description("Authentication token. Requires scope: `chat:write`")),
		mcp.WithString("blocks", mcp.Description("Input parameter: A JSON-based array of structured blocks, presented as a URL-encoded string.")),
		mcp.WithString("channel", mcp.Required(), mcp.Description("Input parameter: Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. See [below](#channels) for more details.")),
		mcp.WithBoolean("link_names", mcp.Description("Input parameter: Find and link channel names and usernames.")),
		mcp.WithString("thread_ts", mcp.Description("Input parameter: Provide another message's `ts` value to make this message a reply. Avoid using a reply's `ts` value; use its parent instead.")),
		mcp.WithString("parse", mcp.Description("Input parameter: Change how messages are treated. Defaults to `none`. See [below](#formatting).")),
		mcp.WithBoolean("mrkdwn", mcp.Description("Input parameter: Disable Slack markup parsing by setting to `false`. Enabled by default.")),
		mcp.WithString("as_user", mcp.Description("Input parameter: Pass true to post the message as the authed user, instead of as a bot. Defaults to false. See [authorship](#authorship) below.")),
		mcp.WithString("text", mcp.Description("Input parameter: How this field works and whether it is required depends on other fields you use in your API call. [See below](#text_usage) for more detail.")),
		mcp.WithBoolean("unfurl_links", mcp.Description("Input parameter: Pass true to enable unfurling of primarily text-based content.")),
		mcp.WithString("icon_emoji", mcp.Description("Input parameter: Emoji to use as the icon for this message. Overrides `icon_url`. Must be used in conjunction with `as_user` set to `false`, otherwise ignored. See [authorship](#authorship) below.")),
		mcp.WithBoolean("unfurl_media", mcp.Description("Input parameter: Pass false to disable unfurling of media content.")),
		mcp.WithBoolean("reply_broadcast", mcp.Description("Input parameter: Used in conjunction with `thread_ts` and indicates whether reply should be made visible to everyone in the channel or conversation. Defaults to `false`.")),
		mcp.WithString("username", mcp.Description("Input parameter: Set your bot's user name. Must be used in conjunction with `as_user` set to false, otherwise ignored. See [authorship](#authorship) below.")),
		mcp.WithString("attachments", mcp.Description("Input parameter: A JSON-based array of structured attachments, presented as a URL-encoded string.")),
		mcp.WithString("icon_url", mcp.Description("Input parameter: URL to an image to use as the icon for this message. Must be used in conjunction with `as_user` set to false, otherwise ignored. See [authorship](#authorship) below.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Chat_postmessageHandler(cfg),
	}
}
