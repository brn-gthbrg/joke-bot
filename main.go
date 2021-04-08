package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/jokeapi-go"
)

// const token = DISCORD_TOKEN

var BotID string

func keyReader(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	// Create a new Discord session using the provided bot token.
	token := keyReader("DISCORD_KEY")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	jt := "single"
	blacklist := []string{"nsfw", "religious", "political", "racist", "sexist", "explicit"}
	ctgs := []string{"Misc", "Pun"}

	api := jokeapi.New()
	api.SetParams(&ctgs, &blacklist, &jt)
	response := api.Fetch()
	joke := strings.Join(response.Joke, "")

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m.Content) == "tell me a joke" {
		s.ChannelMessageSend(m.ChannelID, joke)
	}

	if strings.ToLower(m.Content) == "tell me your best joke" {
		s.ChannelMessageSend(m.ChannelID, "I'm an unskilled bot with no concept of best")
	}

}
