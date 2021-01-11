package pkg

import (
	"github.com/bwmarrin/discordgo"
	l "github.com/rs/zerolog/log"
	"os"
)

var log = l.With().Caller().Logger()

func Run(signal <-chan os.Signal, token string) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create discord session")
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect discord")
	}

	log.Info().Msg("pinpong bot - start")

	<-signal

	log.Info().Msg("pinpong bot - end")

	err = dg.Close()
	if err != nil {
		log.Error().Err(err).Msg("failed to close discord session")
	}
}

func messageCreate(session *discordgo.Session, createMsg *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if createMsg.Author.ID == session.State.User.ID {
		return
	}

	var err error

	if createMsg.Content == "ping" {
		_, err = session.ChannelMessageSend(createMsg.ChannelID, "Pong!")
	} else if createMsg.Content == "pong" {
		_, err = session.ChannelMessageSend(createMsg.ChannelID, "Ping!")
	}

	if err != nil {
		log.Error().Err(err).Msg("failed to send message")
	}
}
