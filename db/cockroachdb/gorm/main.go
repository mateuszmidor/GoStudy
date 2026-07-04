package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type patient struct {
	PatientID    int64      `gorm:"primary_key;auto_increment" `
	Name         string     `gorm:"type:varchar(128)"`
	Age          int        `gorm:"type:int"`
	BloodGroup   bloodGroup `gorm:"foreignKey:BloodGroupID"`
	BloodGroupID int64
}

type bloodGroup struct {
	BloodGroupID int64  `gorm:"primary_key;auto_increment" `
	Type         string `gorm:"type:varchar(6)"`
}

func openConn(connectionString string) *gorm.DB {
	pdb := postgres.Open(connectionString)
	db, err := gorm.Open(pdb, &gorm.Config{})
	panicOnErr(err)

	return db
}

func createTables(db *gorm.DB) {
	err := db.AutoMigrate(&patient{}, &bloodGroup{})
	panicOnErr(err)
}

func writeData(db *gorm.DB) {
	panicOnErr(db.Create(&patient{Name: "Andrzej", Age: 32, BloodGroup: bloodGroup{Type: "A Rh+"}}).Error)
	panicOnErr(db.Create(&patient{Name: "Jola", Age: 22, BloodGroup: bloodGroup{Type: "B Rh+"}}).Error)
	panicOnErr(db.Create(&patient{Name: "Witek", Age: 42, BloodGroup: bloodGroup{Type: "0 Rh-"}}).Error)
}

func readData(db *gorm.DB) {
	var patients []patient
	panicOnErr(db.Preload("BloodGroup").Find(&patients).Error)

	for _, p := range patients {
		fmt.Printf("%+10s  %d  %s\n", p.Name, p.Age, p.BloodGroup.Type)
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db := openConn("postgresql://root@localhost:26257/defaultdb?sslmode=disable") // defaultdb exists in cluster once it has beed initialized
	// defer db.Close() // no Close method

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
