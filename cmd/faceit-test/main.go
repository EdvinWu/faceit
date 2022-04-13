package main

import (
	"faceit-test/internal/app"
	"faceit-test/internal/config"
	"log"
)

func main() {
	conf, err := config.Load(".")
	if err != nil {
		log.Fatalf("failed to load application config")
	}
	app.NewApp(conf).Run()
}
