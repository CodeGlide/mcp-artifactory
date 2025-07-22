package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Objsmessage represents the Objsmessage schema from the OpenAPI specification
type Objsmessage struct {
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Is_intro bool `json:"is_intro,omitempty"`
	Ts string `json:"ts"`
	Parent_user_id string `json:"parent_user_id,omitempty"`
	Files []Objsfile `json:"files,omitempty"`
	Is_delayed_message bool `json:"is_delayed_message,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Upload bool `json:"upload,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	User string `json:"user,omitempty"`
	Bot_profile Objsbotprofile `json:"bot_profile,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Topic string `json:"topic,omitempty"`
	Blocks []map[string]interface{} `json:"blocks,omitempty"` // This is a very loose definition, in the future, we'll populate this with deeper schema in this definition namespace.
	TypeField string `json:"type"`
	File Objsfile `json:"file,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Reply_users_count int `json:"reply_users_count,omitempty"`
	Icons map[string]interface{} `json:"icons,omitempty"`
	Name string `json:"name,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Username string `json:"username,omitempty"`
	Reply_users []string `json:"reply_users,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Reply_count int `json:"reply_count,omitempty"`
	Team string `json:"team,omitempty"`
	Text string `json:"text"`
	Subscribed bool `json:"subscribed,omitempty"`
	Client_msg_id string `json:"client_msg_id,omitempty"`
	Old_name string `json:"old_name,omitempty"`
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
	Comment Objscomment `json:"comment,omitempty"`
	Thread_ts string `json:"thread_ts,omitempty"`
	Bot_id map[string]interface{} `json:"bot_id,omitempty"`
	Inviter string `json:"inviter,omitempty"`
	Latest_reply string `json:"latest_reply,omitempty"`
	User_profile Objsuserprofileshort `json:"user_profile,omitempty"`
	Subtype string `json:"subtype,omitempty"`
}

