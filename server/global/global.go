package global

import (
	"github.com/casbin/casbin/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	Env      *EnvConfig
	Enforcer *casbin.Enforcer
	Client   *mongo.Client
	DB       *mongo.Database
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
