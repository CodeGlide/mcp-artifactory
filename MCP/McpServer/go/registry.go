package main

import (
	"github.com/payroll/mcp-server/config"
	"github.com/payroll/mcp-server/models"
	tools_jobs "github.com/payroll/mcp-server/tools/jobs"
	tools_prompt_values "github.com/payroll/mcp-server/tools/prompt_values"
	tools_paygroupdetails "github.com/payroll/mcp-server/tools/paygroupdetails"
	tools_taxrates "github.com/payroll/mcp-server/tools/taxrates"
	tools_paygroups "github.com/payroll/mcp-server/tools/paygroups"
	tools_payrollinputs "github.com/payroll/mcp-server/tools/payrollinputs"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_jobs.CreateGet_jobs_idTool(cfg),
		tools_prompt_values.CreateGet_values_payrollinputsgroup_worktagsTool(cfg),
		tools_prompt_values.CreateGet_values_payrollinputsgroup_runcategoriesTool(cfg),
		tools_paygroupdetails.CreateGet_paygroupdetails_idTool(cfg),
		tools_prompt_values.CreateGet_values_payrollinputsgroup_positionsTool(cfg),
		tools_jobs.CreateGet_jobs_id_paygroupTool(cfg),
		tools_taxrates.CreateGet_taxratesTool(cfg),
		tools_taxrates.CreatePost_taxratesTool(cfg),
		tools_paygroups.CreateGet_paygroupsTool(cfg),
		tools_prompt_values.CreateGet_values_payrollinputsgroup_paycomponentsTool(cfg),
		tools_paygroups.CreateGet_paygroups_idTool(cfg),
		tools_prompt_values.CreateGet_values_taxratesgroup_companyinstancesTool(cfg),
		tools_prompt_values.CreateGet_values_taxratesgroup_stateinstancesTool(cfg),
		tools_payrollinputs.CreateGet_payrollinputsTool(cfg),
		tools_payrollinputs.CreatePost_payrollinputsTool(cfg),
		tools_jobs.CreateGet_jobsTool(cfg),
		tools_paygroupdetails.CreateGet_paygroupdetailsTool(cfg),
		tools_payrollinputs.CreateDelete_payrollinputs_idTool(cfg),
		tools_payrollinputs.CreateGet_payrollinputs_idTool(cfg),
		tools_payrollinputs.CreatePatch_payrollinputs_idTool(cfg),
		tools_jobs.CreateGet_jobs_id_paygroup_subresourceidTool(cfg),
	}
}
