package logo

import (
	"os"

	"github.com/mbndr/logo"
)

var log *logo.Logger

func NewLog() {
	log = logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)
}

func INFO(str string) {
	log.Info(str)
}
func ERROR(str string) {
	log.Error(str)
}
func WARN(str string) {
	log.Warn(str)
}
