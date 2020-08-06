package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	opt := os.Args[1]

	switch opt {
	case "create":
		CreateTX()
	case "get":
		GetTX()
	}

}
