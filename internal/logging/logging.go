package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type log struct {
	*logrus.Logger
}

var Log *log

func InitLog() error {
	Log = &log{
		logrus.New(),
	}

	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "DEBUG":
		Log.SetLevel(logrus.DebugLevel)
	case "INFO":
		Log.SetLevel(logrus.InfoLevel)
	case "ERROR":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	Log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2023-06-20 14:30:00",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})

	// logFile, err := os.OpenFile("backend.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	return err
	// }
	// // defer logFile.Close()
	// // mw := io.MultiWriter(os.Stdout, logFile)
	// // Log.SetOutput(mw)
	return nil
}
