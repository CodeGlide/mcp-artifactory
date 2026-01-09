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

func Get_locations_location_id_media_recentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		location_idVal, ok := args["location-id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: location-id"), nil
		}
		location_id, ok := location_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: location-id"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["min_timestamp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("min_timestamp=%v", val))
		}
		if val, ok := args["max_timestamp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("max_timestamp=%v", val))
		}
		if val, ok := args["min_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("min_id=%v", val))
		}
		if val, ok := args["max_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("max_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/locations/%s/media/recent%s", cfg.BaseURL, location_id, queryString)
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

func CreateGet_locations_location_id_media_recentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_locations_location-id_media_recent",
		mcp.WithDescription("Get a list of recent media objects from a given location."),
		mcp.WithString("min_timestamp", mcp.Description("Return media after this UNIX timestamp.")),
		mcp.WithString("max_timestamp", mcp.Description("Return media before this UNIX timestamp.")),
		mcp.WithString("min_id", mcp.Description("Return media before this `min_id`.")),
		mcp.WithString("max_id", mcp.Description("Return media after this `max_id`.")),
		mcp.WithString("location-id", mcp.Required(), mcp.Description("The location ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_locations_location_id_media_recentHandler(cfg),
	}
}
