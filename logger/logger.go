package logger

import (
	"fmt"

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
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(UTCFormatter{&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	}})
	// Logger.SetReportCaller(true) // this shows the below functions since it always executes from there
}

func Debug(format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format+"\n", args...))
}
func Info(format string, args ...interface{}) {
	Logger.Info(fmt.Sprintf(format+"\n", args...))
}
func Warn(format string, args ...interface{}) {
	Logger.Warn(fmt.Sprintf(format+"\n", args...))
}
func Error(format string, args ...interface{}) {
	Logger.Error(fmt.Sprintf(format+"\n", args...))
}
