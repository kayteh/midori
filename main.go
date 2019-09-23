package main // import "github.com/kayteh/midori"

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	token         = os.Getenv("DISCORD_BOT_TOKEN")
	ghAccessToken = os.Getenv("GH_ACCESS_TOKEN")
	ghAppKey      = os.Getenv("GH_APP_KEY")
)

func main() {
	if token == "" {
		log.Fatalln("DISCORD_BOT_TOKEN isn't set. Exiting.")
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln(err)
	}

	discord.AddHandler(messageHandler)
}
