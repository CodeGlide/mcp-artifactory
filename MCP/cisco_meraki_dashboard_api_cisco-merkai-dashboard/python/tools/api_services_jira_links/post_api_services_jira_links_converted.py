"""
MCP tool for Create Link
"""

import json
import requests
from typing import Dict, Any, Optional


def post_api_services_jira_links(Content-Type: str, Accept: str) -> str:
    """
    Create Link
    """
    try:
        # Get configuration
        config = get_config()
        
        if not config.base_url or not config.bearer_token:
            return "Error: Missing API configuration. Please set API_BASE_URL and API_BEARER_TOKEN environment variables."
        
        # Build request parameters
        params = {}
        pass
        # Build request body
        request_data = {
        "Content-Type": Content-Type,
        "Accept": Accept
        }
        
        # Remove None values
        request_data = {k: v for k, v in request_data.items() if v is not None}
        
        # Make API call
        url = f"{config.base_url}/api/services/jira/links"
        
        headers = {
            "Authorization": f"Bearer {config.bearer_token}",
            "Accept": "application/json",
            "Content-Type": "application/json"
        }
        
        response = requests.post(url, headers=headers, json=request_data, params=params)
        
        # Handle HTTP errors
        if response.status_code >= 400:
            try:
                error_data = response.json()
                return f"Failed to read response body: {json.dumps(error_data, indent=2)}"
            except json.JSONDecodeError:
                return f"Failed to read response body: {response.text}"
        
        # Parse response
        try:
            result = response.json()
            return json.dumps(result, indent=2)
        except json.JSONDecodeError:
            # Fallback to raw text if JSON parsing fails
            return response.text
            
    except requests.exceptions.ConnectionError as e:
        return f"Failed to create request: Connection error - {str(e)}"
    except requests.exceptions.Timeout as e:
        return f"Failed to create request: Request timeout - {str(e)}"
    except requests.exceptions.RequestException as e:
        return f"Failed to create request: {str(e)}"
    except Exception as e:
        return f"Unexpected error: {str(e)}"


def get_config():
    """Get configuration from environment or config file."""
    import os
    from pathlib import Path
    
    class Config:
        def __init__(self):
            self.base_url = os.getenv("API_BASE_URL")
            self.bearer_token = os.getenv("API_BEARER_TOKEN")
            
            # Try to load from config file if env vars not set
            if not self.base_url or not self.bearer_token:
                config_path = Path.home() / ".api" / "config.json"
                if config_path.exists():
                    with open(config_path, 'r') as f:
                        config_data = json.load(f)
                        self.base_url = self.base_url or config_data.get("baseURL")
                        self.bearer_token = self.bearer_token or config_data.get("bearerToken")
    
    return Config()
