package main

import (
	"github.com/slack-web-api/mcp-server/config"
	"github.com/slack-web-api/mcp-server/models"
	tools_admin_usergroups "github.com/slack-web-api/mcp-server/tools/admin_usergroups"
	tools_chat_scheduledmessages "github.com/slack-web-api/mcp-server/tools/chat_scheduledmessages"
	tools_pins "github.com/slack-web-api/mcp-server/tools/pins"
	tools_usergroups "github.com/slack-web-api/mcp-server/tools/usergroups"
	tools_views "github.com/slack-web-api/mcp-server/tools/views"
	tools_calls "github.com/slack-web-api/mcp-server/tools/calls"
	tools_chat "github.com/slack-web-api/mcp-server/tools/chat"
	tools_apps_permissions_users "github.com/slack-web-api/mcp-server/tools/apps_permissions_users"
	tools_admin_users "github.com/slack-web-api/mcp-server/tools/admin_users"
	tools_dnd "github.com/slack-web-api/mcp-server/tools/dnd"
	tools_admin_conversations "github.com/slack-web-api/mcp-server/tools/admin_conversations"
	tools_admin_teams_admins "github.com/slack-web-api/mcp-server/tools/admin_teams_admins"
	tools_dialog "github.com/slack-web-api/mcp-server/tools/dialog"
	tools_admin_apps "github.com/slack-web-api/mcp-server/tools/admin_apps"
	tools_admin_teams_settings "github.com/slack-web-api/mcp-server/tools/admin_teams_settings"
	tools_conversations "github.com/slack-web-api/mcp-server/tools/conversations"
	tools_admin_inviterequests "github.com/slack-web-api/mcp-server/tools/admin_inviterequests"
	tools_files_remote "github.com/slack-web-api/mcp-server/tools/files_remote"
	tools_admin_users_session "github.com/slack-web-api/mcp-server/tools/admin_users_session"
	tools_users "github.com/slack-web-api/mcp-server/tools/users"
	tools_rtm "github.com/slack-web-api/mcp-server/tools/rtm"
	tools_apps_permissions_scopes "github.com/slack-web-api/mcp-server/tools/apps_permissions_scopes"
	tools_admin_conversations_ekm "github.com/slack-web-api/mcp-server/tools/admin_conversations_ekm"
	tools_reactions "github.com/slack-web-api/mcp-server/tools/reactions"
	tools_workflows "github.com/slack-web-api/mcp-server/tools/workflows"
	tools_search "github.com/slack-web-api/mcp-server/tools/search"
	tools_admin_inviterequests_approved "github.com/slack-web-api/mcp-server/tools/admin_inviterequests_approved"
	tools_reminders "github.com/slack-web-api/mcp-server/tools/reminders"
	tools_apps "github.com/slack-web-api/mcp-server/tools/apps"
	tools_team "github.com/slack-web-api/mcp-server/tools/team"
	tools_admin_conversations_restrictaccess "github.com/slack-web-api/mcp-server/tools/admin_conversations_restrictaccess"
	tools_files "github.com/slack-web-api/mcp-server/tools/files"
	tools_admin_apps_restricted "github.com/slack-web-api/mcp-server/tools/admin_apps_restricted"
	tools_auth "github.com/slack-web-api/mcp-server/tools/auth"
	tools_migration "github.com/slack-web-api/mcp-server/tools/migration"
	tools_emoji "github.com/slack-web-api/mcp-server/tools/emoji"
	tools_usergroups_users "github.com/slack-web-api/mcp-server/tools/usergroups_users"
	tools_bots "github.com/slack-web-api/mcp-server/tools/bots"
	tools_apps_event_authorizations "github.com/slack-web-api/mcp-server/tools/apps_event_authorizations"
	tools_apps_permissions "github.com/slack-web-api/mcp-server/tools/apps_permissions"
	tools_oauth "github.com/slack-web-api/mcp-server/tools/oauth"
	tools_stars "github.com/slack-web-api/mcp-server/tools/stars"
	tools_admin_apps_approved "github.com/slack-web-api/mcp-server/tools/admin_apps_approved"
	tools_admin_teams_owners "github.com/slack-web-api/mcp-server/tools/admin_teams_owners"
	tools_admin_teams "github.com/slack-web-api/mcp-server/tools/admin_teams"
	tools_admin_inviterequests_denied "github.com/slack-web-api/mcp-server/tools/admin_inviterequests_denied"
	tools_oauth_v2 "github.com/slack-web-api/mcp-server/tools/oauth_v2"
	tools_users_profile "github.com/slack-web-api/mcp-server/tools/users_profile"
	tools_calls_participants "github.com/slack-web-api/mcp-server/tools/calls_participants"
	tools_team_profile "github.com/slack-web-api/mcp-server/tools/team_profile"
	tools_apps_permissions_resources "github.com/slack-web-api/mcp-server/tools/apps_permissions_resources"
	tools_admin_emoji "github.com/slack-web-api/mcp-server/tools/admin_emoji"
	tools_files_comments "github.com/slack-web-api/mcp-server/tools/files_comments"
	tools_api "github.com/slack-web-api/mcp-server/tools/api"
	tools_admin_apps_requests "github.com/slack-web-api/mcp-server/tools/admin_apps_requests"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_admin_usergroups.CreateAdmin_usergroups_addteamsTool(cfg),
		tools_chat_scheduledmessages.CreateChat_scheduledmessages_listTool(cfg),
		tools_pins.CreatePins_addTool(cfg),
		tools_usergroups.CreateUsergroups_enableTool(cfg),
		tools_views.CreateViews_pushTool(cfg),
		tools_calls.CreateCalls_infoTool(cfg),
		tools_chat.CreateChat_schedulemessageTool(cfg),
		tools_apps_permissions_users.CreateApps_permissions_users_requestTool(cfg),
		tools_admin_users.CreateAdmin_users_setregularTool(cfg),
		tools_dnd.CreateDnd_teaminfoTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_disconnectsharedTool(cfg),
		tools_admin_teams_admins.CreateAdmin_teams_admins_listTool(cfg),
		tools_dialog.CreateDialog_openTool(cfg),
		tools_admin_users.CreateAdmin_users_removeTool(cfg),
		tools_admin_apps.CreateAdmin_apps_restrictTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_renameTool(cfg),
		tools_admin_teams_settings.CreateAdmin_teams_settings_setnameTool(cfg),
		tools_conversations.CreateConversations_joinTool(cfg),
		tools_admin_inviterequests.CreateAdmin_inviterequests_approveTool(cfg),
		tools_files_remote.CreateFiles_remote_infoTool(cfg),
		tools_admin_users_session.CreateAdmin_users_session_resetTool(cfg),
		tools_users.CreateUsers_listTool(cfg),
		tools_admin_usergroups.CreateAdmin_usergroups_addchannelsTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_deleteTool(cfg),
		tools_conversations.CreateConversations_historyTool(cfg),
		tools_admin_usergroups.CreateAdmin_usergroups_listchannelsTool(cfg),
		tools_chat.CreateChat_deleteTool(cfg),
		tools_usergroups.CreateUsergroups_createTool(cfg),
		tools_rtm.CreateRtm_connectTool(cfg),
		tools_apps_permissions_scopes.CreateApps_permissions_scopes_listTool(cfg),
		tools_admin_conversations_ekm.CreateAdmin_conversations_ekm_listoriginalconnectedchannelinfoTool(cfg),
		tools_reactions.CreateReactions_listTool(cfg),
		tools_conversations.CreateConversations_markTool(cfg),
		tools_workflows.CreateWorkflows_stepcompletedTool(cfg),
		tools_conversations.CreateConversations_listTool(cfg),
		tools_chat.CreateChat_postephemeralTool(cfg),
		tools_chat.CreateChat_updateTool(cfg),
		tools_search.CreateSearch_messagesTool(cfg),
		tools_admin_inviterequests_approved.CreateAdmin_inviterequests_approved_listTool(cfg),
		tools_dnd.CreateDnd_enddndTool(cfg),
		tools_views.CreateViews_publishTool(cfg),
		tools_reminders.CreateReminders_addTool(cfg),
		tools_admin_teams_settings.CreateAdmin_teams_settings_setdiscoverabilityTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_inviteTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_searchTool(cfg),
		tools_conversations.CreateConversations_infoTool(cfg),
		tools_users.CreateUsers_conversationsTool(cfg),
		tools_apps.CreateApps_uninstallTool(cfg),
		tools_conversations.CreateConversations_membersTool(cfg),
		tools_admin_usergroups.CreateAdmin_usergroups_removechannelsTool(cfg),
		tools_reactions.CreateReactions_removeTool(cfg),
		tools_team.CreateTeam_integrationlogsTool(cfg),
		tools_admin_conversations_restrictaccess.CreateAdmin_conversations_restrictaccess_listgroupsTool(cfg),
		tools_reminders.CreateReminders_infoTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_unarchiveTool(cfg),
		tools_conversations.CreateConversations_archiveTool(cfg),
		tools_conversations.CreateConversations_kickTool(cfg),
		tools_files.CreateFiles_sharedpublicurlTool(cfg),
		tools_users.CreateUsers_infoTool(cfg),
		tools_usergroups.CreateUsergroups_listTool(cfg),
		tools_reactions.CreateReactions_addTool(cfg),
		tools_admin_apps_restricted.CreateAdmin_apps_restricted_listTool(cfg),
		tools_reactions.CreateReactions_getTool(cfg),
		tools_files_remote.CreateFiles_remote_shareTool(cfg),
		tools_admin_users.CreateAdmin_users_setexpirationTool(cfg),
		tools_auth.CreateAuth_revokeTool(cfg),
		tools_conversations.CreateConversations_closeTool(cfg),
		tools_users.CreateUsers_lookupbyemailTool(cfg),
		tools_files.CreateFiles_listTool(cfg),
		tools_admin_teams_settings.CreateAdmin_teams_settings_setdescriptionTool(cfg),
		tools_conversations.CreateConversations_inviteTool(cfg),
		tools_migration.CreateMigration_exchangeTool(cfg),
		tools_emoji.CreateEmoji_listTool(cfg),
		tools_conversations.CreateConversations_openTool(cfg),
		tools_conversations.CreateConversations_repliesTool(cfg),
		tools_admin_users_session.CreateAdmin_users_session_invalidateTool(cfg),
		tools_reminders.CreateReminders_listTool(cfg),
		tools_conversations.CreateConversations_settopicTool(cfg),
		tools_calls.CreateCalls_updateTool(cfg),
		tools_chat.CreateChat_postmessageTool(cfg),
		tools_users.CreateUsers_setpresenceTool(cfg),
		tools_usergroups_users.CreateUsergroups_users_updateTool(cfg),
		tools_bots.CreateBots_infoTool(cfg),
		tools_conversations.CreateConversations_unarchiveTool(cfg),
		tools_admin_inviterequests.CreateAdmin_inviterequests_denyTool(cfg),
		tools_admin_users.CreateAdmin_users_setownerTool(cfg),
		tools_apps_event_authorizations.CreateApps_event_authorizations_listTool(cfg),
		tools_chat.CreateChat_unfurlTool(cfg),
		tools_apps_permissions.CreateApps_permissions_infoTool(cfg),
		tools_oauth.CreateOauth_accessTool(cfg),
		tools_stars.CreateStars_removeTool(cfg),
		tools_admin_apps_approved.CreateAdmin_apps_approved_listTool(cfg),
		tools_calls.CreateCalls_endTool(cfg),
		tools_team.CreateTeam_accesslogsTool(cfg),
		tools_admin_teams_owners.CreateAdmin_teams_owners_listTool(cfg),
		tools_admin_teams.CreateAdmin_teams_listTool(cfg),
		tools_usergroups_users.CreateUsergroups_users_listTool(cfg),
		tools_conversations.CreateConversations_leaveTool(cfg),
		tools_admin_inviterequests_denied.CreateAdmin_inviterequests_denied_listTool(cfg),
		tools_workflows.CreateWorkflows_updatestepTool(cfg),
		tools_oauth_v2.CreateOauth_v2_accessTool(cfg),
		tools_chat.CreateChat_getpermalinkTool(cfg),
		tools_team.CreateTeam_billableinfoTool(cfg),
		tools_apps_permissions_users.CreateApps_permissions_users_listTool(cfg),
		tools_pins.CreatePins_removeTool(cfg),
		tools_stars.CreateStars_listTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_converttoprivateTool(cfg),
		tools_users_profile.CreateUsers_profile_getTool(cfg),
		tools_calls_participants.CreateCalls_participants_removeTool(cfg),
		tools_team.CreateTeam_infoTool(cfg),
		tools_team_profile.CreateTeam_profile_getTool(cfg),
		tools_files_remote.CreateFiles_remote_listTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_setteamsTool(cfg),
		tools_apps_permissions_resources.CreateApps_permissions_resources_listTool(cfg),
		tools_reminders.CreateReminders_deleteTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_setconversationprefsTool(cfg),
		tools_auth.CreateAuth_test_handlerTool(cfg),
		tools_admin_emoji.CreateAdmin_emoji_listTool(cfg),
		tools_views.CreateViews_openTool(cfg),
		tools_users.CreateUsers_setactiveTool(cfg),
		tools_files_comments.CreateFiles_comments_deleteTool(cfg),
		tools_users.CreateUsers_getpresenceTool(cfg),
		tools_admin_teams_settings.CreateAdmin_teams_settings_infoTool(cfg),
		tools_chat.CreateChat_memessageTool(cfg),
		tools_files.CreateFiles_infoTool(cfg),
		tools_admin_inviterequests.CreateAdmin_inviterequests_listTool(cfg),
		tools_views.CreateViews_updateTool(cfg),
		tools_files.CreateFiles_deleteTool(cfg),
		tools_admin_users.CreateAdmin_users_inviteTool(cfg),
		tools_usergroups.CreateUsergroups_updateTool(cfg),
		tools_files.CreateFiles_revokepublicurlTool(cfg),
		tools_admin_apps.CreateAdmin_apps_approveTool(cfg),
		tools_chat.CreateChat_deletescheduledmessageTool(cfg),
		tools_api.CreateApi_test_handlerTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_archiveTool(cfg),
		tools_admin_users.CreateAdmin_users_setadminTool(cfg),
		tools_oauth.CreateOauth_tokenTool(cfg),
		tools_admin_users.CreateAdmin_users_listTool(cfg),
		tools_pins.CreatePins_listTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_getteamsTool(cfg),
		tools_apps_permissions.CreateApps_permissions_requestTool(cfg),
		tools_dnd.CreateDnd_infoTool(cfg),
		tools_conversations.CreateConversations_setpurposeTool(cfg),
		tools_workflows.CreateWorkflows_stepfailedTool(cfg),
		tools_stars.CreateStars_addTool(cfg),
		tools_dnd.CreateDnd_endsnoozeTool(cfg),
		tools_usergroups.CreateUsergroups_disableTool(cfg),
		tools_conversations.CreateConversations_renameTool(cfg),
		tools_users.CreateUsers_identityTool(cfg),
		tools_reminders.CreateReminders_completeTool(cfg),
		tools_users_profile.CreateUsers_profile_setTool(cfg),
		tools_admin_users.CreateAdmin_users_assignTool(cfg),
		tools_admin_apps_requests.CreateAdmin_apps_requests_listTool(cfg),
		tools_calls.CreateCalls_addTool(cfg),
		tools_conversations.CreateConversations_createTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_createTool(cfg),
		tools_admin_teams.CreateAdmin_teams_createTool(cfg),
		tools_calls_participants.CreateCalls_participants_addTool(cfg),
		tools_admin_conversations.CreateAdmin_conversations_getconversationprefsTool(cfg),
	}
}
