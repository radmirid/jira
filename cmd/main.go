package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/radmirid/jira/config"
	"github.com/radmirid/jira/db"
	"github.com/radmirid/jira/telegram"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	defer db.Close()

	config.InitJiraClient()

	bot, err := telegram.InitBot()
	if err != nil {
		log.Fatal(err)
	}
	defer bot.Close()

	telegram.InitHandlers(bot)

	bot.Start()
}
