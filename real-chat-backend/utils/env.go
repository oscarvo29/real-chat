package utils

import (
	"fmt"
	"os"
)

func GetEnvValue(name string) string {
	value := os.Getenv(name)
	fmt.Println(value)
	return value
}
