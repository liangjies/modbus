package global

import (
	"modbus-spyder/internal/app/config"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SYS_DB       *gorm.DB
	SYS_MONGODB  *mongo.Database
	SYS_CONFIG   config.Server
	SYS_VIP      *viper.Viper
	SYS_LOG      *zap.Logger
	CollectPoint map[string]string
)
