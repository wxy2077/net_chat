package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net-chat/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
