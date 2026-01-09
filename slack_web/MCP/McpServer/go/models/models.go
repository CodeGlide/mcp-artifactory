package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Objsfile represents the Objsfile schema from the OpenAPI specification
type Objsfile struct {
	Thumb_800_w int `json:"thumb_800_w,omitempty"`
	Name string `json:"name,omitempty"`
	Thumb_1024 string `json:"thumb_1024,omitempty"`
	Last_editor string `json:"last_editor,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Permalink_public string `json:"permalink_public,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Editor string `json:"editor,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Ims []string `json:"ims,omitempty"`
	Thumb_tiny string `json:"thumb_tiny,omitempty"`
	Id string `json:"id,omitempty"`
	Thumb_480_w int `json:"thumb_480_w,omitempty"`
	Thumb_720_w int `json:"thumb_720_w,omitempty"`
	Timestamp int `json:"timestamp,omitempty"`
	Thumb_80 string `json:"thumb_80,omitempty"`
	Size int `json:"size,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Thumb_1024_w int `json:"thumb_1024_w,omitempty"`
	Thumb_1024_h int `json:"thumb_1024_h,omitempty"`
	State string `json:"state,omitempty"`
	Username string `json:"username,omitempty"`
	Created int `json:"created,omitempty"`
	Is_external bool `json:"is_external,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Mode string `json:"mode,omitempty"`
	Thumb_960 string `json:"thumb_960,omitempty"`
	Date_delete int `json:"date_delete,omitempty"`
	Thumb_480 string `json:"thumb_480,omitempty"`
	Thumb_360_w int `json:"thumb_360_w,omitempty"`
	Thumb_64 string `json:"thumb_64,omitempty"`
	Editable bool `json:"editable,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	Thumb_960_w int `json:"thumb_960_w,omitempty"`
	External_id string `json:"external_id,omitempty"`
	Preview string `json:"preview,omitempty"`
	Thumb_480_h int `json:"thumb_480_h,omitempty"`
	Channels []string `json:"channels,omitempty"`
	Updated int `json:"updated,omitempty"`
	Url_private_download string `json:"url_private_download,omitempty"`
	Thumb_960_h int `json:"thumb_960_h,omitempty"`
	Is_public bool `json:"is_public,omitempty"`
	Thumb_160 string `json:"thumb_160,omitempty"`
	Num_stars int `json:"num_stars,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Thumb_720 string `json:"thumb_720,omitempty"`
	Comments_count int `json:"comments_count,omitempty"`
	External_type string `json:"external_type,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Is_tombstoned bool `json:"is_tombstoned,omitempty"`
	Has_rich_preview bool `json:"has_rich_preview,omitempty"`
	Non_owner_editable bool `json:"non_owner_editable,omitempty"`
	Thumb_720_h int `json:"thumb_720_h,omitempty"`
	External_url string `json:"external_url,omitempty"`
	Public_url_shared bool `json:"public_url_shared,omitempty"`
	Original_h int `json:"original_h,omitempty"`
	Thumb_800 string `json:"thumb_800,omitempty"`
	Thumb_360 string `json:"thumb_360,omitempty"`
	User string `json:"user,omitempty"`
	Shares map[string]interface{} `json:"shares,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Url_private string `json:"url_private,omitempty"`
	Image_exif_rotation int `json:"image_exif_rotation,omitempty"`
	Thumb_360_h int `json:"thumb_360_h,omitempty"`
	Thumb_800_h int `json:"thumb_800_h,omitempty"`
	Original_w int `json:"original_w,omitempty"`
	Pretty_type string `json:"pretty_type,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Title string `json:"title,omitempty"`
}

// Objschannel represents the Objschannel schema from the OpenAPI specification
type Objschannel struct {
	Is_read_only bool `json:"is_read_only,omitempty"`
	Is_thread_only bool `json:"is_thread_only,omitempty"`
	Purpose map[string]interface{} `json:"purpose"`
	Members []string `json:"members"`
	Is_general bool `json:"is_general,omitempty"`
	Is_private bool `json:"is_private"`
	Is_shared bool `json:"is_shared"`
	Topic map[string]interface{} `json:"topic"`
	Last_read string `json:"last_read,omitempty"`
	Is_archived bool `json:"is_archived,omitempty"`
	Unread_count_display int `json:"unread_count_display,omitempty"`
	Is_org_shared bool `json:"is_org_shared"`
	Accepted_user string `json:"accepted_user,omitempty"`
	Is_channel bool `json:"is_channel"`
	Is_moved int `json:"is_moved,omitempty"`
	Is_non_threadable bool `json:"is_non_threadable,omitempty"`
	Name_normalized string `json:"name_normalized"`
	Num_members int `json:"num_members,omitempty"`
	Name string `json:"name"`
	Is_mpim bool `json:"is_mpim"`
	Is_pending_ext_shared bool `json:"is_pending_ext_shared,omitempty"`
	Pending_shared []string `json:"pending_shared,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	Previous_names []string `json:"previous_names,omitempty"`
	Created int `json:"created"`
	Creator string `json:"creator"`
	Latest map[string]interface{} `json:"latest,omitempty"`
	Is_member bool `json:"is_member,omitempty"`
	Priority float64 `json:"priority,omitempty"`
	Unlinked int `json:"unlinked,omitempty"`
	Id string `json:"id"`
	Is_frozen bool `json:"is_frozen,omitempty"`
}

