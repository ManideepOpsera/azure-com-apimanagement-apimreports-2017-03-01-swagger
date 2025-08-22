package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_reports "github.com/apimanagementclient/mcp-server/tools/reports"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_reports.CreateReports_listbyapiTool(cfg),
		tools_reports.CreateReports_listbygeoTool(cfg),
		tools_reports.CreateReports_listbyoperationTool(cfg),
		tools_reports.CreateReports_listbyproductTool(cfg),
		tools_reports.CreateReports_listbyrequestTool(cfg),
		tools_reports.CreateReports_listbysubscriptionTool(cfg),
		tools_reports.CreateReports_listbytimeTool(cfg),
		tools_reports.CreateReports_listbyuserTool(cfg),
	}
}
