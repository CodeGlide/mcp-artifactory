/**
 * MCP Server function for Search for a location by geographic coordinate.
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

class Get_Locations_SearchMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Locations_SearchHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("distance")) {
            queryParams.add("distance=" + args.get("distance"));
        }
        if (args.containsKey("facebook_places_id")) {
            queryParams.add("facebook_places_id=" + args.get("facebook_places_id"));
        }
        if (args.containsKey("foursquare_id")) {
            queryParams.add("foursquare_id=" + args.get("foursquare_id"));
        }
        if (args.containsKey("lat")) {
            queryParams.add("lat=" + args.get("lat"));
        }
        if (args.containsKey("lng")) {
            queryParams.add("lng=" + args.get("lng"));
        }
        if (args.containsKey("foursquare_v2_id")) {
            queryParams.add("foursquare_v2_id=" + args.get("foursquare_v2_id"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_locations_search" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Locations_SearchTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> distanceProperty = new HashMap<>();
        distanceProperty.put("type", "string");
        distanceProperty.put("required", false);
        distanceProperty.put("description", "Default is 1000m (distance=1000), max distance is 5000.");
        properties.put("distance", distanceProperty);
        Map<String, Object> facebook_places_idProperty = new HashMap<>();
        facebook_places_idProperty.put("type", "string");
        facebook_places_idProperty.put("required", false);
        facebook_places_idProperty.put("description", "Returns a location mapped off of a Facebook places id. If used, a Foursquare id and `lat`, `lng` are not required.");
        properties.put("facebook_places_id", facebook_places_idProperty);
        Map<String, Object> foursquare_idProperty = new HashMap<>();
        foursquare_idProperty.put("type", "string");
        foursquare_idProperty.put("required", false);
        foursquare_idProperty.put("description", "Returns a location mapped off of a foursquare v1 api location id. If used, you are not required to use `lat` and `lng`. Note that this method is deprecated; you should use the new foursquare IDs with V2 of their API.");
        properties.put("foursquare_id", foursquare_idProperty);
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
        Map<String, Object> foursquare_v2_idProperty = new HashMap<>();
        foursquare_v2_idProperty.put("type", "string");
        foursquare_v2_idProperty.put("required", false);
        foursquare_v2_idProperty.put("description", "Returns a location mapped off of a foursquare v2 api location id. If used, you are not required to use `lat` and `lng`.");
        properties.put("foursquare_v2_id", foursquare_v2_idProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_locations_search",
            "Search for a location by geographic coordinate.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Locations_SearchHandler(config));
    }
    
}