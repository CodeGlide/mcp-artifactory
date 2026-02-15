/**
 * Retrieves a collection of pay group details.
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

export async function get_pay_group_details(limit, offset, runCategories) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (limit) params.append("limit", limit);
      if (offset) params.append("offset", offset);
      if (runCategories) params.append("runCategories", runCategories);
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

export function createGetPayGroupDetailsTool() {
  return {
    definition: {
      name: 'get-pay-group-details',
      description: 'Retrieves a collection of pay group details.',
      inputSchema: {
        type: 'object',
        properties: {
          limit: {
            type: 'number',
            description: 'The maximum number of objects in a single response. The default is 20. The maximum is 100.'
          },
          offset: {
            type: 'number',
            description: 'The zero-based index of the first object in a response collection. The default is 0. Use offset with the limit parameter to control paging of a response collection. Example: If limit is 5 and offset is 9, the response returns a collection of 5 objects starting with the 10th object.'
          },
          runCategories: {
            type: 'string',
            description: 'One or more Workday IDs of run categories for the pay group. You can use returned ids from GET /values/payrollInputsGroup/runCategories.You can specify 1 or more runCategories query parameters, example: runCategories=category1&runCategories=category2'
          }
        },
        required: []
      }
    },
    handler: get_pay_group_details
  };
}