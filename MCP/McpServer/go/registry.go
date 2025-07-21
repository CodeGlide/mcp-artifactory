package main

import (
	"github.com/openai/mcp-server/config"
	"github.com/openai/mcp-server/models"
	tools_moderations "github.com/openai/mcp-server/tools/moderations"
	tools_files "github.com/openai/mcp-server/tools/files"
	tools_completions "github.com/openai/mcp-server/tools/completions"
	tools_fine_tunes "github.com/openai/mcp-server/tools/fine_tunes"
	tools_edits "github.com/openai/mcp-server/tools/edits"
	tools_engines "github.com/openai/mcp-server/tools/engines"
	tools_models "github.com/openai/mcp-server/tools/models"
	tools_embeddings "github.com/openai/mcp-server/tools/embeddings"
	tools_images "github.com/openai/mcp-server/tools/images"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_moderations.CreatePost_moderationsTool(cfg),
		tools_files.CreateGet_files_file_idTool(cfg),
		tools_files.CreateDelete_files_file_idTool(cfg),
		tools_completions.CreatePost_completionsTool(cfg),
		tools_fine_tunes.CreatePost_fine_tunes_fine_tune_id_cancelTool(cfg),
		tools_edits.CreatePost_editsTool(cfg),
		tools_engines.CreateGet_engines_engineidTool(cfg),
		tools_files.CreateGet_filesTool(cfg),
		tools_fine_tunes.CreateDelete_models_modelTool(cfg),
		tools_fine_tunes.CreateGet_fine_tunes_fine_tune_idTool(cfg),
		tools_models.CreateGet_models_modelidTool(cfg),
		tools_fine_tunes.CreateGet_fine_tunes_fine_tune_id_eventsTool(cfg),
		tools_embeddings.CreatePost_embeddingsTool(cfg),
		tools_files.CreateGet_files_file_id_contentTool(cfg),
		tools_images.CreatePost_images_generationsTool(cfg),
		tools_engines.CreateGet_enginesTool(cfg),
		tools_fine_tunes.CreateGet_fine_tunesTool(cfg),
		tools_fine_tunes.CreatePost_fine_tunesTool(cfg),
		tools_models.CreateGet_modelsTool(cfg),
	}
}
