package main

import (
	"github.com/access-approval-api/mcp-server/config"
	"github.com/access-approval-api/mcp-server/models"
	tools_projects "github.com/access-approval-api/mcp-server/tools/projects"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_projects.CreateAccessapproval_projects_approvalrequests_approveTool(cfg),
		tools_projects.CreateAccessapproval_projects_approvalrequests_dismissTool(cfg),
		tools_projects.CreateAccessapproval_projects_approvalrequests_invalidateTool(cfg),
		tools_projects.CreateAccessapproval_projects_approvalrequests_listTool(cfg),
		tools_projects.CreateAccessapproval_projects_deleteaccessapprovalsettingsTool(cfg),
		tools_projects.CreateAccessapproval_projects_approvalrequests_getTool(cfg),
		tools_projects.CreateAccessapproval_projects_updateaccessapprovalsettingsTool(cfg),
	}
}
