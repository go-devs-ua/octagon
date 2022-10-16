package pg

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func loadEnvVar() error {
	f, err := os.Open(".env")
	if err != nil {
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("%s", err)
		}
	}()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		if len(pair) != 2 {
			return errors.New("not enough data for the configuration in .env file")
		}

		os.Setenv(pair[0], pair[1])
	}

	return nil
}

type DBOptions struct {
	DB DB
}

func GetConfig() (DBOptions, error) {
	if err := loadEnvVar(); err != nil {
		return DBOptions{}, err
	}

	return DBOptions{
		DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}, nil
}

func ConnectDB(opt DB) (*sql.DB, error) {
	str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		opt.Host, opt.Port, opt.Username, opt.Password, opt.DBName)

	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %w", err)
	}

	return db, nil
}
