package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	"log/slog"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func main() {
	// default dsn for embeddedpostgres
	const dsn = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	var noLogger io.Writer = nil // to avoid printing lots of postgre slog
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Logger(noLogger))
	panicOnErr(postgres.Start())
	defer postgres.Stop()

	panicOnErr(applyMigrations("./migrations/", dsn))

	db := openConn(dsn)
	writeData(db)
	readData(db)
}

// applyMigrations will migrate-up the running database with migration files found at provided path
func applyMigrations(migrationFilesRelativePath string, dsn string) error {
	// open new bare SQL connection to configured DB
	db, errOpen := sql.Open("postgres", dsn)
	if errOpen != nil {
		return errors.Wrap(errOpen, "error opening connection to postgres")
	}

	// initialize new postgres driver to use for applying migrations
	driver, errDriver := postgres.WithInstance(db, &postgres.Config{})
	if errDriver != nil {
		return errors.Wrap(errDriver, "error initializing postgres driver")
	}

	// gracefully close driver instances
	defer func(driver database.Driver) {
		errClose := driver.Close()
		slog.Info("closing postgres driver")
		if errClose != nil {
			slog.Error("migration: failed to gracefully close postgres driver", slog.String("error", errClose.Error()))
		}
	}(driver)

	// initialize migration instance with given driver
	m, errNewMigration := migrate.NewWithDatabaseInstance("file://"+migrationFilesRelativePath, "postgres", driver)
	if errNewMigration != nil {
		return errors.Wrap(errNewMigration, "error initializing migration instance")
	}

	// display schema version before running migration
	errVersion := checkVersion(m)
	if errVersion != nil {
		return errVersion
	}

	slog.Info("migration: starting to apply migrations")

	// apply migrations up to latest
	if errUp := m.Up(); errUp != nil {
		if errUp.Error() != "no change" {
			return errors.Wrap(errUp, "error running migration over postgres")
		}
	}

	slog.Info("migration: migrations have been successfully applied")

	// display schema version again after completion
	errVersion = checkVersion(m)
	if errVersion != nil {
		return errVersion
	}

	return nil
}

// version means how many migrations have been applied, db starts with version 0, after first migration file processd version is bumped to 1
// dirty means previous migration failed and this state must be resolved manually
func checkVersion(m *migrate.Migrate) error {
	// fetch stored schema migration information
	version, dirty, errVersion := m.Version()
	if errVersion != nil && !errors.Is(errVersion, migrate.ErrNilVersion) {
		return errors.Wrap(errVersion, "failed to fetch current schema info from db")
	}

	slog.Info("migration: got schema info",
		slog.Int64("version", int64(version)),
		slog.Bool("dirty", dirty))

	// quit out early if the schema is dirty
	if dirty {
		slog.Error("migration: schema is actively dirty, user interaction is required!")
		os.Exit(1)
	}

	return nil
}

func openConn(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	panicOnErr(err)

	return db
}

func writeData(db *sql.DB) {
	const sqlWriteData string = `
	INSERT INTO 
	people (name, age, height) 
	VALUES ($1, $2, $3);
	`

	_, err := db.Exec(sqlWriteData, "Andrzej", 32, 179) // supply arguments for placeholders $1, $2
	panicOnErr(err)

	_, err = db.Exec(sqlWriteData, "Jola", 24, 164) // supply arguments for placeholders $1, $2
	panicOnErr(err)
}

func readData(db *sql.DB) {
	const sqlReadData string = `
	SELECT name, age, height
	FROM people
	`

	rows, err := db.Query(sqlReadData)
	panicOnErr(err)

	for rows.Next() {
		var name string
		var age int
		var height int
		rows.Scan(&name, &age, &height)
		fmt.Println(name, age, height)
	}
	panicOnErr(rows.Close())
}

func panicOnErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
