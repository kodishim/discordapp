package models

type GuildPreview struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	Icon                     string   `json:"icon"`
	Splash                   any      `json:"splash"`
	DiscoverySplash          any      `json:"discovery_splash"`
	Emojis                   []any    `json:"emojis"`
	Features                 []string `json:"features"`
	ApproximateMemberCount   int      `json:"approximate_member_count"`
	ApproximatePresenceCount int      `json:"approximate_presence_count"`
	Description              string   `json:"description"`
	Stickers                 []any    `json:"stickers"`
}

type Guild struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	Icon                     string   `json:"icon"`
	Description              *string  `json:"description"`
	Splash                   string   `json:"splash"`
	DiscoverySplash          *string  `json:"discovery_splash"`
	ApproximateMemberCount   int      `json:"approximate_member_count"`
	ApproximatePresenceCount int      `json:"approximate_presence_count"`
	Features                 []string `json:"features"`
	Emojis                   []struct {
		Name          string   `json:"name"`
		Roles         []string `json:"roles"`
		ID            string   `json:"id"`
		RequireColons bool     `json:"require_colons"`
		Managed       bool     `json:"managed"`
		Animated      bool     `json:"animated"`
		Available     bool     `json:"available"`
	} `json:"emojis"`
	Banner            string  `json:"banner"`
	OwnerID           string  `json:"owner_id"`
	ApplicationID     *string `json:"application_id"`
	Region            *string `json:"region"`
	AfkChannelID      *string `json:"afk_channel_id"`
	AfkTimeout        int     `json:"afk_timeout"`
	SystemChannelID   *string `json:"system_channel_id"`
	WidgetEnabled     bool    `json:"widget_enabled"`
	WidgetChannelID   string  `json:"widget_channel_id"`
	VerificationLevel int     `json:"verification_level"`
	Roles             []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Permissions int    `json:"permissions"`
		Position    int    `json:"position"`
		Color       int    `json:"color"`
		Hoist       bool   `json:"hoist"`
		Managed     bool   `json:"managed"`
		Mentionable bool   `json:"mentionable"`
	} `json:"roles"`
	DefaultMessageNotifications int     `json:"default_message_notifications"`
	MfaLevel                    int     `json:"mfa_level"`
	ExplicitContentFilter       int     `json:"explicit_content_filter"`
	MaxPresences                *int    `json:"max_presences"`
	MaxMembers                  int     `json:"max_members"`
	MaxVideoChannelUsers        int     `json:"max_video_channel_users"`
	VanityURLCode               string  `json:"vanity_url_code"`
	PremiumTier                 int     `json:"premium_tier"`
	PremiumSubscriptionCount    int     `json:"premium_subscription_count"`
	SystemChannelFlags          int     `json:"system_channel_flags"`
	PreferredLocale             string  `json:"preferred_locale"`
	RulesChannelID              *string `json:"rules_channel_id"`
	PublicUpdatesChannelID      *string `json:"public_updates_channel_id"`
	SafetyAlertsChannelID       *string `json:"safety_alerts_channel_id"`
}
