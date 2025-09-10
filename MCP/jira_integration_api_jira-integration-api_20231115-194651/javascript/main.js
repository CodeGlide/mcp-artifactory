/**
 * MCP Server - JavaScript Implementation
 */

import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import { ListToolsRequestSchema, CallToolRequestSchema } from '@modelcontextprotocol/sdk/types.js';
import fs from 'fs';
import path from 'path';
import os from 'os';

import { get_api_services_jira_links, createGetApiServicesJiraLinksTool } from './tools/api_services_jira_links/get_api_services_jira_links_converted.js';
import { post_api_services_jira_links, createPostApiServicesJiraLinksTool } from './tools/api_services_jira_links/post_api_services_jira_links_converted.js';
import { delete_api_services_jira_links_link_id, createDeleteApiServicesJiraLinksLinkIdTool } from './tools/api_services_jira_links_link_id/delete_api_services_jira_links_link_id_converted.js';
import { get_api_services_jira_links_link_id, createGetApiServicesJiraLinksLinkIdTool } from './tools/api_services_jira_links_link_id/get_api_services_jira_links_link_id_converted.js';
import { get_api_v2_jira_links, createGetApiV2JiraLinksTool } from './tools/api/get_api_v2_jira_links_converted.js';

// Create MCP server
const server = new Server({
  name: 'MCP Server',
  version: '1.0.0'
}, {
  capabilities: {
    tools: {}
  }
});

function getConfig() {
  const baseURL = process.env.API_BASE_URL;
  const bearerToken = process.env.API_BEARER_TOKEN;
  
  if (!baseURL || !bearerToken) {
    const configPath = path.join(os.homedir(), '.api', 'config.json');
    try {
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      return {
        baseURL: baseURL || configData.baseURL,
        bearerToken: bearerToken || configData.bearerToken
      };
    } catch (e) {
      throw new Error('Configuration not found. Please set API_BASE_URL and API_BEARER_TOKEN environment variables or create config file at ~/.api/config.json');
    }
  }
  
  return { baseURL, bearerToken };
}

// Register all tools
const tools = [
  createGetApiServicesJiraLinksTool(),
  createPostApiServicesJiraLinksTool(),
  createDeleteApiServicesJiraLinksLinkIdTool(),
  createGetApiServicesJiraLinksLinkIdTool(),
  createGetApiV2JiraLinksTool()
];

// List all available tools
function listToolsHandler() {
  return { tools: tools.map(tool => tool.definition) };
}

// Handle tool calls
function createCallToolHandler(toolMap) {
  return async (request) => {
    const { name, arguments: args } = request.params;
    
    const tool = toolMap.find(t => t.definition.name === name);
    if (!tool) {
      throw new Error(`Unknown tool: ${name}`);
    }
    
    try {
      const result = await tool.handler(args);
      return {
        content: [{
          type: 'text',
          text: result
        }]
      };
    } catch (error) {
      throw new Error(`Tool execution failed: ${error.message}`);
    }
  };
}

// Setup request handlers
server.setRequestHandler(ListToolsRequestSchema, listToolsHandler);
server.setRequestHandler(CallToolRequestSchema, createCallToolHandler(tools));

async function main() {
  try {
    const config = getConfig();
    console.error('MCP Server started successfully');
    
    const transport = new StdioServerTransport();
    await server.connect(transport);
  } catch (error) {
    console.error('Failed to start server:', error);
    process.exit(1);
  }
}

if (import.meta.url === `file://${process.argv[1]}`) {
  main().catch(console.error);
}