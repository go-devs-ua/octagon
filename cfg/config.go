// Package cfg contains structs
// that will hold on all needful parameters for our app
// that will be retrieved from  .env or ./cfg/config.yml
package cfg

import "time"

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
func NewOptions() Options {
	return Options{}
}
