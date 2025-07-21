package main

import (
	"github.com/input-api/mcp-server/config"
	"github.com/input-api/mcp-server/models"
	tools_config "github.com/input-api/mcp-server/tools/config"
	tools_auth "github.com/input-api/mcp-server/tools/auth"
	tools_api "github.com/input-api/mcp-server/tools/api"
	tools_mcp "github.com/input-api/mcp-server/tools/mcp"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_config.CreateHead_api_base_urlTool(cfg),
		tools_auth.CreateHead_api_keyTool(cfg),
		tools_auth.CreateHead_basic_authTool(cfg),
		tools_auth.CreateHead_bearer_tokenTool(cfg),
		tools_api.CreateGetTool(cfg),
		tools_mcp.CreateGet_mcpTool(cfg),
	}
}
