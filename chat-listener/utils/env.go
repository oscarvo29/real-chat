package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(name string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("ERR:", err)
		return ""
	}
	return os.Getenv(name)
}
