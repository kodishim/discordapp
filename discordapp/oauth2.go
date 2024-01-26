package discordapp

import (
	"fmt"
	"net/http"

	"github.com/kodishim/discordapp/discordapp/models"
	"github.com/kodishim/discordapp/discordapp/util"
)

// FetchAuthInfo fetches the authorization info using the passed access token.
//
// Possible Errors:
//   - ErrInvalidAccessToken: Returned if the access token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func FetchAuthInfo(accessToken string) (*models.AuthInfo, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/oauth2/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	var authInfo models.AuthInfo
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
func FetchAuthUser(accessToken string) (*models.AuthorizedUser, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/users/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	var user models.AuthorizedUser
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
