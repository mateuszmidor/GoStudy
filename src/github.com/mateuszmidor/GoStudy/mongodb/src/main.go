package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 10
	connectionStringTemplate = "mongodb://%s:%s@%s"
	username                 = "myuser"
	password                 = "mypass"
	clusterEndpoint          = "localhost"
)

type item struct {
	ID    primitive.ObjectID
	Title string
	Price uint
}

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	return client, ctx, cancel
}

//Create creating a item in a mongo or document db
func Create(item *item) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	item.ID = primitive.NewObjectID()

	result, err := client.Database("shopping").Collection("coffee-shop").InsertOne(ctx, item)
	if err != nil {
		log.Fatalf("Could not create item: %v", err)
		// return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)

	return oid, nil
}

func List() {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	findAll := bson.D{}
	cursor, err := client.Database("shopping").Collection("coffee-shop").Find(ctx, findAll)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var t item
	for cursor.Next(ctx) {
		cursor.Decode(&t)
		fmt.Printf("%-10s%2d PLN\n", t.Title, t.Price)
	}
}

func main() {
	Create(&item{Title: "Coffee", Price: 9})
	Create(&item{Title: "Cake", Price: 12})
	Create(&item{Title: "Icecream", Price: 5})

	List()
}
