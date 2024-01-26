package models

import "time"

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
