package pkg

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

type GomudBot struct {
	session *discordgo.Session
	token   string
}

func NewBot(token string) *GomudBot {
	bot := new(GomudBot)
	var err error
	bot.session, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Could not create new Discord session, ", err)
		return nil
	}
	bot.token = token

	return bot
}

func (b *GomudBot) InstallRouter(router *GomudRouter) *GomudBot {
	createHandler := func(s *discordgo.Session, m *discordgo.MessageCreate) {
		router.Match(m.Content)
	}
	b.session.AddHandler(createHandler)
	return b
}

func (b *GomudBot) Start() {
	var err error
	err = b.session.Open()
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
	_ = b.session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
