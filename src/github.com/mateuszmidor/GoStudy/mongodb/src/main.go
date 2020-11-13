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
	databaseCoffeeShop       = "coffee-shop"
	databaseGrocery          = "grocery"
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

func ListDatabases() {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	dbs, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		log.Fatalf("Error listing databases: %v", err)
	}

	fmt.Println(dbs)
}

//Create creating a item in a mongo or document db
func Create(item *item, database string) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	item.ID = primitive.NewObjectID()

	result, err := client.Database(database).Collection("items").InsertOne(ctx, item)
	if err != nil {
		log.Fatalf("Could not create item: %v", err)
	}
	oid := result.InsertedID.(primitive.ObjectID)

	return oid, nil
}

func List(database string) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	findAll := bson.D{}
	cursor, err := client.Database(database).Collection("items").Find(ctx, findAll)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Database", database, ":")
	var t item
	for cursor.Next(ctx) {
		cursor.Decode(&t)
		fmt.Printf("%-10s%2d PLN\n", t.Title, t.Price)
	}
}

func main() {
	fmt.Println("Listing existing databases:")
	ListDatabases()
	fmt.Println("Done.")

	fmt.Println("\nAdding documents to db...")
	Create(&item{Title: "Coffee", Price: 9}, databaseCoffeeShop)
	Create(&item{Title: "Cake", Price: 12}, databaseCoffeeShop)
	Create(&item{Title: "Icecream", Price: 5}, databaseCoffeeShop)
	Create(&item{Title: "Bread", Price: 4}, databaseGrocery)
	Create(&item{Title: "Milk", Price: 3}, databaseGrocery)
	Create(&item{Title: "Guacamole", Price: 23}, databaseGrocery)
	fmt.Println("Done.")

	fmt.Println("\nListing documents in db:")
	List(databaseCoffeeShop)
	fmt.Println()
	List(databaseGrocery)
	fmt.Println("Done.")

	fmt.Println("\nListing existing databases:")
	ListDatabases()
	fmt.Println("Done.")
}