// Objschannel represents the Objschannel schema from the OpenAPI specification
type Objschannel struct {
	Creator string `json:"creator"`
	Is_mpim bool `json:"is_mpim"`
	Topic map[string]interface{} `json:"topic"`
	Unlinked int `json:"unlinked,omitempty"`
	Is_read_only bool `json:"is_read_only,omitempty"`
	Name string `json:"name"`
	Is_frozen bool `json:"is_frozen,omitempty"`
	Is_thread_only bool `json:"is_thread_only,omitempty"`
	Id string `json:"id"`
	Is_channel bool `json:"is_channel"`
	Pending_shared []string `json:"pending_shared,omitempty"`
	Name_normalized string `json:"name_normalized"`
	Unread_count int `json:"unread_count,omitempty"`
	Unread_count_display int `json:"unread_count_display,omitempty"`
	Is_moved int `json:"is_moved,omitempty"`
	Latest map[string]interface{} `json:"latest,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Previous_names []string `json:"previous_names,omitempty"`
	Purpose map[string]interface{} `json:"purpose"`
	Is_member bool `json:"is_member,omitempty"`
	Num_members int `json:"num_members,omitempty"`
	Is_org_shared bool `json:"is_org_shared"`
	Is_pending_ext_shared bool `json:"is_pending_ext_shared,omitempty"`
	Created int `json:"created"`
	Priority float64 `json:"priority,omitempty"`
	Is_archived bool `json:"is_archived,omitempty"`
	Members []string `json:"members"`
	Is_private bool `json:"is_private"`
	Is_non_threadable bool `json:"is_non_threadable,omitempty"`
	Accepted_user string `json:"accepted_user,omitempty"`
	Is_general bool `json:"is_general,omitempty"`
	Is_shared bool `json:"is_shared"`
}

// Objsconversation represents the Objsconversation schema from the OpenAPI specification
type Objsconversation struct {
}

// Objsuserprofile represents the Objsuserprofile schema from the OpenAPI specification
type Objsuserprofile struct {
	Updated int `json:"updated,omitempty"`
	Skype string `json:"skype"`
	Is_app_user bool `json:"is_app_user,omitempty"`
	Is_custom_image bool `json:"is_custom_image,omitempty"`
	Real_name_normalized string `json:"real_name_normalized"`
	Always_active bool `json:"always_active,omitempty"`
	Bot_id string `json:"bot_id,omitempty"`
	Team string `json:"team,omitempty"`
	Last_avatar_image_hash string `json:"last_avatar_image_hash,omitempty"`
	Phone string `json:"phone"`
	Real_name string `json:"real_name"`
	Status_expiration int `json:"status_expiration,omitempty"`
	Status_emoji string `json:"status_emoji"`
	Title string `json:"title"`
	Display_name_normalized string `json:"display_name_normalized"`
	Memberships_count int `json:"memberships_count,omitempty"`
	Pronouns string `json:"pronouns,omitempty"`
	Display_name string `json:"display_name"`
	Avatar_hash string `json:"avatar_hash"`
	Status_default_emoji string `json:"status_default_emoji,omitempty"`
	Api_app_id string `json:"api_app_id,omitempty"`
	Status_text string `json:"status_text"`
	Status_default_text string `json:"status_default_text,omitempty"`
	User_id string `json:"user_id,omitempty"`
}

// Objsresources represents the Objsresources schema from the OpenAPI specification
type Objsresources struct {
	Excluded_ids []map[string]interface{} `json:"excluded_ids,omitempty"`
	Ids []map[string]interface{} `json:"ids"`
	Wildcard bool `json:"wildcard,omitempty"`
}

// Objsresponsemetadata represents the Objsresponsemetadata schema from the OpenAPI specification
type Objsresponsemetadata struct {
}

// Objsbotprofile represents the Objsbotprofile schema from the OpenAPI specification
type Objsbotprofile struct {
	Team_id string `json:"team_id"`
	Updated int `json:"updated"`
	App_id string `json:"app_id"`
	Deleted bool `json:"deleted"`
	Icons map[string]interface{} `json:"icons"`
	Id string `json:"id"`
	Name string `json:"name"`
}

// Objsteamprofilefieldoption represents the Objsteamprofilefieldoption schema from the OpenAPI specification
type Objsteamprofilefieldoption struct {
}

// Objsuser represents the Objsuser schema from the OpenAPI specification
type Objsuser struct {
}

// Objsprimaryowner represents the Objsprimaryowner schema from the OpenAPI specification
type Objsprimaryowner struct {
	Email string `json:"email"`
	Id string `json:"id"`
}

// Objsuserprofileshort represents the Objsuserprofileshort schema from the OpenAPI specification
type Objsuserprofileshort struct {
	Name string `json:"name"`
	Real_name string `json:"real_name"`
	Real_name_normalized string `json:"real_name_normalized,omitempty"`
	Avatar_hash string `json:"avatar_hash"`
	Display_name string `json:"display_name"`
	Display_name_normalized string `json:"display_name_normalized,omitempty"`
	Is_restricted bool `json:"is_restricted"`
	Image_72 string `json:"image_72"`
	Is_ultra_restricted bool `json:"is_ultra_restricted"`
	Team string `json:"team"`
}

// Objscomment represents the Objscomment schema from the OpenAPI specification
type Objscomment struct {
	Created int `json:"created"`
	Timestamp int `json:"timestamp"`
	Is_intro bool `json:"is_intro"`
	Id string `json:"id"`
	Is_starred bool `json:"is_starred,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	User string `json:"user"`
	Num_stars int `json:"num_stars,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Comment string `json:"comment"`
}

// Objsicon represents the Objsicon schema from the OpenAPI specification
type Objsicon struct {
	Image_68 string `json:"image_68,omitempty"`
	Image_88 string `json:"image_88,omitempty"`
	Image_default bool `json:"image_default,omitempty"`
	Image_102 string `json:"image_102,omitempty"`
	Image_132 string `json:"image_132,omitempty"`
	Image_230 string `json:"image_230,omitempty"`
	Image_34 string `json:"image_34,omitempty"`
	Image_44 string `json:"image_44,omitempty"`
}

// Objssubteam represents the Objssubteam schema from the OpenAPI specification
type Objssubteam struct {
	Enterprise_subteam_id string `json:"enterprise_subteam_id"`
	Auto_type map[string]interface{} `json:"auto_type"`
	Handle string `json:"handle"`
	Id string `json:"id"`
	Name string `json:"name"`
	Channel_count int `json:"channel_count,omitempty"`
	Is_external bool `json:"is_external"`
	Auto_provision bool `json:"auto_provision"`
	Users []string `json:"users,omitempty"`
	Prefs map[string]interface{} `json:"prefs"`
	Updated_by string `json:"updated_by"`
	Team_id string `json:"team_id"`
	Is_usergroup bool `json:"is_usergroup"`
	Date_create int `json:"date_create"`
	Date_update int `json:"date_update"`
	Date_delete int `json:"date_delete"`
	Is_subteam bool `json:"is_subteam"`
	User_count int `json:"user_count,omitempty"`
	Deleted_by map[string]interface{} `json:"deleted_by"`
	Description string `json:"description"`
	Created_by string `json:"created_by"`
}

// Objsteam represents the Objsteam schema from the OpenAPI specification
type Objsteam struct {
	Name string `json:"name"`
	Locale string `json:"locale,omitempty"`
	Created int `json:"created,omitempty"`
	Over_storage_limit bool `json:"over_storage_limit,omitempty"`
	Plan string `json:"plan,omitempty"`
	Msg_edit_window_mins int `json:"msg_edit_window_mins,omitempty"`
	Domain string `json:"domain"`
	Over_integrations_limit bool `json:"over_integrations_limit,omitempty"`
	Date_create int `json:"date_create,omitempty"`
	External_org_migrations Objsexternalorgmigrations `json:"external_org_migrations,omitempty"`
	Id string `json:"id"`
	Is_over_storage_limit bool `json:"is_over_storage_limit,omitempty"`
	Messages_count int `json:"messages_count,omitempty"`
	Avatar_base_url string `json:"avatar_base_url,omitempty"`
	Enterprise_name string `json:"enterprise_name,omitempty"`
	Icon Objsicon `json:"icon"`
	Pay_prod_cur string `json:"pay_prod_cur,omitempty"`
	Discoverable map[string]interface{} `json:"discoverable,omitempty"`
	Has_compliance_export bool `json:"has_compliance_export,omitempty"`
	Limit_ts int `json:"limit_ts,omitempty"`
	Enterprise_id string `json:"enterprise_id,omitempty"`
	Is_assigned bool `json:"is_assigned,omitempty"`
	Primary_owner Objsprimaryowner `json:"primary_owner,omitempty"`
	Email_domain string `json:"email_domain"`
	Sso_provider map[string]interface{} `json:"sso_provider,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Is_enterprise int `json:"is_enterprise,omitempty"`
}

