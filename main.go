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

	_tgClient := telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
