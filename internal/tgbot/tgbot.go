package tgbot

import (
	"log"
	"tgbot/internal/app"
)

var gfsd = make(chan bool)

func Run() {
	if err := app.InitBotApi(); err != nil {
		log.Fatalln("Бот не запустился блет")
	}
	log.Println("@taptykovo_bot запустился")

	<-gfsd
}
