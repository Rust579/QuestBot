package tgbot

import (
	"log"
	"tgbot/internal/app"
	"tgbot/internal/tgmanager"
)

var gfsd = make(chan bool)

func Run() {
	c := make(chan string)

	if err := app.InitBotApi(c); err != nil {
		log.Fatalln("taptykovo_bot не запустился блет")
	}
	if err := tgmanager.InitTgApi(c); err != nil {
		log.Fatalln("manager_bot не запустился блет")
	}
	log.Println("@taptykovo_bot запустился")

	<-gfsd
}
