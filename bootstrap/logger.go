package bootstrap

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func InitializeLogger() {
	Log = logrus.New()

	Log.SetLevel(logrus.DebugLevel)

	Log.SetOutput(os.Stdout)
}