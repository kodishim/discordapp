package discordapp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kodishim/discordapp/discordapp/util"
)

// A bot represents a Discord bot.
type Bot struct {
	Token       string
	Application *ApplicationInfo
}

// ApplicationInfo represents an application object returned by Discord's API
type ApplicationInfo struct {
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

// DiscordErrorResponse represents an error response returned by the Discord API.
type DiscordErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// NewBot creates & returns a pointer to a bot using the passed token.
//
// A bot's token can be found at https://discord.com/developers/applications.
//
// Possible Errors:
//   - ErrUnauthorized - Returned if the token is invalid.
func NewBot(token string) (*Bot, error) {
	bot := &Bot{
		Token:       token,
		Application: nil,
	}
	var err error
	bot.Application, err = bot.FetchApplicationInfo()
	if err != nil {
		return nil, fmt.Errorf("error fetching bot's application object: %w", err)
	}
	return bot, nil
}

// Request makes a request using the Bot's token for authentication & unmarshals the response into unmarshalTo if unmarshalTo is not nil.
//
// unmarshalTo should be a pointer or nil.
//
// If a response with status code less than 200 or greater than 299 is received an error is returned.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - DiscordError: Returned if a non-200 response is received & there is an error code in the body.
//   - UnexpectedResponseError: Returned if a non-200 response is received without a code in the body.
func (b *Bot) Request(req *http.Request, unmarshalTo any) (*util.Response, error) {
	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set("Authorization", "Bot "+b.Token)
	resp, err := util.MakeRequest(req, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status < 200 || resp.Status > 299 {
		if resp.Status == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		}
		var discordErrorResp struct {
			Message string `json:"mesage"`
			Code    *int   `json:"code"`
		}
		err = json.Unmarshal(resp.Body, &discordErrorResp)
		if err != nil || discordErrorResp.Code == nil {
			return nil, &UnexpectedResponseError{resp}
		}
		return nil, &DiscordError{
			response: resp,
			code:     *discordErrorResp.Code,
			message:  discordErrorResp.Message,
		}
	}
	return resp, nil
}

// FetchApplication fetches the bot's application object.
//
// unmarshalTo should be a pointer or nil.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
func (b *Bot) FetchApplicationInfo() (*ApplicationInfo, error) {
	req, err := http.NewRequest(http.MethodGet, BaseDiscordAPIURL+"/oauth2/applications/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var application ApplicationInfo
	resp, err := b.Request(req, &application)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		return nil, &UnexpectedResponseError{
			response: resp,
		}
	}
	err = json.Unmarshal(resp.Body, &application)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return &application, nil
}
