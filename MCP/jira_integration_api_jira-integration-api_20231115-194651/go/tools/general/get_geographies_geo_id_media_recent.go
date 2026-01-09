package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/instagram/mcp-server/config"
	"github.com/instagram/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_geographies_geo_id_media_recentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		geo_idVal, ok := args["geo-id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: geo-id"), nil
		}
		geo_id, ok := geo_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: geo-id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["min_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("min_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/geographies/%s/media/recent%s", cfg.BaseURL, geo_id, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No specific authentication scheme defined - add fallback authentication
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", "Bearer "+cfg.BearerToken)
		} else if cfg.APIKey != "" {
			req.Header.Set("Authorization", "Bearer "+cfg.APIKey)
		} else if cfg.BasicAuth != "" {
			req.Header.Set("Authorization", "Basic "+cfg.BasicAuth)
		}
		// Note: If no auth tokens provided, requests will be made without authentication
		
		// Add custom headers if provided
		
		// Set client identification headers
		req.Header.Set("X-Request-Source", "Codeglide-MCP-generator")
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

func CreateGet_geographies_geo_id_media_recentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_geographies_geo-id_media_recent",
		mcp.WithDescription("Get recent media from a custom geo-id."),
		mcp.WithString("count", mcp.Description("Max number of media to return.")),
		mcp.WithString("min_id", mcp.Description("Return media before this `min_id`.")),
		mcp.WithString("geo-id", mcp.Required(), mcp.Description("The geography ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_geographies_geo_id_media_recentHandler(cfg),
	}
}
