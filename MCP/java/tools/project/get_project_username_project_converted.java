/**
 * MCP Server function for Build summary for each of the last 30 builds for a single git repo.
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

class Get_Project_Username_ProjectMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_Project_Username_ProjectHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("username")) {
            queryParams.add("username=" + args.get("username"));
        }
        if (args.containsKey("project")) {
            queryParams.add("project=" + args.get("project"));
        }
        if (args.containsKey("filter")) {
            queryParams.add("filter=" + args.get("filter"));
        }
        if (args.containsKey("limit")) {
            queryParams.add("limit=" + args.get("limit"));
        }
        if (args.containsKey("offset")) {
            queryParams.add("offset=" + args.get("offset"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_project_username_project" + queryString;
                
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
    
    public static MCPServer.Tool createGet_Project_Username_ProjectTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> usernameProperty = new HashMap<>();
        usernameProperty.put("type", "string");
        usernameProperty.put("required", true);
        usernameProperty.put("description", "XXXXXXXXX");
        properties.put("username", usernameProperty);
        Map<String, Object> projectProperty = new HashMap<>();
        projectProperty.put("type", "string");
        projectProperty.put("required", true);
        projectProperty.put("description", "XXXXXXXXX");
        properties.put("project", projectProperty);
        Map<String, Object> filterProperty = new HashMap<>();
        filterProperty.put("type", "string");
        filterProperty.put("required", false);
        filterProperty.put("description", "Restricts which builds are returned. Set to \"completed\", \"successful\", \"failed\", \"running\", or defaults to no filter.");
        properties.put("filter", filterProperty);
        Map<String, Object> limitProperty = new HashMap<>();
        limitProperty.put("type", "string");
        limitProperty.put("required", false);
        limitProperty.put("description", "The number of builds to return. Maximum 100, defaults to 30.");
        properties.put("limit", limitProperty);
        Map<String, Object> offsetProperty = new HashMap<>();
        offsetProperty.put("type", "string");
        offsetProperty.put("required", false);
        offsetProperty.put("description", "The API returns builds starting from this offset, defaults to 0.");
        properties.put("offset", offsetProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_project_username_project",
            "Build summary for each of the last 30 builds for a single git repo.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_Project_Username_ProjectHandler(config));
    }
    
}