package models

type AuthorizedUser struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`
	Banner        string `json:"banner"`
	AccentColor   int    `json:"accent_color"`
	PremiumType   int    `json:"premium_type"`
	PublicFlags   int    `json:"public_flags"`
}
