/**
 * Creates a single or a collection of Company SUI Rates.
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

export async function post_tax_rates(endDate, id, taxCode, applicableRate, ein, startDate, exempt, stateInstance, companyInstance) {
  try {
    const config = getConfig();
    const requestBody = {
      endDate,
      id,
      taxCode,
      applicableRate,
      ein,
      startDate,
      exempt,
      stateInstance,
      companyInstance
    };
    
    const url = `${config.baseURL}/taxRates`;
    
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

export function createPostTaxRatesTool() {
  return {
    definition: {
      name: 'post-tax-rates',
      description: 'Creates a single or a collection of Company SUI Rates.',
      inputSchema: {
        type: 'object',
        properties: {
          endDate: {
            type: 'string',
            description: 'Input parameter: The end date for company tax reporting.'
          },
          id: {
            type: 'string',
            description: 'Input parameter: Id of the instance'
          },
          taxCode: {
            type: 'string',
            description: 'Input parameter: The deduction for company tax reporting.'
          },
          applicableRate: {
            type: 'string',
            description: 'Input parameter: The tax override rate for company tax reporting.'
          },
          ein: {
            type: 'string',
            description: 'Input parameter: The payroll tax authority EIN field for company tax reporting.'
          },
          startDate: {
            type: 'string',
            description: 'Input parameter: The start date for company tax reporting.'
          },
          exempt: {
            type: 'boolean',
            description: 'Input parameter: If true, the SUI rate is exempt.'
          },
          stateInstance: {
            type: 'object',
            description: 'Input parameter: The payroll tax authority object for company tax reporting.'
          },
          companyInstance: {
            type: 'object',
            description: 'Input parameter: The company object for company tax reporting.'
          }
        },
        required: []
      }
    },
    handler: post_tax_rates
  };
}