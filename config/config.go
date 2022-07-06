package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDb(database string) *sql.DB {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// load environment variables
	err := godotenv.Load(".env")

	// handle error, if any
	if err != nil {
		_ = fmt.Errorf("error loading .env-example file")
	}

	// format connection string from environment variables
	databaseURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	// create a database object which can be used to communicate with database
	db, err := sql.Open(database, databaseURL)

	// handle error, if any
	if err != nil {
		log.Fatal(database+"failed to start", err)
	}

	// test database connection
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("could not ping the database", err)
	}

	// return database object
	return db
}
