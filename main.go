package main

import (
	"flag"
	"log"
	"title-save-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	_ := telegram.New(tgBotHost, mustTokern())

}

func mustTokern() string {
	token := flag.String(
		"tg-bot-tokern",
		"",
		"telegram access token",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
