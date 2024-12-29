package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slainless/markxus/cli/markxus/command"
)

func main() {
	_ = godotenv.Load(
		".env",
		".env.local",
		".env.markxus",
		".markxus",
	)
	if err := command.Main.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
