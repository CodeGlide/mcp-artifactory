package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/openai/mcp-server/config"
	"github.com/openai/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_files_file_idHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		file_idVal, ok := args["file_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: file_id"), nil
		}
		file_id, ok := file_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: file_id"), nil
		}
		url := fmt.Sprintf("%s/files/%s", cfg.BaseURL, file_id)
		req, err := http.NewRequest("GET", url, nil)
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

func CreateGet_files_file_idTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_files_file_id",
		mcp.WithDescription("Retrieve files"),
		mcp.WithString("file_id", mcp.Required(), mcp.Description("The ID of the file to use for this request.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_files_file_idHandler(cfg),
	}
}
