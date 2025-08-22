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

func CreateopsitemHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateOpsItemRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=AmazonSSM.CreateOpsItem", cfg.BaseURL)
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
		var result models.CreateOpsItemResponse
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

func CreateCreateopsitemTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AmazonSSM_CreateOpsItem",
		mcp.WithDescription("<p>Creates a new OpsItem. You must have permission in Identity and Access Management (IAM) to create a new OpsItem. For more information, see <a href="https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-setup.html">Set up OpsCenter</a> in the <i>Amazon Web Services Systems Manager User Guide</i>.</p> <p>Operations engineers and IT professionals use Amazon Web Services Systems Manager OpsCenter to view, investigate, and remediate operational issues impacting the performance and health of their Amazon Web Services resources. For more information, see <a href="https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html">Amazon Web Services Systems Manager OpsCenter</a> in the <i>Amazon Web Services Systems Manager User Guide</i>. </p>"),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("PlannedStartTime", mcp.Description("")),
		mcp.WithString("RelatedOpsItems", mcp.Description("")),
		mcp.WithString("Tags", mcp.Description("")),
		mcp.WithString("Notifications", mcp.Description("")),
		mcp.WithString("PlannedEndTime", mcp.Description("")),
		mcp.WithString("OpsItemType", mcp.Description("")),
		mcp.WithString("AccountId", mcp.Description("")),
		mcp.WithString("OperationalData", mcp.Description("")),
		mcp.WithString("Priority", mcp.Description("")),
		mcp.WithString("Severity", mcp.Description("")),
		mcp.WithString("ActualEndTime", mcp.Description("")),
		mcp.WithString("Title", mcp.Required(), mcp.Description("")),
		mcp.WithString("Description", mcp.Required(), mcp.Description("")),
		mcp.WithString("Source", mcp.Required(), mcp.Description("")),
		mcp.WithString("ActualStartTime", mcp.Description("")),
		mcp.WithString("Category", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateopsitemHandler(cfg),
	}
}
