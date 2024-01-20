package integration_test

import (
	"os"
	"testing"

	"github.com/kodishim/gocord/pkg/application"
)

func TestNewBot(t *testing.T) {
	_, err := application.NewBot(os.Getenv("TEST_TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
}

func TestFetchApplication(t *testing.T) {
	bot, err := application.NewBot(os.Getenv("TEST_TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	_, err = bot.FetchApplication()
	if err != nil {
		t.Fatalf("Error fetching application: %s", err)
	}
}

func TestFetchGuildPreview(t *testing.T) {
	bot, err := application.NewBot(os.Getenv("TEST_TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	_, err = bot.FetchGuildPreview(os.Getenv("TEST_GUILD"))
	if err != nil {
		t.Fatalf("Error fetching guild preview: %s", err)
	}
}

func TestFetchGuildMember(t *testing.T) {
	bot, err := application.NewBot(os.Getenv("TEST_TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}

	_, err = bot.FetchGuildMember(os.Getenv("TEST_GUILD"), os.Getenv("TEST_MEMBER"))
	if err != nil {
		t.Fatalf("Error fetching member: %s", err)
	}
}
