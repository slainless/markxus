package generate

import (
	"context"
	"text/template"

	"github.com/slainless/markxus"
	"github.com/slainless/markxus/cli/markxus/config"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
	"github.com/slainless/markxus/resty"
	"github.com/urfave/cli/v3"
)

func createClient(ctx context.Context, c *cli.Command) (*markxus.Markxus, error) {
	var err error

	var headerTemplate *template.Template
	if config.Config.Generation.HeaderFormat == markxus.DefaultMarkdownHeaderFormat {
		headerTemplate = markxus.DefaultMarkdownHeaderTemplate
	} else {
		headerTemplate, err = template.New("markxus.header").Parse(config.Config.Generation.HeaderFormat)
		if err != nil {
			return nil, err
		}
	}

	resty := resty.NewRestyClient()
	nexusClient, err := nexus.NewClient(
		nexus.WithApiKey(config.Config.Nexus.ApiKey),
		nexus.WithHTTPDriver(resty),
		nexus.WithUrlGetModFormat(config.Config.Nexus.Url.GetModFormat),
	)
	if err != nil {
		return nil, err
	}

	genaiClient, err := genai.NewGenAiClient(ctx,
		genai.WithApiKey(config.Config.GenAi.ApiKey),
		genai.WithModelName(config.Config.GenAi.ModelName),
	)
	if err != nil {
		return nil, err
	}

	return markxus.NewMarkxus(nexusClient, genaiClient,
		markxus.WithPromptFormat(config.Config.GenAi.Prompt),
		markxus.WithUrlModPageFormat(config.Config.Nexus.Url.ModPageFormat),
		markxus.WithMarkdownHeaderTemplate(headerTemplate),
	), nil
}
