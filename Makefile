DEBUG ?= false

build:
	go build -o discordgo-example

run:
	go run main.go -debug=$(DEBUG) -bot.token=$(BOT_TOKEN)
