package logo

import (
	"os"

	"github.com/mbndr/logo"
)

var log *logo.Logger

func NewLog() {
	log = logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
}

func INFO(str ...interface{}) {
	log.Info(str)
}
func ERROR(str ...interface{}) {
	log.Error(str)
}
func WARN(str ...interface{}) {
	log.Warn(str)
}