// Objsenterpriseuser represents the Objsenterpriseuser schema from the OpenAPI specification
type Objsenterpriseuser struct {
	Id string `json:"id"`
	Is_admin bool `json:"is_admin"`
	Is_owner bool `json:"is_owner"`
	Teams []string `json:"teams"`
	Enterprise_id string `json:"enterprise_id"`
	Enterprise_name string `json:"enterprise_name"`
}

// Objsicon represents the Objsicon schema from the OpenAPI specification
type Objsicon struct {
	Image_88 string `json:"image_88,omitempty"`
	Image_default bool `json:"image_default,omitempty"`
	Image_102 string `json:"image_102,omitempty"`
	Image_132 string `json:"image_132,omitempty"`
	Image_230 string `json:"image_230,omitempty"`
	Image_34 string `json:"image_34,omitempty"`
	Image_44 string `json:"image_44,omitempty"`
	Image_68 string `json:"image_68,omitempty"`
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

// Objsuser represents the Objsuser schema from the OpenAPI specification
type Objsuser struct {
}

// Objsresources represents the Objsresources schema from the OpenAPI specification
type Objsresources struct {
	Excluded_ids []map[string]interface{} `json:"excluded_ids,omitempty"`
	Ids []map[string]interface{} `json:"ids"`
	Wildcard bool `json:"wildcard,omitempty"`
}

// Objspaging represents the Objspaging schema from the OpenAPI specification
type Objspaging struct {
	Count int `json:"count,omitempty"`
	Page int `json:"page"`
	Pages int `json:"pages,omitempty"`
	Per_page int `json:"per_page,omitempty"`
	Spill int `json:"spill,omitempty"`
	Total int `json:"total"`
}

// Objsteamprofilefieldoption represents the Objsteamprofilefieldoption schema from the OpenAPI specification
type Objsteamprofilefieldoption struct {
}

// Objsexternalorgmigrations represents the Objsexternalorgmigrations schema from the OpenAPI specification
type Objsexternalorgmigrations struct {
	Current []map[string]interface{} `json:"current"`
	Date_updated int `json:"date_updated"`
}

// Objssubteam represents the Objssubteam schema from the OpenAPI specification
type Objssubteam struct {
	Channel_count int `json:"channel_count,omitempty"`
	Created_by string `json:"created_by"`
	Auto_provision bool `json:"auto_provision"`
	Auto_type map[string]interface{} `json:"auto_type"`
	Date_delete int `json:"date_delete"`
	Team_id string `json:"team_id"`
	Is_usergroup bool `json:"is_usergroup"`
	Is_external bool `json:"is_external"`
	Name string `json:"name"`
	Is_subteam bool `json:"is_subteam"`
	Updated_by string `json:"updated_by"`
	Users []string `json:"users,omitempty"`
	Handle string `json:"handle"`
	User_count int `json:"user_count,omitempty"`
	Id string `json:"id"`
	Description string `json:"description"`
	Enterprise_subteam_id string `json:"enterprise_subteam_id"`
	Date_update int `json:"date_update"`
	Deleted_by map[string]interface{} `json:"deleted_by"`
	Date_create int `json:"date_create"`
	Prefs map[string]interface{} `json:"prefs"`
}

// Objsuserprofileshort represents the Objsuserprofileshort schema from the OpenAPI specification
type Objsuserprofileshort struct {
	Real_name string `json:"real_name"`
	Is_restricted bool `json:"is_restricted"`
	Avatar_hash string `json:"avatar_hash"`
	Display_name string `json:"display_name"`
	Display_name_normalized string `json:"display_name_normalized,omitempty"`
	Real_name_normalized string `json:"real_name_normalized,omitempty"`
	Team string `json:"team"`
	Image_72 string `json:"image_72"`
	Is_ultra_restricted bool `json:"is_ultra_restricted"`
	Name string `json:"name"`
}

// Defspinnedinfo represents the Defspinnedinfo schema from the OpenAPI specification
type Defspinnedinfo struct {
}

// Objsmessage represents the Objsmessage schema from the OpenAPI specification
type Objsmessage struct {
	Upload bool `json:"upload,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Icons map[string]interface{} `json:"icons,omitempty"`
	File Objsfile `json:"file,omitempty"`
	Reply_count int `json:"reply_count,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Client_msg_id string `json:"client_msg_id,omitempty"`
	Text string `json:"text"`
	Subtype string `json:"subtype,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Comment Objscomment `json:"comment,omitempty"`
	Is_delayed_message bool `json:"is_delayed_message,omitempty"`
	Topic string `json:"topic,omitempty"`
	Bot_profile Objsbotprofile `json:"bot_profile,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Reply_users []string `json:"reply_users,omitempty"`
	Parent_user_id string `json:"parent_user_id,omitempty"`
	Subscribed bool `json:"subscribed,omitempty"`
	Latest_reply string `json:"latest_reply,omitempty"`
	Name string `json:"name,omitempty"`
	Ts string `json:"ts"`
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
	Files []Objsfile `json:"files,omitempty"`
	TypeField string `json:"type"`
	Is_starred bool `json:"is_starred,omitempty"`
	Team string `json:"team,omitempty"`
	Old_name string `json:"old_name,omitempty"`
	Blocks []map[string]interface{} `json:"blocks,omitempty"` // This is a very loose definition, in the future, we'll populate this with deeper schema in this definition namespace.
	User_profile Objsuserprofileshort `json:"user_profile,omitempty"`
	Reply_users_count int `json:"reply_users_count,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	User string `json:"user,omitempty"`
	Inviter string `json:"inviter,omitempty"`
	Bot_id map[string]interface{} `json:"bot_id,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Username string `json:"username,omitempty"`
	Thread_ts string `json:"thread_ts,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Is_intro bool `json:"is_intro,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
}

// Objsreminder represents the Objsreminder schema from the OpenAPI specification
type Objsreminder struct {
	Text string `json:"text"`
	Time int `json:"time,omitempty"`
	User string `json:"user"`
	Complete_ts int `json:"complete_ts,omitempty"`
	Creator string `json:"creator"`
	Id string `json:"id"`
	Recurring bool `json:"recurring"`
}

// Objsconversation represents the Objsconversation schema from the OpenAPI specification
type Objsconversation struct {
}

// Objsreaction represents the Objsreaction schema from the OpenAPI specification
type Objsreaction struct {
	Users []string `json:"users"`
	Count int `json:"count"`
	Name string `json:"name"`
}

// Objsteam represents the Objsteam schema from the OpenAPI specification
type Objsteam struct {
	Limit_ts int `json:"limit_ts,omitempty"`
	Created int `json:"created,omitempty"`
	Enterprise_id string `json:"enterprise_id,omitempty"`
	Enterprise_name string `json:"enterprise_name,omitempty"`
	External_org_migrations Objsexternalorgmigrations `json:"external_org_migrations,omitempty"`
	Is_assigned bool `json:"is_assigned,omitempty"`
	Over_storage_limit bool `json:"over_storage_limit,omitempty"`
	Email_domain string `json:"email_domain"`
	Primary_owner Objsprimaryowner `json:"primary_owner,omitempty"`
	Avatar_base_url string `json:"avatar_base_url,omitempty"`
	Msg_edit_window_mins int `json:"msg_edit_window_mins,omitempty"`
	Discoverable map[string]interface{} `json:"discoverable,omitempty"`
	Is_over_storage_limit bool `json:"is_over_storage_limit,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Sso_provider map[string]interface{} `json:"sso_provider,omitempty"`
	Locale string `json:"locale,omitempty"`
	Pay_prod_cur string `json:"pay_prod_cur,omitempty"`
	Date_create int `json:"date_create,omitempty"`
	Domain string `json:"domain"`
	Messages_count int `json:"messages_count,omitempty"`
	Name string `json:"name"`
	Has_compliance_export bool `json:"has_compliance_export,omitempty"`
	Over_integrations_limit bool `json:"over_integrations_limit,omitempty"`
	Plan string `json:"plan,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Icon Objsicon `json:"icon"`
	Id string `json:"id"`
	Is_enterprise int `json:"is_enterprise,omitempty"`
}

// Objscomment represents the Objscomment schema from the OpenAPI specification
type Objscomment struct {
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	User string `json:"user"`
	Comment string `json:"comment"`
	Id string `json:"id"`
	Num_stars int `json:"num_stars,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Created int `json:"created"`
	Is_intro bool `json:"is_intro"`
	Timestamp int `json:"timestamp"`
	Is_starred bool `json:"is_starred,omitempty"`
}

// Objsbotprofile represents the Objsbotprofile schema from the OpenAPI specification
type Objsbotprofile struct {
	Name string `json:"name"`
	Team_id string `json:"team_id"`
	Updated int `json:"updated"`
	App_id string `json:"app_id"`
	Deleted bool `json:"deleted"`
	Icons map[string]interface{} `json:"icons"`
	Id string `json:"id"`
}

// Objsuserprofile represents the Objsuserprofile schema from the OpenAPI specification
type Objsuserprofile struct {
	Updated int `json:"updated,omitempty"`
	Phone string `json:"phone"`
	Last_avatar_image_hash string `json:"last_avatar_image_hash,omitempty"`
	Status_text string `json:"status_text"`
	Team string `json:"team,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Is_custom_image bool `json:"is_custom_image,omitempty"`
	Memberships_count int `json:"memberships_count,omitempty"`
	Real_name string `json:"real_name"`
	Bot_id string `json:"bot_id,omitempty"`
	Always_active bool `json:"always_active,omitempty"`
	Display_name string `json:"display_name"`
	Skype string `json:"skype"`
	Api_app_id string `json:"api_app_id,omitempty"`
	Avatar_hash string `json:"avatar_hash"`
	Status_emoji string `json:"status_emoji"`
	Title string `json:"title"`
	Real_name_normalized string `json:"real_name_normalized"`
	Status_expiration int `json:"status_expiration,omitempty"`
	Pronouns string `json:"pronouns,omitempty"`
	Display_name_normalized string `json:"display_name_normalized"`
	Status_default_emoji string `json:"status_default_emoji,omitempty"`
	Is_app_user bool `json:"is_app_user,omitempty"`
	Status_default_text string `json:"status_default_text,omitempty"`
}

// Objsprimaryowner represents the Objsprimaryowner schema from the OpenAPI specification
type Objsprimaryowner struct {
	Email string `json:"email"`
	Id string `json:"id"`
}

// Objsresponsemetadata represents the Objsresponsemetadata schema from the OpenAPI specification
type Objsresponsemetadata struct {
}
