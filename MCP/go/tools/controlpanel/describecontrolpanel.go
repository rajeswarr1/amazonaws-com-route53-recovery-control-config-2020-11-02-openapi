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

func DescribecontrolpanelHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		ControlPanelArnVal, ok := args["ControlPanelArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: ControlPanelArn"), nil
		}
		ControlPanelArn, ok := ControlPanelArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: ControlPanelArn"), nil
		}
		url := fmt.Sprintf("%s/controlpanel/%s", cfg.BaseURL, ControlPanelArn)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.DescribeControlPanelResponse
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

func CreateDescribecontrolpanelTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_controlpanel_ControlPanelArn",
		mcp.WithDescription("Displays details about a control panel."),
		mcp.WithString("ControlPanelArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the control panel.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribecontrolpanelHandler(cfg),
	}
}
