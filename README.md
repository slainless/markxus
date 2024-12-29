# Markxus

<img alt="Custom badge" src="https://img.shields.io/endpoint?url=https%3A%2F%2Fgist.githubusercontent.com%2Fradj307%2Fe9a80731ee236cc67fb00b698e75201e%2Fraw%2F5230074dfb1a60fba917a1232f9382fa5cfec5db%2Fendpoint.json&style=for-the-badge"> ![Google Gemini](https://img.shields.io/badge/google%20gemini-8E75B2?style=for-the-badge&logo=google%20gemini&logoColor=white) ![Markdown](https://img.shields.io/badge/markdown-%23000000.svg?style=for-the-badge&logo=markdown&logoColor=white) ![Obsidian](https://img.shields.io/badge/Obsidian-%23483699.svg?style=for-the-badge&logo=obsidian&logoColor=white)

Markxus is an LLM-powered markdown converter targeted for [Nexus Mods](https://www.nexusmods.com/)'s mod page.
By default, Markxus is opinionated and it converts nexusmods HTML+BB format into markdown that is catered for Obsidian markdown format.
The `cli` package will use [resty](https://github.com/go-resty/resty) as its HTTP client 
and [google-generative-ai](https://github.com/google/generative-ai-go) as its LLM client. 

This package provides built-in supports for both [Resty (./resty)](./resty) and [Google Generative AI (./genai)](./genai/).

However, the main package by design is client-agnostic and any combination of HTTP and LLM client can be used, but
a dedicated `cli` package must be created instead to make use of these interfaces.

## Obsidian markdown

The default markdown header format contains frontmatter and only a handful number of editors are capable of making use of these data.
Obsidian is prominently known its for features and its rich community support. Obsidian by default support showing and editing frontmatter. 
Obsidian community also has many plugins that can handle these data.

I planned to support these plugins:
  - [Banners](https://github.com/noatpad/obsidian-banners)
  - [Iconize](https://github.com/FlorianWoelki/obsidian-iconize)

More support to be added later.

## CLI Usage

### Installation

Go is needed to install the CLI. I will make an executable build pipeline later if I feel like it.

```sh
go install github.com/slainless/markxus/cli/markxus@latest
```

### Configuration

Most of the functionality of the apps relies on API key of Nexus Mods and Google Generative AI so 
it is important to set up global configuration (or local configuration) or both vars will need to be supplied to the binary per usage.

To initiate global (or local) configuration, run:

```sh
markxus config init
```

It will create `.markxus.yaml` to user directory (or cwd, depending on config type). It will also prompt you with both API key and store it to OS keyring.
It is recommended to store the API key in OS-managed credential storage rather
than disk.

To change API key later without going through config init, run:

```sh
markxus config set NEXUS_API_KEY $new_key
```

or

```sh
markxus config set GEN_AI_API_KEY $new_key
```

Command flags is available to `config init` to alter its behaviour:
- `--force` or `FORCE_OVERWRITE`: Force overwrite to file if exist, skipping prompt
- `--type` or `FORCE_CONFIG_TYPE`: Set config type, whether to generate to `global` or `local` (cwd)

#### Available Options

All these options can also be set from env vars or CLI flag.

##### Google Generative AI

- **API key** 

	```yaml
	Flag: --genai-key, --gk
	Env: GEN_AI_API_KEY
	YAML: genai_api_key
	```

	Self-explanatory. **Required**.

- **Model name**

	```yaml
	Flag: --model, --m
	Env: GEN_AI_API_KEY
	YAML: genai_api_key
	```

	LLM model used for conversion. Defaults to `gemini-1.5-flash`.

- **Prompt format** 

	```yaml
	Flag: --prompt, --p
	Env: GEN_AI_PROMPT_FORMAT
	YAML: genai_prompt_format
	```

	Prompt format used to query the LLM. Defaults to [prompt.txt](./prompt.txt).

	Prompt can be changed the direction of the resulting markdown.

##### Nexus Mods

- **API key**

	```yaml
	Flag: --nexus-key, --nk
	Env: NEXUS_API_KEY
	YAML: nexus_api_key
	```

	Self-explanatory. **Required**.

- **API mod url format** 

	```yaml
	Flag: --api-url-format, --af
	Env: NEXUS_URL_GET_MOD_FORMAT
	YAML: nexus_url_get_mod_format
	```

	Url format to be used when querying for mod data. Defaults to:

	`https://api.nexusmods.com/v1/games/%v/mods/%v.json`

- **Mod page url format** 

	```yaml
	Flag: --page-url-format, --pf
	Env: NEXUS_URL_MOD_PAGE_FORMAT
	YAML: nexus_url_mod_page_format
	```

	Mod page Url format to be used in markdown header generation. Defaults to:

	`https://nexusmods.com/%v/mods/%v`

#### Generation

- **Markdown header format** 

	```yaml
	Flag: --header-format, --hf
	Env: MARKDOWN_HEADER_FORMAT
	YAML: markdown_header_format
	```

	Prompt format used to generate header of the markdown. Defaults to [header.txt](./header.txt)

	Header can be used to alter the resulting header of the markdown, in particular, in how the frontmatter is generated.

#### Helper

- **Fallback game code**

	```yaml
	Flag: --game-code, --gc
	Env: FALLBACK_GAME_CODE
	YAML: fallback_game_code
	```

	Game code to be used when not supplied in generation command.

	The command will fail when no game code found both in args or config.

### Generate markdown

The most basic command to generate the markdown is:

```sh
markxus generate [$game_code] $mod_id
```

By default, it will output the resulting markdown to current working directory with format: `{sanitized_mod_name}.md`. It will also prompt permission to overwrite when mod already exist.

If `$game_code` is not supplied, it will fallback to config & env vars. If none are available, the program will exit.

Aside from configuration flags, these flags are also available to alter the command's behaviour:
- `--force` or `FORCE_OVERWRITE`: Force overwrite conflicting file, if exist, skipping prompt.

### Generate multiple markdowns

TBA

## Programmatical Usage

Simplest example of usage:

```go
package main

import (
	"context"
	"fmt"

	"github.com/slainless/markxus"
	"github.com/slainless/markxus/genai"
	"github.com/slainless/markxus/nexus"
	"github.com/slainless/markxus/resty"
)

func main() {
	nexusKey := nexus.WithApiKey("{{NEXUS_API_KEY}}")
	nexus, err := nexus.NewClient(nexusKey, nexus.WithHTTPDriver(resty.NewRestyClient()))
	if err != nil {
		panic(err)
	}

	genaiKey := genai.WithApiKey("{{GENERATIVE_AI_API_KEY}}")
	genai, err := genai.NewGenAiClient(context.TODO(), genaiKey)
	if err != nil {
		panic(err)
	}

	markxus := markxus.NewMarkxus(nexus, genai)
	mod, err := markxus.Generate(context.TODO(), "skyrimspecialedition", "68068")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", mod)
}
```

To run example above, you will need nexusmods API key and Google generative AI API key.