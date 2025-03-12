package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database interface composed of the VacationRepository interface,
// meaning if any type definition implements VacationRepository, it will satisfy the Database interface as well
// This abstraction allows us to define database operations without tying it to a specific implementation like db * of type MongoDB,
// so that i can also use a mocked database for testing
type Database interface {
	VacationRepository
}

// MongoDB is a struct that will implement all the functions defined in VacationRepository interface
type MongoDB struct{}

// compile-time check to make sure MongoDB implements the Database interface (pulls in VacationRepository functions)
var _ Database = (*MongoDB)(nil)

var db *MongoDB
var once sync.Once

// GetDatabase thread safe implementation of MongoDB so that only one instance is initialized
func GetDatabase() Database {
	once.Do(func() {
		db = &MongoDB{}
	})
	return db
}

var onceDbClient sync.Once
var dbClient *mongo.Collection

// GetNoCloseClientCollection Creates a thread safe connection and returns the collection specified
func GetNoCloseClientCollection(db string, col string) (*mongo.Collection, error) {
	var err error
	onceDbClient.Do(func() {
		dbClient, err = mongodbConnect(db, col)
	})
	return dbClient, err
}

// MongodbConnect connects to mongodb usign the specified database and collection parameters and returns the collection object of the client
func mongodbConnect(db string, col string) (*mongo.Collection, error) {

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
		return nil, err
	}

	collection := client.Database(db).Collection(col)
	return collection, nil

}
