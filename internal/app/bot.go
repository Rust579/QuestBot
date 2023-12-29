package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"strconv"
	"tgbot/internal/model"
	"tgbot/internal/service/logic"
)

const (
	// Кнопки в боте
	CommandReference = "Текущее задание"

	r10m  = 0.000139
	r20m  = 0.000278
	r40m  = 0.000556
	r60m  = 0.000834
	r80m  = 0.001112
	r100m = 0.001139

	//msg10m    = "Раз два три четыре пять я иду штурвал считать."
	msg20m    = "Горячо"
	msg40m    = "Очень тепло"
	msg60m    = "Тепло"
	msg80m    = "Свежо"
	msg100m   = "Прохладно"
	msg10000m = "Капец как холодно"

	// Эталонные координаты детской площадки
	refLat1 = 54.595999
	refLon1 = 55.802448
)

type Bot struct {
	bot    tgbotapi.BotAPI
	Msgs   chan string
	SendTo chan model.SendTo
}

var BotApi Bot

func InitBotApi(msgs chan string, sendTo chan model.SendTo) error {

	bot, err := tgbotapi.NewBotAPI("6583722718:AAE9b84iNSj_YHFEOBad1P_8my7IgwyD7gg")
	if err != nil {
		return err
	}

	bot.Debug = false
	BotApi.bot = *bot
	BotApi.Msgs = msgs
	BotApi.SendTo = sendTo

	// Bot Start
	go BotApi.Start()
	println("telegram bot start: ")

	return nil
}

func (b *Bot) Start() {

	go func() {
		log.Println("start reading Send to")
		for st := range b.SendTo {
			err := b.SendMsg(st.ChatId, st.Msg)
			if err != nil {
				log.Println(err)
			}
		}
	}()

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

	msg := logic.ProcessMessagesCommand(message.Command(), message, model.PullUsers.Stage)

	_, err := model.PullUsers.GetUser(message.From.UserName)

	if err != nil {
		model.PullUsers.AddUser(message.From.ID, message.From.UserName, msg.Stage)
	}

	fmt.Println(model.PullUsers.Stage)
	fmt.Println(model.PullUsers.P)

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
			msg := logic.ProcessMessagesText(update.Message.Text, model.PullUsers.Stage)
			if msg.Stage != 0 && msg.Stage > model.PullUsers.Stage {
				model.PullUsers.IncStage(msg.Stage)
				b.Msgs <- "Игроки начали этап " + strconv.Itoa(model.PullUsers.Stage)
			}

			if msg.Type == logic.TypeImg {
				b.sendPhoto(update.Message, msg)
			}
			if msg.Type == logic.TypeImgs {
				b.sendPhotos(update.Message, msg)
			}
			if msg.Type == logic.TypeAudio {
				b.sendAudio(update.Message, msg)
			}
			if msg.Type == logic.TypeStr {
				b.SendTxtEveryOne(update.Message, msg)
			}
		}

		if update.Message.Location != nil && model.PullUsers.Stage == 2 {
			b.CheckLocation(update.Message)
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

	userIds := model.PullUsers.GetAllUserIds()

	if msg.ReferenceStartMsg != logic.ReqReference {
		for _, userId := range userIds {
			msgBot := tgbotapi.NewPhoto(userId, msgf)

			if _, err := b.bot.Send(msgBot); err != nil {
				b.Msgs <- "Ошибка отправки фото " + "@" + user.From.UserName + " " + err.Error()
			}

			b.SendTxt(userId, msg)
		}
	} else {
		msgBot := tgbotapi.NewPhoto(user.From.ID, msgf)

		if _, err := b.bot.Send(msgBot); err != nil {
			b.Msgs <- "Ошибка отправки фото " + "@" + user.From.UserName + " " + err.Error()
		}

		b.SendTxt(user.From.ID, msg)
	}
}

func (b *Bot) sendPhotos(user *tgbotapi.Message, msg logic.RespMsg) {

	userIds := model.PullUsers.GetAllUserIds()

	if msg.ReferenceStartMsg != logic.ReqReference {
		for _, userId := range userIds {
			b.SendTxt(userId, msg)

			for _, im := range msg.Images {
				fileBytes, er := ioutil.ReadFile(im)
				if er != nil {
					b.Msgs <- "Ошибка чтения файла фото " + "@" + user.From.UserName + " " + er.Error()
				}

				msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

				msgBot := tgbotapi.NewPhoto(userId, msgf)

				if _, err := b.bot.Send(msgBot); err != nil {
					b.Msgs <- "Ошибка отправки фото " + "@" + user.From.UserName + " " + err.Error()
				}
			}
		}
	} else {
		b.SendTxt(user.From.ID, msg)

		for _, im := range msg.Images {
			fileBytes, er := ioutil.ReadFile(im)
			if er != nil {
				b.Msgs <- "Ошибка чтения файла фото " + "@" + user.From.UserName + " " + er.Error()
			}

			msgf := tgbotapi.FileBytes{Name: "111", Bytes: fileBytes}

			msgBot := tgbotapi.NewPhoto(user.From.ID, msgf)

			if _, err := b.bot.Send(msgBot); err != nil {
				b.Msgs <- "Ошибка отправки фото " + "@" + user.From.UserName + " " + err.Error()
			}
		}
	}
}

