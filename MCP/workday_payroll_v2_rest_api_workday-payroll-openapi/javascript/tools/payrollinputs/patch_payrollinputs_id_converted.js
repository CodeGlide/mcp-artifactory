/**
 * Partially updates an existing payroll input instance.
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

export async function patch_payroll_inputs_id(ID, fieldEditability, endDate, comment, id, startDate, descriptor, ongoing, adjustment, payComponent, currency, position, worker, runCategories, inputDetails, worktags) {
  try {
    const config = getConfig();
    const requestBody = {
      ID,
      fieldEditability,
      endDate,
      comment,
      id,
      startDate,
      descriptor,
      ongoing,
      adjustment,
      payComponent,
      currency,
      position,
      worker,
      runCategories,
      inputDetails,
      worktags
    };
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(url, {
      method: 'PATCH',
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

export function createPatchPayrollInputsIdTool() {
  return {
    definition: {
      name: 'patch-payroll-inputs-id',
      description: 'Partially updates an existing payroll input instance.',
      inputSchema: {
        type: 'object',
        properties: {
          ID: {
            type: 'string',
            description: 'The Workday ID of the resource.'
          },
          fieldEditability: {
            type: 'string',
            description: 'Input parameter: The editability status indicating the fields that can be updated in the payroll input request. Possible values: all, none, endDateOnly'
          },
          endDate: {
            type: 'string',
            description: 'Input parameter: The end date after which this input does not apply.'
          },
          comment: {
            type: 'string',
            description: 'Input parameter: The text comment for this input.'
          },
          id: {
            type: 'string',
            description: 'Input parameter: Id of the instance'
          },
          startDate: {
            type: 'string',
            description: 'Input parameter: The start date before which this input does not apply.'
          },
          descriptor: {
            type: 'string',
            description: 'Input parameter: A preview of the instance'
          },
          ongoing: {
            type: 'boolean',
            description: 'Input parameter: If true, the payroll input is ongoing.'
          },
          adjustment: {
            type: 'boolean',
            description: 'Input parameter: If true, the input is for an adjustment as opposed to an override.'
          },
          payComponent: {
            type: 'object',
            description: 'Input parameter: The pay component for this payroll input.'
          },
          currency: {
            type: 'object',
            description: 'Input parameter: The currency for the payroll input. If no currency exists, the system assumes the Pay Group currency. The Pay Group currency is derived from the default currency for the Pay Group country.'
          },
          position: {
            type: 'object',
            description: 'Input parameter: The worker's position the payroll input applies to if Multi Job Payroll is used.'
          },
          worker: {
            type: 'object',
            description: 'Input parameter: The worker for this payroll input.'
          },
          runCategories: {
            type: 'string',
            description: 'Input parameter: The run category for the payroll input.'
          },
          inputDetails: {
            type: 'string',
            description: 'Input parameter: The details for this payroll input.'
          },
          worktags: {
            type: 'string',
            description: 'Input parameter: The worktags associated with the payroll input.'
          }
        },
        required: ["ID"]
      }
    },
    handler: patch_payroll_inputs_id
  };
}