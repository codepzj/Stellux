package global

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	. "server/internal/pkg/logger"
)

var (
	Env      *EnvConfig
	Enforcer *casbin.Enforcer
)

type EnvConfig struct {
	URL               string `json:"URL"`
	MongoDatabase     string `json:"MONGO_INITDB_DATABASE"`
	MongoRootUsername string `json:"MONGO_INITDB_ROOT_USERNAME"`
	MongoRootPassword string `json:"MONGO_INITDB_ROOT_PASSWORD"`
	MongoUserName     string `json:"MONGO_USERNAME"`
	MongoPassword     string `json:"MONGO_PASSWORD"`
	Port              string `json:"PORT"`
}

func init() {
	viper.SetDefault("mongodb.URL", "mongodb://admin:123456@localhost:27017/?authSource=admin")
	viper.SetDefault("mongodb.MONGO_INITDB_DATABASE", "stellux")
	viper.SetDefault("mongodb.MONGO_INITDB_ROOT_USERNAME", "admin")
	viper.SetDefault("mongodb.MONGO_INITDB_ROOT_PASSWORD", "123456")
	viper.SetDefault("mongodb.MONGO_USERNAME", "codepzj")
	viper.SetDefault("mongodb.MONGO_PASSWORD", "123456")
	viper.SetDefault("server.PORT", "9001")

	viper.SetConfigFile("./config/stellux.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Error("找不到环境配置文件", zap.Error(err))
		Logger.Info("将使用默认配置")
	}

	// 加载配置
	Env = &EnvConfig{
		URL:               viper.GetString("mongodb.URL"),
		MongoDatabase:     viper.GetString("mongodb.MONGO_INITDB_DATABASE"),
		MongoRootUsername: viper.GetString("mongodb.MONGO_INITDB_ROOT_USERNAME"),
		MongoRootPassword: viper.GetString("mongodb.MONGO_INITDB_ROOT_PASSWORD"),
		MongoUserName:     viper.GetString("mongodb.MONGO_USERNAME"),
		MongoPassword:     viper.GetString("mongodb.MONGO_PASSWORD"),
		Port:              viper.GetString("server.PORT"),
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件发生改变:", e.Name)
	})

	mongoClientOption := mongooptions.Client().ApplyURI(Env.URL)
	databaseName := Env.MongoDatabase
	a, err := mongodbadapter.NewAdapterWithClientOption(mongoClientOption, databaseName)
	if err != nil {
		panic(err)
	}
	Enforcer, err = casbin.NewEnforcer("config/policy.conf", a)
	if err != nil {
		panic(err)
	}
	err = Enforcer.LoadPolicy()
	if err != nil {
		Logger.Error(fmt.Sprintf("权限加载失败：%s", err.Error()))
	}
}
