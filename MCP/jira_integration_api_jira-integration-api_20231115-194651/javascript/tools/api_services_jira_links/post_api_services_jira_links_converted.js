/**
 * Create Link
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

export async function post_api_services_jira_links(Content-Type, Accept) {
  try {
    const config = getConfig();
    const requestBody = {
      Content-Type,
      Accept
    };
    
    const url = `${config.baseURL}/api/services/jira/links`;
    
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestBody)
    });
    
    if (!response.ok) {
      return `Failed to read response body: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Failed to create request: ${error.message}`;
  }
}

export function createPostApiServicesJiraLinksTool() {
  return {
    definition: {
      name: 'post-api-services-jira-links',
      description: 'Create Link',
      inputSchema: {
        type: 'object',
        properties: {
          Content-Type: {
            type: 'string',
            description: ''
          },
          Accept: {
            type: 'string',
            description: ''
          }
        },
        required: []
      }
    },
    handler: post_api_services_jira_links
  };
}