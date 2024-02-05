package discordapp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kodishim/discordapp/discordapp/util"
)

// AuthorizedUser represents the object of a user authorized to an application.
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

// AuthInfo represents an authorization object returned by Discord's API.
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

// CreateAuthLink creates an authorization link. State can be "" for no state. Scope can be nil for no scopes.
//
// The redirectURI must be configured on the Discord application at https://discord.com/developers/applications.
func (a *Application) CreateAuthLink(redirectURI string, state string, scopes []string) string {
	link := BaseDiscordAPIURL + "/oauth2/authorize"
	link += "?client_id=" + a.Bot.Application.ID
	if scopes != nil {
		link += "&scope=" + strings.Join(scopes, "+")
	}
	link += "&response_type=code"
	link += "&redirect_uri=" + url.QueryEscape(redirectURI)
	if state != "" {
		link += "&state=" + state
	}
	return link
}

// FetchAuthInfo fetches the authorization info using the passed access token.
//
// Possible Errors:
//   - ErrInvalidAccessToken: Returned if the access token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func FetchAuthInfo(accessToken string) (*AuthInfo, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/oauth2/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	var authInfo AuthInfo
	resp, err := util.MakeRequest(req, http.DefaultClient, &authInfo)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		if resp.Status == http.StatusUnauthorized {
			return nil, ErrInvalidAccessToken
		}
		return nil, &UnexpectedResponseError{resp}
	}
	return &authInfo, nil
}

// FetchAuthUser fetches a user object from an access token.
//
// Possible Errors:
//   - ErrInvalidAccessToken: Returned if the access token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func FetchAuthUser(accessToken string) (*AuthorizedUser, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/users/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	var user AuthorizedUser
	resp, err := util.MakeRequest(req, http.DefaultClient, &user)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		if resp.Status == http.StatusUnauthorized {
			return nil, ErrInvalidAccessToken
		}
		return nil, &UnexpectedResponseError{resp}
	}
	return &user, nil
}
