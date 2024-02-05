package discordapp

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// GuildPreview represents a guild preview object returned by Discord's API
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

// Guild represents a guild object returned by Discord's API
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

// Member represents the object of a user within the context of a guild returned by Discord's API.
type Member struct {
	Avatar                     string     `json:"avatar"`
	CommunicationDisabledUntil time.Time  `json:"communication_disabled_until"`
	Flags                      int        `json:"flags"`
	JoinedAt                   time.Time  `json:"joined_at"`
	Nick                       string     `json:"nick"`
	Pending                    bool       `json:"pending"`
	PremiumSince               time.Time  `json:"premium_since"`
	Roles                      []string   `json:"roles"`
	UnusualDmActivityUntil     time.Time  `json:"unusual_dm_activity_until"`
	User                       MemberUser `json:"user"`
	Mute                       bool       `json:"mute"`
	Deaf                       bool       `json:"deaf"`
}

type MemberUser struct {
	ID                   string            `json:"id"`
	Username             string            `json:"username"`
	Avatar               string            `json:"avatar"`
	Discriminator        string            `json:"discriminator"`
	PublicFlags          int               `json:"public_flags"`
	PremiumType          int               `json:"premium_type"`
	Flags                int               `json:"flags"`
	Banner               string            `json:"banner"`
	AccentColor          string            `json:"accent_color"`
	GlobalName           string            `json:"global_name"`
	AvatarDecorationData map[string]string `json:"avatar_decoration_data"`
	BannerColor          string            `json:"banner_color"`
}

// FetchGuildPreview fetches the guild preview of the guild with the passed ID.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - DiscordError: Returned if a non-200 response is received & there is an error code in the body.
//   - ErrGuildNotFound: Returned if the guild does not exist or the bot is not in the guild.
func (b *Bot) FetchGuildPreview(guildID string) (*GuildPreview, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/guilds/"+guildID+"/preview", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var guildPreview GuildPreview
	resp, err := b.Request(req, &guildPreview)
	if err != nil {
		var discordErr *DiscordError
		if errors.As(err, &discordErr) {
			if discordErr.code == 10004 {
				return nil, ErrGuildNotFound
			}
		}
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		return nil, &UnexpectedResponseError{resp}
	}
	return &guildPreview, nil
}

// FetchGuild fetches the guild object of the guild with the passed ID.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - DiscordError: Returned if a non-200 response is received & there is an error code in the body.
//   - ErrGuildNotFound: Returned if the guild does not exist or the bot is not in the guild.
func (b *Bot) FetchGuild(guildID string) (*Guild, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/guilds/"+guildID, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var guild Guild
	resp, err := b.Request(req, &guild)
	if err != nil {
		var discordErr *DiscordError
		if errors.As(err, &discordErr) {
			if discordErr.code == 10004 {
				return nil, ErrGuildNotFound
			}
		}
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		return nil, &UnexpectedResponseError{resp}
	}
	return &guild, nil
}

// FetchGuildMember feches a member based on the passed member ID. The member must be in the guild with the passed guild ID.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - DiscordError: Returned if a non-200 response is received & there is an error code in the body.
//   - ErrUserNotFound: Returned if a user with the passed member ID could not be found in the guild.
//   - ErrGuildNotFound: Returned if the guild does not exist or the bot is not in the guild.
func (b *Bot) FetchGuildMember(guildID string, memberID string) (*Member, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/guilds/"+guildID+"/members/"+memberID, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var member Member
	resp, err := b.Request(req, &member)
	if err != nil {
		var discordErr *DiscordError
		if errors.As(err, &discordErr) {
			if discordErr.code == 10004 {
				return nil, ErrGuildNotFound
			}
			if discordErr.code == 10013 {
				return nil, ErrUserNotFound
			}
		}
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		return nil, &UnexpectedResponseError{resp}
	}
	return &member, nil
}

// AddMemberToGuild joins the user with the passed user ID to the guild with the passed guild ID using the access token.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - DiscordError: Returned if a non-200 response is received & there is an error code in the body.
//   - ErrAlreadyInGuild: Returned if the user is already in the guild.
//   - ErrMaxGuilds: Returned if the user is in max servers.
//   - ErrGuildNotFound: Returned if the bot is not in the guild.
//   - ErrInvalidAccessToken: Returned if the access token is invalid.
//   - ErrUserNotFound: Returned if a user with the passed member ID could not be found in the guild.
//   - ErrMissingPermissions: Returned if the bot does not have the permission to create invites.
func (b *Bot) AddMemberToGuild(accessToken string, userID string, guildID string) error {
	body := fmt.Sprintf(`{
		"access_token": "%s"
	}`, accessToken)
	req, err := http.NewRequest(http.MethodPut, BaseDiscordAPIURL+"/guilds/"+guildID+"/members/"+userID, strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("error forming request: %w", err)
	}
	resp, err := b.Request(req, nil)
	if err != nil {
		var discordErr *DiscordError
		if errors.As(err, &discordErr) {
			if discordErr.code == 10004 {
				return ErrGuildNotFound
			}
			if discordErr.code == 10013 {
				return ErrUserNotFound
			}
			if discordErr.code == 30001 {
				return ErrMaxGuilds
			}
			if discordErr.code == 50013 {
				return ErrMissingPermissions
			}
			if discordErr.code == 50025 {
				return ErrInvalidAccessToken
			}
		}
		return fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusCreated {
		if resp.Status == http.StatusNoContent {
			return ErrAlreadyInGuild
		}
		return &UnexpectedResponseError{resp}
	}
	return nil
}
