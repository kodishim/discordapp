package discordapp

import (
	"errors"
	"fmt"

	"github.com/kodishim/discordapp/discordapp/util"
)

var ErrGuildNotFound = errors.New("guild_not_found")
var ErrUserNotFound = errors.New("user_not_found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrAlreadyInGuild = errors.New("already_in_guild")
var ErrMaxGuilds = errors.New("max_guilds")
var ErrInvalidAccessToken = errors.New("invalid_access_token")
var ErrMissingPermissions = errors.New("missing_permissions")

type UnexpectedResponseError struct {
	response *util.Response
}

func (e *UnexpectedResponseError) Error() string {
	return fmt.Sprintf("unexpected response: %d %s", e.response.Status, e.response.Body)
}

type DiscordError struct {
	response *util.Response
	code     int
	message  string
}

func (e *DiscordError) Error() string {
	return fmt.Sprintf("error from discord api: %d %d %s", e.response.Status, e.code, e.message)
}
