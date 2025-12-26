/**
 * MCP Server function for List Links (deprecated)
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Get_Api_Services_Jira_LinksMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Api_Services_Jira_LinksHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("Accept")) {
            queryParams.add("Accept=" + args.get("Accept"));
        }
        if (args.containsKey("since_id")) {
            queryParams.add("since_id=" + args.get("since_id"));
        }
        if (args.containsKey("ticket_id")) {
            queryParams.add("ticket_id=" + args.get("ticket_id"));
        }
        if (args.containsKey("limit")) {
            queryParams.add("limit=" + args.get("limit"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_api_services_jira_links" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createGet_Api_Services_Jira_LinksTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> AcceptProperty = new HashMap<>();
        AcceptProperty.put("type", "string");
        AcceptProperty.put("required", false);
        AcceptProperty.put("description", "");
        properties.put("Accept", AcceptProperty);
        Map<String, Object> since_idProperty = new HashMap<>();
        since_idProperty.put("type", "string");
        since_idProperty.put("required", false);
        since_idProperty.put("description", "The start id of the collection");
        properties.put("since_id", since_idProperty);
        Map<String, Object> ticket_idProperty = new HashMap<>();
        ticket_idProperty.put("type", "string");
        ticket_idProperty.put("required", false);
        ticket_idProperty.put("description", "List links for a specific Zendesk Ticket or Jira issue by providing `ticket_id` and/or `issue_id` param. We currently do not support `issue_key` param.");
        properties.put("ticket_id", ticket_idProperty);
        Map<String, Object> limitProperty = new HashMap<>();
        limitProperty.put("type", "string");
        limitProperty.put("required", false);
        limitProperty.put("description", "The number of entries that will be returned");
        properties.put("limit", limitProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_api_services_jira_links",
            "List Links (deprecated)",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Api_Services_Jira_LinksHandler(config));
    }
    
}