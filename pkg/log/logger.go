package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/everettraven/logmd/pkg/renderer"
)

type Logger interface {
	Info(message string)
	Debug(message string)
	Warn(message string)
	Error(message string)
	Print(message string)
}

// MarkdownLogger for logging things in markdown
type MarkdownLogger struct {
	condensed     bool
	infoRenderer  renderer.Renderer
	debugRenderer renderer.Renderer
	warnRenderer  renderer.Renderer
	errorRenderer renderer.Renderer
	printRenderer renderer.Renderer
}

type MarkdownLoggerOptions func(ml *MarkdownLogger)

func WithInfoRenderer(r renderer.Renderer) MarkdownLoggerOptions {
	return func(ml *MarkdownLogger) {
		ml.infoRenderer = r
	}
}

func WithDebugRenderer(r renderer.Renderer) MarkdownLoggerOptions {
	return func(ml *MarkdownLogger) {
		ml.debugRenderer = r
	}
}

func WithWarnRenderer(r renderer.Renderer) MarkdownLoggerOptions {
	return func(ml *MarkdownLogger) {
		ml.warnRenderer = r
	}
}

func WithErrorRenderer(r renderer.Renderer) MarkdownLoggerOptions {
	return func(ml *MarkdownLogger) {
		ml.errorRenderer = r
	}
}

func WithPrintRenderer(r renderer.Renderer) MarkdownLoggerOptions {
	return func(ml *MarkdownLogger) {
		ml.printRenderer = r
	}
}

func WithCondensedLogging() MarkdownLoggerOptions {
	return func(ml *MarkdownLogger) {
		ml.condensed = true
	}
}

func NewMarkdownLogger(opts ...MarkdownLoggerOptions) *MarkdownLogger {
	ml := &MarkdownLogger{
		condensed: false,
		infoRenderer: renderer.NewBasicRenderer(
			renderer.WithPrimaryColor("45"),
			renderer.WithSecondaryColor("231"),
		),
		debugRenderer: renderer.NewBasicRenderer(
			renderer.WithPrimaryColor("9"),
			renderer.WithSecondaryColor("231"),
		),
		warnRenderer: renderer.NewBasicRenderer(
			renderer.WithPrimaryColor("226"),
			renderer.WithSecondaryColor("232"),
		),
		errorRenderer: renderer.NewBasicRenderer(
			renderer.WithPrimaryColor("196"),
			renderer.WithSecondaryColor("231"),
		),
		printRenderer: renderer.NewBasicRenderer(),
	}

	for _, opt := range opts {
		opt(ml)
	}

	return ml
}

func (ml *MarkdownLogger) Info(message string) {
	var output string
	caller := ""
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller += details.Name()
	}
	heading := `# INFO 🛈` + "\n"
	timestamp := time.Now()
	output = fmt.Sprintf("%s---\n**TIMESTAMP**: %s\n\n**CALLER**: %s\n\n**MESSAGE**: %s\n\n---\n", heading, timestamp, caller, message)
	if ml.condensed {
		output = fmt.Sprintf("%s **%s** -- **%s** -- %s", heading, timestamp, caller, message)
	}

	out := ml.infoRenderer.Render(output, ml.condensed)

	if ml.condensed {
		out = condenseOutput(out)
	}

	fmt.Println(out)
}

func (ml *MarkdownLogger) Debug(message string) {
	var output string
	caller := ""
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller += details.Name()
	}
	heading := `# DEBUG 🕷` + "\n"
	timestamp := time.Now()
	output = fmt.Sprintf("%s---\n**TIMESTAMP**: %s\n\n**CALLER**: %s\n\n**MESSAGE**: %s\n\n---\n", heading, timestamp, caller, message)
	if ml.condensed {
		output = fmt.Sprintf("%s **%s** -- **%s** -- %s", heading, timestamp, caller, message)
	}

	out := ml.debugRenderer.Render(output, ml.condensed)

	if ml.condensed {
		out = condenseOutput(out)
	}

	fmt.Println(out)
}

func (ml *MarkdownLogger) Warn(message string) {
	var output string
	caller := ""
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller += details.Name()
	}
	heading := `# WARNING ⚠` + "\n"
	timestamp := time.Now()
	output = fmt.Sprintf("%s---\n**TIMESTAMP**: %s\n\n**CALLER**: %s\n\n**MESSAGE**: %s\n\n---\n", heading, timestamp, caller, message)
	if ml.condensed {
		output = fmt.Sprintf("%s **%s** -- **%s** -- %s", heading, timestamp, caller, message)
	}

	out := ml.warnRenderer.Render(output, ml.condensed)

	if ml.condensed {
		out = condenseOutput(out)
	}

	fmt.Println(out)
}

func (ml *MarkdownLogger) Error(message string) {
	var output string
	caller := ""
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller += details.Name()
	}
	heading := `# ERROR ✖` + "\n"
	timestamp := time.Now()
	output = fmt.Sprintf("%s---\n**TIMESTAMP**: %s\n\n**CALLER**: %s\n\n**MESSAGE**: %s\n\n---\n", heading, timestamp, caller, message)
	if ml.condensed {
		output = fmt.Sprintf("%s **%s** -- **%s** -- %s", heading, timestamp, caller, message)
	}

	out := ml.errorRenderer.Render(output, ml.condensed)

	if ml.condensed {
		out = condenseOutput(out)
	}

	fmt.Println(out)
}

func (ml *MarkdownLogger) Print(message string) {
	out := ml.printRenderer.Render(message, ml.condensed)

	fmt.Println(out)
}

func condenseOutput(output string) string {
	var condensedString string

	outParts := strings.Split(output, "\n")
	for _, outPart := range outParts {
		condensedString += strings.TrimSpace(outPart) + " "
	}

	return condensedString + "\n"
}
