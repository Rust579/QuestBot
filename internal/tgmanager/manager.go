package tgmanager

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const CommandStart = "start"
const CommandGetCode = "cas"
const token = "6694897672:AAEkkL38aHyei2_YkeOYN47D12bgwMkGIHA"
const chatId = -4081879081

var TelegramApi *CoreTgApi

type CoreTgApi struct {
	bot    *tgbotapi.BotAPI
	ChatID int64
	msgs   chan string
}

func InitTgApi(msgs chan string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = false

	TelegramApi = &CoreTgApi{
		bot:    bot,
		ChatID: chatId,
		msgs:   msgs,
	}

	go TelegramApi.Start()

	return nil
}

func (b *CoreTgApi) Start() {

	go func() {
		for msg := range b.msgs {
			b.SendCode(b.ChatID, msg)
		}
	}()

	updates := b.initUpdatesChan()
	b.HandleUpdates(updates)

	return
}

func (b *CoreTgApi) initUpdatesChan() tgbotapi.UpdatesChannel {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *CoreTgApi) SendCode(ChatID int64, input string) error {

	msg := tgbotapi.NewMessage(ChatID, input)
	//msg.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *CoreTgApi) HandleCommand(message *tgbotapi.Message) error {

	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandGetCode),
		),
	)
	if message.Command() == CommandStart {

		msg := tgbotapi.NewMessage(b.ChatID, "Здарова")
		msg.ReplyMarkup = numericKeyboard
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		if _, err := b.bot.Send(msg); err != nil {
			return err

		}
	}

	return nil
}

func (b *CoreTgApi) HandleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.HandleCommand(update.Message); err != nil {
			}
			continue
		}

		if update.Message.Text == CommandGetCode {

			msg := tgbotapi.NewMessage(b.ChatID, "MessageAlreadyExist")
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			if _, err := b.bot.Send(msg); err != nil {
			}
			continue

		}
	}
}
