/**
 * Retrieves a single pay group instance.
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

export async function get_jobs_id_pay_group_subresource_id(ID, subresourceID) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (ID) params.append("ID", ID);
      if (subresourceID) params.append("subresourceID", subresourceID);
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

export function createGetJobsIdPayGroupSubresourceIdTool() {
  return {
    definition: {
      name: 'get-jobs-id-pay-group-subresource-id',
      description: 'Retrieves a single pay group instance.',
      inputSchema: {
        type: 'object',
        properties: {
          ID: {
            type: 'string',
            description: 'The Workday ID of the resource.'
          },
          subresourceID: {
            type: 'string',
            description: 'The Workday ID of the subresource.'
          }
        },
        required: ["ID", "subresourceID"]
      }
    },
    handler: get_jobs_id_pay_group_subresource_id
  };
}