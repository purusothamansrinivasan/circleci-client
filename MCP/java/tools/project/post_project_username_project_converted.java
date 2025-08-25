/**
 * MCP Server function for Triggers a new build, returns a summary of the build.
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

class Post_Project_Username_ProjectMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Project_Username_ProjectHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("revision")) {
            queryParams.add("revision=" + args.get("revision"));
        }
        if (args.containsKey("tag")) {
            queryParams.add("tag=" + args.get("tag"));
        }
        if (args.containsKey("parallel")) {
            queryParams.add("parallel=" + args.get("parallel"));
        }
        if (args.containsKey("build_parameters")) {
            queryParams.add("build_parameters=" + args.get("build_parameters"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_project_username_project" + queryString;
                
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
    
    public static MCPServer.Tool createPost_Project_Username_ProjectTool(MCPServer.APIConfig config) {
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
        Map<String, Object> revisionProperty = new HashMap<>();
        revisionProperty.put("type", "string");
        revisionProperty.put("required", false);
        revisionProperty.put("description", "Input parameter: The specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.");
        properties.put("revision", revisionProperty);
        Map<String, Object> tagProperty = new HashMap<>();
        tagProperty.put("type", "string");
        tagProperty.put("required", false);
        tagProperty.put("description", "Input parameter: The tag to build. Default is null. Cannot be used with revision parameter.");
        properties.put("tag", tagProperty);
        Map<String, Object> parallelProperty = new HashMap<>();
        parallelProperty.put("type", "string");
        parallelProperty.put("required", false);
        parallelProperty.put("description", "Input parameter: The number of containers to use to run the build. Default is null and the project default is used.");
        properties.put("parallel", parallelProperty);
        Map<String, Object> build_parametersProperty = new HashMap<>();
        build_parametersProperty.put("type", "string");
        build_parametersProperty.put("required", false);
        build_parametersProperty.put("description", "Input parameter: Additional environment variables to inject into the build environment. A map of names to values.");
        properties.put("build_parameters", build_parametersProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_project_username_project",
            "Triggers a new build, returns a summary of the build.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Project_Username_ProjectHandler(config));
    }
    
}