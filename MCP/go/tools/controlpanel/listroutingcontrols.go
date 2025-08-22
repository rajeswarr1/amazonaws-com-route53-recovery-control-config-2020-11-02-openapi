package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws-route53-recovery-control-config/mcp-server/config"
	"github.com/aws-route53-recovery-control-config/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListroutingcontrolsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/controlpanel/%s/routingcontrols%s", cfg.BaseURL, ControlPanelArn, queryString)
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
		var result models.ListRoutingControlsResponse
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

func CreateListroutingcontrolsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_controlpanel_ControlPanelArn_routingcontrols",
		mcp.WithDescription("Returns an array of routing controls for a control panel. A routing control is an Amazon Route 53 Application Recovery Controller construct that has one of two states: ON and OFF. You can map the routing control state to the state of an Amazon Route 53 health check, which can be used to control routing."),
		mcp.WithString("ControlPanelArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the control panel.")),
		mcp.WithNumber("MaxResults", mcp.Description("The number of objects that you want to return with this call.")),
		mcp.WithString("NextToken", mcp.Description("The token that identifies which batch of results you want to see.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListroutingcontrolsHandler(cfg),
	}
}
