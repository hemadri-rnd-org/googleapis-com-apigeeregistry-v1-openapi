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

func Apigeeregistry_projects_locations_apis_versions_specs_createHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["apiSpecId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiSpecId=%v", val))
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
		var requestBody models.ApiSpec
		
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
		url := fmt.Sprintf("%s/v1/%s/specs%s", cfg.BaseURL, parent, queryString)
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
		var result models.ApiSpec
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

func CreateApigeeregistry_projects_locations_apis_versions_specs_createTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_v1_parent_specs",
		mcp.WithDescription("Creates a specified spec."),
		mcp.WithString("parent", mcp.Required(), mcp.Description("Required. The parent, which owns this collection of specs. Format: `projects/*/locations/*/apis/*/versions/*`")),
		mcp.WithString("apiSpecId", mcp.Description("Required. The ID to use for the spec, which will become the final component of the spec's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.")),
		mcp.WithObject("annotations", mcp.Description("Input parameter: Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.")),
		mcp.WithString("mimeType", mcp.Description("Input parameter: A style (format) descriptor for this spec that is specified as a [Media Type](https://en.wikipedia.org/wiki/Media_type). Possible values include `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and `application/vnd.apigee.graphql`, with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with \"+gzip\").")),
		mcp.WithString("contents", mcp.Description("Input parameter: Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.")),
		mcp.WithString("createTime", mcp.Description("Input parameter: Output only. Creation timestamp; when the spec resource was created.")),
		mcp.WithString("description", mcp.Description("Input parameter: A detailed description.")),
		mcp.WithString("filename", mcp.Description("Input parameter: A possibly-hierarchical name used to refer to the spec from other specs.")),
		mcp.WithObject("labels", mcp.Description("Input parameter: Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.")),
		mcp.WithString("name", mcp.Description("Input parameter: Resource name.")),
		mcp.WithString("revisionCreateTime", mcp.Description("Input parameter: Output only. Revision creation timestamp; when the represented revision was created.")),
		mcp.WithString("hash", mcp.Description("Input parameter: Output only. A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.")),
		mcp.WithString("revisionUpdateTime", mcp.Description("Input parameter: Output only. Last update timestamp: when the represented revision was last modified.")),
		mcp.WithNumber("sizeBytes", mcp.Description("Input parameter: Output only. The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.")),
		mcp.WithString("sourceUri", mcp.Description("Input parameter: The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.")),
		mcp.WithString("revisionId", mcp.Description("Input parameter: Output only. Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Apigeeregistry_projects_locations_apis_versions_specs_createHandler(cfg),
	}
}
