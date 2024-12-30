package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slainless/markxus/cli/markxus/command"
	"github.com/slainless/markxus/cli/markxus/internal/style"
)

func main() {
	_ = godotenv.Load(
		".env",
		".env.local",
		".env.markxus",
		".markxus",
	)
	if err := command.Main.Run(context.Background(), os.Args); err != nil {
		fmt.Println(style.Card().Render(
			fmt.Sprintf("Error\n%s", style.GetTheme().Focused.ErrorMessage.Render(err.Error())),
		))
	}
}
