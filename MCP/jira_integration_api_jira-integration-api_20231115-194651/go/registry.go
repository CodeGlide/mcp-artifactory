package main

import (
	"github.com/instagram/mcp-server/config"
	"github.com/instagram/mcp-server/models"
	tools_tags_tag_name "github.com/instagram/mcp-server/tools/tags_tag_name"
	tools_media_media_id_comments "github.com/instagram/mcp-server/tools/media_media_id_comments"
	tools_users_user_id "github.com/instagram/mcp-server/tools/users_user_id"
	tools_locations "github.com/instagram/mcp-server/tools/locations"
	tools_tags "github.com/instagram/mcp-server/tools/tags"
	tools_users "github.com/instagram/mcp-server/tools/users"
	tools_locations_location_id "github.com/instagram/mcp-server/tools/locations_location_id"
	tools_media_media_id "github.com/instagram/mcp-server/tools/media_media_id"
	tools_users_self "github.com/instagram/mcp-server/tools/users_self"
	tools_media "github.com/instagram/mcp-server/tools/media"
	tools_general "github.com/instagram/mcp-server/tools/general"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tags_tag_name.CreateGet_tags_tag_nameTool(cfg),
		tools_media_media_id_comments.CreateDelete_media_media_id_comments_comment_idTool(cfg),
		tools_tags_tag_name.CreateGet_tags_tag_name_media_recentTool(cfg),
		tools_users_user_id.CreateGet_users_user_id_followed_byTool(cfg),
		tools_media_media_id_comments.CreateGet_media_media_id_commentsTool(cfg),
		tools_locations.CreateGet_locations_searchTool(cfg),
		tools_tags.CreateGet_tags_searchTool(cfg),
		tools_users.CreateGet_users_searchTool(cfg),
		tools_locations_location_id.CreateGet_locations_location_idTool(cfg),
		tools_users_user_id.CreateGet_users_user_id_followsTool(cfg),
		tools_users_user_id.CreateGet_users_user_idTool(cfg),
		tools_locations_location_id.CreateGet_locations_location_id_media_recentTool(cfg),
		tools_users_user_id.CreateGet_users_user_id_media_recentTool(cfg),
		tools_media_media_id.CreateGet_media_media_idTool(cfg),
		tools_users_self.CreateGet_users_self_feedTool(cfg),
		tools_users_self.CreateGet_users_self_media_likedTool(cfg),
		tools_media.CreateGet_media_popularTool(cfg),
		tools_general.CreateGet_geographies_geo_id_media_recentTool(cfg),
		tools_media.CreateGet_media_searchTool(cfg),
		tools_media_media_id.CreateDelete_media_media_id_likesTool(cfg),
		tools_users_self.CreateGet_users_self_requested_byTool(cfg),
		tools_users_user_id.CreateGet_users_user_id_relationshipTool(cfg),
		tools_media.CreateGet_media_shortcode_shortcodeTool(cfg),
	}
}
