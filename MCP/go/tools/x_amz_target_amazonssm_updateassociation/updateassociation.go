package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-simple-systems-manager-ssm/mcp-server/config"
	"github.com/amazon-simple-systems-manager-ssm/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdateassociationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.UpdateAssociationRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=AmazonSSM.UpdateAssociation", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

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
		var result models.UpdateAssociationResult
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

func CreateUpdateassociationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AmazonSSM_UpdateAssociation",
		mcp.WithDescription("<p>Updates an association. You can update the association name and version, the document version, schedule, parameters, and Amazon Simple Storage Service (Amazon S3) output. When you call <code>UpdateAssociation</code>, the system removes all optional parameters from the request and overwrites the association with null values for those parameters. This is by design. You must specify all optional parameters in the call, even if you are not changing the parameters. This includes the <code>Name</code> parameter. Before calling this API action, we recommend that you call the <a>DescribeAssociation</a> API operation and make a note of all optional parameters required for your <code>UpdateAssociation</code> call.</p> <p>In order to call this API operation, a user, group, or role must be granted permission to call the <a>DescribeAssociation</a> API operation. If you don't have permission to call <code>DescribeAssociation</code>, then you receive the following error: <code>An error occurred (AccessDeniedException) when calling the UpdateAssociation operation: User: &lt;user_arn&gt; isn't authorized to perform: ssm:DescribeAssociation on resource: &lt;resource_arn&gt;</code> </p> <important> <p>When you update an association, the association immediately runs against the specified targets. You can add the <code>ApplyOnlyAtCronInterval</code> parameter to run the association during the next schedule run.</p> </important>"),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithObject("AlarmConfiguration", mcp.Description("Input parameter: The details for the CloudWatch alarm you want to apply to an automation or command.")),
		mcp.WithString("ScheduleOffset", mcp.Description("")),
		mcp.WithString("AutomationTargetParameterName", mcp.Description("")),
		mcp.WithString("Targets", mcp.Description("")),
		mcp.WithString("TargetLocations", mcp.Description("")),
		mcp.WithString("MaxConcurrency", mcp.Description("")),
		mcp.WithString("Parameters", mcp.Description("")),
		mcp.WithString("AssociationId", mcp.Required(), mcp.Description("")),
		mcp.WithString("AssociationName", mcp.Description("")),
		mcp.WithString("DocumentVersion", mcp.Description("")),
		mcp.WithString("Name", mcp.Description("")),
		mcp.WithString("ScheduleExpression", mcp.Description("")),
		mcp.WithString("ApplyOnlyAtCronInterval", mcp.Description("")),
		mcp.WithString("ComplianceSeverity", mcp.Description("")),
		mcp.WithString("TargetMaps", mcp.Description("")),
		mcp.WithString("AssociationVersion", mcp.Description("")),
		mcp.WithString("MaxErrors", mcp.Description("")),
		mcp.WithString("OutputLocation", mcp.Description("")),
		mcp.WithString("CalendarNames", mcp.Description("")),
		mcp.WithString("SyncCompliance", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateassociationHandler(cfg),
	}
}
