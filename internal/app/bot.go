package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"tgbot/internal/model"
	"tgbot/internal/service/logic"
)

const (
	// Кнопки в боте
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

// Обработчик команд
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
	}

	msgBot := tgbotapi.NewMessage(message.From.ID, msg.Message)
	msgBot.ReplyMarkup = numericKeyboard
	msgBot.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msgBot); err != nil {
		return err
	}

	return nil
}

// Обработчик сообщений
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

		if len(update.Message.Text) != 0 {
			msg := logic.ProcessMessagesText(update.Message.Text, model.PullUsers.P[update.Message.From.ID].Stage)
			if msg.Stage != 0 && msg.Stage > model.PullUsers.P[update.Message.From.ID].Stage {
				model.PullUsers.IncStage(update.Message.From.ID, msg.Stage)
			}

			if msg.Type == logic.TypeImg {
				b.sendPhoto(update.Message, msg)
			}
			if msg.Type == logic.TypeAudio {
				b.sendAudio(update.Message, msg)
			}
			if msg.Type == logic.TypeStr {
				b.SendTxt(update.Message, msg)
			}
		}

		if update.Message.Location != nil {
			refLoc := logic.ProcessLocation(model.PullUsers.P[update.Message.From.ID].Stage)
			b.CheckLocation(update.Message, refLoc)
		}

		continue
	}
}

func (b *Bot) sendPhoto(user *tgbotapi.Message, msg logic.RespMsg) {

	fileBytes, er := ioutil.ReadFile(msg.FilePath)
	if er != nil {
		b.Msgs <- "Ошибка чтения файла фото " + "@" + user.From.UserName + " " + er.Error()
	}

	msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

	msgBot := tgbotapi.NewPhoto(user.From.ID, msgf)

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки фото " + "@" + user.From.UserName + " " + err.Error()
	}

	b.SendTxt(user, msg)

}

func (b *Bot) sendAudio(user *tgbotapi.Message, msg logic.RespMsg) {

	fileBytes, er := ioutil.ReadFile(msg.FilePath)
	if er != nil {
		b.Msgs <- "Ошибка чтения файла аудио " + "@" + user.From.UserName + " " + er.Error()
	}

	msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

	msgBot := tgbotapi.NewAudio(user.From.ID, msgf)

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки файла аудио " + "@" + user.From.UserName + " " + er.Error()
	}

}

func (b *Bot) SendTxt(user *tgbotapi.Message, msg logic.RespMsg) {

	msgBot := tgbotapi.NewMessage(user.From.ID, msg.Message)
	msgBot.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки текстового сообщения " + "@" + user.From.UserName + " " + err.Error()
	}
}

func (b *Bot) CheckLocation(user *tgbotapi.Message, refLoc logic.RefLocation) {

	var msgBot tgbotapi.MessageConfig
	if user.Location.Latitude > refLoc.Latitude+0.000278 || user.Location.Latitude < refLoc.Latitude-0.000278 ||
		user.Location.Longitude > refLoc.Longitude+0.000278 || user.Location.Longitude < refLoc.Longitude-0.000278 {

		msgBot = tgbotapi.NewMessage(user.From.ID, refLoc.IncorrectMsg)
	} else {
		msgBot = tgbotapi.NewMessage(user.From.ID, refLoc.CorrectMsg)
	}

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки текстового сообщения " + "@" + user.From.UserName + " " + err.Error()
	}
}
