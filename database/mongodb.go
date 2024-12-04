package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database ...
type Database interface {
	vacationRepository
}

// MongoDB ...
type MongoDB struct{}

var _ Database = (*MongoDB)(nil)

var db *MongoDB
var once sync.Once

// GetDatabase ...
func GetDatabase() Database {
	once.Do(func() {
		db = &MongoDB{}
	})
	return db
}

// MongodbConnect ...
func MongodbConnect() *mongo.Collection {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("vacationsApi").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your cluster. You successfully connected to MongoDB!")

	collection := client.Database("vacationsApi").Collection("vacations")
	return collection

}
