package integration

import (
	"os"
	"testing"

	"github.com/kodishim/discordapp/discordapp"
)

func TestNewApplication(t *testing.T) {
	_, err := discordapp.NewApplication(os.Getenv("TOKEN"), os.Getenv("SECRET"))
	if err != nil {
		t.Fatalf("Error creating new application: %s", err)
	}
}
