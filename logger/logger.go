package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()

	return u.Formatter.Format(e)
}

var (
	Logger        = logrus.New()
	HTTPLogFormat = fmt.Sprintf("%s\n", `{"time": ${time_unix_nano}, }`) // Annoying \n hack
)

func ConfigureLogger() {
	if os.Getenv("DEBUG") == "1" {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.InfoLevel)
	}
	Logger.SetFormatter(UTCFormatter{&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	}})
	// Logger.SetReportCaller(true) // this shows the below functions since it always executes from there
	// https://stackoverflow.com/questions/63658002/is-it-possible-to-wrap-logrus-logger-functions-without-losing-the-line-number-pr
}

func Debug(format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, args...))
}
func Info(format string, args ...interface{}) {
	Logger.Info(fmt.Sprintf(format, args...))
}
func Warn(format string, args ...interface{}) {
	Logger.Warn(fmt.Sprintf(format, args...))
}
func Error(format string, args ...interface{}) {
	Logger.Error(fmt.Sprintf(format, args...))
}