func (b *Bot) sendAudio(user *tgbotapi.Message, msg logic.RespMsg) {

	fileBytes, er := ioutil.ReadFile(msg.FilePath)
	if er != nil {
		b.Msgs <- "Ошибка чтения файла аудио " + "@" + user.From.UserName + " " + er.Error()
	}

	msgf := tgbotapi.FileBytes{Name: msg.FileName, Bytes: fileBytes}

	userIds := model.PullUsers.GetAllUserIds()

	if msg.ReferenceStartMsg != logic.ReqReference {
		for _, userId := range userIds {
			msgBot := tgbotapi.NewAudio(userId, msgf)

			if _, err := b.bot.Send(msgBot); err != nil {
				b.Msgs <- "Ошибка отправки файла аудио " + "@" + user.From.UserName + " " + er.Error()
			}
		}
	} else {
		msgBot := tgbotapi.NewAudio(user.From.ID, msgf)

		if _, err := b.bot.Send(msgBot); err != nil {
			b.Msgs <- "Ошибка отправки файла аудио " + "@" + user.From.UserName + " " + er.Error()
		}
	}
}

func (b *Bot) SendTxt(userId int64, msg logic.RespMsg) {

	msgBot := tgbotapi.NewMessage(userId, msg.Message)
	msgBot.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки текстового сообщения " + "@" + strconv.Itoa(int(userId)) + " " + err.Error()
	}
}

func (b *Bot) SendTxtEveryOne(user *tgbotapi.Message, msg logic.RespMsg) {

	userIds := model.PullUsers.GetAllUserIds()

	if msg.ReferenceStartMsg != logic.ReqReference {
		for _, userId := range userIds {
			msgBot := tgbotapi.NewMessage(userId, msg.Message)
			msgBot.ParseMode = tgbotapi.ModeMarkdownV2

			if _, err := b.bot.Send(msgBot); err != nil {
				b.Msgs <- "Ошибка отправки текстового сообщения " + "@" + user.From.UserName + " " + err.Error()
			}
		}
	} else {
		msgBot := tgbotapi.NewMessage(user.From.ID, msg.Message)
		msgBot.ParseMode = tgbotapi.ModeMarkdownV2

		if _, err := b.bot.Send(msgBot); err != nil {
			b.Msgs <- "Ошибка отправки текстового сообщения " + "@" + user.From.UserName + " " + err.Error()
		}
	}
}

func (b *Bot) SendMsg(ChatID int64, input string) error {

	msg := tgbotapi.NewMessage(ChatID, input)
	//msg.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) CheckLocation(user *tgbotapi.Message) {

	var msgBot tgbotapi.MessageConfig

	usLat := user.Location.Latitude
	usLon := user.Location.Longitude
	refLat := refLat1
	refLon := refLon1

	if usLat < refLat+r20m && usLat > refLat-r20m &&
		usLon < refLon+r20m && usLon > refLon-r20m {

		b.sendAudio(user, logic.RespMsg{FilePath: logic.AudioFileShturval, FileName: logic.AudioFileShturvalName})
		return

	} else if usLat < refLat+r40m && usLat > refLat-r40m &&
		usLon < refLon+r40m && usLon > refLon-r40m {

		msgBot = tgbotapi.NewMessage(user.From.ID, msg40m)
	} else if usLat < refLat+r60m && usLat > refLat-r60m &&
		usLon < refLon+r60m && usLon > refLon-r60m {

		msgBot = tgbotapi.NewMessage(user.From.ID, msg60m)
	} else if usLat < refLat+r80m && usLat > refLat-r80m &&
		usLon < refLon+r80m && usLon > refLon-r80m {

		msgBot = tgbotapi.NewMessage(user.From.ID, msg80m)
	} else if usLat < refLat+r100m && usLat > refLat-r100m &&
		usLon < refLon+r100m && usLon > refLon-r100m {

		msgBot = tgbotapi.NewMessage(user.From.ID, msg100m)
	} else if usLat > refLat+r100m || usLat < refLat-r100m ||
		usLon > refLon+r100m || usLon < refLon-r100m {

		msgBot = tgbotapi.NewMessage(user.From.ID, msg10000m)
	}

	if _, err := b.bot.Send(msgBot); err != nil {
		b.Msgs <- "Ошибка отправки текстового сообщения " + "@" + user.From.UserName + " " + err.Error()
	}
}