// Objsfile represents the Objsfile schema from the OpenAPI specification
type Objsfile struct {
	Size int `json:"size,omitempty"`
	Comments_count int `json:"comments_count,omitempty"`
	Original_h int `json:"original_h,omitempty"`
	Username string `json:"username,omitempty"`
	Is_external bool `json:"is_external,omitempty"`
	Timestamp int `json:"timestamp,omitempty"`
	Editable bool `json:"editable,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Editor string `json:"editor,omitempty"`
	Thumb_tiny string `json:"thumb_tiny,omitempty"`
	Thumb_360_h int `json:"thumb_360_h,omitempty"`
	Thumb_720_h int `json:"thumb_720_h,omitempty"`
	Date_delete int `json:"date_delete,omitempty"`
	Thumb_800 string `json:"thumb_800,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Pretty_type string `json:"pretty_type,omitempty"`
	Url_private string `json:"url_private,omitempty"`
	Created int `json:"created,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Name string `json:"name,omitempty"`
	Thumb_960_w int `json:"thumb_960_w,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Ims []string `json:"ims,omitempty"`
	Preview string `json:"preview,omitempty"`
	Updated int `json:"updated,omitempty"`
	Thumb_960 string `json:"thumb_960,omitempty"`
	Thumb_480 string `json:"thumb_480,omitempty"`
	Thumb_720 string `json:"thumb_720,omitempty"`
	Shares map[string]interface{} `json:"shares,omitempty"`
	User string `json:"user,omitempty"`
	State string `json:"state,omitempty"`
	Has_rich_preview bool `json:"has_rich_preview,omitempty"`
	Thumb_1024 string `json:"thumb_1024,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Is_public bool `json:"is_public,omitempty"`
	Thumb_360 string `json:"thumb_360,omitempty"`
	Title string `json:"title,omitempty"`
	Num_stars int `json:"num_stars,omitempty"`
	Is_tombstoned bool `json:"is_tombstoned,omitempty"`
	Public_url_shared bool `json:"public_url_shared,omitempty"`
	Id string `json:"id,omitempty"`
	Thumb_480_w int `json:"thumb_480_w,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Thumb_800_w int `json:"thumb_800_w,omitempty"`
	Thumb_160 string `json:"thumb_160,omitempty"`
	Permalink_public string `json:"permalink_public,omitempty"`
	Thumb_800_h int `json:"thumb_800_h,omitempty"`
	Channels []string `json:"channels,omitempty"`
	External_url string `json:"external_url,omitempty"`
	Thumb_1024_h int `json:"thumb_1024_h,omitempty"`
	Url_private_download string `json:"url_private_download,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Last_editor string `json:"last_editor,omitempty"`
	Groups []string `json:"groups,omitempty"`
	External_id string `json:"external_id,omitempty"`
	Non_owner_editable bool `json:"non_owner_editable,omitempty"`
	Thumb_480_h int `json:"thumb_480_h,omitempty"`
	Original_w int `json:"original_w,omitempty"`
	Thumb_64 string `json:"thumb_64,omitempty"`
	Thumb_1024_w int `json:"thumb_1024_w,omitempty"`
	Image_exif_rotation int `json:"image_exif_rotation,omitempty"`
	Thumb_960_h int `json:"thumb_960_h,omitempty"`
	External_type string `json:"external_type,omitempty"`
	Thumb_360_w int `json:"thumb_360_w,omitempty"`
	Thumb_720_w int `json:"thumb_720_w,omitempty"`
	Mode string `json:"mode,omitempty"`
	Thumb_80 string `json:"thumb_80,omitempty"`
}

// Objsteamprofilefield represents the Objsteamprofilefield schema from the OpenAPI specification
type Objsteamprofilefield struct {
	TypeField string `json:"type"`
	Hint string `json:"hint"`
	Id string `json:"id"`
	Is_hidden bool `json:"is_hidden,omitempty"`
	Label string `json:"label"`
	Options map[string]interface{} `json:"options,omitempty"`
	Ordering float64 `json:"ordering"`
}

// Objsreminder represents the Objsreminder schema from the OpenAPI specification
type Objsreminder struct {
	Time int `json:"time,omitempty"`
	User string `json:"user"`
	Complete_ts int `json:"complete_ts,omitempty"`
	Creator string `json:"creator"`
	Id string `json:"id"`
	Recurring bool `json:"recurring"`
	Text string `json:"text"`
}

// Defspinnedinfo represents the Defspinnedinfo schema from the OpenAPI specification
type Defspinnedinfo struct {
}

// Objsexternalorgmigrations represents the Objsexternalorgmigrations schema from the OpenAPI specification
type Objsexternalorgmigrations struct {
	Date_updated int `json:"date_updated"`
	Current []map[string]interface{} `json:"current"`
}

// Objsenterpriseuser represents the Objsenterpriseuser schema from the OpenAPI specification
type Objsenterpriseuser struct {
	Is_owner bool `json:"is_owner"`
	Teams []string `json:"teams"`
	Enterprise_id string `json:"enterprise_id"`
	Enterprise_name string `json:"enterprise_name"`
	Id string `json:"id"`
	Is_admin bool `json:"is_admin"`
}

// Objspaging represents the Objspaging schema from the OpenAPI specification
type Objspaging struct {
	Per_page int `json:"per_page,omitempty"`
	Spill int `json:"spill,omitempty"`
	Total int `json:"total"`
	Count int `json:"count,omitempty"`
	Page int `json:"page"`
	Pages int `json:"pages,omitempty"`
}

// Objsreaction represents the Objsreaction schema from the OpenAPI specification
type Objsreaction struct {
	Users []string `json:"users"`
	Count int `json:"count"`
	Name string `json:"name"`
}
