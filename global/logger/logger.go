package logger

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var l *logrus.Logger

func init() {
	once := sync.Once{}

	once.Do(func() {
		logrus.SetLevel(logrus.DebugLevel)
		l = logrus.New()
	})
}

func Logger() *logrus.Logger {
	return l
}