package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/openai/mcp-server/config"
	"github.com/openai/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_fine_tunes_fine_tune_id_cancelHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		fine_tune_idVal, ok := args["fine_tune_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: fine_tune_id"), nil
		}
		fine_tune_id, ok := fine_tune_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: fine_tune_id"), nil
		}
		// Fallback to map[string]any for untyped request body
		bodyMap := map[string]any{}
		bodyBytes, err := json.Marshal(bodyMap)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/fine-tunes/%s/cancel", cfg.BaseURL, fine_tune_id)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication header based on auth type
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
		}
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

func CreatePost_fine_tunes_fine_tune_id_cancelTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_fine-tunes_fine_tune_id_cancel",
		mcp.WithDescription("Cancel fine-tune"),
		mcp.WithString("fine_tune_id", mcp.Required(), mcp.Description("The ID of the fine-tune job.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_fine_tunes_fine_tune_id_cancelHandler(cfg),
	}
}
