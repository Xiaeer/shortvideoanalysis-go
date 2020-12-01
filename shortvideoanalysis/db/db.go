package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DbManager 全局的dbManager，方便controller那边操作数据库（增删改查）
var DbManager dbManager

func init() {
	// 初始化dbManager
	DbManager.InitDatabase()
}

// dbManager DB的管理器结构体
type dbManager struct {
	DBClient             *mongo.Client
	CommonDatabase       *mongo.Database
	ShortVideoCollection *mongo.Collection
}

// InitDatabase ...
func (manager *dbManager) InitDatabase() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions.SetMinPoolSize(10)
	clientOptions.SetMaxPoolSize(20)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	manager.DBClient = client
	manager.CommonDatabase = client.Database("common")
	fmt.Println("Connected to database: common.")
}

func (manager *dbManager) CloseDB() {
	fmt.Println("closing db client connection...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := manager.DBClient.Disconnect(ctx); err != nil {
		panic(err)
	}
	fmt.Println("closed db client connection")
}

func (manager *dbManager) SetShortVideoCollection(collection string) {
	manager.ShortVideoCollection = manager.CommonDatabase.Collection(collection)
	fmt.Println("Set Collection:ShortVideoCollection success!")
}

func (manager *dbManager) FindOneShortVideo(filter interface{}, result interface{}) {
	if err := manager.ShortVideoCollection.FindOne(context.Background(), filter).Decode(result); err != nil {
		fmt.Println("manager.ShortVideoCollection.FindOne.Decode error:", err.Error())
		return
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func (manager *dbManager) FindOne(collection string, filter interface{}, result interface{}) {
	if err := manager.CommonDatabase.Collection(collection).FindOne(context.Background(), filter).Decode(result); err != nil {
		panic(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func (manager *dbManager) InsertOneShortVideo(data interface{}) (interface{}, error) {
	return manager.ShortVideoCollection.InsertOne(context.Background(), data)
}

func (manager *dbManager) InsertOne(collection string, data interface{}) (interface{}, error) {
	return manager.CommonDatabase.Collection(collection).InsertOne(context.Background(), data)
}
