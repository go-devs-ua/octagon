package migration

import (
	"database/sql"
	"fmt"
	"os"
)

// Execute migrations.
// Use "migrate-up" argument after run of program. It will create new table
// Or use "migrate-down" - that will remove table
// Example: go run main.go migrate-up
func Migrate(db *sql.DB) error {
	for _, arg := range os.Args {
		if arg == "migrate-up" {
			fmt.Println("Starting up")
			execMigrationQuery(db, "./migration/migrationsUp.sql")
		}

		if arg == "migrate-down" {
			fmt.Println("Starting down")
			execMigrationQuery(db, "./migration/migrationsDown.sql")
		}
	}

	return nil
}

func execMigrationQuery(db *sql.DB, fileName string) error {
	byteQuery, err := os.ReadFile(fileName)
	fmt.Println(string(byteQuery))
	if err != nil {
		return fmt.Errorf("read migration file: %w", err)
	}

	if err != nil {
		return err
	}

	if _, err := db.Exec(string(byteQuery)); err != nil {
		return fmt.Errorf("cannot exec query: %w", err)
	}
	return nil
}
