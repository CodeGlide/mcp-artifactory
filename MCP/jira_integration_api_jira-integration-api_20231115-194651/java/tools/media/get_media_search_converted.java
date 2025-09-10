/**
 * MCP Server function for Search for media in a given area.
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

class Get_Media_SearchMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Media_SearchHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("lat")) {
            queryParams.add("lat=" + args.get("lat"));
        }
        if (args.containsKey("lng")) {
            queryParams.add("lng=" + args.get("lng"));
        }
        if (args.containsKey("min_timestamp")) {
            queryParams.add("min_timestamp=" + args.get("min_timestamp"));
        }
        if (args.containsKey("max_timestamp")) {
            queryParams.add("max_timestamp=" + args.get("max_timestamp"));
        }
        if (args.containsKey("distance")) {
            queryParams.add("distance=" + args.get("distance"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_media_search" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Media_SearchTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> latProperty = new HashMap<>();
        latProperty.put("type", "string");
        latProperty.put("required", false);
        latProperty.put("description", "Latitude of the center search coordinate. If used, `lng` is required.");
        properties.put("lat", latProperty);
        Map<String, Object> lngProperty = new HashMap<>();
        lngProperty.put("type", "string");
        lngProperty.put("required", false);
        lngProperty.put("description", "Longitude of the center search coordinate. If used, `lat` is required.");
        properties.put("lng", lngProperty);
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
        Map<String, Object> distanceProperty = new HashMap<>();
        distanceProperty.put("type", "string");
        distanceProperty.put("required", false);
        distanceProperty.put("description", "Default is 1000m (distance=1000), max distance is 5000.");
        properties.put("distance", distanceProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_media_search",
            "Search for media in a given area.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Media_SearchHandler(config));
    }
    
}