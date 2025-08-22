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

func DescriberoutingcontrolHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		RoutingControlArnVal, ok := args["RoutingControlArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: RoutingControlArn"), nil
		}
		RoutingControlArn, ok := RoutingControlArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: RoutingControlArn"), nil
		}
		url := fmt.Sprintf("%s/routingcontrol/%s", cfg.BaseURL, RoutingControlArn)
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
		var result models.DescribeRoutingControlResponse
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

func CreateDescriberoutingcontrolTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_routingcontrol_RoutingControlArn",
		mcp.WithDescription("<p>Displays details about a routing control. A routing control has one of two states: ON and OFF. You can map the routing control state to the state of an Amazon Route 53 health check, which can be used to control routing.</p> <p>To get or update the routing control state, see the Recovery Cluster (data plane) API actions for Amazon Route 53 Application Recovery Controller.</p>"),
		mcp.WithString("RoutingControlArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the routing control.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescriberoutingcontrolHandler(cfg),
	}
}
