package telegram

import "title-save-bot/clients/telegram"

type Processor struct {
	tg    *telegram.Client
	ofset int
	//storage
}
