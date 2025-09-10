/**
 * Show Link
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

export async function get_api_services_jira_links_link_id(Accept, link_id) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (Accept) params.append("Accept", Accept);
      if (link_id) params.append("link_id", link_id);
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

export function createGetApiServicesJiraLinksLinkIdTool() {
  return {
    definition: {
      name: 'get-api-services-jira-links-link-id',
      description: 'Show Link',
      inputSchema: {
        type: 'object',
        properties: {
          Accept: {
            type: 'string',
            description: ''
          },
          link_id: {
            type: 'string',
            description: '(Required) The id of the link'
          }
        },
        required: ["link_id"]
      }
    },
    handler: get_api_services_jira_links_link_id
  };
}