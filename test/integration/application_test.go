package integration

import (
	"os"
	"testing"

	"github.com/kodishim/gocord/pkg/application"
)

func TestNewApplication(t *testing.T) {
	_, err := application.NewApplication(os.Getenv("TEST_TOKEN"), os.Getenv("TEST_SECRET"))
	if err != nil {
		t.Fatalf("Error creating new application: %s", err)
	}
}
