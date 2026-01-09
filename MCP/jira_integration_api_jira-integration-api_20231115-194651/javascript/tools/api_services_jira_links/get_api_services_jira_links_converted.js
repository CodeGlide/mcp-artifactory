/**
 * List Links (deprecated)
 */

import fs from 'fs';
import path from 'path';
import os from 'os';

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

export async function get_api_services_jira_links(Accept, since_id, ticket_id, limit) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (Accept) params.append("Accept", Accept);
      if (since_id) params.append("since_id", since_id);
      if (ticket_id) params.append("ticket_id", ticket_id);
      if (limit) params.append("limit", limit);
    const queryString = params.toString();
    const finalUrl = queryString ? `${url}?${queryString}` : url;
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(finalUrl, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json'
      }
    });
    
    if (!response.ok) {
      return `Failed to format JSON: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Request failed: ${error.message}`;
  }
}

export function createGetApiServicesJiraLinksTool() {
  return {
    definition: {
      name: 'get-api-services-jira-links',
      description: 'List Links (deprecated)',
      inputSchema: {
        type: 'object',
        properties: {
          Accept: {
            type: 'string',
            description: ''
          },
          since_id: {
            type: 'string',
            description: 'The start id of the collection'
          },
          ticket_id: {
            type: 'string',
            description: 'List links for a specific Zendesk Ticket or Jira issue by providing `ticket_id` and/or `issue_id` param. We currently do not support `issue_key` param.'
          },
          limit: {
            type: 'number',
            description: 'The number of entries that will be returned'
          }
        },
        required: []
      }
    },
    handler: get_api_services_jira_links
  };
}