/**
 * MCP Server function for Get a list of recent media objects from a given location.
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

class Get_Locations_Location_Id_Media_RecentMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Locations_Location_Id_Media_RecentHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("min_timestamp")) {
            queryParams.add("min_timestamp=" + args.get("min_timestamp"));
        }
        if (args.containsKey("max_timestamp")) {
            queryParams.add("max_timestamp=" + args.get("max_timestamp"));
        }
        if (args.containsKey("min_id")) {
            queryParams.add("min_id=" + args.get("min_id"));
        }
        if (args.containsKey("max_id")) {
            queryParams.add("max_id=" + args.get("max_id"));
        }
        if (args.containsKey("location-id")) {
            queryParams.add("location-id=" + args.get("location-id"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_locations_location_id_media_recent" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Locations_Location_Id_Media_RecentTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> min_timestampProperty = new HashMap<>();
        min_timestampProperty.put("type", "string");
        min_timestampProperty.put("required", false);
        min_timestampProperty.put("description", "Return media after this UNIX timestamp.");
        properties.put("min_timestamp", min_timestampProperty);
        Map<String, Object> max_timestampProperty = new HashMap<>();
        max_timestampProperty.put("type", "string");
        max_timestampProperty.put("required", false);
        max_timestampProperty.put("description", "Return media before this UNIX timestamp.");
        properties.put("max_timestamp", max_timestampProperty);
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
        Map<String, Object> location-idProperty = new HashMap<>();
        location-idProperty.put("type", "string");
        location-idProperty.put("required", true);
        location-idProperty.put("description", "The location ID.");
        properties.put("location-id", location-idProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_locations_location_id_media_recent",
            "Get a list of recent media objects from a given location.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Locations_Location_Id_Media_RecentHandler(config));
    }
    
}