package mongo

import (
	"context"
	"github.com/mokaz111/candy-server/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	DBClient *mongo.Client
	err      error
	MGDB     *mongo.Database
)

func Init() {
	uri := conf.GetConf().Mongo.URI
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	DBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = DBClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	dbname := conf.GetConf().Mongo.DBName
	if dbname == "" {
		log.Fatal("dbname is empty")
	}
	MGDB = DBClient.Database(dbname)

	log.Println("Connected to MongoDB!")

}
