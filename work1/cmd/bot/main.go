package main

import (
	"flonsay/workspace_go/tech/ozon_tech/work1/internal/clients/tg"
	"flonsay/workspace_go/tech/ozon_tech/work1/internal/config"
	"flonsay/workspace_go/tech/ozon_tech/work1/internal/model/messages"
	"log"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed: ", err)
	}
	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed: ", err)
	}
	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}