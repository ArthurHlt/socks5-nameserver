package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Version string

func main() {
	if Version == "" {
		Version = "dev"
	}
	app := NewApp(Version)
	err := app.Run(os.Args)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
