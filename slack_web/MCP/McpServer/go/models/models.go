package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Objspaging represents the Objspaging schema from the OpenAPI specification
type Objspaging struct {
	Page int `json:"page"`
	Pages int `json:"pages,omitempty"`
	Per_page int `json:"per_page,omitempty"`
	Spill int `json:"spill,omitempty"`
	Total int `json:"total"`
	Count int `json:"count,omitempty"`
}

// Objsreaction represents the Objsreaction schema from the OpenAPI specification
type Objsreaction struct {
	Count int `json:"count"`
	Name string `json:"name"`
	Users []string `json:"users"`
}

// Objssubteam represents the Objssubteam schema from the OpenAPI specification
type Objssubteam struct {
	User_count int `json:"user_count,omitempty"`
	Handle string `json:"handle"`
	Id string `json:"id"`
	Channel_count int `json:"channel_count,omitempty"`
	Date_update int `json:"date_update"`
	Prefs map[string]interface{} `json:"prefs"`
	Is_subteam bool `json:"is_subteam"`
	Updated_by string `json:"updated_by"`
	Deleted_by map[string]interface{} `json:"deleted_by"`
	Description string `json:"description"`
	Created_by string `json:"created_by"`
	Team_id string `json:"team_id"`
	Auto_type map[string]interface{} `json:"auto_type"`
	Date_delete int `json:"date_delete"`
	Enterprise_subteam_id string `json:"enterprise_subteam_id"`
	Is_external bool `json:"is_external"`
	Is_usergroup bool `json:"is_usergroup"`
	Date_create int `json:"date_create"`
	Users []string `json:"users,omitempty"`
	Auto_provision bool `json:"auto_provision"`
	Name string `json:"name"`
}

// Objsreminder represents the Objsreminder schema from the OpenAPI specification
type Objsreminder struct {
	Recurring bool `json:"recurring"`
	Text string `json:"text"`
	Time int `json:"time,omitempty"`
	User string `json:"user"`
	Complete_ts int `json:"complete_ts,omitempty"`
	Creator string `json:"creator"`
	Id string `json:"id"`
}

// Objsuser represents the Objsuser schema from the OpenAPI specification
type Objsuser struct {
}

// Objsteamprofilefield represents the Objsteamprofilefield schema from the OpenAPI specification
type Objsteamprofilefield struct {
	Options map[string]interface{} `json:"options,omitempty"`
	Ordering float64 `json:"ordering"`
	TypeField string `json:"type"`
	Hint string `json:"hint"`
	Id string `json:"id"`
	Is_hidden bool `json:"is_hidden,omitempty"`
	Label string `json:"label"`
}

// Objscomment represents the Objscomment schema from the OpenAPI specification
type Objscomment struct {
	Num_stars int `json:"num_stars,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Timestamp int `json:"timestamp"`
	Comment string `json:"comment"`
	Is_starred bool `json:"is_starred,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	User string `json:"user"`
	Created int `json:"created"`
	Id string `json:"id"`
	Is_intro bool `json:"is_intro"`
}

// Objsenterpriseuser represents the Objsenterpriseuser schema from the OpenAPI specification
type Objsenterpriseuser struct {
	Teams []string `json:"teams"`
	Enterprise_id string `json:"enterprise_id"`
	Enterprise_name string `json:"enterprise_name"`
	Id string `json:"id"`
	Is_admin bool `json:"is_admin"`
	Is_owner bool `json:"is_owner"`
}

// Objsexternalorgmigrations represents the Objsexternalorgmigrations schema from the OpenAPI specification
type Objsexternalorgmigrations struct {
	Current []map[string]interface{} `json:"current"`
	Date_updated int `json:"date_updated"`
}

// Defspinnedinfo represents the Defspinnedinfo schema from the OpenAPI specification
type Defspinnedinfo struct {
}

// Objsteamprofilefieldoption represents the Objsteamprofilefieldoption schema from the OpenAPI specification
type Objsteamprofilefieldoption struct {
}

