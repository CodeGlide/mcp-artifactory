/**
 * MCP Server function for List Links
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

class Get_Api_V2_Jira_LinksMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Api_V2_Jira_LinksHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("page[after]")) {
            queryParams.add("page[after]=" + args.get("page[after]"));
        }
        if (args.containsKey("filter[ticket_id]")) {
            queryParams.add("filter[ticket_id]=" + args.get("filter[ticket_id]"));
        }
        if (args.containsKey("page[size]")) {
            queryParams.add("page[size]=" + args.get("page[size]"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_api_v2_jira_links" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Api_V2_Jira_LinksTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> AcceptProperty = new HashMap<>();
        AcceptProperty.put("type", "string");
        AcceptProperty.put("required", false);
        AcceptProperty.put("description", "");
        properties.put("Accept", AcceptProperty);
        Map<String, Object> page[after]Property = new HashMap<>();
        page[after]Property.put("type", "string");
        page[after]Property.put("required", false);
        page[after]Property.put("description", "When provided, the returned paginated data must have as its first item the item that is immediately after the cursor in the results list. Exception: If there are no items in the results list that fall after the cursor, the returned paginated data must be an empty array");
        properties.put("page[after]", page[after]Property);
        Map<String, Object> filter[ticket_id]Property = new HashMap<>();
        filter[ticket_id]Property.put("type", "string");
        filter[ticket_id]Property.put("required", false);
        filter[ticket_id]Property.put("description", "List links for a specific Zendesk ticket or Jira issue by specifying a ticket id or issue id. Filtering by issue key is not currently supported");
        properties.put("filter[ticket_id]", filter[ticket_id]Property);
        Map<String, Object> page[size]Property = new HashMap<>();
        page[size]Property.put("type", "string");
        page[size]Property.put("required", false);
        page[size]Property.put("description", "The number of entries that will be returned");
        properties.put("page[size]", page[size]Property);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_api_v2_jira_links",
            "List Links",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Api_V2_Jira_LinksHandler(config));
    }
    
}