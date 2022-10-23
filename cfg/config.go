// Package cfg contains structs
// that will hold on all needful parameters for our app
// that will be retrieved from  .env or ./cfg/config.yml.
package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Allowed logger levels & config key.
const (
	DebugLogLvl     = "DEBUG"
	InfoLogLvl      = "INFO"
	ErrorLogLvl     = "ERROR"
	LogLvlConfigKey = "LOG_LEVEL"
	lenOfLines      = 2
	envFileName     = ".env"
)

// Load configs from a env file & sets them in environment variables.
func loadEnvVar() error {
	f, err := os.Open(envFileName)
	if err != nil {
		return fmt.Errorf("error while opening %s file: %w", envFileName, err)
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
		return fmt.Errorf("error while scaning %s file: %w", envFileName, err)
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		if len(pair) != lenOfLines {
			return errors.New("not enough data for the configuration at the config file")
		}

		os.Setenv(pair[0], pair[1])
	}

	return nil
}

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// Server configuration description.
type Server struct {
	Host string
	Port string
}

// Options will keep all needful configs.
type Options struct {
	LogLevel string
	Server   Server
	DB       DB
}

// GetConfig will create instance of Options
// that will be used im main package.
func GetConfig() (Options, error) {
	if err := loadEnvVar(); err != nil {
		return Options{}, err
	}

	opt := Options{
		LogLevel: os.Getenv(LogLvlConfigKey),
		Server: Server{
			Host: os.Getenv("SERV_HOST"),
			Port: os.Getenv("SERV_PORT"),
		},
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}

	if err := opt.validate(); err != nil {
		return Options{}, fmt.Errorf("validation failed: %w", err)
	}

	return opt, nil
}

func (opt Options) validate() error {
	if strings.ToUpper(opt.LogLevel) != DebugLogLvl &&
		strings.ToUpper(opt.LogLevel) != ErrorLogLvl &&
		strings.ToUpper(opt.LogLevel) != InfoLogLvl {
		return fmt.Errorf("\"%v\" is not allowed logger level", opt.LogLevel)
	}

	return nil
}
