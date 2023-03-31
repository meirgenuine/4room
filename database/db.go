package database

import (
	"4room/config"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func Initialize() {
	var err error

	dsn := fmt.Sprintf("%s", config.DBConfig.Database)
	DB, err = sql.Open(config.DBConfig.Driver, dsn)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	// Run migrations
	Migrate()
}

func Migrate() {
	migrationFiles, err := filepath.Glob("database/migrations/*.sql")
	if err != nil {
		log.Fatalf("Error reading migration files: %v", err)
	}

	for _, file := range migrationFiles {
		migration, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading migration file %s: %v", file, err)
		}

		_, err = DB.Exec(string(migration))
		if err != nil {
			log.Fatalf("Error executing migration for file %s: %v", file, err)
		}
	}

	log.Println("Database migrations applied successfully")
}
