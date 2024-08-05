package messages

import (
	"regexp"
	"strings"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
	db       map[int]Cost
	id       int
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
		db:       make(map[int]Cost),
		id:       0,
	}
}

type Message struct {
	Text   string
	UserID int64
}

// структура трат
type Cost struct {
	date     string
	category string
	price    string
}

var re, _ = regexp.Compile(`\s*\d\d[/]\d\d[/]\d\d\d\d\s*,\s*[a-zA-Zа-яА-Я]+\s*,\s*\d+`)

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		s.tgClient.SendMessage("Отправьте новый расход в формате:\n Дата(14/04/2024), Категория трат, Сумма (123231, без копеек)", msg.UserID)
		return nil
	} else if cost := re.FindString(msg.Text); cost != "" {
		fields := strings.Split(cost, ",")
		date := strings.TrimSpace(fields[0])
		category := strings.TrimSpace(fields[1])
		price := strings.TrimSpace(fields[2])
		s.db[s.id] = Cost{date, category, price}
		s.id += 1
		s.tgClient.SendMessage("Расход добавлен", msg.UserID)
	} else {
		s.tgClient.SendMessage("unknown command", msg.UserID)
	}
	return nil
}
