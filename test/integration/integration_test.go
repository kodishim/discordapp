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
	variablesToCheck := []string{"TEST_TOKEN", "TEST_SECRET", "TEST_GUILD", "TEST_MEMBER"}
	for _, v := range variablesToCheck {
		checkVariable(v)
	}
}
