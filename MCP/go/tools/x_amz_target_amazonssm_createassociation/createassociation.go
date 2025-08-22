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

func CreateassociationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateAssociationRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=AmazonSSM.CreateAssociation", cfg.BaseURL)
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
		var result models.CreateAssociationResult
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

func CreateCreateassociationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AmazonSSM_CreateAssociation",
		mcp.WithDescription("A State Manager association defines the state that you want to maintain on your managed nodes. For example, an association can specify that anti-virus software must be installed and running on your managed nodes, or that certain ports must be closed. For static targets, the association specifies a schedule for when the configuration is reapplied. For dynamic targets, such as an Amazon Web Services resource group or an Amazon Web Services autoscaling group, State Manager, a capability of Amazon Web Services Systems Manager applies the configuration when new managed nodes are added to the group. The association also specifies actions to take when applying the configuration. For example, an association for anti-virus software might run once a day. If the software isn't installed, then State Manager installs it. If the software is installed, but the service isn't running, then the association might instruct State Manager to start the service. "),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("Parameters", mcp.Description("")),
		mcp.WithString("MaxConcurrency", mcp.Description("")),
		mcp.WithString("OutputLocation", mcp.Description("")),
		mcp.WithString("ApplyOnlyAtCronInterval", mcp.Description("")),
		mcp.WithString("AutomationTargetParameterName", mcp.Description("")),
		mcp.WithString("Tags", mcp.Description("")),
		mcp.WithString("TargetMaps", mcp.Description("")),
		mcp.WithString("ScheduleExpression", mcp.Description("")),
		mcp.WithString("ScheduleOffset", mcp.Description("")),
		mcp.WithString("Targets", mcp.Description("")),
		mcp.WithString("AssociationName", mcp.Description("")),
		mcp.WithString("MaxErrors", mcp.Description("")),
		mcp.WithString("DocumentVersion", mcp.Description("")),
		mcp.WithString("ComplianceSeverity", mcp.Description("")),
		mcp.WithString("CalendarNames", mcp.Description("")),
		mcp.WithString("SyncCompliance", mcp.Description("")),
		mcp.WithObject("AlarmConfiguration", mcp.Description("Input parameter: The details for the CloudWatch alarm you want to apply to an automation or command.")),
		mcp.WithString("InstanceId", mcp.Description("")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("")),
		mcp.WithString("TargetLocations", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateassociationHandler(cfg),
	}
}
