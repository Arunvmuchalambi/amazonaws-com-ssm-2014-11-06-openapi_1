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

func DeleteresourcedatasyncHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.DeleteResourceDataSyncRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=AmazonSSM.DeleteResourceDataSync", cfg.BaseURL)
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
		var result models.DeleteResourceDataSyncResult
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

func CreateDeleteresourcedatasyncTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AmazonSSM_DeleteResourceDataSync",
		mcp.WithDescription("Deletes a resource data sync configuration. After the configuration is deleted, changes to data on managed nodes are no longer synced to or from the target. Deleting a sync configuration doesn't delete data."),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("SyncName", mcp.Required(), mcp.Description("")),
		mcp.WithString("SyncType", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteresourcedatasyncHandler(cfg),
	}
}
