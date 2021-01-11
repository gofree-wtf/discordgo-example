package main

import (
	"flag"
	"github.com/gofree-wtf/discordgo-example/pkg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

var (
	debug bool
	token string
)

func main() {
	flag.BoolVar(&debug, "debug", false, "set log level to debug")
	flag.StringVar(&token, "bot.token", "", "bot authentication token")
	flag.Parse()

	initLogger()

	log.Debug().Bool("-debug", debug).Msg("")
	log.Debug().Str("-bot.token", token).Msg("")

	validateFlags()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	pkg.Run(sc, token)
}

func initLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func validateFlags() {
	if token == "" {
		log.Panic().Msg("-bot.token 을 입력하세요")
	}
}
