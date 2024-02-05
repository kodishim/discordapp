package integration_test

import (
	"os"
	"testing"

	"github.com/kodishim/discordapp/discordapp"
)

func TestNewBot(t *testing.T) {
	_, err := discordapp.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
}

func TestFetchApplicationInfo(t *testing.T) {
	bot, err := discordapp.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	_, err = bot.FetchApplicationInfo()
	if err != nil {
		t.Fatalf("Error fetching application info: %s", err)
	}
}

func TestFetchGuildPreview(t *testing.T) {
	bot, err := discordapp.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	_, err = bot.FetchGuildPreview(os.Getenv("GUILD"))
	if err != nil {
		t.Fatalf("Error fetching guild preview: %s", err)
	}
	_, err = bot.FetchGuildPreview("111")
	if err != discordapp.ErrGuildNotFound {
		t.Fatalf("Expected ErrGuildNotFound: %s", err)
	}
}

func TestFetchGuild(t *testing.T) {
	bot, err := discordapp.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	_, err = bot.FetchGuild(os.Getenv("GUILD"))
	if err != nil {
		t.Fatalf("Error fetching guild preview: %s", err)
	}
	_, err = bot.FetchGuild("111")
	if err != discordapp.ErrGuildNotFound {
		t.Fatalf("Expected ErrGuildNotFound: %s", err)
	}
}

func TestFetchGuildMember(t *testing.T) {
	bot, err := discordapp.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	_, err = bot.FetchGuildMember(os.Getenv("GUILD"), os.Getenv("MEMBER"))
	if err != nil {
		t.Fatalf("Error fetching member: %s", err)
	}
	_, err = bot.FetchGuildMember(("111"), os.Getenv("MEMBER"))
	if err != discordapp.ErrGuildNotFound {
		t.Fatalf("Error ErrGuildNotFound: %s", err)
	}
	_, err = bot.FetchGuildMember(os.Getenv("GUILD"), "111")
	if err != discordapp.ErrUserNotFound {
		t.Fatalf("Expected ErrUserNotFound: %s", err)
	}
}

func TestAddMemberToGuild(t *testing.T) {
	bot, err := discordapp.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating new bot: %s", err)
	}
	err = bot.AddMemberToGuild("111", os.Getenv("AUTHORIZED_USER"), os.Getenv("GUILD"))
	if err != discordapp.ErrInvalidAccessToken {
		t.Fatalf("Expected ErrInvalidAccessToken: %s", err)
	}
	err = bot.AddMemberToGuild(os.Getenv("ACCESS_TOKEN"), "111", os.Getenv("GUILD"))
	if err != discordapp.ErrInvalidAccessToken {
		t.Fatalf("Expected ErrInvalidAccessToken: %s", err)
	}
	err = bot.AddMemberToGuild(os.Getenv("ACCESS_TOKEN"), os.Getenv("AUTHORIZED_USER"), "111")
	if err != discordapp.ErrGuildNotFound {
		t.Fatalf("Error ErrGuildNotFound: %s", err)
	}
	err = bot.AddMemberToGuild(os.Getenv("ACCESS_TOKEN"), os.Getenv("AUTHORIZED_USER"), os.Getenv("GUILD"))
	if err != nil {
		t.Fatalf("Error adding user to guild: %s", err)
	}
	err = bot.AddMemberToGuild(os.Getenv("ACCESS_TOKEN"), os.Getenv("AUTHORIZED_USER"), os.Getenv("GUILD"))
	if err != discordapp.ErrAlreadyInGuild {
		t.Fatalf("Expected ErrAlreadyInGuild: %s", err)
	}
}
