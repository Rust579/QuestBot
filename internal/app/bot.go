package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"tgbot/internal/service/logic"
)

const (
	//Command
	CommandGetCode   = "Получить код"
	CommandReference = "Справка"
)

type Bot struct {
	bot  tgbotapi.BotAPI
	msgs chan string
}

var BotApi Bot

func InitBotApi(msgs chan string) error {

	bot, err := tgbotapi.NewBotAPI("6583722718:AAE9b84iNSj_YHFEOBad1P_8my7IgwyD7gg")
	if err != nil {
		return err
	}

	bot.Debug = false
	BotApi.bot = *bot
	BotApi.msgs = msgs

	// Bot Start
	go BotApi.Start()
	println("telegram bot start: ")

	return nil
}

func (b *Bot) Start() {

	updates := b.initUpdatesChan()
	b.HandleUpdates(updates)

	return
}

func (b *Bot) initUpdatesChan() tgbotapi.UpdatesChannel {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) HandleCommand(message *tgbotapi.Message) error {

	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandGetCode),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandReference),
		),
	)

	b.msgs <- "@" + message.From.UserName + " запустил бота"
	fmt.Println("@" + message.From.UserName + " запустил бота")

	msg := logic.ProcessMessagesCommand(message.Command())

	msgBot := tgbotapi.NewMessage(message.From.ID, msg)
	msgBot.ReplyMarkup = numericKeyboard
	//msgBot.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msgBot); err != nil {
		return err
	}

	return nil
}

func (b *Bot) HandleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.HandleCommand(update.Message); err != nil {
				b.msgs <- "Ошибка HandleCommand у " + "@" + update.Message.From.UserName + " " + err.Error()
			}
			continue
		}

		b.msgs <- "Сообщение от " + "@" + update.Message.From.UserName + " : " + update.Message.Text
		fmt.Println("Сообщение от " + "@" + update.Message.From.UserName + " : " + update.Message.Text)

		msg := logic.ProcessMessagesText(update.Message.Text)

		if msg == "файл" {
			b.sendPhoto(update.Message.From.ID)
		}

		msgBot := tgbotapi.NewMessage(update.Message.From.ID, msg)
		msgBot.ParseMode = tgbotapi.ModeMarkdownV2

		if _, err := b.bot.Send(msgBot); err != nil {
			b.msgs <- "Ошибка отправки сообщения " + "@" + update.Message.From.UserName + " " + err.Error()
		}
		continue

	}
}

func (b *Bot) sendPhoto(chatId int64) {

	filePath := "files/alt.mp3"

	fileBytes, er := ioutil.ReadFile(filePath)
	if er != nil {
		log.Panic(er)
	}

	msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

	msgBot := tgbotapi.NewPhoto(chatId, msgf)

	if _, err := b.bot.Send(msgBot); err != nil {
	}

}

func (b *Bot) sendAudio(chatId int64) {

	filePath := "files/alt.mp3"

	fileBytes, er := ioutil.ReadFile(filePath)
	if er != nil {
		log.Panic(er)
	}

	msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

	msgBot := tgbotapi.NewAudio(chatId, msgf)

	if _, err := b.bot.Send(msgBot); err != nil {
	}

}

func (b *Bot) SendCode(ChatID int64, msg string) error {

	msgBot := tgbotapi.NewMessage(ChatID, msg)
	msgBot.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msgBot); err != nil {
		return err
	}
	return nil
}
