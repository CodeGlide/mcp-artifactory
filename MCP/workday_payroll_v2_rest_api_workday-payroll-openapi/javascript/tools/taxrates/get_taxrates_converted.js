/**
 * Retrieves a single or a collection of company SUI rates.
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

export async function get_tax_rates(company, effective, payrollStateAuthorityTaxCode, limit, offset) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (company) params.append("company", company);
      if (effective) params.append("effective", effective);
      if (payrollStateAuthorityTaxCode) params.append("payrollStateAuthorityTaxCode", payrollStateAuthorityTaxCode);
      if (limit) params.append("limit", limit);
      if (offset) params.append("offset", offset);
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

export function createGetTaxRatesTool() {
  return {
    definition: {
      name: 'get-tax-rates',
      description: 'Retrieves a single or a collection of company SUI rates.',
      inputSchema: {
        type: 'object',
        properties: {
          company: {
            type: 'string',
            description: 'The company reference ID or WID that represents 1 or more companies. Example: company=comp1&company=comp2&company=cb550da820584750aae8f807882fa79a'
          },
          effective: {
            type: 'string',
            description: 'The effective date for the SUI rate, using the yyyy-mm-dd format.'
          },
          payrollStateAuthorityTaxCode: {
            type: 'string',
            description: 'The FIPS code or WID that represents 1 or more states. Example: payrollStateAuthorityTaxCode=06&payrollStateAuthorityTaxCode=3b3d378d5f4a48b8b3ac46fee0703226&payrollStateAuthorityTaxCode=48'
          },
          limit: {
            type: 'number',
            description: 'The maximum number of objects in a single response. The default is 20. The maximum is 100.'
          },
          offset: {
            type: 'number',
            description: 'The zero-based index of the first object in a response collection. The default is 0. Use offset with the limit parameter to control paging of a response collection. Example: If limit is 5 and offset is 9, the response returns a collection of 5 objects starting with the 10th object.'
          }
        },
        required: []
      }
    },
    handler: get_tax_rates
  };
}