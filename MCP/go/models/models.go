package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListSafetyRulesResponse represents the ListSafetyRulesResponse schema from the OpenAPI specification
type ListSafetyRulesResponse struct {
	Safetyrules interface{} `json:"SafetyRules,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeControlPanelResponse represents the DescribeControlPanelResponse schema from the OpenAPI specification
type DescribeControlPanelResponse struct {
	Controlpanel interface{} `json:"ControlPanel,omitempty"`
}

// Rule represents the Rule schema from the OpenAPI specification
type Rule struct {
	Gating interface{} `json:"GATING,omitempty"`
	Assertion interface{} `json:"ASSERTION,omitempty"`
}

// ListClustersResponse represents the ListClustersResponse schema from the OpenAPI specification
type ListClustersResponse struct {
	Clusters interface{} `json:"Clusters,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateClusterResponse represents the CreateClusterResponse schema from the OpenAPI specification
type CreateClusterResponse struct {
	Cluster interface{} `json:"Cluster,omitempty"`
}

// DescribeControlPanelRequest represents the DescribeControlPanelRequest schema from the OpenAPI specification
type DescribeControlPanelRequest struct {
}

// ListControlPanelsRequest represents the ListControlPanelsRequest schema from the OpenAPI specification
type ListControlPanelsRequest struct {
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// ListAssociatedRoute53HealthChecksRequest represents the ListAssociatedRoute53HealthChecksRequest schema from the OpenAPI specification
type ListAssociatedRoute53HealthChecksRequest struct {
}

// NewAssertionRule represents the NewAssertionRule schema from the OpenAPI specification
type NewAssertionRule struct {
	Ruleconfig interface{} `json:"RuleConfig"`
	Waitperiodms interface{} `json:"WaitPeriodMs"`
	Assertedcontrols interface{} `json:"AssertedControls"`
	Controlpanelarn interface{} `json:"ControlPanelArn"`
	Name interface{} `json:"Name"`
}

// DescribeClusterRequest represents the DescribeClusterRequest schema from the OpenAPI specification
type DescribeClusterRequest struct {
}

// GatingRuleUpdate represents the GatingRuleUpdate schema from the OpenAPI specification
type GatingRuleUpdate struct {
	Name interface{} `json:"Name"`
	Safetyrulearn interface{} `json:"SafetyRuleArn"`
	Waitperiodms interface{} `json:"WaitPeriodMs"`
}

// UpdateControlPanelRequest represents the UpdateControlPanelRequest schema from the OpenAPI specification
type UpdateControlPanelRequest struct {
	Controlpanelarn interface{} `json:"ControlPanelArn"`
	Controlpanelname interface{} `json:"ControlPanelName"`
}

// MapOfstringMin0Max256PatternS represents the MapOfstringMin0Max256PatternS schema from the OpenAPI specification
type MapOfstringMin0Max256PatternS struct {
}

// Cluster represents the Cluster schema from the OpenAPI specification
type Cluster struct {
	Status interface{} `json:"Status,omitempty"`
	Clusterarn interface{} `json:"ClusterArn,omitempty"`
	Clusterendpoints interface{} `json:"ClusterEndpoints,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// AssertionRuleUpdate represents the AssertionRuleUpdate schema from the OpenAPI specification
type AssertionRuleUpdate struct {
	Name interface{} `json:"Name"`
	Safetyrulearn interface{} `json:"SafetyRuleArn"`
	Waitperiodms interface{} `json:"WaitPeriodMs"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// DeleteSafetyRuleResponse represents the DeleteSafetyRuleResponse schema from the OpenAPI specification
type DeleteSafetyRuleResponse struct {
}

// DeleteClusterRequest represents the DeleteClusterRequest schema from the OpenAPI specification
type DeleteClusterRequest struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// ListAssociatedRoute53HealthChecksResponse represents the ListAssociatedRoute53HealthChecksResponse schema from the OpenAPI specification
type ListAssociatedRoute53HealthChecksResponse struct {
	Healthcheckids interface{} `json:"HealthCheckIds,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListSafetyRulesRequest represents the ListSafetyRulesRequest schema from the OpenAPI specification
type ListSafetyRulesRequest struct {
}

// CreateControlPanelResponse represents the CreateControlPanelResponse schema from the OpenAPI specification
type CreateControlPanelResponse struct {
	Controlpanel interface{} `json:"ControlPanel,omitempty"`
}

// DeleteClusterResponse represents the DeleteClusterResponse schema from the OpenAPI specification
type DeleteClusterResponse struct {
}

// DescribeSafetyRuleRequest represents the DescribeSafetyRuleRequest schema from the OpenAPI specification
type DescribeSafetyRuleRequest struct {
}

// ListRoutingControlsRequest represents the ListRoutingControlsRequest schema from the OpenAPI specification
type ListRoutingControlsRequest struct {
}

// UpdateControlPanelResponse represents the UpdateControlPanelResponse schema from the OpenAPI specification
type UpdateControlPanelResponse struct {
	Controlpanel interface{} `json:"ControlPanel,omitempty"`
}

// UpdateSafetyRuleResponse represents the UpdateSafetyRuleResponse schema from the OpenAPI specification
type UpdateSafetyRuleResponse struct {
	Assertionrule interface{} `json:"AssertionRule,omitempty"`
	Gatingrule interface{} `json:"GatingRule,omitempty"`
}

// RoutingControl represents the RoutingControl schema from the OpenAPI specification
type RoutingControl struct {
	Routingcontrolarn interface{} `json:"RoutingControlArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Controlpanelarn interface{} `json:"ControlPanelArn,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DeleteControlPanelRequest represents the DeleteControlPanelRequest schema from the OpenAPI specification
type DeleteControlPanelRequest struct {
}

// NewGatingRule represents the NewGatingRule schema from the OpenAPI specification
type NewGatingRule struct {
	Controlpanelarn interface{} `json:"ControlPanelArn"`
	Gatingcontrols interface{} `json:"GatingControls"`
	Name interface{} `json:"Name"`
	Ruleconfig interface{} `json:"RuleConfig"`
	Targetcontrols interface{} `json:"TargetControls"`
	Waitperiodms interface{} `json:"WaitPeriodMs"`
}

// ControlPanel represents the ControlPanel schema from the OpenAPI specification
type ControlPanel struct {
	Defaultcontrolpanel interface{} `json:"DefaultControlPanel,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Routingcontrolcount interface{} `json:"RoutingControlCount,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Clusterarn interface{} `json:"ClusterArn,omitempty"`
	Controlpanelarn interface{} `json:"ControlPanelArn,omitempty"`
}

// CreateRoutingControlResponse represents the CreateRoutingControlResponse schema from the OpenAPI specification
type CreateRoutingControlResponse struct {
	Routingcontrol interface{} `json:"RoutingControl,omitempty"`
}

// RuleConfig represents the RuleConfig schema from the OpenAPI specification
type RuleConfig struct {
	Inverted interface{} `json:"Inverted"`
	Threshold interface{} `json:"Threshold"`
	TypeField interface{} `json:"Type"`
}

// UpdateRoutingControlRequest represents the UpdateRoutingControlRequest schema from the OpenAPI specification
type UpdateRoutingControlRequest struct {
	Routingcontrolarn interface{} `json:"RoutingControlArn"`
	Routingcontrolname interface{} `json:"RoutingControlName"`
}

// DescribeSafetyRuleResponse represents the DescribeSafetyRuleResponse schema from the OpenAPI specification
type DescribeSafetyRuleResponse struct {
	Assertionrule interface{} `json:"AssertionRule,omitempty"`
	Gatingrule interface{} `json:"GatingRule,omitempty"`
}

// DescribeClusterResponse represents the DescribeClusterResponse schema from the OpenAPI specification
type DescribeClusterResponse struct {
	Cluster interface{} `json:"Cluster,omitempty"`
}

// CreateClusterRequest represents the CreateClusterRequest schema from the OpenAPI specification
type CreateClusterRequest struct {
	Clustername interface{} `json:"ClusterName"`
	Tags interface{} `json:"Tags,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
}

// CreateSafetyRuleRequest represents the CreateSafetyRuleRequest schema from the OpenAPI specification
type CreateSafetyRuleRequest struct {
	Assertionrule interface{} `json:"AssertionRule,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Gatingrule interface{} `json:"GatingRule,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DescribeRoutingControlRequest represents the DescribeRoutingControlRequest schema from the OpenAPI specification
type DescribeRoutingControlRequest struct {
}

// CreateSafetyRuleResponse represents the CreateSafetyRuleResponse schema from the OpenAPI specification
type CreateSafetyRuleResponse struct {
	Assertionrule interface{} `json:"AssertionRule,omitempty"`
	Gatingrule interface{} `json:"GatingRule,omitempty"`
}

// ListRoutingControlsResponse represents the ListRoutingControlsResponse schema from the OpenAPI specification
type ListRoutingControlsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Routingcontrols interface{} `json:"RoutingControls,omitempty"`
}

// ListControlPanelsResponse represents the ListControlPanelsResponse schema from the OpenAPI specification
type ListControlPanelsResponse struct {
	Controlpanels interface{} `json:"ControlPanels,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateControlPanelRequest represents the CreateControlPanelRequest schema from the OpenAPI specification
type CreateControlPanelRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Clusterarn interface{} `json:"ClusterArn"`
	Controlpanelname interface{} `json:"ControlPanelName"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DeleteSafetyRuleRequest represents the DeleteSafetyRuleRequest schema from the OpenAPI specification
type DeleteSafetyRuleRequest struct {
}

// AssertionRule represents the AssertionRule schema from the OpenAPI specification
type AssertionRule struct {
	Ruleconfig interface{} `json:"RuleConfig"`
	Safetyrulearn interface{} `json:"SafetyRuleArn"`
	Status interface{} `json:"Status"`
	Waitperiodms interface{} `json:"WaitPeriodMs"`
	Assertedcontrols interface{} `json:"AssertedControls"`
	Controlpanelarn interface{} `json:"ControlPanelArn"`
	Name interface{} `json:"Name"`
}

// DeleteRoutingControlRequest represents the DeleteRoutingControlRequest schema from the OpenAPI specification
type DeleteRoutingControlRequest struct {
}

// ListClustersRequest represents the ListClustersRequest schema from the OpenAPI specification
type ListClustersRequest struct {
}

// DeleteControlPanelResponse represents the DeleteControlPanelResponse schema from the OpenAPI specification
type DeleteControlPanelResponse struct {
}

// DeleteRoutingControlResponse represents the DeleteRoutingControlResponse schema from the OpenAPI specification
type DeleteRoutingControlResponse struct {
}

// GatingRule represents the GatingRule schema from the OpenAPI specification
type GatingRule struct {
	Waitperiodms interface{} `json:"WaitPeriodMs"`
	Controlpanelarn interface{} `json:"ControlPanelArn"`
	Gatingcontrols interface{} `json:"GatingControls"`
	Name interface{} `json:"Name"`
	Ruleconfig interface{} `json:"RuleConfig"`
	Safetyrulearn interface{} `json:"SafetyRuleArn"`
	Status interface{} `json:"Status"`
	Targetcontrols interface{} `json:"TargetControls"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// UpdateSafetyRuleRequest represents the UpdateSafetyRuleRequest schema from the OpenAPI specification
type UpdateSafetyRuleRequest struct {
	Assertionruleupdate interface{} `json:"AssertionRuleUpdate,omitempty"`
	Gatingruleupdate interface{} `json:"GatingRuleUpdate,omitempty"`
}

// DescribeRoutingControlResponse represents the DescribeRoutingControlResponse schema from the OpenAPI specification
type DescribeRoutingControlResponse struct {
	Routingcontrol interface{} `json:"RoutingControl,omitempty"`
}

// UpdateRoutingControlResponse represents the UpdateRoutingControlResponse schema from the OpenAPI specification
type UpdateRoutingControlResponse struct {
	Routingcontrol interface{} `json:"RoutingControl,omitempty"`
}

// CreateRoutingControlRequest represents the CreateRoutingControlRequest schema from the OpenAPI specification
type CreateRoutingControlRequest struct {
	Controlpanelarn interface{} `json:"ControlPanelArn,omitempty"`
	Routingcontrolname interface{} `json:"RoutingControlName"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Clusterarn interface{} `json:"ClusterArn"`
}

// ClusterEndpoint represents the ClusterEndpoint schema from the OpenAPI specification
type ClusterEndpoint struct {
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Region interface{} `json:"Region,omitempty"`
}
