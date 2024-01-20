package models

type Application struct {
	BotPublic           bool   `json:"bot_public"`
	BotRequireCodeGrant bool   `json:"bot_require_code_grant"`
	CoverImage          string `json:"cover_image"`
	Description         string `json:"description"`
	GuildID             string `json:"guild_id"`
	Icon                string `json:"icon"`
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Owner               struct {
		Avatar        string `json:"avatar"`
		Discriminator string `json:"discriminator"`
		Flags         int    `json:"flags"`
		ID            string `json:"id"`
		Username      string `json:"username"`
	} `json:"owner"`
	PrimarySkuID string `json:"primary_sku_id"`
	Slug         string `json:"slug"`
	Summary      string `json:"summary"`
	Team         struct {
		Icon    string `json:"icon"`
		ID      string `json:"id"`
		Members []struct {
			MembershipState int      `json:"membership_state"`
			Permissions     []string `json:"permissions"`
			TeamID          string   `json:"team_id"`
			User            struct {
				Avatar        string `json:"avatar"`
				Discriminator string `json:"discriminator"`
				ID            string `json:"id"`
				Username      string `json:"username"`
			} `json:"user"`
		} `json:"members"`
	} `json:"team"`
	VerifyKey string `json:"verify_key"`
}
