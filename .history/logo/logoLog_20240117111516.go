package logo

import (
	"os"

	"github.com/mbndr/logo"
)

// var log
func NewLog() {

	log := logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
}
