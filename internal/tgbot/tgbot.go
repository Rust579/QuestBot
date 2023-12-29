package tgbot

import (
	"log"
	"tgbot/internal/app"
	"tgbot/internal/model"
	"tgbot/internal/tgmanager"
)

var gfsd = make(chan bool)

func Run() {
	fromBot2Manager := make(chan string)
	fromMan2Bot := make(chan model.SendTo)

	if err := app.InitBotApi(fromBot2Manager, fromMan2Bot); err != nil {
		log.Fatalln("taptykovo_bot не запустился блет")
	}
	if err := tgmanager.InitTgApi(fromBot2Manager, fromMan2Bot); err != nil {
		log.Fatalln("manager_bot не запустился блет")
	}
	log.Println("Боты запустились")

	<-gfsd
}