// Objsicon represents the Objsicon schema from the OpenAPI specification
type Objsicon struct {
	Image_102 string `json:"image_102,omitempty"`
	Image_132 string `json:"image_132,omitempty"`
	Image_230 string `json:"image_230,omitempty"`
	Image_34 string `json:"image_34,omitempty"`
	Image_44 string `json:"image_44,omitempty"`
	Image_68 string `json:"image_68,omitempty"`
	Image_88 string `json:"image_88,omitempty"`
	Image_default bool `json:"image_default,omitempty"`
}

// Objsbotprofile represents the Objsbotprofile schema from the OpenAPI specification
type Objsbotprofile struct {
	Icons map[string]interface{} `json:"icons"`
	Id string `json:"id"`
	Name string `json:"name"`
	Team_id string `json:"team_id"`
	Updated int `json:"updated"`
	App_id string `json:"app_id"`
	Deleted bool `json:"deleted"`
}

// Objsprimaryowner represents the Objsprimaryowner schema from the OpenAPI specification
type Objsprimaryowner struct {
	Id string `json:"id"`
	Email string `json:"email"`
}

// Objsuserprofileshort represents the Objsuserprofileshort schema from the OpenAPI specification
type Objsuserprofileshort struct {
	Real_name string `json:"real_name"`
	Is_restricted bool `json:"is_restricted"`
	Is_ultra_restricted bool `json:"is_ultra_restricted"`
	Display_name string `json:"display_name"`
	Display_name_normalized string `json:"display_name_normalized,omitempty"`
	Name string `json:"name"`
	Avatar_hash string `json:"avatar_hash"`
	Real_name_normalized string `json:"real_name_normalized,omitempty"`
	Team string `json:"team"`
	Image_72 string `json:"image_72"`
}

// Objsconversation represents the Objsconversation schema from the OpenAPI specification
type Objsconversation struct {
}

// Objsresources represents the Objsresources schema from the OpenAPI specification
type Objsresources struct {
	Wildcard bool `json:"wildcard,omitempty"`
	Excluded_ids []map[string]interface{} `json:"excluded_ids,omitempty"`
	Ids []map[string]interface{} `json:"ids"`
}

