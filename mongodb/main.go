package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 10
	databaseConnectionString = "mongodb://myuser:mypass@localhost"
	databaseCoffeeShop       = "coffee-shop"
	databaseGrocery          = "grocery"
)

type item struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Price uint               `bson:"price"`
}

// getConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseConnectionString))
	panicOnErr("Failed to createItem client", err)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	panicOnErr("Failed to connect to cluster", err)

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	panicOnErr("Failed to ping cluster", err)

	return client, ctx, cancel
}

func listDatabases() {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	dbs, err := client.ListDatabaseNames(ctx, bson.D{})
	panicOnErr("Error listing databases", err)

	fmt.Println(dbs)
}

func create(database string, name string, price uint) primitive.ObjectID {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	item := &item{Name: name, Price: price}
	result, err := client.Database(database).Collection("items").InsertOne(ctx, &item)
	panicOnErr("Could not create item", err)
	fmt.Printf("created %s\n", name)

	return result.InsertedID.(primitive.ObjectID)
}

func listAllAtOnce(database string) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	findAll := bson.D{} // empty document means no filtering
	cursor, err := client.Database(database).Collection("items").Find(ctx, findAll)
	panicOnErr("Error finding documents", err)

	// get all documents into a slice
	var items []bson.M
	err = cursor.All(ctx, &items)
	panicOnErr("Error unmarshalling documents", err)

	// print documents
	for _, item := range items {
		fmt.Printf("%-10s%2d PLN\n", item["name"], item["price"])
	}
}

func listOneByOne(database string) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	findAll := bson.D{} // D for Document, that is defined as pairs: bson.D {{"name", "Andrzej"}, {"age", 33}}
	cursor, err := client.Database(database).Collection("items").Find(ctx, findAll)
	panicOnErr("Error finding documents", err)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var t item
		cursor.Decode(&t)
		fmt.Printf("%-10s%2d PLN\n", t.Name, t.Price)
	}
}

func listSorted(database string) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	findAll := bson.D{}
	sortDescending := options.Find().SetSort(bson.D{{"price", -1}})
	cursor, err := client.Database(database).Collection("items").Find(ctx, findAll, sortDescending)
	panicOnErr("Error finding documents", err)

	defer cursor.Close(ctx)
	var t item
	for cursor.Next(ctx) {
		cursor.Decode(&t)
		fmt.Printf("%-10s%2d PLN\n", t.Name, t.Price)
	}
}

func listCheaperThan7(database string) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// findCheapest := bson.D{{"price", 3}} // price == 3
	findCheap := bson.D{{"price", bson.D{{"$lt", 7}}}} // price < 7
	cursor, err := client.Database(database).Collection("items").Find(ctx, findCheap)
	panicOnErr("Error finding documents", err)

	defer cursor.Close(ctx)
	var t item
	for cursor.Next(ctx) {
		cursor.Decode(&t)
		fmt.Printf("%-10s%2d PLN\n", t.Name, t.Price)
	}
}

func updatePrice(database string, name string, newPrice uint) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	filterByName := bson.D{{"name", name}}
	updatePrice := bson.D{{"$set", bson.D{{"price", newPrice}}}}
	_, err := client.Database(database).Collection("items").UpdateOne(ctx, filterByName, updatePrice)
	panicOnErr("Error updating price", err)
}

// replace given item with completely new one but keep ObjectID "_id"
func replaceEntireItem(database string, oldName string, newName string, newPrice uint) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	filterByName := bson.D{{"name", oldName}}
	replacement := item{Name: newName, Price: newPrice}
	_, err := client.Database(database).Collection("items").ReplaceOne(ctx, filterByName, replacement)
	panicOnErr("Error replacing item", err)
}

func delete(database string, name string) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	filterByName := bson.D{{"name", name}}
	_, err := client.Database(database).Collection("items").DeleteOne(ctx, filterByName)
	panicOnErr("Error replacing item", err)
}

func panicOnErr(msg string, err error) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}

func main() {
	// Check connection to DB
	fmt.Println("Connecting to mongodb...")
	getConnection()
	fmt.Print("Connected\n\n")

	// Print available databases, at this point only the default ones are available
	fmt.Println("Listing existing databases:")
	listDatabases()
	fmt.Print("\n\n")

	// Create documents under collections "items" under databases "grocery" and "coffee-shop"
	fmt.Println("Adding documents to db...")
	create(databaseCoffeeShop, "Coffee", 9)
	create(databaseCoffeeShop, "Cake", 12)
	create(databaseCoffeeShop, "Icecream", 6)
	create(databaseGrocery, "Bread", 4)
	create(databaseGrocery, "Milk", 3)
	create(databaseGrocery, "Cheese", 8)
	create(databaseGrocery, "Guacamole", 23)
	fmt.Print("\n\n")

	// Print documents in coffee-shop
	fmt.Println("Listing items in DB coffee-shop:")
	listOneByOne(databaseCoffeeShop)
	fmt.Print("\n\n")

	// Print documents in grocery
	fmt.Println("Listing items in DB grocery:")
	listAllAtOnce(databaseGrocery)
	fmt.Print("\n\n")

	// Print sorted documents in grocery (by price, descendig)
	fmt.Println("Listing items in DB grocery sorted descending:")
	listSorted(databaseGrocery)
	fmt.Print("\n\n")

	// Print documents with price < 7
	fmt.Println("Listing items in DB grocery cheaper than 7")
	listCheaperThan7(databaseGrocery)
	fmt.Print("\n\n")

	// Change Milk price to 7
	fmt.Println("Updating Milk price to 7")
	updatePrice(databaseGrocery, "Milk", 7)
	listAllAtOnce(databaseGrocery)
	fmt.Print("\n\n")

	// Replace Guacamole with Pasztet
	fmt.Println("Replacing Guacamole for Pasztet")
	replaceEntireItem(databaseGrocery, "Guacamole", "Pasztet", 1)
	listAllAtOnce(databaseGrocery)
	fmt.Print("\n\n")

	// Delete Cheese
	fmt.Println("Deleting Cheese")
	delete(databaseGrocery, "Cheese")
	listAllAtOnce(databaseGrocery)
	fmt.Print("\n\n")

	fmt.Println("Listing existing databases:")
	listDatabases()
}
