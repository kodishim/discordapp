package application

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"example.com/pkg/util"
)

// An application represents a Discord application.
type Application struct {
	Bot    *Bot
	Secret string
}

// NewApplication creates & returns a pointer to a Discord application using the passed token & secret.
//
// An application's token and secret can be found at https://discord.com/developers/applications.
//
// Possible Errors:
//   - ErrUnauthorized - Returned if the passed token is invalid.
func NewApplication(token string, secret string) (*Application, error) {
	bot, err := NewBot(token)
	if err != nil {
		return nil, fmt.Errorf("error creating new bot: %w", err)
	}
	return &Application{Bot: bot, Secret: secret}, nil
}

// RefreshAccessToken exchanges the passed refresh token for a new access token & refresh token.
// Expires in represents the seconds until the token expires.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func (a *Application) RefreshAccessToken(refreshToken string) (newAccessToken string, newRefreshToken string, expiresIn int, err error) {
	formData := url.Values{}
	formData.Set("client_id", a.Bot.Application.ID)
	formData.Set("client_secret", a.Secret)
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", refreshToken)
	req, err := http.NewRequest("POST", BASE_DISCORD_API_URL+"/oauth2/token", strings.NewReader(formData.Encode()))
	if err != nil {
		err = fmt.Errorf("error forming request: %w", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var respBody struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	resp, err := util.MakeRequest(req, nil, &respBody)
	if err != nil {
		err = fmt.Errorf("error making request: %w", err)
		return
	}
	if resp.Status != http.StatusOK {
		err = &UnexpectedResponseError{resp}
		return
	}
	newAccessToken = respBody.AccessToken
	newRefreshToken = respBody.RefreshToken
	expiresIn = respBody.ExpiresIn
	return
}
