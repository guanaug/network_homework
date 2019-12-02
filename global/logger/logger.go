package logger

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var l *logrus.Logger

func init() {
	once := sync.Once{}

	once.Do(func() {
		l = logrus.New()
		l.SetLevel(logrus.DebugLevel)
	})
}

func Logger() *logrus.Logger {
	return l
}
