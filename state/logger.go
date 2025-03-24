package state

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func InitializeLogger() {
	Log = logrus.New()

	Log.SetLevel(logrus.DebugLevel)

	Log.SetOutput(os.Stdout)

	{
		// at last print out some log to check if logger works
		Log.Info("Logger initialized!")
	}
}