// Objsfile represents the Objsfile schema from the OpenAPI specification
type Objsfile struct {
	Ims []string `json:"ims,omitempty"`
	Thumb_64 string `json:"thumb_64,omitempty"`
	External_id string `json:"external_id,omitempty"`
	Thumb_1024 string `json:"thumb_1024,omitempty"`
	Thumb_960_h int `json:"thumb_960_h,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Is_tombstoned bool `json:"is_tombstoned,omitempty"`
	Thumb_480_h int `json:"thumb_480_h,omitempty"`
	Mode string `json:"mode,omitempty"`
	Thumb_960_w int `json:"thumb_960_w,omitempty"`
	Url_private string `json:"url_private,omitempty"`
	Username string `json:"username,omitempty"`
	Timestamp int `json:"timestamp,omitempty"`
	Editable bool `json:"editable,omitempty"`
	Thumb_80 string `json:"thumb_80,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Thumb_tiny string `json:"thumb_tiny,omitempty"`
	Thumb_360_h int `json:"thumb_360_h,omitempty"`
	Id string `json:"id,omitempty"`
	Thumb_960 string `json:"thumb_960,omitempty"`
	User string `json:"user,omitempty"`
	Thumb_720_h int `json:"thumb_720_h,omitempty"`
	Original_h int `json:"original_h,omitempty"`
	Public_url_shared bool `json:"public_url_shared,omitempty"`
	Comments_count int `json:"comments_count,omitempty"`
	Size int `json:"size,omitempty"`
	Has_rich_preview bool `json:"has_rich_preview,omitempty"`
	Image_exif_rotation int `json:"image_exif_rotation,omitempty"`
	Channels []string `json:"channels,omitempty"`
	Non_owner_editable bool `json:"non_owner_editable,omitempty"`
	Is_external bool `json:"is_external,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
	Thumb_480_w int `json:"thumb_480_w,omitempty"`
	Shares map[string]interface{} `json:"shares,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	Thumb_800 string `json:"thumb_800,omitempty"`
	External_url string `json:"external_url,omitempty"`
	Name string `json:"name,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Thumb_800_h int `json:"thumb_800_h,omitempty"`
	Thumb_1024_h int `json:"thumb_1024_h,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Created int `json:"created,omitempty"`
	Permalink_public string `json:"permalink_public,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Thumb_480 string `json:"thumb_480,omitempty"`
	Num_stars int `json:"num_stars,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Preview string `json:"preview,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Original_w int `json:"original_w,omitempty"`
	Thumb_160 string `json:"thumb_160,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	External_type string `json:"external_type,omitempty"`
	Pretty_type string `json:"pretty_type,omitempty"`
	Updated int `json:"updated,omitempty"`
	Thumb_800_w int `json:"thumb_800_w,omitempty"`
	Is_public bool `json:"is_public,omitempty"`
	Thumb_360 string `json:"thumb_360,omitempty"`
	Url_private_download string `json:"url_private_download,omitempty"`
	Thumb_360_w int `json:"thumb_360_w,omitempty"`
	Thumb_720_w int `json:"thumb_720_w,omitempty"`
	Last_editor string `json:"last_editor,omitempty"`
	Thumb_720 string `json:"thumb_720,omitempty"`
	Thumb_1024_w int `json:"thumb_1024_w,omitempty"`
	Date_delete int `json:"date_delete,omitempty"`
	Editor string `json:"editor,omitempty"`
}

// Objsresponsemetadata represents the Objsresponsemetadata schema from the OpenAPI specification
type Objsresponsemetadata struct {
}

// Objsmessage represents the Objsmessage schema from the OpenAPI specification
type Objsmessage struct {
	User_team string `json:"user_team,omitempty"`
	Parent_user_id string `json:"parent_user_id,omitempty"`
	Bot_profile Objsbotprofile `json:"bot_profile,omitempty"`
	Upload bool `json:"upload,omitempty"`
	TypeField string `json:"type"`
	Unread_count int `json:"unread_count,omitempty"`
	Icons map[string]interface{} `json:"icons,omitempty"`
	User_profile Objsuserprofileshort `json:"user_profile,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Files []Objsfile `json:"files,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Ts string `json:"ts"`
	Latest_reply string `json:"latest_reply,omitempty"`
	Subscribed bool `json:"subscribed,omitempty"`
	Comment Objscomment `json:"comment,omitempty"`
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
	Team string `json:"team,omitempty"`
	Username string `json:"username,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Name string `json:"name,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Thread_ts string `json:"thread_ts,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Blocks []map[string]interface{} `json:"blocks,omitempty"` // This is a very loose definition, in the future, we'll populate this with deeper schema in this definition namespace.
	Is_delayed_message bool `json:"is_delayed_message,omitempty"`
	Is_intro bool `json:"is_intro,omitempty"`
	Reply_users []string `json:"reply_users,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Subtype string `json:"subtype,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Old_name string `json:"old_name,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Inviter string `json:"inviter,omitempty"`
	Reply_count int `json:"reply_count,omitempty"`
	Bot_id map[string]interface{} `json:"bot_id,omitempty"`
	Text string `json:"text"`
	Topic string `json:"topic,omitempty"`
	File Objsfile `json:"file,omitempty"`
	Reply_users_count int `json:"reply_users_count,omitempty"`
	Client_msg_id string `json:"client_msg_id,omitempty"`
	User string `json:"user,omitempty"`
}

