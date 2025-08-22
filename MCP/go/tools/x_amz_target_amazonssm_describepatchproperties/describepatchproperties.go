package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/amazon-simple-systems-manager-ssm/mcp-server/config"
	"github.com/amazon-simple-systems-manager-ssm/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DescribepatchpropertiesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.DescribePatchPropertiesRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=AmazonSSM.DescribePatchProperties%s", cfg.BaseURL, queryString)
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
		var result models.DescribePatchPropertiesResult
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

func CreateDescribepatchpropertiesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AmazonSSM_DescribePatchProperties",
		mcp.WithDescription("<p>Lists the properties of available patches organized by product, product family, classification, severity, and other properties of available patches. You can use the reported properties in the filters you specify in requests for operations such as <a>CreatePatchBaseline</a>, <a>UpdatePatchBaseline</a>, <a>DescribeAvailablePatches</a>, and <a>DescribePatchBaselines</a>.</p> <p>The following section lists the properties that can be used in filters for each major operating system type:</p> <dl> <dt>AMAZON_LINUX</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> | <code>SEVERITY</code> </p> </dd> <dt>AMAZON_LINUX_2</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> | <code>SEVERITY</code> </p> </dd> <dt>CENTOS</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> | <code>SEVERITY</code> </p> </dd> <dt>DEBIAN</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>PRIORITY</code> </p> </dd> <dt>MACOS</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> </p> </dd> <dt>ORACLE_LINUX</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> | <code>SEVERITY</code> </p> </dd> <dt>REDHAT_ENTERPRISE_LINUX</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> | <code>SEVERITY</code> </p> </dd> <dt>SUSE</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>CLASSIFICATION</code> | <code>SEVERITY</code> </p> </dd> <dt>UBUNTU</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>PRIORITY</code> </p> </dd> <dt>WINDOWS</dt> <dd> <p>Valid properties: <code>PRODUCT</code> | <code>PRODUCT_FAMILY</code> | <code>CLASSIFICATION</code> | <code>MSRC_SEVERITY</code> </p> </dd> </dl>"),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("NextToken", mcp.Description("")),
		mcp.WithString("OperatingSystem", mcp.Required(), mcp.Description("")),
		mcp.WithString("PatchSet", mcp.Description("")),
		mcp.WithString("Property", mcp.Required(), mcp.Description("")),
		mcp.WithString("MaxResults", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribepatchpropertiesHandler(cfg),
	}
}
