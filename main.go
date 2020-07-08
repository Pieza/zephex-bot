package main

import (
	"fmt"
	"log"
	"os"
	"./actions/init"
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

	// register events
	//bot.AddHandler(ready)
	bot.AddHandler(messageCreate)

	err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	for {
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}
