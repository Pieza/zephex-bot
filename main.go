package main

import (
	"fmt"
	"log"
	"main/db"
	"main/events"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		panic(err)
	}

	db.SetUpDB()

	// register events
	bot.AddHandler(events.Ready)
	bot.AddHandler(events.MessageCreate)

	err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	for {
		// wait to close
	}
}
