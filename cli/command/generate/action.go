package generate

import (
	"context"
	"fmt"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/components/generate_progress"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/cli/markxus/internal/style"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
	"github.com/slainless/markxus/resty"
	"github.com/urfave/cli/v3"
)

func action(ctx context.Context, c *cli.Command) error {
	gameCode, modId, err := resolveModId(c)
	if err != nil {
		return err
	}

	var headerTemplate *template.Template
	if config.Config.Generation.HeaderFormat != markxus.DefaultMarkdownHeaderFormat {
		headerTemplate, err = template.New("markxus.header").Parse(config.Config.Generation.HeaderFormat)
		if err != nil {
			return err
		}
	}

	resty := resty.NewRestyClient()
	nexusClient, err := nexus.NewClient(
		nexus.WithApiKey(config.Config.Nexus.ApiKey),
		nexus.WithHTTPDriver(resty),
		nexus.WithUrlGetModFormat(config.Config.Nexus.Url.GetModFormat),
	)
	if err != nil {
		return err
	}

	genaiClient, err := genai.NewGenAiClient(ctx,
		genai.WithApiKey(config.Config.GenAi.ApiKey),
		genai.WithModelName(config.Config.GenAi.ModelName),
	)
	if err != nil {
		return err
	}

	app := markxus.NewMarkxus(nexusClient, genaiClient,
		markxus.WithPromptFormat(config.Config.GenAi.Prompt),
		markxus.WithUrlModPageFormat(config.Config.Nexus.Url.ModPageFormat),
		markxus.WithMarkdownHeaderTemplate(headerTemplate),
	)

	theme := style.GetTheme()
	progress := generate_progress.New(gameCode, modId).SetStatus(generate_progress.StatusStarted)
	progress.Update(generate_progress.StartMsg(0))

	program := tea.NewProgram(progress)

	derivedCtx, cancel := context.WithCancel(ctx)
	var teaError error
	go func() {
		if _, err := program.Run(); err != nil {
			fmt.Println(
				style.Card().Render(
					theme.Focused.ErrorMessage.Render(err.Error()),
				),
			)
			teaError = err
			cancel()
		}
	}()

	generated, err := app.Generate(derivedCtx, gameCode, modId,
		markxus.WithOnModFetched(func(ctx context.Context, mod *nexus.SchemaMod) error {
			program.Send(generate_progress.ModDiscoveredMsg(mod))
			return nil
		}),
		markxus.WithOnLlmStreamConsuming(func(ctx context.Context, streamData any, currentOutput *string) error {
			program.Send(generate_progress.GenerationProgressMsg(0))
			return nil
		}),
	)

	if err == derivedCtx.Err() {
		return teaError
	} else if err != nil {
		return err
	}

	program.ReleaseTerminal()
	program.Quit()
	return nil
	fmt.Println(generated)
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
