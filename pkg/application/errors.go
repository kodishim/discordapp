package application

import (
	"errors"
	"fmt"

	"example.com/pkg/util"
)

var ErrGuildNotFound = errors.New("guild_not_found")
var ErrUserNotFound = errors.New("user_not_found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrAlreadyInGuild = errors.New("already_in_guild")
var ErrMaxServers = errors.New("max_servers")
var ErrInvalidAccessToken = errors.New("invalid_access_token")

type UnexpectedResponseError struct {
	response *util.Response
}

func (e *UnexpectedResponseError) Error() string {
	return fmt.Sprintf("unexpectes response: %d %s", e.response.Status, e.response.Body)
}
