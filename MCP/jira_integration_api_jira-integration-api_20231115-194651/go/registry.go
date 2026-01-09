package main

import (
	"github.com/jira-integration-api/mcp-server/config"
	"github.com/jira-integration-api/mcp-server/models"
	tools_api "github.com/jira-integration-api/mcp-server/tools/api"
	tools_api_services_jira_links "github.com/jira-integration-api/mcp-server/tools/api_services_jira_links"
	tools_api_services_jira_links_link_id "github.com/jira-integration-api/mcp-server/tools/api_services_jira_links_link_id"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_api.CreateGet_api_v2_jira_linksTool(cfg),
		tools_api_services_jira_links.CreateGet_api_services_jira_linksTool(cfg),
		tools_api_services_jira_links.CreatePost_api_services_jira_linksTool(cfg),
		tools_api_services_jira_links_link_id.CreateDelete_api_services_jira_links_link_idTool(cfg),
		tools_api_services_jira_links_link_id.CreateGet_api_services_jira_links_link_idTool(cfg),
	}
}
