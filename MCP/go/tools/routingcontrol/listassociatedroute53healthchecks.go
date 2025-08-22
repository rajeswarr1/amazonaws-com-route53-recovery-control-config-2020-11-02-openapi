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

func Listassociatedroute53healthchecksHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/routingcontrol/%s/associatedRoute53HealthChecks%s", cfg.BaseURL, RoutingControlArn, queryString)
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
		var result models.ListAssociatedRoute53HealthChecksResponse
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

func CreateListassociatedroute53healthchecksTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_routingcontrol_RoutingControlArn_associatedRoute53HealthChecks",
		mcp.WithDescription("Returns an array of all Amazon Route 53 health checks associated with a specific routing control."),
		mcp.WithNumber("MaxResults", mcp.Description("The number of objects that you want to return with this call.")),
		mcp.WithString("NextToken", mcp.Description("The token that identifies which batch of results you want to see.")),
		mcp.WithString("RoutingControlArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the routing control.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Listassociatedroute53healthchecksHandler(cfg),
	}
}
