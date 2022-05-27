package main

import (
	"github.com/charmbracelet/glamour"
	"github.com/everettraven/logmd/pkg/log"
	"github.com/everettraven/logmd/pkg/renderer"
)

func main() {

	condensedLogger := log.NewMarkdownLogger(log.WithCondensedLogging())

	// Condensed logs
	condensedLogger.Info("*hello condensed info world!*")
	condensedLogger.Debug("*hello condensed debug world!*")
	condensedLogger.Warn("*hello condensed warn world!*")
	condensedLogger.Error("*hello condensed error world!*")

	expandedLogger := log.NewMarkdownLogger()
	// Expanded logs
	expandedLogger.Info("*hello expanded info world!*")
	expandedLogger.Debug("*hello expanded debug world!*")
	expandedLogger.Warn("*hello expanded warn world!*")
	expandedLogger.Error("*hello expanded error world!*")

	draculaRenderer := renderer.NewBasicRenderer(renderer.WithStyle(glamour.DraculaStyleConfig))

	customLogger := log.NewMarkdownLogger(
		log.WithPrintRenderer(draculaRenderer),
	)

	// Custom log print -- example showing a command line help example
	output := `# logmd
logmd is your logging doctor. It will make sure your logs are in tip top shape!

It can:

- Extract error messages from log files

- Sanitize log messages in log files

## Usage

~~~
$ logmd [SUBCOMMAND] [OPTIONS]
~~~

## Global CLI Options
**-h, --help**	*show help messages for subcommands or the parent logmd command*

**-c, --config**	*set the JSON configuration for logmd*

## Subcommands
- ` + "`diagnose`" + ` - extract all errors from your logs so that you can diagnose them easier. Example:
~~~
$ logmd diagnose logs.txt
~~~
- ` + "`sterilize`" + ` - clean up the logs and make them easier to read. Example:
~~~
$ logmd sterilize logs.txt
~~~
`
	customLogger.Print(output)

}
