/**
 * MCP Server function for Remove a comment.
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

class Delete_Media_Media_Id_Comments_Comment_IdMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getDelete_Media_Media_Id_Comments_Comment_IdHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("media-id")) {
            queryParams.add("media-id=" + args.get("media-id"));
        }
        if (args.containsKey("comment-id")) {
            queryParams.add("comment-id=" + args.get("comment-id"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/delete_media_media_id_comments_comment_id" + queryString;
                
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
    
    public static MCPServer.Tool createDelete_Media_Media_Id_Comments_Comment_IdTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> media-idProperty = new HashMap<>();
        media-idProperty.put("type", "string");
        media-idProperty.put("required", true);
        media-idProperty.put("description", "The ID of the media resource.");
        properties.put("media-id", media-idProperty);
        Map<String, Object> comment-idProperty = new HashMap<>();
        comment-idProperty.put("type", "string");
        comment-idProperty.put("required", true);
        comment-idProperty.put("description", "The ID of the comment entry.");
        properties.put("comment-id", comment-idProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "delete_media_media_id_comments_comment_id",
            "Remove a comment.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getDelete_Media_Media_Id_Comments_Comment_IdHandler(config));
    }
    
}