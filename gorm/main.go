package main

import (
	"fmt"
	"io"
	"log/slog"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	dialector "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User represents PostgreSQL table "user"
type User struct {
	ID   int64  `gorm:"primaryKey"` // instead, we can embed gorm.Model that includes ID, CreatedAt, UpdatedAt, DeletedAt
	Name string `gorm:"index"`
	Age  int    `gorm:"default:18"`
}

// BeforeCreate hook will print user as it was passed in db.Create call (no ID, no default values)
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	slog.Info("creating", slog.Any("user", *u))
	return
}

// AfterCreate hook will print user with generated ID and assigned default Age (if wasn't explicitly provided)
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	slog.Info("created", slog.Any("user", *u))
	return
}

// threre are more hooks that can be used: https://medium.com/@itskenzylimon/getting-started-on-golang-gorm-af49381caf3f

func main() {
	var noLogger io.Writer = nil // to avoid printing lots of postgre logs
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Logger(noLogger))
	panicOnErr(postgres.Start())
	defer postgres.Stop()
	db := openConn("host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable") // use embedded-postgres default credentials

	createTables(db)
	writeData(db)
	readData(db)
	deleteData(db)
}

func openConn(connectionString string) *gorm.DB {
	db, err := gorm.Open(dialector.Open(connectionString), &gorm.Config{})
	panicOnErr(err)
	return db
}

func createTables(db *gorm.DB) {
	err := db.AutoMigrate(&User{}) // create table if missing, update it if exists but does not remove any columns
	panicOnErr(err)
}

func writeData(db *gorm.DB) {
	err := db.Create(&User{Name: "Andrzej", Age: 34}).Error // PrimaryKey will be auto-assigned
	panicOnErr(err)
	err = db.Create(&User{Name: "Jola"}).Error // Age will be auto-assigned a default value of 18
	panicOnErr(err)
}

func readData(db *gorm.DB) {
	var users []User
	err := db.Model(&User{}).Find(&users).Error // find all users in table User
	panicOnErr(err)

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func deleteData(db *gorm.DB) {
	db.Where("1 = 1").Delete(&User{}) // delete all users in table User. GORM prevents bulk delete, so this dummy Where clause is needed
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
