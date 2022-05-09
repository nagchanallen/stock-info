package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/nagchanallen/stock-info/backend/utils"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	discord, err := discordgo.New("Bot " + config.DISCORD_TOKEN)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Gracefully shutting down")

	defer discord.Close()
}
