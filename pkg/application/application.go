package application

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/kodishim/gocord/pkg/util"
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

var ScopeIdentify = "identify"
var ScopeEmail = "email"
var ScopeConnections = "connections"
var ScopeGuildsMembersRead = "guilds.members.read"
var ScopeRpcNotificationsRead = "rpc.notifications.read"
var ScopeRpcVideoRead = "rpc.video.read"
var ScopeRpcScreenshareWrite = "rpc.screenshare.write"
var ScopeWebhookIncoming = "webhook.incoming"
var ScopeApplicationsBuildsRead = "applications.builds.read"
var ScopeApplicationsEntitlements = "applications.entitlements"
var ScopeRelationshipsRead = "relationships.read"
var ScopeRoleConnectionsWrite = "role_connections.write"
var ScopeDmChannelsRead = "dm_channels.read"
var ScopeActivitiesWrite = "activities.write"
var ScopeApplicationsStoreUpdate = "applications.store.update"
var ScopeApplicationsBuildsUpload = "applications.builds.upload"
var ScopeRpcScreenshareRead = "rpc.screenshare.read"
var ScopeRpcVoiceWrite = "rpc.voice.write"
var ScopeRpcGuildsJoin = "rpc.guilds.join"
var ScopeGuilds = "guilds"
var ScopeGdmJoin = "gdm.join"
var ScopeRpcVoiceRead = "rpc.voice.read"
var ScopeRpcVideoWrite = "rpc.video.write"
var ScopeRpcActivitiesWrite = "rpc.activities.write"
var ScopeMessagesRead = "messages.read"
var ScopeApplicationsCommands = "applications.commands"
var ScopeActivitiesRead = "activities.read"
var ScopeVoice = "voice"
var ScopeApplicationsCommandsPermissionsUpdate = "applications.commands.permissions.update"

// CreateAuthLink creates an authorization link. State can be "" for no state. Scope can be nil for no scopes.
//
// The redirectURI must be configured on the Discord application at https://discord.com/developers/applications.
func (a *Application) CreateAuthLink(redirectURI string, state string, scopes []string) string {
	link := BASE_DISCORD_API_URL + "/oauth2/authorize"
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

// FetchAccessToken fetches an access & refresh token using the passed code.
//
// A code can be found in the payload of a Discord callback request during the OAuth2 process.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if authentication failed.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func (a *Application) FetchAccessToken(code string, redirectURI string) (accessToken string, refreshToken string, expiresIn int, err error) {
	credByte := []byte(fmt.Sprintf("%s:%s", a.Bot.Application.ID, a.Secret))
	cred := base64.StdEncoding.EncodeToString(credByte)
	formData := url.Values{}
	formData.Set("client_id", a.Bot.Application.ID)
	formData.Set("client_secret", a.Secret)
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("redirect_uri", redirectURI)
	req, err := http.NewRequest("POST", BASE_DISCORD_API_URL+"/api/oauth2/token", strings.NewReader(formData.Encode()))
	if err != nil {
		err = fmt.Errorf("error forming request: %w", err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", cred))
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
		if resp.Status == http.StatusUnauthorized {
			err = ErrUnauthorized
			return
		}
		err = &UnexpectedResponseError{resp}
		return
	}
	accessToken = respBody.AccessToken
	refreshToken = respBody.RefreshToken
	expiresIn = respBody.ExpiresIn
	return
}

// RefreshAccessToken exchanges the passed refresh token for a new access token & refresh token.
// Expires in represents the seconds until the token expires.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if authentication failed.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func (a *Application) RefreshAccessToken(refreshToken string) (newAccessToken string, newRefreshToken string, expiresIn int, err error) {
	formData := url.Values{}
	formData.Set("client_id", a.Bot.Application.ID)
	formData.Set("client_secret", a.Secret)
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", refreshToken)
	req, err := http.NewRequest(http.MethodPost, BASE_DISCORD_API_URL+"/oauth2/token", strings.NewReader(formData.Encode()))
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
		if resp.Status == http.StatusUnauthorized {
			err = ErrUnauthorized
			return
		}
		err = &UnexpectedResponseError{resp}
		return
	}
	newAccessToken = respBody.AccessToken
	newRefreshToken = respBody.RefreshToken
	expiresIn = respBody.ExpiresIn
	return
}
