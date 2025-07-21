package main

import (
	"github.com/square-connect-api/mcp-server/config"
	"github.com/square-connect-api/mcp-server/models"
	tools_customers "github.com/square-connect-api/mcp-server/tools/customers"
	tools_orders "github.com/square-connect-api/mcp-server/tools/orders"
	tools_webhooks "github.com/square-connect-api/mcp-server/tools/webhooks"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_customers.CreateRetrievecustomerTool(cfg),
		tools_orders.CreateRetrieveorderTool(cfg),
		tools_webhooks.CreateCreatewebhooksubscriptionTool(cfg),
	}
}
