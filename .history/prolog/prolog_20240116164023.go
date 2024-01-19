package prolog

import (
	"os"

	"github.com/mbndr/logo"
)

func newLog() *logo.Logger {
	return logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "prefix ", true)
}
