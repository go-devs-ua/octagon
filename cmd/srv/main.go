// Package main is key entrypoint
// of our awesome application
package main

import (
	"log"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	return nil
}
