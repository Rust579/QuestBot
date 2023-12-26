package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"tgbot/internal/model"
	"tgbot/internal/service/logic"
)

const (
	//Command
	CommandGetCode   = "Получить код"
	CommandReference = "Справка"
)

type Bot struct {
	bot  tgbotapi.BotAPI
	Msgs chan string
}

var BotApi Bot

func InitBotApi(msgs chan string) error {

	bot, err := tgbotapi.NewBotAPI("6583722718:AAE9b84iNSj_YHFEOBad1P_8my7IgwyD7gg")
	if err != nil {
		return err
	}

	bot.Debug = false
	BotApi.bot = *bot
	BotApi.Msgs = msgs

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
			tgbotapi.NewKeyboardButton(CommandReference),
		),
	)

	b.Msgs <- "@" + message.From.UserName + " запустил бота"

	msg := logic.ProcessMessagesCommand(message.Command(), model.PullUsers.P[message.From.ID].Stage)
	if msg.Stage != 0 && model.PullUsers.P[message.From.ID].Stage == 0 {
		model.PullUsers.AddUser(message.From.ID, message.From.UserName, 1)

		fmt.Println(model.PullUsers.P[message.From.ID])
	}

	msgBot := tgbotapi.NewMessage(message.From.ID, msg.Message)
	msgBot.ReplyMarkup = numericKeyboard

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
				b.Msgs <- "Ошибка HandleCommand у " + "@" + update.Message.From.UserName + " " + err.Error()
			}
			continue
		}

		b.Msgs <- "Сообщение от " + "@" + update.Message.From.UserName + ": " + update.Message.Text

		msg := logic.ProcessMessagesText(update.Message.Text, model.PullUsers.P[update.Message.From.ID].Stage)
		if msg.Stage != 0 && msg.Stage > model.PullUsers.P[update.Message.From.ID].Stage {
			model.PullUsers.IncStage(update.Message.From.ID, msg.Stage)

			fmt.Println(model.PullUsers.P[update.Message.From.ID])
		}

		if msg.Type == logic.TypeImg {
			b.sendPhoto(update.Message.From, msg)
		}
		if msg.Type == logic.TypeAudio {
			b.sendAudio(update.Message.From, msg)
		}
		if msg.Type == logic.TypeStr {
			b.SendTxt(update.Message.From, msg)
		}
		continue
	}
}

func (b *Bot) sendPhoto(user *tgbotapi.User, msg logic.RespMsg) {

	fileBytes, er := ioutil.ReadFile(msg.FilePath)
	if er != nil {
		b.Msgs <- "Ошибка чтения файла фото " + "@" + user.UserName + " " + er.Error()
	}

	msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

	msgBot := tgbotapi.NewPhoto(user.ID, msgf)

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки фото " + "@" + user.UserName + " " + err.Error()
	}

}

func (b *Bot) sendAudio(user *tgbotapi.User, msg logic.RespMsg) {

	fileBytes, er := ioutil.ReadFile(msg.FilePath)
	if er != nil {
		b.Msgs <- "Ошибка чтения файла аудио " + "@" + user.UserName + " " + er.Error()
	}

	msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

	msgBot := tgbotapi.NewAudio(user.ID, msgf)

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка чтения файла аудио " + "@" + user.UserName + " " + er.Error()
	}

}

func (b *Bot) SendTxt(user *tgbotapi.User, msg logic.RespMsg) error {

	msgBot := tgbotapi.NewMessage(user.ID, msg.Message)
	//msgBot.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msgBot); err != nil {
		return err
	}
	return nil
}
