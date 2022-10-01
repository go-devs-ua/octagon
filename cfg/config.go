// Package cfg contains structs
// that will hold on all needful parameters for our app
// that will be retrieved from  .env or ./cfg/config.yml
package cfg

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// Load configs from a env file & sets them in environment variables .
func LoadEnvVar() error {

	f, err := os.Open(".env")

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("%s", err)
		return err
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
	return nil
}

var (
	serverHost = os.Getenv("SERVHOST")
	serverPort = os.Getenv("SERVPORT")
	DBName     = os.Getenv("DBNAME")
	DBUser     = os.Getenv("DDBUSER")
	DBPassword = os.Getenv("DBPASSWORD")
	DBPort     = os.Getenv("DBPORT")
	DBHost     = os.Getenv("DBHOST")
)

// Server configuration description
type Server struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
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

// NewOptions will create instance of Options
// that will be used im main package
func NewOptions() *Options {
	return &Options{
		Server{
			Host: serverHost,
			Port: ":" + serverPort,
		},
		DB{
			Host:     DBHost,
			Port:     DBPort,
			Username: DBUser,
			Password: DBPassword,
			DBName:   DBName,
		},
	}
}
