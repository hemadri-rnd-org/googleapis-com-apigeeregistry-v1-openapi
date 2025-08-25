package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListApiVersionsResponse represents the ListApiVersionsResponse schema from the OpenAPI specification
type ListApiVersionsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	Apiversions []ApiVersion `json:"apiVersions,omitempty"` // The versions from the specified publisher.
}

// ListArtifactsResponse represents the ListArtifactsResponse schema from the OpenAPI specification
type ListArtifactsResponse struct {
	Artifacts []Artifact `json:"artifacts,omitempty"` // The artifacts from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// TagApiDeploymentRevisionRequest represents the TagApiDeploymentRevisionRequest schema from the OpenAPI specification
type TagApiDeploymentRevisionRequest struct {
	Tag string `json:"tag,omitempty"` // Required. The tag to apply. The tag should be at most 40 characters, and match `a-z{3,39}`.
}

// TestIamPermissionsRequest represents the TestIamPermissionsRequest schema from the OpenAPI specification
type TestIamPermissionsRequest struct {
	Permissions []string `json:"permissions,omitempty"` // The set of permissions to check for the `resource`. Permissions with wildcards (such as `*` or `storage.*`) are not allowed. For more information see [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
}

// CancelOperationRequest represents the CancelOperationRequest schema from the OpenAPI specification
type CancelOperationRequest struct {
}

// RollbackApiDeploymentRequest represents the RollbackApiDeploymentRequest schema from the OpenAPI specification
type RollbackApiDeploymentRequest struct {
	Revisionid string `json:"revisionId,omitempty"` // Required. The revision ID to roll back to. It must be a revision of the same deployment. Example: `c7cfa2a8`
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Name string `json:"name,omitempty"` // Resource name for the location, which may vary between implementations. For example: `"projects/example-project/locations/us-east1"`
	Displayname string `json:"displayName,omitempty"` // The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	Labels map[string]interface{} `json:"labels,omitempty"` // Cross-service attributes for the location. For example {"cloud.googleapis.com/region": "us-east1"}
	Locationid string `json:"locationId,omitempty"` // The canonical id for this location. For example: `"us-east1"`.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata. For example the available capacity at the given location.
}

// RollbackApiSpecRequest represents the RollbackApiSpecRequest schema from the OpenAPI specification
type RollbackApiSpecRequest struct {
	Revisionid string `json:"revisionId,omitempty"` // Required. The revision ID to roll back to. It must be a revision of the same spec. Example: `c7cfa2a8`
}

// OperationMetadata represents the OperationMetadata schema from the OpenAPI specification
type OperationMetadata struct {
	Createtime string `json:"createTime,omitempty"` // The time the operation was created.
	Endtime string `json:"endTime,omitempty"` // The time the operation finished running.
	Statusmessage string `json:"statusMessage,omitempty"` // Human-readable status of the operation, if any.
	Target string `json:"target,omitempty"` // Server-defined resource path for the target of the operation.
	Verb string `json:"verb,omitempty"` // Name of the verb executed by the operation.
	Apiversion string `json:"apiVersion,omitempty"` // API version used to start the operation.
	Cancellationrequested bool `json:"cancellationRequested,omitempty"` // Identifies whether the user has requested cancellation of the operation. Operations that have successfully been cancelled have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`.
}

// HttpBody represents the HttpBody schema from the OpenAPI specification
type HttpBody struct {
	Data string `json:"data,omitempty"` // The HTTP request/response body as raw binary.
	Extensions []map[string]interface{} `json:"extensions,omitempty"` // Application specific response metadata. Must be set in the first response for streaming APIs.
	Contenttype string `json:"contentType,omitempty"` // The HTTP Content-Type header value specifying the content type of the body.
}

// Config represents the Config schema from the OpenAPI specification
type Config struct {
	Cmekkeyname string `json:"cmekKeyName,omitempty"` // Required. The Customer Managed Encryption Key (CMEK) used for data encryption. The CMEK name should follow the format of `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`, where the `location` must match InstanceConfig.location.
	Location string `json:"location,omitempty"` // Output only. The GCP location where the Instance resides.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// ListLocationsResponse represents the ListLocationsResponse schema from the OpenAPI specification
type ListLocationsResponse struct {
	Locations []Location `json:"locations,omitempty"` // A list of locations that matches the specified filter in the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
}

// ApiVersion represents the ApiVersion schema from the OpenAPI specification
type ApiVersion struct {
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	State string `json:"state,omitempty"` // A user-definable description of the lifecycle phase of this API version. Format: free-form, but we expect single words that describe API maturity, e.g., "CONCEPT", "DESIGN", "DEVELOPMENT", "STAGING", "PRODUCTION", "DEPRECATED", "RETIRED".
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
	Description string `json:"description,omitempty"` // A detailed description.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Name string `json:"name,omitempty"` // Resource name.
	Primaryspec string `json:"primarySpec,omitempty"` // The primary spec for this version. Format: projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
	Displayname string `json:"displayName,omitempty"` // Human-meaningful name.
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Operations []Operation `json:"operations,omitempty"` // A list of operations that matches the specified filter in the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
}

// TagApiSpecRevisionRequest represents the TagApiSpecRevisionRequest schema from the OpenAPI specification
type TagApiSpecRevisionRequest struct {
	Tag string `json:"tag,omitempty"` // Required. The tag to apply. The tag should be at most 40 characters, and match `a-z{3,39}`.
}

// ListApiSpecRevisionsResponse represents the ListApiSpecRevisionsResponse schema from the OpenAPI specification
type ListApiSpecRevisionsResponse struct {
	Apispecs []ApiSpec `json:"apiSpecs,omitempty"` // The revisions of the spec.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// ListApiDeploymentsResponse represents the ListApiDeploymentsResponse schema from the OpenAPI specification
type ListApiDeploymentsResponse struct {
	Apideployments []ApiDeployment `json:"apiDeployments,omitempty"` // The deployments from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// Api represents the Api schema from the OpenAPI specification
type Api struct {
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
	Description string `json:"description,omitempty"` // A detailed description.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Recommendedversion string `json:"recommendedVersion,omitempty"` // The recommended version of the API. Format: `projects/{project}/locations/{location}/apis/{api}/versions/{version}`
	Displayname string `json:"displayName,omitempty"` // Human-meaningful name.
	Name string `json:"name,omitempty"` // Resource name.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
	Availability string `json:"availability,omitempty"` // A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g., "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores, and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Recommendeddeployment string `json:"recommendedDeployment,omitempty"` // The recommended deployment of the API. Format: `projects/{project}/locations/{location}/apis/{api}/deployments/{deployment}`
}

// ListApiSpecsResponse represents the ListApiSpecsResponse schema from the OpenAPI specification
type ListApiSpecsResponse struct {
	Apispecs []ApiSpec `json:"apiSpecs,omitempty"` // The specs from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// Binding represents the Binding schema from the OpenAPI specification
type Binding struct {
	Members []string `json:"members,omitempty"` // Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
	Role string `json:"role,omitempty"` // Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Condition Expr `json:"condition,omitempty"` // Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
}

// Expr represents the Expr schema from the OpenAPI specification
type Expr struct {
	Description string `json:"description,omitempty"` // Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Expression string `json:"expression,omitempty"` // Textual representation of an expression in Common Expression Language syntax.
	Location string `json:"location,omitempty"` // Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Title string `json:"title,omitempty"` // Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
}

// ListApisResponse represents the ListApisResponse schema from the OpenAPI specification
type ListApisResponse struct {
	Apis []Api `json:"apis,omitempty"` // The APIs from the specified publisher.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
}

// TestIamPermissionsResponse represents the TestIamPermissionsResponse schema from the OpenAPI specification
type TestIamPermissionsResponse struct {
	Permissions []string `json:"permissions,omitempty"` // A subset of `TestPermissionsRequest.permissions` that the caller is allowed.
}

// SetIamPolicyRequest represents the SetIamPolicyRequest schema from the OpenAPI specification
type SetIamPolicyRequest struct {
	Policy Policy `json:"policy,omitempty"` // An Identity and Access Management (IAM) policy, which specifies access controls for Google Cloud resources. A `Policy` is a collection of `bindings`. A `binding` binds one or more `members`, or principals, to a single `role`. Principals can be user accounts, service accounts, Google groups, and domains (such as G Suite). A `role` is a named list of permissions; each `role` can be an IAM predefined role or a user-created custom role. For some types of Google Cloud resources, a `binding` can also specify a `condition`, which is a logical expression that allows access to a resource only if the expression evaluates to `true`. A condition can add constraints based on attributes of the request, the resource, or both. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies). **JSON example:** ``` { "bindings": [ { "role": "roles/resourcemanager.organizationAdmin", "members": [ "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role": "roles/resourcemanager.organizationViewer", "members": [ "user:eve@example.com" ], "condition": { "title": "expirable access", "description": "Does not grant access after Sep 2020", "expression": "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag": "BwWWja0YfJA=", "version": 3 } ``` **YAML example:** ``` bindings: - members: - user:mike@example.com - group:admins@example.com - domain:google.com - serviceAccount:my-project-id@appspot.gserviceaccount.com role: roles/resourcemanager.organizationAdmin - members: - user:eve@example.com role: roles/resourcemanager.organizationViewer condition: title: expirable access description: Does not grant access after Sep 2020 expression: request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA= version: 3 ``` For a description of IAM and its features, see the [IAM documentation](https://cloud.google.com/iam/docs/).
}

// ApiDeployment represents the ApiDeployment schema from the OpenAPI specification
type ApiDeployment struct {
	Externalchanneluri string `json:"externalChannelUri,omitempty"` // The address of the external channel of the API (e.g., the Developer Portal). Changes to this value will not affect the revision.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp; when the deployment resource was created.
	Description string `json:"description,omitempty"` // A detailed description.
	Displayname string `json:"displayName,omitempty"` // Human-meaningful name.
	Name string `json:"name,omitempty"` // Resource name.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Intendedaudience string `json:"intendedAudience,omitempty"` // Text briefly identifying the intended audience of the API. Changes to this value will not affect the revision.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Revisionupdatetime string `json:"revisionUpdateTime,omitempty"` // Output only. Last update timestamp: when the represented revision was last modified.
	Accessguidance string `json:"accessGuidance,omitempty"` // Text briefly describing how to access the endpoint. Changes to this value will not affect the revision.
	Revisioncreatetime string `json:"revisionCreateTime,omitempty"` // Output only. Revision creation timestamp; when the represented revision was created.
	Revisionid string `json:"revisionId,omitempty"` // Output only. Immutable. The revision ID of the deployment. A new revision is committed whenever the deployment contents are changed. The format is an 8-character hexadecimal string.
	Apispecrevision string `json:"apiSpecRevision,omitempty"` // The full resource name (including revision ID) of the spec of the API being served by the deployment. Changes to this value will update the revision. Format: `projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec@revision}`
	Endpointuri string `json:"endpointUri,omitempty"` // The address where the deployment is serving. Changes to this value will update the revision.
}

// ListApiDeploymentRevisionsResponse represents the ListApiDeploymentRevisionsResponse schema from the OpenAPI specification
type ListApiDeploymentRevisionsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	Apideployments []ApiDeployment `json:"apiDeployments,omitempty"` // The revisions of the deployment.
}

// Policy represents the Policy schema from the OpenAPI specification
type Policy struct {
	Bindings []Binding `json:"bindings,omitempty"` // Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Etag string `json:"etag,omitempty"` // `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Version int `json:"version,omitempty"` // Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
}

// ApiSpec represents the ApiSpec schema from the OpenAPI specification
type ApiSpec struct {
	Mimetype string `json:"mimeType,omitempty"` // A style (format) descriptor for this spec that is specified as a [Media Type](https://en.wikipedia.org/wiki/Media_type). Possible values include `application/vnd.apigee.proto`, `application/vnd.apigee.openapi`, and `application/vnd.apigee.graphql`, with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	Contents string `json:"contents,omitempty"` // Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp; when the spec resource was created.
	Description string `json:"description,omitempty"` // A detailed description.
	Filename string `json:"filename,omitempty"` // A possibly-hierarchical name used to refer to the spec from other specs.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
	Name string `json:"name,omitempty"` // Resource name.
	Revisioncreatetime string `json:"revisionCreateTime,omitempty"` // Output only. Revision creation timestamp; when the represented revision was created.
	Hash string `json:"hash,omitempty"` // Output only. A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.
	Revisionupdatetime string `json:"revisionUpdateTime,omitempty"` // Output only. Last update timestamp: when the represented revision was last modified.
	Sizebytes int `json:"sizeBytes,omitempty"` // Output only. The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.
	Sourceuri string `json:"sourceUri,omitempty"` // The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
	Revisionid string `json:"revisionId,omitempty"` // Output only. Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
}

// Build represents the Build schema from the OpenAPI specification
type Build struct {
	Committime string `json:"commitTime,omitempty"` // Output only. Commit time of the latest commit in the build.
	Repo string `json:"repo,omitempty"` // Output only. Path of the open source repository: github.com/apigee/registry.
	Commitid string `json:"commitId,omitempty"` // Output only. Commit ID of the latest commit in the build.
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Annotations map[string]interface{} `json:"annotations,omitempty"` // Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Hash string `json:"hash,omitempty"` // Output only. A SHA-256 hash of the artifact's contents. If the artifact is gzipped, this is the hash of the uncompressed artifact.
	Name string `json:"name,omitempty"` // Resource name.
	Sizebytes int `json:"sizeBytes,omitempty"` // Output only. The size of the artifact in bytes. If the artifact is gzipped, this is the size of the uncompressed artifact.
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
	Labels map[string]interface{} `json:"labels,omitempty"` // Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "registry.googleapis.com/" and cannot be changed.
	Contents string `json:"contents,omitempty"` // Input only. The contents of the artifact. Provided by API callers when artifacts are created or replaced. To access the contents of an artifact, use GetArtifactContents.
	Mimetype string `json:"mimeType,omitempty"` // A content type specifier for the artifact. Content type specifiers are Media Types (https://en.wikipedia.org/wiki/Media_type) with a possible "schema" parameter that specifies a schema for the stored information. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
}

// Instance represents the Instance schema from the OpenAPI specification
type Instance struct {
	Createtime string `json:"createTime,omitempty"` // Output only. Creation timestamp.
	Name string `json:"name,omitempty"` // Format: `projects/*/locations/*/instance`. Currently only `locations/global` is supported.
	State string `json:"state,omitempty"` // Output only. The current state of the Instance.
	Statemessage string `json:"stateMessage,omitempty"` // Output only. Extra information of Instance.State if the state is `FAILED`.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Last update timestamp.
	Build Build `json:"build,omitempty"` // Build information of the Instance if it's in `ACTIVE` state.
	Config Config `json:"config,omitempty"` // Available configurations to provision an Instance.
}
