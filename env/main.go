package main

import (
	"log"
	"os"
)

func GetEnv(k string, d string) string {

	if val := os.Getenv(k); val != "" {
		return val
	}

	return d
}

func main() {

	log.Println(
		"Value Of KEY",
		// Will look for .env value Key
		// and return "FOO" if not found.
		GetEnv("KEY", "FOO"),
	)
}
