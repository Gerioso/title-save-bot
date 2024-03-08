package telegram

import (
	"errors"
	"title-save-bot/clients/telegram"
	"title-save-bot/events"
	"title-save-bot/lib/e"
	"title-save-bot/storage"
)

type Processor struct {
	tg      *telegram.Client
	ofset   int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

var ErrUnknownEventType = errors.New("unknown event type")
var ErrUnknownMeta = errors.New("unknown meta")

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.ofset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}
	if len(updates) == 0 {
		return nil, nil
	}
	res := make([]events.Event, len(updates))
	for _, u := range updates {
		res = append(res, event(u))
	}
	p.ofset = updates[len(updates)-1].ID + 1
	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("can't process message", err)
	}

}
func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("can't get meta", ErrUnknownMeta)
	}
	return res, nil
}

func event(upd telegram.Update) events.Event {
	updateType := fetchType(upd)
	res := events.Event{
		Type: updateType,
		Text: fetchText(upd),
	}
	if updateType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.Id,
			Username: upd.Message.From.Username,
		}
	}
	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}
