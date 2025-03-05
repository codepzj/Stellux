package global

import (
	"context"
	"fmt"
	. "server/internal/pkg/logger"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	mongooptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
)

func init() {
	Env = NewEnv()
	// 初始化全局变量
	Enforcer = NewEnforcer(Env.URL, Env.MongoDatabase)
	Client = NewMongoClient()
	DB = NewMongoDB(Client)
}

func NewEnv() *EnvConfig {
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
		Logger.Error("找不到环境配置文件，将使用默认配置", zap.Error(err))
	}

	// 加载配置
	return &EnvConfig{
		URL:               viper.GetString("mongodb.URL"),
		MongoDatabase:     viper.GetString("mongodb.MONGO_INITDB_DATABASE"),
		MongoRootUsername: viper.GetString("mongodb.MONGO_INITDB_ROOT_USERNAME"),
		MongoRootPassword: viper.GetString("mongodb.MONGO_INITDB_ROOT_PASSWORD"),
		MongoUserName:     viper.GetString("mongodb.MONGO_USERNAME"),
		MongoPassword:     viper.GetString("mongodb.MONGO_PASSWORD"),
		Port:              viper.GetString("server.PORT"),
	}
}

func NewEnforcer(url string, dbName string) *casbin.Enforcer {
	mongoClientOption := mongooptions.Client().ApplyURI(url)
	a, err := mongodbadapter.NewAdapterWithClientOption(mongoClientOption, dbName)
	if err != nil {
		panic(err)
	}
	enforcer, err := casbin.NewEnforcer("config/policy.conf", a)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		panic(fmt.Sprintf("权限加载失败：%s", err.Error()))
	}
	return enforcer
}

func NewMongoClient() *mongo.Client {
	MongoClient, err := mongo.Connect(options.Client().ApplyURI(Env.URL))
	if err != nil {
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	return MongoClient
}

func NewMongoDB(MongoClient *mongo.Client) *mongo.Database {
	if err := MongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "数据库无法ping通"))
	}
	return MongoClient.Database(Env.MongoDatabase)
}
