/**
 * Main MCP Server - Handles MCP JSON-RPC protocol
 */

import java.io.*;
import java.util.*;
import java.util.function.Function;
import java.util.concurrent.CompletableFuture;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.fasterxml.jackson.databind.node.ArrayNode;

// Import Delete_Media_Media_Id_Comments_Comment_IdMCPTool - included in same package
// Import Get_Media_Media_Id_CommentsMCPTool - included in same package
// Import Get_Geographies_Geo_Id_Media_RecentMCPTool - included in same package
// Import Get_Users_SearchMCPTool - included in same package
// Import Get_Tags_Tag_Name_Media_RecentMCPTool - included in same package
// Import Get_Tags_Tag_NameMCPTool - included in same package
// Import Get_Users_User_Id_RelationshipMCPTool - included in same package
// Import Get_Users_User_IdMCPTool - included in same package
// Import Get_Users_User_Id_Followed_ByMCPTool - included in same package
// Import Get_Users_User_Id_FollowsMCPTool - included in same package
// Import Get_Users_User_Id_Media_RecentMCPTool - included in same package
// Import Get_Users_Self_Media_LikedMCPTool - included in same package
// Import Get_Users_Self_Requested_ByMCPTool - included in same package
// Import Get_Users_Self_FeedMCPTool - included in same package
// Import Get_Locations_Location_Id_Media_RecentMCPTool - included in same package
// Import Get_Locations_Location_IdMCPTool - included in same package
// Import Get_Media_ShortcodeMCPTool - included in same package
// Import Get_Media_SearchMCPTool - included in same package
// Import Get_Media_PopularMCPTool - included in same package
// Import Get_Tags_SearchMCPTool - included in same package
// Import Get_Locations_SearchMCPTool - included in same package
// Import Delete_Media_Media_Id_LikesMCPTool - included in same package
// Import Get_Media_Media_IdMCPTool - included in same package

public class MCPServer {
    
    // Common classes that all tool classes use
    public static class APIConfig {
        private final String baseUrl;
        private final String apiKey;
        
        public APIConfig(String baseUrl, String apiKey) {
            this.baseUrl = baseUrl;
            this.apiKey = apiKey;
        }
        
        public String getBaseUrl() { return baseUrl; }
        public String getApiKey() { return apiKey; }
    }
    
    public static class MCPRequest {
        private Map<String, Object> params;
        
        public Map<String, Object> getParams() { return params; }
        public void setParams(Map<String, Object> params) { this.params = params; }
        
        @SuppressWarnings("unchecked")
        public Map<String, Object> getArguments() {
            if (params != null && params.containsKey("arguments")) {
                return (Map<String, Object>) params.get("arguments");
            }
            return new HashMap<>();
        }
    }
    
    public static class MCPToolResult {
        private final String content;
        private final boolean isError;
        
        public MCPToolResult(String content, boolean isError) {
            this.content = content;
            this.isError = isError;
        }
        
        public MCPToolResult(String content) {
            this(content, false);
        }
        
        public String getContent() { return content; }
        public boolean isError() { return isError; }
    }
    
    public static class ToolDefinition {
        private final String name;
        private final String description;
        private final Map<String, Object> parameters;
        
        public ToolDefinition(String name, String description, Map<String, Object> parameters) {
            this.name = name;
            this.description = description;
            this.parameters = parameters;
        }
        
        public String getName() { return name; }
        public String getDescription() { return description; }
        public Map<String, Object> getParameters() { return parameters; }
    }
    
    public static class Tool {
        private final ToolDefinition definition;
        private final Function<MCPRequest, MCPToolResult> handler;
        
        public Tool(ToolDefinition definition, Function<MCPRequest, MCPToolResult> handler) {
            this.definition = definition;
            this.handler = handler;
        }
        
        public ToolDefinition getDefinition() { return definition; }
        public Function<MCPRequest, MCPToolResult> getHandler() { return handler; }
    }
    
    private static final ObjectMapper mapper = new ObjectMapper();
    private static List<Tool> tools = new ArrayList<>();
    private static APIConfig config;
    
