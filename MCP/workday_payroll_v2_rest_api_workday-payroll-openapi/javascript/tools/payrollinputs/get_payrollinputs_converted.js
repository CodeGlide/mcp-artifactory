/**
 * Retrieves a collection of payroll inputs. At least 1 query parameter is required to get data.
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

export async function get_payroll_inputs(endDate, startDate, limit, offset, payComponent, worker) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (endDate) params.append("endDate", endDate);
      if (startDate) params.append("startDate", startDate);
      if (limit) params.append("limit", limit);
      if (offset) params.append("offset", offset);
      if (payComponent) params.append("payComponent", payComponent);
      if (worker) params.append("worker", worker);
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

export function createGetPayrollInputsTool() {
  return {
    definition: {
      name: 'get-payroll-inputs',
      description: 'Retrieves a collection of payroll inputs. At least 1 query parameter is required to get data.',
      inputSchema: {
        type: 'object',
        properties: {
          endDate: {
            type: 'string',
            description: 'The end date filter which returns payroll inputs that are active on or before the end date, using the date format yyyy-mm-dd.'
          },
          startDate: {
            type: 'string',
            description: 'The start date filter which returns payroll inputs that are active on or after the start date, using the date format yyyy-mm-dd.'
          },
          limit: {
            type: 'number',
            description: 'The maximum number of objects in a single response. The default is 20. The maximum is 100.'
          },
          offset: {
            type: 'number',
            description: 'The zero-based index of the first object in a response collection. The default is 0. Use offset with the limit parameter to control paging of a response collection. Example: If limit is 5 and offset is 9, the response returns a collection of 5 objects starting with the 10th object.'
          },
          payComponent: {
            type: 'string',
            description: 'The Workday ID of the pay component for the payroll input. You can use a returned id from GET /values/payrollInputsGroup/payComponents.'
          },
          worker: {
            type: 'string',
            description: 'The Workday ID of worker for the payroll input. You can use a returned id from GET /workers in the Staffing service.'
          }
        },
        required: []
      }
    },
    handler: get_payroll_inputs
  };
}