package database

import(
	"fmt"
	"log"
	"os"
	"time"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading the file")
	}

	MongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil{
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB!")
	return client
}


var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("").Collection(collectionName)
	return collection
}