    public static void main(String[] args) {
        try {
            // Initialize configuration
            String baseUrl = System.getenv("API_BASE_URL");
            String apiKey = System.getenv("API_KEY");
            
            if (baseUrl == null || apiKey == null) {
                System.err.println("Error: Please set API_BASE_URL and API_KEY environment variables");
                System.exit(1);
            }
            
            config = new APIConfig(baseUrl, apiKey);
            
            // Register all tools
            tools = Arrays.asList(
            Delete_Media_Media_Id_Comments_Comment_IdMCPTool.createDelete_Media_Media_Id_Comments_Comment_IdTool(config),
            Get_Media_Media_Id_CommentsMCPTool.createGet_Media_Media_Id_CommentsTool(config),
            Get_Geographies_Geo_Id_Media_RecentMCPTool.createGet_Geographies_Geo_Id_Media_RecentTool(config),
            Get_Users_SearchMCPTool.createGet_Users_SearchTool(config),
            Get_Tags_Tag_Name_Media_RecentMCPTool.createGet_Tags_Tag_Name_Media_RecentTool(config),
            Get_Tags_Tag_NameMCPTool.createGet_Tags_Tag_NameTool(config),
            Get_Users_User_Id_RelationshipMCPTool.createGet_Users_User_Id_RelationshipTool(config),
            Get_Users_User_IdMCPTool.createGet_Users_User_IdTool(config),
            Get_Users_User_Id_Followed_ByMCPTool.createGet_Users_User_Id_Followed_ByTool(config),
            Get_Users_User_Id_FollowsMCPTool.createGet_Users_User_Id_FollowsTool(config),
            Get_Users_User_Id_Media_RecentMCPTool.createGet_Users_User_Id_Media_RecentTool(config),
            Get_Users_Self_Media_LikedMCPTool.createGet_Users_Self_Media_LikedTool(config),
            Get_Users_Self_Requested_ByMCPTool.createGet_Users_Self_Requested_ByTool(config),
            Get_Users_Self_FeedMCPTool.createGet_Users_Self_FeedTool(config),
            Get_Locations_Location_Id_Media_RecentMCPTool.createGet_Locations_Location_Id_Media_RecentTool(config),
            Get_Locations_Location_IdMCPTool.createGet_Locations_Location_IdTool(config),
            Get_Media_ShortcodeMCPTool.createGet_Media_ShortcodeTool(config),
            Get_Media_SearchMCPTool.createGet_Media_SearchTool(config),
            Get_Media_PopularMCPTool.createGet_Media_PopularTool(config),
            Get_Tags_SearchMCPTool.createGet_Tags_SearchTool(config),
            Get_Locations_SearchMCPTool.createGet_Locations_SearchTool(config),
            Delete_Media_Media_Id_LikesMCPTool.createDelete_Media_Media_Id_LikesTool(config),
            Get_Media_Media_IdMCPTool.createGet_Media_Media_IdTool(config)
            );
            
            System.err.println("MCP Server started with " + tools.size() + " tools");
            
            // Start JSON-RPC server
            runServer();
            
        } catch (Exception e) {
            System.err.println("Failed to start MCP server: " + e.getMessage());
            e.printStackTrace();
            System.exit(1);
        }
    }
    
    private static void runServer() throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String line;
        
