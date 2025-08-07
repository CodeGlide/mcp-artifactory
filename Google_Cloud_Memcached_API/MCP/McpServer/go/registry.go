package main

import (
	"github.com/cloud-memorystore-for-memcached-api/mcp-server/config"
	"github.com/cloud-memorystore-for-memcached-api/mcp-server/models"
	tools_projects "github.com/cloud-memorystore-for-memcached-api/mcp-server/tools/projects"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_projects.CreateMemcache_projects_locations_operations_cancelTool(cfg),
		tools_projects.CreateMemcache_projects_locations_instances_updateparametersTool(cfg),
		tools_projects.CreateMemcache_projects_locations_instances_listTool(cfg),
		tools_projects.CreateMemcache_projects_locations_instances_createTool(cfg),
		tools_projects.CreateMemcache_projects_locations_instances_reschedulemaintenanceTool(cfg),
		tools_projects.CreateMemcache_projects_locations_operations_deleteTool(cfg),
		tools_projects.CreateMemcache_projects_locations_operations_getTool(cfg),
		tools_projects.CreateMemcache_projects_locations_instances_patchTool(cfg),
		tools_projects.CreateMemcache_projects_locations_listTool(cfg),
		tools_projects.CreateMemcache_projects_locations_operations_listTool(cfg),
		tools_projects.CreateMemcache_projects_locations_instances_applyparametersTool(cfg),
	}
}
