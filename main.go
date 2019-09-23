package main // import "github.com/kayteh/midori"

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kayteh/midori/chatops"
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

	chatops, err := chatops.NewChatOps(&chatops.ChatOpsConfig{})

	handler := &messageHandler{
		chatops: chatops,
	}

	discord.AddHandler(handler.Handle)

	err = discord.Open()
	if err != nil {
		log.Fatalln(err)
	}

	handler.createRegexpFromUser(discord.State.User)

	// todo: start http handler

	fmt.Println("midori: started bot")

	syscallExit := make(chan os.Signal, 1)
	signal.Notify(
		syscallExit,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
		os.Kill,
	)
	<-syscallExit

	discord.Close()
}
