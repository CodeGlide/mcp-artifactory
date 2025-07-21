package main

import (
	"github.com/nowpayments-api/mcp-server/config"
	"github.com/nowpayments-api/mcp-server/models"
	tools_v1 "github.com/nowpayments-api/mcp-server/tools/v1"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v1.CreateGetestimatedpriceTool(cfg),
		tools_v1.CreateGettheminimumpaymentamountTool(cfg),
		tools_v1.CreateGetlistofpaymentsTool(cfg),
	}
}
