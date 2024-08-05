package messages

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	date     time.Time
	category string
	price    int
}

var re, _ = regexp.Compile(`\s*\d\d\d\d[/]\d\d[/]\d\d\s*,\s*[a-zA-Zа-яА-Я]+\s*,\s*\d+`)

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		s.tgClient.SendMessage("Отправьте новый расход в формате:\n Дата(2024/03/25), Категория трат, Сумма (123231, без копеек)", msg.UserID)
		return nil
	} else if cost := re.FindString(msg.Text); cost != "" {
		fields := strings.Split(cost, ",")
		date := strings.TrimSpace(fields[0])
		t, _ := time.Parse("2006/01/02", date)
		category := strings.TrimSpace(fields[1])
		price, _ := strconv.Atoi(strings.TrimSpace(fields[2]))
		s.db[s.id] = Cost{t, category, price}
		s.id += 1
		s.tgClient.SendMessage("Расход добавлен", msg.UserID)
	} else if msg.Text == "/weeklyreport" {
		var sum int
		for _, c := range s.db {
			if int(time.Now().Sub(c.date)/(24*time.Hour)) <= 7 {
				sum += c.price
			}
		}
		var mesg string = fmt.Sprintf("Всего потрачено за последнюю неделю: %d", sum)
		s.tgClient.SendMessage(mesg, msg.UserID)
	} else {
		s.tgClient.SendMessage("unknown command", msg.UserID)
	}
	return nil
}
