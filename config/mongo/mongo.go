package mongo

import (
	"context"
	"fmt"
	"github.com/wegoteam/wepkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const (
	MONGO = "mongo"
)

var (
	MongoClient *mongo.Client
	once        sync.Once
)

func init() {
	once.Do(func() {
		initMongoDBConfig()
	})
}

// initMongoDBConfig
// @Description: 初始化MongoDB配置
//mongodb://user:password@localhost:27017/?authSource=admin
func initMongoDBConfig() {
	var mongoConfig = &config.Mongo{}
	c := config.GetConfig()
	c.Load(MONGO, mongoConfig)
	url := fmt.Sprintf("mongodb://%s", mongoConfig.Address)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Errorf("MongoDB connect failed: %v", err)
	}
	MongoClient = client
}
