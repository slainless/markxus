package markxus

import (
	_ "embed"
	"text/template"
)

type MarkxusOptions struct {
	UrlModPageFormat       string
	MarkdownHeaderTemplate *template.Template
	LlmPromptTemplate      *template.Template
}

//go:embed prompt.txt
var DefaultLlmPromptFormat string
var DefaultLlmPromptTemplate *template.Template

//go:embed header.txt
var DefaultMarkdownHeaderFormat string
var DefaultMarkdownHeaderTemplate *template.Template

func init() {
	template, err := template.New("markxus.header").Parse(DefaultMarkdownHeaderFormat)
	if err != nil {
		panic(err)
	}

	DefaultMarkdownHeaderTemplate = template

	template, err = template.New("markxus.prompt").Parse(DefaultLlmPromptFormat)
	if err != nil {
		panic(err)
	}

	DefaultLlmPromptTemplate = template
}

const DefaultLlmModelName = "gemini-1.5-flash"
const DefaultUrlModPageFormat = "https://nexusmods.com/%v/mods/%v"

type MarkxusOption func(*MarkxusOptions)

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
func WithPromptTemplate(template *template.Template) MarkxusOption {
	return func(mo *MarkxusOptions) {
		mo.LlmPromptTemplate = template
	}
}

// Template will be exposed to [[nexus.SchemaMod]]
func WithMarkdownHeaderTemplate(format *template.Template) MarkxusOption {
	return func(mo *MarkxusOptions) {
		mo.MarkdownHeaderTemplate = format
	}
}
