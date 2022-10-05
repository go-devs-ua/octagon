// Package cfg contains structs
// that will hold on all needful parameters for our app
// that will be retrieved from  .env or ./cfg/config.yml
package cfg

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

// Load configs from a env file & sets them in environment variables
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

// Server configuration description
type Server struct {
	Host string
	Port string
}

// Database configuration description
type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// Options will keep all needful configs
type Options struct {
	Server Server
	DB     DB
}

// GetConfig will create instance of Options
// that will be used im main package
func GetConfig() (Options, error) {
	err := loadEnvVar(); if err != nil {
		return Options{}, err
	}

	return Options{
		Server{
			Host: os.Getenv("SERV_HOST"),
			Port: os.Getenv("SERV_PORT"),
		},
		DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}, nil
}
