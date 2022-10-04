package logs

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func New() log.Logger {
	return log.Logger{
		Handler: cli.New(os.Stderr),
		Level:   log.DebugLevel,
	}
}
