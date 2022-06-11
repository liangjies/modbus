// initialize
package initialize

import (
	"context"
	"modbus-spyder/internal/app/global"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func MongoDB() *mongo.Database {
	// 获取配置文件
	m := global.SYS_CONFIG.MongoDB
	dsn := "mongodb://" + m.Path
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(dsn)
	// 设置最大连接池
	clientOptions.SetMaxPoolSize(m.MaxPoolSize)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		global.SYS_LOG.Error("MongoDB连接异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	}
	// 判断服务是不是可用
	if err = client.Ping(context.TODO(), nil); err != nil {
		global.SYS_LOG.Error("MongoDB服务不可用", zap.Any("err", err))
		os.Exit(0)
		return nil
	}
	return client.Database(m.Dbname)
}
