/**
 * MCP Server function for Get a list of recently tagged media.
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

class Get_Tags_Tag_Name_Media_RecentMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Tags_Tag_Name_Media_RecentHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("min_tag_id")) {
            queryParams.add("min_tag_id=" + args.get("min_tag_id"));
        }
        if (args.containsKey("max_tag_id")) {
            queryParams.add("max_tag_id=" + args.get("max_tag_id"));
        }
        if (args.containsKey("tag-name")) {
            queryParams.add("tag-name=" + args.get("tag-name"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_tags_tag_name_media_recent" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Tags_Tag_Name_Media_RecentTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> countProperty = new HashMap<>();
        countProperty.put("type", "string");
        countProperty.put("required", false);
        countProperty.put("description", "Max number of media to return.");
        properties.put("count", countProperty);
        Map<String, Object> min_tag_idProperty = new HashMap<>();
        min_tag_idProperty.put("type", "string");
        min_tag_idProperty.put("required", false);
        min_tag_idProperty.put("description", "Return media before this `min_tag_id`.");
        properties.put("min_tag_id", min_tag_idProperty);
        Map<String, Object> max_tag_idProperty = new HashMap<>();
        max_tag_idProperty.put("type", "string");
        max_tag_idProperty.put("required", false);
        max_tag_idProperty.put("description", "Return media after this `max_tag_id`.");
        properties.put("max_tag_id", max_tag_idProperty);
        Map<String, Object> tag-nameProperty = new HashMap<>();
        tag-nameProperty.put("type", "string");
        tag-nameProperty.put("required", true);
        tag-nameProperty.put("description", "The tag name.");
        properties.put("tag-name", tag-nameProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_tags_tag_name_media_recent",
            "Get a list of recently tagged media.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Tags_Tag_Name_Media_RecentHandler(config));
    }
    
}