        while ((line = reader.readLine()) != null) {
            JsonNode request = null;
            try {
                request = mapper.readTree(line);
                JsonNode response = handleRequest(request);
                
                if (response != null) {
                    System.out.println(mapper.writeValueAsString(response));
                }
                
            } catch (Exception e) {
                // Send error response
                ObjectNode errorResponse = mapper.createObjectNode();
                errorResponse.put("jsonrpc", "2.0");
                if (request != null && request.has("id")) {
                    errorResponse.put("id", request.get("id").asText());
                } else {
                    errorResponse.putNull("id");
                }
                
                ObjectNode error = mapper.createObjectNode();
                error.put("code", -32603);
                error.put("message", "Internal error: " + e.getMessage());
                errorResponse.set("error", error);
                
                System.out.println(mapper.writeValueAsString(errorResponse));
            }
        }
    }
    
    private static JsonNode handleRequest(JsonNode request) {
        if (!request.has("method")) {
            return createErrorResponse(request, -32600, "Invalid Request");
        }
        
        String method = request.get("method").asText();
        JsonNode params = request.get("params");
        String id = request.has("id") ? request.get("id").asText() : null;
        
        switch (method) {
            case "initialize":
                return handleInitialize(id);
            case "tools/list":
                return handleToolsList(id);
            case "tools/call":
                return handleToolsCall(id, params);
            default:
                return createErrorResponse(request, -32601, "Method not found");
        }
    }
    
    private static JsonNode handleInitialize(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        result.put("protocolVersion", "2024-11-05");
        
        ObjectNode capabilities = mapper.createObjectNode();
        ObjectNode toolsCapability = mapper.createObjectNode();
        toolsCapability.put("listChanged", false);
        capabilities.set("tools", toolsCapability);
        result.set("capabilities", capabilities);
        
        ObjectNode serverInfo = mapper.createObjectNode();
        serverInfo.put("name", "Opsera MCP Server");
        serverInfo.put("version", "1.0.0");
        result.set("serverInfo", serverInfo);
        
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsList(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        ArrayNode toolsArray = mapper.createArrayNode();
        
        for (Tool tool : tools) {
            ObjectNode toolDef = mapper.createObjectNode();
            toolDef.put("name", tool.getDefinition().getName());
            toolDef.put("description", tool.getDefinition().getDescription());
            
            // Convert parameters to JSON
            ObjectNode inputSchema = mapper.valueToTree(tool.getDefinition().getParameters());
            toolDef.set("inputSchema", inputSchema);
            
            toolsArray.add(toolDef);
        }
        
        result.set("tools", toolsArray);
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsCall(String id, JsonNode params) {
        if (!params.has("name")) {
            return createErrorResponse(null, -32602, "Invalid params: missing 'name'");
        }
        
        String toolName = params.get("name").asText();
        JsonNode arguments = params.has("arguments") ? params.get("arguments") : mapper.createObjectNode();
        
        // Find the tool
        Tool tool = null;
        for (Tool t : tools) {
            if (t.getDefinition().getName().equals(toolName)) {
                tool = t;
                break;
            }
        }
        
        if (tool == null) {
            return createErrorResponse(null, -32602, "Tool not found: " + toolName);
        }
        
        try {
            // Convert arguments to Map
            Map<String, Object> argsMap = mapper.convertValue(arguments, Map.class);
            
            // Create MCP request
            MCPRequest mcpRequest = new MCPRequest();
            Map<String, Object> requestParams = new HashMap<>();
            requestParams.put("arguments", argsMap);
            mcpRequest.setParams(requestParams);
            
            // Execute tool
            MCPToolResult result = tool.getHandler().apply(mcpRequest);
            
            // Create response
            ObjectNode response = mapper.createObjectNode();
            response.put("jsonrpc", "2.0");
            response.put("id", id);
            
            ObjectNode resultObj = mapper.createObjectNode();
            ArrayNode content = mapper.createArrayNode();
            
            ObjectNode textContent = mapper.createObjectNode();
            textContent.put("type", "text");
            textContent.put("text", result.getContent());
            content.add(textContent);
            
            resultObj.set("content", content);
            resultObj.put("isError", result.isError());
            
            response.set("result", resultObj);
            return response;
            
        } catch (Exception e) {
            return createErrorResponse(null, -32603, "Tool execution failed: " + e.getMessage());
        }
    }
    
    private static JsonNode createErrorResponse(JsonNode request, int code, String message) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", request != null && request.has("id") ? request.get("id").asText() : null);
        
        ObjectNode error = mapper.createObjectNode();
        error.put("code", code);
        error.put("message", message);
        response.set("error", error);
        
        return response;
    }
}