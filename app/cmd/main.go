package main

import (
	"app/internal/client"
	"app/internal/config"
	"app/internal/handlers"
	"app/pkg/database"
)

func main() {
	config, err := config.Load()
	if !err.IsNil() {
		err.Fatal()
	}
	
	err = database.InitDb()
	if !err.IsNil() {
		err.Fatal()
	}

	bot, err := client.SetupWebhook(config)
	if !err.IsNil() {
		err.Fatal()
	}

	err = handlers.LoadHandlers()
	if !err.IsNil() {
		err.Fatal()
	}

	bot.Start()
}
