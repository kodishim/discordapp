package integration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %s", err)
	}
	checkVariable := func(variableName string) {
		if os.Getenv(variableName) == "" {
			log.Fatalf("%s env variable is missing", variableName)
		}
	}
	variablesToCheck := []string{"TOKEN", "SECRET", "GUILD", "MEMBER", "ACCESS_TOKEN", "AUTHORIZED_USER"}
	for _, v := range variablesToCheck {
		checkVariable(v)
	}
}
