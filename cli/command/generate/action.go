package generate

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/components/generate_progress"
	"github.com/slainless/markxus/cli/markxus/components/mod_name"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/slainless/markxus/nexus"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	gameCode, modId, err := resolveModId(c)
	if err != nil {
		return err
	}

	app, err := createClient(ctx, c)
	if err != nil {
		return err
	}

	progress := generate_progress.New(gameCode, modId).SetStatus(generate_progress.StatusStarted)
	progress.Update(generate_progress.StartMsg(0))

	program := tea.NewProgram(view{
		progress: progress,
	})

	var generated *markxus.Generated
	var generateError error
	go func() {
		defer program.Quit()
		defer program.Send(DoneMsg(0))

		generated, err = app.Generate(ctx, gameCode, modId,
			markxus.WithOnModFetched(func(ctx context.Context, mod *nexus.SchemaMod) error {
				if err := checkMarkdown(createOutputPath(mod.Name)); err != nil {
					return err
				}

				program.Send(generate_progress.ModDiscoveredMsg(mod))
				return nil
			}),
			markxus.WithOnLlmStreamConsuming(func(ctx context.Context, streamData any, currentOutput *string) error {
				program.Send(generate_progress.GenerationProgressMsg(0))
				return nil
			}),
		)
		if err != nil {
			generateError = err
			return
		}

		program.Send(generate_progress.DoneMsg(0))

		err = checkMarkdown(createOutputPath(generated.Mod.Name))
		if err != nil {
			generateError = err
			return
		}

		err = writeMarkdown(generated)
		if err != nil {
			generateError = err
			return
		}
	}()

	if _, err := program.Run(); err != nil {
		return err
	}

	if generateError != nil {
		return err
	}

	theme := style.GetTheme()
	fmt.Println(
		style.Card().Render(
			fmt.Sprintf("%s\n%s\n\n%s\n%s",
				mod_name.View(fmt.Sprint(gameCode), fmt.Sprint(generated.Mod.ModId)),
				generated.Mod.Name,
				"Saved to:",
				theme.Focused.Description.Render(createOutputPath(generated.Mod.Name)),
			),
		),
	)

	return nil
}

func resolveModId(c *cli.Command) (string, string, error) {
	fallbackGameCode := config.Config.Helper.FallbackGameCode
	gameCode := c.Args().Get(0)
	modId := c.Args().Get(1)

	if modId == "" && fallbackGameCode == "" {
		return "", "", fmt.Errorf(
			"%s\n%s",
			"no mod id given or fallback game code is not set.",
			"either set fallback game code or use 2 arguments (game code and mod id)",
		)
	}

	if modId == "" {
		modId = gameCode
		gameCode = fallbackGameCode
	}

	return gameCode, modId, nil
}
