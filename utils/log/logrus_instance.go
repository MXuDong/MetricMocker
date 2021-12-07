package log

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func InitLogrus() {
	Logger = logrus.New()
}
