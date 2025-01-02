# Markxus

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/slainless/markxus)
![Nexus Mods](https://img.shields.io/endpoint?url=https%3A%2F%2Fgist.githubusercontent.com%2Fradj307%2Fe9a80731ee236cc67fb00b698e75201e%2Fraw%2F5230074dfb1a60fba917a1232f9382fa5cfec5db%2Fendpoint.json&style=for-the-badge) ![Markdown](https://img.shields.io/badge/markdown-%23000000.svg?style=for-the-badge&logo=markdown&logoColor=white) ![Obsidian](https://img.shields.io/badge/Obsidian-%23483699.svg?style=for-the-badge&logo=obsidian&logoColor=white)

![Google Gemini](https://img.shields.io/badge/google%20gemini-8E75B2?style=for-the-badge&logo=google%20gemini&logoColor=white)  ![OpenAI Badge](https://img.shields.io/badge/OpenAI-412991?logo=openai&logoColor=fff&style=for-the-badge)


<img src="https://github.com/user-attachments/assets/738471dd-7a6f-4d88-879f-aab5a591486b" alt="usage-demo"/>
<p>
	<sub>The GIF is fast-forwarded</sub>
</p>

> [!TIP]
> Use a suitable model to generate more accurate result. Example GIF above used `gemini-1.5-pro`.
>
> Using flash model (default option) will result in inaccurate mod referencing result but generate faster. Instead, pro model
> generate a very accurate result with tradeoff being slower generation. OpenAI's GPT-4o-mini is also pretty capable
> with this particular task while generating much faster than Google's gemini-1.5-pro.
>
> Recommended: Use `gemini-1.5-pro` or `gpt-4o-mini` (but `gpt-4o` is highly recommended)

Markxus is an LLM-powered markdown converter targeted for [Nexus Mods](https://www.nexusmods.com/)'s mod page.
By default, Markxus is opinionated and it converts nexusmods HTML + [BBCode](https://en.wikipedia.org/wiki/BBCode) 
into markdown that is catered for Obsidian markdown format.
The `cli` package will use [resty](https://github.com/go-resty/resty) as its HTTP client 
and [google-generative-ai](https://github.com/google/generative-ai-go) as its LLM client by default.

This package provides support for: 
- [Resty](./resty) as Nexusmods HTTP client
- [Google Generative AI](./genai/) as default LLM provider
- [OpenAI GPT](./openai/) as alternative LLM provider

The main package by design is client-agnostic so extending support to other combination of HTTP and LLM client
is pretty simple.

> [!NOTE]
> Switch to Open AI if `FinishReasonRecitation` occurs when using generative AI model.
> This is a known issue, thanks to Google's usual antic:
> 
> https://github.com/google/generative-ai-docs/issues/257

## Obsidian markdown

![obsidian](https://github.com/user-attachments/assets/6ce507bc-2e25-4143-8d02-e378340b13a0)

The default markdown header format contains frontmatter and only a handful number of editors are capable of making use of these data.
Obsidian is prominently known its for features and its rich community support. Obsidian by default support showing and editing frontmatter. 
Obsidian community also has many plugins that can handle these data.

Supported plugins:
  - [🚩 Pixel Banner for Obsidian](https://github.com/jparkerweb/pixel-banner)
  - [Obsidian Iconize](https://github.com/FlorianWoelki/obsidian-iconize)

More support to be added later.

## Background

Recently, I tried to seriously rebuild my Skyrim mod list for the latest version. Obsidian is very handy in organizing pile ton of
mods information, and with its graph view, it helps a lot in quickly finding relevant mods, compatibility with each other, etc.
However, I dont have that much time to index every mods I use to Obsidian, so here we are.

![graph-view](https://github.com/user-attachments/assets/d53aa089-c3aa-4482-b105-3043f5be858a)
<p>
	<sub>All the markdowns are AI-generated. Now its easier to see relationship between each mods.</sub>
</p>

Now, LLM is needed in this case since I need to annotate all cross-mod references with `[[Obsidian internal linking]]` and LLM is intelligent 
enough to guess whether some tokens are actually a mention or title of another mod. 
Nexusmods page conversion to markdown is also not that straightforward since it contains mix of HTML and BBCode,
so relying on LLM to intuit the correct tag usage is the best choice here.

For now, this package can only support Google generative AI and Open AI since that is my only choice at the moment.

## CLI Usage

### Installation

Go is needed to install the CLI. I will make an executable build pipeline later if I feel like it.

```sh
go install github.com/slainless/markxus/cli/markxus@latest
```

### Configuration

Most of the functionality of the apps relies on API key of Nexus Mods and Google Generative AI so 
it is important to set up global configuration (or local configuration) or both vars will need to be supplied to the binary per usage.

To initialize global (or local) configuration, run:

```sh
markxus config init
```

It will create `.markxus.yaml` to user directory (or cwd, depending on config type). It will also prompt you with both API key and store it to OS keyring.
It is recommended to store the API key in OS-managed credential storage rather
than disk.

Edit command can also be used to initialize the file then directly open file editor against the config, by running:

```sh
markxus config edit -i
```

To change API key later without going through config init, run:

```sh
markxus config set NEXUS_API_KEY <new_key>
```

or

```sh
markxus config set LLM_API_KEY <new_key>
```

Command flags is available to `config init` to alter its behaviour:
- `--force` or `FORCE_OVERWRITE`: Force overwrite to file if exist
- `--type` or `CONFIG_TYPE`: Set config type, whether to generate to `global` or `local` (cwd)

### Changing Configuration

`config edit` can be used to open file editor for the config, using OS preferred text editor:

```sh
markxus config edit
```

Command flags is available to `config edit` to alter its behaviour:
- `--i` or `INIT_CONFIG`: Initialize config if not exist
- `--type` or `CONFIG_TYPE`: Set config type, whether to edit `global` or `local` config (cwd)

Individual config key can also be set using:

```sh
markxus config set <yaml_or_env_key> <value>
```

Must be noted, however, that `NEXUS_API_KEY` and `LLM_API_KEY` cannot be set to config via this command
and must be edited manually. Setting either field will configure OS credential storage instead of config 
(and obviously, `--type` will be ignored when setting either fields). 

Command flags is available to `config set` to alter its behaviour:
- `--type` or `CONFIG_TYPE`: Set config type, whether to set value to `global` or `local` config (cwd)


#### Available Options

All these options can also be set from env vars or CLI flag.

##### LLM

- **Provider**

	```yaml
	Flag: --llm-provider, --gp
	Env: LLM_PROVIDER
	YAML: llm_provider
	```

	LLM provider to be used, either `open_ai` or `gen_ai`. Defaults to `gen_ai`.

- **API key** 

	```yaml
	Flag: --llm-key, --gk
	Env: LLM_API_KEY
	YAML: llm_key
	```

	Self-explanatory. **Required**.

- **Model name**

	```yaml
	Flag: --model, --m
	Env: LLM_API_KEY
	YAML: llm_key
	```

	LLM model used for conversion. Defaults to `gemini-1.5-flash`.

- **Prompt format** 

	```yaml
	Flag: --prompt, --p
	Env: LLM_PROMPT_FORMAT
	YAML: llm_prompt_format
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

##### Generation

- **Markdown header format** 

	```yaml
	Flag: --header-format, --hf
	Env: MARKDOWN_HEADER_FORMAT
	YAML: markdown_header_format
	```

	Prompt format used to generate header of the markdown. Defaults to [header.txt](./header.txt)

	Header can be used to alter the resulting header of the markdown, in particular, in how the frontmatter is generated.

##### Helper

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

## Developer Notes

This is my first time building a CLI using [charmbracelet](https://github.com/charmbracelet)'s [bubbletea](https://github.com/charmbracelet/bubbletea)
and I admit that the framework is a bit overwhelming for user that didn't come from Elm background, since I'm too used to React's reactive paradigm.

Building the TUI is actually the most time consuming part of this project, largely caused by my own inexperience in handling these libraries.
I managed to make it work nonetheless 😁. But some parts of the CLI need some refactoring and improvement, specifically at some parts that still use
pure [huh](https://github.com/charmbracelet/huh) form. Ideally, it should be incorporated into bubbletea modeling framework and make it work seamlessly
with [urfave/cli](https://github.com/urfave/cli) flow.
