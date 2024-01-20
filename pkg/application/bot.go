package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kodishim/gocord/pkg/models"
	"github.com/kodishim/gocord/pkg/util"
)

// A bot represents a Discord bot.
type Bot struct {
	Token       string
	Application *models.Application
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
	bot.Application, err = bot.FetchApplication()
	if err != nil {
		return nil, fmt.Errorf("error fetching bot's application object: %w", err)
	}
	return bot, nil
}

// Request makes a request using the Bot's token for authentication & unmarshals the response into unmarshalTo if unmarshalTo is not nil.
//
// unmarshalTo should be a pointer or nil.
//
// Possible Errors:
//   - ErrUnauthorized - Returned if the bot's token is invalid.
func (b *Bot) Request(req *http.Request, unmarshalTo any) (*util.Response, error) {
	if req.Header == nil {
		req.Header = http.Header{}
	}
	req.Header.Set("Authorization", "Bot "+b.Token)
	resp, err := util.MakeRequest(req, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status == http.StatusUnauthorized {
		return nil, ErrUnauthorized
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
func (b *Bot) FetchApplication() (*models.Application, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_DISCORD_API_URL+"/oauth2/applications/@me", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var application models.Application
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

// FetchGuildPreview fetches the guild preview of the guild with the passed ID.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - ErrGuildNotFound: Returned if the guild does not exist or the bot is not in the guild.
func (b *Bot) FetchGuildPreview(guildID string) (*models.GuildPreview, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_DISCORD_API_URL+"/guilds/"+guildID+"/preview", nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var guildPreview models.GuildPreview
	resp, err := b.Request(req, &guildPreview)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		if resp.Status == http.StatusNotFound {
			return nil, ErrGuildNotFound
		}
		return nil, &UnexpectedResponseError{response: resp}
	}
	return &guildPreview, nil
}

// FetchGuild fetches the guild object of the guild with the passed ID.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - ErrGuildNotFound: Returned if the guild does not exist or the bot is not in the guild.
func (b *Bot) FetchGuild(guildID string) (*models.Guild, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_DISCORD_API_URL+"/guilds/"+guildID, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var guild models.Guild
	resp, err := b.Request(req, &guild)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		if resp.Status == http.StatusNotFound {
			return nil, ErrGuildNotFound
		}
		return nil, &UnexpectedResponseError{response: resp}
	}
	return &guild, nil
}

// FetchGuildMember feches a member based on the passed member ID. The member must be in the guild with the passed guild ID.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - ErrUserNotFound: Returned if a user with the passed member ID could not be found in the guild.
//   - ErrGuildNotFound: Returned if the guild does not exist or the bot is not in the guild.
func (b *Bot) FetchGuildMember(guildID string, memberID string) (*models.Member, error) {
	req, err := http.NewRequest(http.MethodGet, BASE_DISCORD_API_URL+"/guilds/"+guildID+"/members/"+memberID, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming request: %w", err)
	}
	var member models.Member
	resp, err := b.Request(req, &member)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		if resp.Status == http.StatusNotFound {
			if strings.Contains(string(resp.Body), "Unknown User") {
				return nil, ErrUserNotFound
			}
			return nil, ErrGuildNotFound
		}
		return nil, &UnexpectedResponseError{response: resp}
	}
	err = json.Unmarshal(resp.Body, &member)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}
	return &member, nil
}

// AddMemberToGuild joins the user with the passed user ID to the guild with the passed guild ID using the access token.
//
// Possible Errors:
//   - ErrUnauthorized: Returned if the bot's token is invalid.
//   - UnexpectedResponseError: Returned if an unexpected response was received.
//   - ErrAlreadyInGuild: Returned if the user is already in the guild.
//   - ErrMaxServers: Returned if the user is in max servers.
//   - ErrGuildNotFound: Returned if the bot is not in the guild.
//   - ErrInvalidAccessToken: Returned if the access token is invalid.
func (b *Bot) AddMemberToGuild(accessToken string, userID string, guildID string) error {
	body := fmt.Sprintf(`{
		"access_token": "%s"
	}`, accessToken)
	req, err := http.NewRequest(http.MethodPut, BASE_DISCORD_API_URL+"/guilds/"+guildID+"/member/"+userID, strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("error forming request: %w", err)
	}
	resp, err := b.Request(req, nil)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	if resp.Status != http.StatusOK {
		if resp.Status == http.StatusNoContent {
			return ErrAlreadyInGuild
		}
		if resp.Status == http.StatusBadRequest {
			return ErrMaxServers
		}
		if resp.Status == http.StatusForbidden {
			_, err = b.FetchGuildPreview(guildID)
			if err != nil {
				return ErrGuildNotFound
			}
			return ErrInvalidAccessToken
		}
		return &UnexpectedResponseError{resp}
	}
	return nil
}
