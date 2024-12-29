package markxus

import (
	_ "embed"
	"text/template"
)

type MarkxusOptions struct {
	GenAiPromptFormat      string
	UrlModPageFormat       string
	MarkdownHeaderTemplate *template.Template
}

//go:embed prompt.txt
var DefaultGenAiPromptFormat string

//go:embed header.txt
var DefaultMarkdownHeaderFormat string
var DefaultMarkdownHeaderTemplate *template.Template

func init() {
	template, err := template.New("markxus.header").Parse(DefaultMarkdownHeaderFormat)
	if err != nil {
		panic(err)
	}

	DefaultMarkdownHeaderTemplate = template
}

const DefaultGenAiModelName = "gemini-1.5-flash"
const DefaultUrlModPageFormat = "https://nexusmods.com/%v/mods/%v"

type MarkxusOption func(*MarkxusOptions)

// Format should contains placeholder that will be filled with
// these parameters in sequence:
//
//   - Mod description
//
// Defaults to [[DefaultGenAiPromptFormat]]
func WithPromptFormat(prompt string) MarkxusOption {
	return func(mo *MarkxusOptions) {
		mo.GenAiPromptFormat = prompt
	}
}

// Format should contains 2 placeholder in this sequence:
//   - Game code
//   - Mod ID
//
// Defaults to: [[DefaultUrlModPageFormat]]
func WithUrlModPageFormat(format string) MarkxusOption {
	return func(mo *MarkxusOptions) {
		mo.UrlModPageFormat = format
	}
}

// Template will be exposed to [[nexus.SchemaMod]]
func WithMarkdownHeaderTemplate(format *template.Template) MarkxusOption {
	return func(mo *MarkxusOptions) {
		mo.MarkdownHeaderTemplate = format
	}
}
