package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"xlblog/config"
)

var (
	Db     *gorm.DB
	Redis  *redis.Client
	Config config.Config
	Viper  *viper.Viper
	Logger *zap.Logger
)
