package database

import (
	"context"
	"fmt"
	"log"

	setting "github.com/circleyu/GameServer/setting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v@%v:%v", setting.DBUser(), setting.DBPassword(), setting.DBServer(), setting.DBPort()))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database("demo")
}

func InsertOne(data interface{}) (*mongo.InsertOneResult, error) {
	userColl := db.Collection("user")
	return userColl.InsertOne(context.Background(), data)
}

func InsertMany(datas []interface{}) (*mongo.InsertManyResult, error) {
	userColl := db.Collection("user")
	return userColl.InsertMany(context.Background(), datas)
}

// 回傳之後自己去處理Decode
func FindOne(condition bson.M) *mongo.SingleResult {
	userColl := db.Collection("user")
	return userColl.FindOne(context.Background(), condition)
}

// 回傳之後自己去處理Decode
func Find(condition bson.M) (*mongo.Cursor, error) {
	userColl := db.Collection("user")
	return userColl.Find(context.Background(), condition)
	/*
		if cur, err := userColl.Find(ctx, condition); err == nil {
			defer cur.Close(ctx)
			for cur.Next(ctx) {
				var user interface{}
				if err := cur.Decode(&user); err != nil {
					log.Fatal(err)
				}
				log.Println(user)
			}
		} else {
			log.Fatal(err)
		}
	*/
}

func Disconnect() {
	//defer cancel()
	client.Disconnect(context.Background())
}
