/**
 * Creates payroll inputs.
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

export async function post_payroll_inputs(descriptor, fieldEditability, endDate, comment, id, startDate, ongoing, adjustment, position, worker, payComponent, currency, worktags, runCategories, inputDetails) {
  try {
    const config = getConfig();
    const requestBody = {
      descriptor,
      fieldEditability,
      endDate,
      comment,
      id,
      startDate,
      ongoing,
      adjustment,
      position,
      worker,
      payComponent,
      currency,
      worktags,
      runCategories,
      inputDetails
    };
    
    const url = `${config.baseURL}/payrollInputs`;
    
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

export function createPostPayrollInputsTool() {
  return {
    definition: {
      name: 'post-payroll-inputs',
      description: 'Creates payroll inputs.',
      inputSchema: {
        type: 'object',
        properties: {
          descriptor: {
            type: 'string',
            description: 'Input parameter: A preview of the instance'
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
          ongoing: {
            type: 'boolean',
            description: 'Input parameter: If true, the payroll input is ongoing.'
          },
          adjustment: {
            type: 'boolean',
            description: 'Input parameter: If true, the input is for an adjustment as opposed to an override.'
          },
          position: {
            type: 'object',
            description: 'Input parameter: The worker's position the payroll input applies to if Multi Job Payroll is used.'
          },
          worker: {
            type: 'object',
            description: 'Input parameter: The worker for this payroll input.'
          },
          payComponent: {
            type: 'object',
            description: 'Input parameter: The pay component for this payroll input.'
          },
          currency: {
            type: 'object',
            description: 'Input parameter: The currency for the payroll input. If no currency exists, the system assumes the Pay Group currency. The Pay Group currency is derived from the default currency for the Pay Group country.'
          },
          worktags: {
            type: 'string',
            description: 'Input parameter: The worktags associated with the payroll input.'
          },
          runCategories: {
            type: 'string',
            description: 'Input parameter: The run category for the payroll input.'
          },
          inputDetails: {
            type: 'string',
            description: 'Input parameter: The details for this payroll input.'
          }
        },
        required: []
      }
    },
    handler: post_payroll_inputs
  };
}