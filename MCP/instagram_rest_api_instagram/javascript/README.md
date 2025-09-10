# MCP Server (JavaScript)

This is a MCP Server converted from Go to JavaScript.

## Setup

### Prerequisites
- Node.js 18 or higher
- npm or yarn

### Installation
```bash
npm install
```

### Configuration
Set the following environment variables:
- `API_BASE_URL`: Your API base URL
- `API_BEARER_TOKEN`: Your API bearer token

Alternatively, create a config file at `~/.api/config.json`:
```json
{
  "baseURL": "https://your-api-instance.com",
  "bearerToken": "your-bearer-token"
}
```

### Running the Server
```bash
npm start
```

### Development
```bash
npm run dev
```

## Usage

This MCP server provides access to various API endpoints through the MCP protocol.

## Available Tools

The server provides access to various API endpoints. Each tool corresponds to a specific API endpoint and allows you to interact with different parts of the platform.

## License

MIT
