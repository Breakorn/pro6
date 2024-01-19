package prolog

import (
	"os"

	"github.com/mbndr/logo"
)

func newLog() {
	return logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "prefix ", true)
}
