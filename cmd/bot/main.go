package main

import (
	"log"

	"github.com/nqrm/go_bot/internal/clients/tg"
	"github.com/nqrm/go_bot/internal/config"
	"github.com/nqrm/go_bot/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)

}
