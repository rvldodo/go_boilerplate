package log

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type customFieldsHook struct{}

func init() {
	logrus.SetReportCaller(true)
	formater := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", ""
		},
	}
	logrus.SetFormatter(formater)
	logrus.AddHook(&customFieldsHook{})
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func (c *customFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (c *customFieldsHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = "boilerplate"

	if entry.HasCaller() {
		entry.Data["file"] = fmt.Sprintf(
			"%s:%d",
			formatFilePath(entry.Caller.File),
			entry.Caller.Line,
		)
	}

	return nil
}

func Info(value ...interface{}) {
	logrus.Info(value...)
}

func Infof(format string, value ...interface{}) {
	logrus.Infof(format, value...)
}

func Error(value ...interface{}) {
	logrus.Error(value...)
}

func Errorf(format string, value ...interface{}) {
	logrus.Errorf(format, value...)
}

func Warn(value ...interface{}) {
	logrus.Warn(value...)
}

func Warnf(format string, value ...interface{}) {
	logrus.Warnf(format, value...)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"url":    r.URL.String(),
			"time":   time.Since(start),
		}).Info("Endpoint hit")
	})
}
