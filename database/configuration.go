package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/HasnathJami/go-myWallet-mux-mongoDb/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Collection {
    var collection *mongo.Collection

    err := godotenv.Load(".env")
	utils.CheckSimpleError(err)
    
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	utils.CheckSimpleError(err)
	fmt.Println("Connected to MongoDB")
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instance created")


	return collection
}