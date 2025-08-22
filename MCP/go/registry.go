package main

import (
	"github.com/aws-route53-recovery-control-config/mcp-server/config"
	"github.com/aws-route53-recovery-control-config/mcp-server/models"
	tools_safetyrule "github.com/aws-route53-recovery-control-config/mcp-server/tools/safetyrule"
	tools_cluster "github.com/aws-route53-recovery-control-config/mcp-server/tools/cluster"
	tools_controlpanel "github.com/aws-route53-recovery-control-config/mcp-server/tools/controlpanel"
	tools_tags "github.com/aws-route53-recovery-control-config/mcp-server/tools/tags"
	tools_controlpanels "github.com/aws-route53-recovery-control-config/mcp-server/tools/controlpanels"
	tools_routingcontrol "github.com/aws-route53-recovery-control-config/mcp-server/tools/routingcontrol"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_safetyrule.CreateDeletesafetyruleTool(cfg),
		tools_safetyrule.CreateDescribesafetyruleTool(cfg),
		tools_cluster.CreateListclustersTool(cfg),
		tools_cluster.CreateCreateclusterTool(cfg),
		tools_controlpanel.CreateCreatecontrolpanelTool(cfg),
		tools_controlpanel.CreateUpdatecontrolpanelTool(cfg),
		tools_controlpanel.CreateListroutingcontrolsTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_controlpanel.CreateListsafetyrulesTool(cfg),
		tools_safetyrule.CreateCreatesafetyruleTool(cfg),
		tools_safetyrule.CreateUpdatesafetyruleTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_cluster.CreateDeleteclusterTool(cfg),
		tools_cluster.CreateDescribeclusterTool(cfg),
		tools_controlpanels.CreateListcontrolpanelsTool(cfg),
		tools_routingcontrol.CreateCreateroutingcontrolTool(cfg),
		tools_routingcontrol.CreateUpdateroutingcontrolTool(cfg),
		tools_controlpanel.CreateDeletecontrolpanelTool(cfg),
		tools_controlpanel.CreateDescribecontrolpanelTool(cfg),
		tools_routingcontrol.CreateDescriberoutingcontrolTool(cfg),
		tools_routingcontrol.CreateDeleteroutingcontrolTool(cfg),
		tools_routingcontrol.CreateListassociatedroute53healthchecksTool(cfg),
	}
}
