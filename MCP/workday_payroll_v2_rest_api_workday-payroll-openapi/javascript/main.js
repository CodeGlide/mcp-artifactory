/**
 * MCP Server - JavaScript Implementation
 */

import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import { ListToolsRequestSchema, CallToolRequestSchema } from '@modelcontextprotocol/sdk/types.js';
import fs from 'fs';
import path from 'path';
import os from 'os';

import { get_pay_groups_id, createGetPayGroupsIdTool } from './tools/paygroups/get_paygroups_id_converted.js';
import { get_pay_groups, createGetPayGroupsTool } from './tools/paygroups/get_paygroups_converted.js';
import { get_tax_rates, createGetTaxRatesTool } from './tools/taxrates/get_taxrates_converted.js';
import { post_tax_rates, createPostTaxRatesTool } from './tools/taxrates/post_taxrates_converted.js';
import { get_jobs_id_pay_group_subresource_id, createGetJobsIdPayGroupSubresourceIdTool } from './tools/jobs/get_jobs_id_paygroup_subresourceid_converted.js';
import { get_jobs_id, createGetJobsIdTool } from './tools/jobs/get_jobs_id_converted.js';
import { get_jobs_id_pay_group, createGetJobsIdPayGroupTool } from './tools/jobs/get_jobs_id_paygroup_converted.js';
import { get_jobs, createGetJobsTool } from './tools/jobs/get_jobs_converted.js';
import { delete_payroll_inputs_id, createDeletePayrollInputsIdTool } from './tools/payrollinputs/delete_payrollinputs_id_converted.js';
import { get_payroll_inputs_id, createGetPayrollInputsIdTool } from './tools/payrollinputs/get_payrollinputs_id_converted.js';
import { get_payroll_inputs, createGetPayrollInputsTool } from './tools/payrollinputs/get_payrollinputs_converted.js';
import { patch_payroll_inputs_id, createPatchPayrollInputsIdTool } from './tools/payrollinputs/patch_payrollinputs_id_converted.js';
import { post_payroll_inputs, createPostPayrollInputsTool } from './tools/payrollinputs/post_payrollinputs_converted.js';
import { get_pay_group_details_id, createGetPayGroupDetailsIdTool } from './tools/paygroupdetails/get_paygroupdetails_id_converted.js';
import { get_pay_group_details, createGetPayGroupDetailsTool } from './tools/paygroupdetails/get_paygroupdetails_converted.js';
import { get_values_payroll_inputs_group_run_categories, createGetValuesPayrollInputsGroupRunCategoriesTool } from './tools/prompt_values/get_values_payrollinputsgroup_runcategories_converted.js';
import { get_values_payroll_inputs_group_pay_components, createGetValuesPayrollInputsGroupPayComponentsTool } from './tools/prompt_values/get_values_payrollinputsgroup_paycomponents_converted.js';
import { get_values_tax_rates_group_state_instances, createGetValuesTaxRatesGroupStateInstancesTool } from './tools/prompt_values/get_values_taxratesgroup_stateinstances_converted.js';
import { get_values_payroll_inputs_group_worktags, createGetValuesPayrollInputsGroupWorktagsTool } from './tools/prompt_values/get_values_payrollinputsgroup_worktags_converted.js';
import { get_values_payroll_inputs_group_positions, createGetValuesPayrollInputsGroupPositionsTool } from './tools/prompt_values/get_values_payrollinputsgroup_positions_converted.js';
import { get_values_tax_rates_group_company_instances, createGetValuesTaxRatesGroupCompanyInstancesTool } from './tools/prompt_values/get_values_taxratesgroup_companyinstances_converted.js';

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
  createGetPayGroupsIdTool(),
  createGetPayGroupsTool(),
  createGetTaxRatesTool(),
  createPostTaxRatesTool(),
  createGetJobsIdPayGroupSubresourceIdTool(),
  createGetJobsIdTool(),
  createGetJobsIdPayGroupTool(),
  createGetJobsTool(),
  createDeletePayrollInputsIdTool(),
  createGetPayrollInputsIdTool(),
  createGetPayrollInputsTool(),
  createPatchPayrollInputsIdTool(),
  createPostPayrollInputsTool(),
  createGetPayGroupDetailsIdTool(),
  createGetPayGroupDetailsTool(),
  createGetValuesPayrollInputsGroupRunCategoriesTool(),
  createGetValuesPayrollInputsGroupPayComponentsTool(),
  createGetValuesTaxRatesGroupStateInstancesTool(),
  createGetValuesPayrollInputsGroupWorktagsTool(),
  createGetValuesPayrollInputsGroupPositionsTool(),
  createGetValuesTaxRatesGroupCompanyInstancesTool()
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