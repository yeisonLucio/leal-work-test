package database

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func GetTestDB(t *testing.T) *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("failed to connect database: %s", err)
	}

	err = RunMigrations(db)
	if err != nil {
		t.Fatalf("failed to run migrations: %s", err)
	}

	return db
}
