/**
 * MCP Server function for See the authenticated user's feed.
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

class Get_Users_Self_FeedMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Users_Self_FeedHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("count")) {
            queryParams.add("count=" + args.get("count"));
        }
        if (args.containsKey("min_id")) {
            queryParams.add("min_id=" + args.get("min_id"));
        }
        if (args.containsKey("max_id")) {
            queryParams.add("max_id=" + args.get("max_id"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_users_self_feed" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Users_Self_FeedTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> countProperty = new HashMap<>();
        countProperty.put("type", "string");
        countProperty.put("required", false);
        countProperty.put("description", "Max number of media to return.");
        properties.put("count", countProperty);
        Map<String, Object> min_idProperty = new HashMap<>();
        min_idProperty.put("type", "string");
        min_idProperty.put("required", false);
        min_idProperty.put("description", "Return media before this `min_id`.");
        properties.put("min_id", min_idProperty);
        Map<String, Object> max_idProperty = new HashMap<>();
        max_idProperty.put("type", "string");
        max_idProperty.put("required", false);
        max_idProperty.put("description", "Return media after this `max_id`.");
        properties.put("max_id", max_idProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_users_self_feed",
            "See the authenticated user's feed.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Users_Self_FeedHandler(config));
    }
    
}