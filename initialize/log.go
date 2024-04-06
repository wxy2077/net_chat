package initialize

import (
	"github.com/sirupsen/logrus"
	"net-chat/global"
	"net-chat/pkg"
)

func InitLogger(logLevel string) *logrus.Logger {

	return pkg.NewLog(logLevel, global.DefaultLogFormatterJSON, global.DefaultLogPath)
}