// Objsteam represents the Objsteam schema from the OpenAPI specification
type Objsteam struct {
	Id string `json:"id"`
	Discoverable map[string]interface{} `json:"discoverable,omitempty"`
	Domain string `json:"domain"`
	Email_domain string `json:"email_domain"`
	Has_compliance_export bool `json:"has_compliance_export,omitempty"`
	Msg_edit_window_mins int `json:"msg_edit_window_mins,omitempty"`
	Created int `json:"created,omitempty"`
	Is_over_storage_limit bool `json:"is_over_storage_limit,omitempty"`
	Enterprise_id string `json:"enterprise_id,omitempty"`
	Name string `json:"name"`
	Date_create int `json:"date_create,omitempty"`
	Over_integrations_limit bool `json:"over_integrations_limit,omitempty"`
	Locale string `json:"locale,omitempty"`
	Pay_prod_cur string `json:"pay_prod_cur,omitempty"`
	Primary_owner Objsprimaryowner `json:"primary_owner,omitempty"`
	Avatar_base_url string `json:"avatar_base_url,omitempty"`
	Icon Objsicon `json:"icon"`
	Is_enterprise int `json:"is_enterprise,omitempty"`
	Over_storage_limit bool `json:"over_storage_limit,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Is_assigned bool `json:"is_assigned,omitempty"`
	Limit_ts int `json:"limit_ts,omitempty"`
	Plan string `json:"plan,omitempty"`
	Enterprise_name string `json:"enterprise_name,omitempty"`
	Sso_provider map[string]interface{} `json:"sso_provider,omitempty"`
	Messages_count int `json:"messages_count,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	External_org_migrations Objsexternalorgmigrations `json:"external_org_migrations,omitempty"`
}

// Objschannel represents the Objschannel schema from the OpenAPI specification
type Objschannel struct {
	Id string `json:"id"`
	Last_read string `json:"last_read,omitempty"`
	Unread_count_display int `json:"unread_count_display,omitempty"`
	Is_private bool `json:"is_private"`
	Is_thread_only bool `json:"is_thread_only,omitempty"`
	Is_shared bool `json:"is_shared"`
	Is_read_only bool `json:"is_read_only,omitempty"`
	Is_pending_ext_shared bool `json:"is_pending_ext_shared,omitempty"`
	Creator string `json:"creator"`
	Num_members int `json:"num_members,omitempty"`
	Priority float64 `json:"priority,omitempty"`
	Latest map[string]interface{} `json:"latest,omitempty"`
	Name_normalized string `json:"name_normalized"`
	Previous_names []string `json:"previous_names,omitempty"`
	Accepted_user string `json:"accepted_user,omitempty"`
	Pending_shared []string `json:"pending_shared,omitempty"`
	Topic map[string]interface{} `json:"topic"`
	Is_channel bool `json:"is_channel"`
	Is_non_threadable bool `json:"is_non_threadable,omitempty"`
	Unlinked int `json:"unlinked,omitempty"`
	Is_general bool `json:"is_general,omitempty"`
	Purpose map[string]interface{} `json:"purpose"`
	Is_archived bool `json:"is_archived,omitempty"`
	Is_frozen bool `json:"is_frozen,omitempty"`
	Members []string `json:"members"`
	Is_mpim bool `json:"is_mpim"`
	Created int `json:"created"`
	Name string `json:"name"`
	Is_moved int `json:"is_moved,omitempty"`
	Is_org_shared bool `json:"is_org_shared"`
	Unread_count int `json:"unread_count,omitempty"`
	Is_member bool `json:"is_member,omitempty"`
}

// Objsuserprofile represents the Objsuserprofile schema from the OpenAPI specification
type Objsuserprofile struct {
	Title string `json:"title"`
	Display_name_normalized string `json:"display_name_normalized"`
	Is_app_user bool `json:"is_app_user,omitempty"`
	Api_app_id string `json:"api_app_id,omitempty"`
	Updated int `json:"updated,omitempty"`
	Team string `json:"team,omitempty"`
	Real_name_normalized string `json:"real_name_normalized"`
	Skype string `json:"skype"`
	Status_expiration int `json:"status_expiration,omitempty"`
	Avatar_hash string `json:"avatar_hash"`
	Is_custom_image bool `json:"is_custom_image,omitempty"`
	Status_text string `json:"status_text"`
	Display_name string `json:"display_name"`
	Status_emoji string `json:"status_emoji"`
	Phone string `json:"phone"`
	Real_name string `json:"real_name"`
	Pronouns string `json:"pronouns,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Memberships_count int `json:"memberships_count,omitempty"`
	Bot_id string `json:"bot_id,omitempty"`
	Status_default_text string `json:"status_default_text,omitempty"`
	Status_default_emoji string `json:"status_default_emoji,omitempty"`
	Always_active bool `json:"always_active,omitempty"`
	Last_avatar_image_hash string `json:"last_avatar_image_hash,omitempty"`
}
