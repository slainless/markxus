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
Obsidian is prominently known for features and its rich community support. Obsidian by default support showing and editing frontmatter. 
Obsidian community also has many plugins that can handle these data.

I planned to support these plugins:
  - [Banners](https://github.com/noatpad/obsidian-banners)
  - [Iconize](https://github.com/FlorianWoelki/obsidian-iconize)

More support to be added later.

## CLI Usage

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