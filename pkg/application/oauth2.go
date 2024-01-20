package application

import (
	"fmt"
	"net/http"

	"github.com/kodishim/gocord/pkg/models"
	"github.com/kodishim/gocord/pkg/util"
)

// FetchAuthInfo fetches the authorization info using the passed access token.
//
// Possible Errors:
//   - ErrInvalidAccessToken: Returned if the access token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func FetchAuthInfo(accessToken string) (*models.AuthInfo, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_DISCORD_API_URL+"/oauth2/@me", nil)
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