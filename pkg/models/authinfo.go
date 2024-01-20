package models

import "time"

type AuthInfo struct {
	Application struct {
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Icon                string `json:"icon"`
		Description         string `json:"description"`
		Hook                bool   `json:"hook"`
		BotPublic           bool   `json:"bot_public"`
		BotRequireCodeGrant bool   `json:"bot_require_code_grant"`
		VerifyKey           string `json:"verify_key"`
	} `json:"application"`
	Scopes  []string  `json:"scopes"`
	Expires time.Time `json:"expires"`
	User    struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Avatar        string `json:"avatar"`
		Discriminator string `json:"discriminator"`
		GlobalName    string `json:"global_name"`
		PublicFlags   int    `json:"public_flags"`
	} `json:"user"`
}
