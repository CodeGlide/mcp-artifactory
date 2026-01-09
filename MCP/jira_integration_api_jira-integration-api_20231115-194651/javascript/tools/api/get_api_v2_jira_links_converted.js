/**
 * List Links
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

export async function get_api_v2_jira_links(Accept, page[after], filter[ticket_id], page[size]) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (Accept) params.append("Accept", Accept);
      if (page[after]) params.append("page[after]", page[after]);
      if (filter[ticket_id]) params.append("filter[ticket_id]", filter[ticket_id]);
      if (page[size]) params.append("page[size]", page[size]);
    const queryString = params.toString();
    const finalUrl = queryString ? `${url}?${queryString}` : url;
    
    const url = `${config.baseURL}/%s/api/v2/jira/links%s`;
    
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

export function createGetApiV2JiraLinksTool() {
  return {
    definition: {
      name: 'get-api-v2-jira-links',
      description: 'List Links',
      inputSchema: {
        type: 'object',
        properties: {
          Accept: {
            type: 'string',
            description: ''
          },
          page[after]: {
            type: 'string',
            description: 'When provided, the returned paginated data must have as its first item the item that is immediately after the cursor in the results list. Exception: If there are no items in the results list that fall after the cursor, the returned paginated data must be an empty array'
          },
          filter[ticket_id]: {
            type: 'string',
            description: 'List links for a specific Zendesk ticket or Jira issue by specifying a ticket id or issue id. Filtering by issue key is not currently supported'
          },
          page[size]: {
            type: 'number',
            description: 'The number of entries that will be returned'
          }
        },
        required: []
      }
    },
    handler: get_api_v2_jira_links
  };
}