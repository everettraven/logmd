package renderer

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

// Renderer is an interface that is used to render output for a logger
type Renderer interface {
	// Render takes in a string to render and output to the console
	Render(output string, condensed bool) string
}

// BasicRenderer is an implementation of the Renderer interface
type BasicRenderer struct {
	primaryColor    string
	secondaryColor  string
	style           ansi.StyleConfig
	useDefaultStyle bool
}

type BasicRendererOptions func(br *BasicRenderer)

func WithStyle(style ansi.StyleConfig) BasicRendererOptions {
	return func(br *BasicRenderer) {
		br.style = style
		br.useDefaultStyle = false
	}
}
func WithPrimaryColor(color string) BasicRendererOptions {
	return func(br *BasicRenderer) {
		br.primaryColor = color
	}
}

func WithSecondaryColor(color string) BasicRendererOptions {
	return func(br *BasicRenderer) {
		br.secondaryColor = color
	}
}

func NewBasicRenderer(opts ...BasicRendererOptions) *BasicRenderer {
	br := &BasicRenderer{
		primaryColor:    "45",
		secondaryColor:  "231",
		useDefaultStyle: true,
	}

	for _, opt := range opts {
		opt(br)
	}

	return br
}

func (br *BasicRenderer) Render(output string, condensed bool) string {
	style := getDefaultStyle(br.primaryColor, br.secondaryColor, condensed)

	if !br.useDefaultStyle {
		style = br.style
	}

	termRender, _ := glamour.NewTermRenderer(
		glamour.WithStyles(style),
	)

	out, _ := termRender.Render(output)

	return out
}

func getDefaultStyle(primaryColor string, secondaryColor string, condensed bool) ansi.StyleConfig {
	bold := true
	headingSuffix := ""
	italic := true
	codeMargin := uint(2)
	codeBg := "236"

	if !condensed {
		headingSuffix = "\n"
	}

	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Suffix: "\n",
			},
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:       &secondaryColor,
				Bold:        &bold,
				BlockSuffix: headingSuffix,
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BackgroundColor: &primaryColor,
				Prefix:          " ",
				Suffix:          " ",
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "## ",
				Bold:   &bold,
			},
		},
		Strong: ansi.StylePrimitive{
			Bold: &bold,
		},
		Emph: ansi.StylePrimitive{
			Italic: &italic,
		},
		HorizontalRule: ansi.StylePrimitive{
			Format: "------------------------------",
		},
		CodeBlock: ansi.StyleCodeBlock{
			StyleBlock: ansi.StyleBlock{
				Margin: &codeMargin,
			},
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BackgroundColor: &codeBg,
				Prefix:          " ",
				Suffix:          " ",
			},
		},
		List: glamour.DarkStyleConfig.List,
		Item: glamour.DarkStyleConfig.Item,
	}
}
