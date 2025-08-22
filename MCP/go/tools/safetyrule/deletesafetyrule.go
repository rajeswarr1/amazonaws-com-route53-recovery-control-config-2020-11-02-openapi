package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws-route53-recovery-control-config/mcp-server/config"
	"github.com/aws-route53-recovery-control-config/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeletesafetyruleHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		SafetyRuleArnVal, ok := args["SafetyRuleArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: SafetyRuleArn"), nil
		}
		SafetyRuleArn, ok := SafetyRuleArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: SafetyRuleArn"), nil
		}
		url := fmt.Sprintf("%s/safetyrule/%s", cfg.BaseURL, SafetyRuleArn)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
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
		var result models.DeleteSafetyRuleResponse
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

func CreateDeletesafetyruleTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_safetyrule_SafetyRuleArn",
		mcp.WithDescription("<p>Deletes a safety rule.</p>/&gt;"),
		mcp.WithString("SafetyRuleArn", mcp.Required(), mcp.Description("The ARN of the safety rule.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletesafetyruleHandler(cfg),
	}
}
