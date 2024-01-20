package application

import "fmt"

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
