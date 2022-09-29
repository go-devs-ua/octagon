package migration

import (
	"fmt"
	"os"

	"github.com/go-devs-ua/octagon/app/repository/pg"
)

// Execute migrations.
// Use "migrate-up" argument after run of program. It will create new table
// Or use "migrate-down" - that will remove table
// Example: go run main.go migrate-up
func Migrate(repo *pg.Repo) error {
	for _, arg := range os.Args {
		if arg == "migrate-up" {
			fmt.Println("Starting up")
			execMigrationQuery(repo, "./migration/migrationsUp.sql")
		}

		if arg == "migrate-down" {
			fmt.Println("Starting down")
			execMigrationQuery(repo, "./migration/migrationsDown.sql")
		}
	}

	return nil
}

func execMigrationQuery(repo *pg.Repo, fileName string) error {
	byteQuery, err := os.ReadFile(fileName)
	fmt.Println(string(byteQuery))
	if err != nil {
		return fmt.Errorf("read migration file: %w", err)
	}

	if err != nil {
		return err
	}

	if _, err := repo.DB.Exec(string(byteQuery)); err != nil {
		return fmt.Errorf("cannot exec query: %w", err)
	}
	return nil
}
