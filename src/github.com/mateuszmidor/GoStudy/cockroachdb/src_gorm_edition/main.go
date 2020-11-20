package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type person struct {
	PersonID int64  `gorm:"primary_key;auto_increment" `
	Name     string `gorm:"type:varchar(32)"`
	Age      int    `gorm:"type:int"`
	Room     room   `gorm:"foreignKey:RoomID"`
	RoomID   int64
}

type room struct {
	RoomID int64  `gorm:"primary_key;auto_increment" `
	Type   string `gorm:"type:varchar(32)"`
}

func openConn(connectionString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	panicOnErr(err)

	return db
}

func createTables(db *gorm.DB) {
	err := db.AutoMigrate(&person{}, &room{})
	panicOnErr(err)
}

func writeData(db *gorm.DB) {
	panicOnErr(db.Create(&person{Name: "Andrzej", Age: 32, Room: room{Type: "Kuchnia"}}).Error)
	panicOnErr(db.Create(&person{Name: "Jola", Age: 22, Room: room{Type: "WC"}}).Error)
	panicOnErr(db.Create(&person{Name: "Witek", Age: 42, Room: room{Type: "Weranda"}}).Error)
}

func readData(db *gorm.DB) {
	var people []person
	panicOnErr(db.Preload("Room").Find(&people).Error)

	for _, p := range people {
		fmt.Println(p.Name, p.Age, p.Room.Type)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db := openConn("postgresql://root@localhost:26257/defaultdb?sslmode=disable") // defaultdb exists in cluster once it has beed initialized
	// defer db.Close()

	fmt.Println("Creating tables")
	createTables(db)
	fmt.Print("Done\n\n")

	fmt.Println("Adding records")
	writeData(db)
	fmt.Print("Done\n\n")

	fmt.Println("Reading records")
	readData(db)
	fmt.Print("Done\n\n")
}
