package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/apigee-registry-api/mcp-server/config"
	"github.com/apigee-registry-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Apigeeregistry_projects_locations_instances_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		parentVal, ok := args["parent"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: parent"), nil
		}
		parent, ok := parentVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: parent"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["instanceId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("instanceId=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Instance
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/v1/%s/instances%s", cfg.BaseURL, parent, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Operation
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateApigeeregistry_projects_locations_instances_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_instances",
		mcp.WithDescription("Provisions instance resources for the Registry."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. Parent resource of the Instance, of the form: `projects/*/locations/*`")),
		mcp.WithString("instanceId", mcp.Description("Required. Identifier to assign to the Instance. Must be unique within scope of the parent resource.")),
		mcp.WithString("updateTime", mcp.Description("Input parameter: Output only. Last update timestamp.")),
		mcp.WithObject("build", mcp.Description("Input parameter: Build information of the Instance if it's in `ACTIVE` state.")),
		mcp.WithObject("config", mcp.Description("Input parameter: Available configurations to provision an Instance.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Creation timestamp.")),
		mcp.WithString("name", mcp.Description("Input parameter: Format: `projects/*/locations/*/instance`. Currently only `locations/global` is supported.")),
		mcp.WithString("state", mcp.Description("Input parameter: Output only. The current state of the Instance.")),
		mcp.WithString("stateMessage", mcp.Description("Input parameter: Output only. Extra information of Instance.State if the state is `FAILED`.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apigeeregistry_projects_locations_instances_createHandler(cfg),
	}
}
