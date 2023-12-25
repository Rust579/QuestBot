package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	// Messages

	//Command
	CommandGetCode = "Получить код"
	CommandStart   = "start"
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

/*func (b *Bot) SendCode(ChatID int64, Code string) error {

	msg := tgbotapi.NewMessage(ChatID, MessageConfirmCodeHead+fmt.Sprintf("\n`%v`\n", Code)+MessageConfirmCodePost)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) SendBindCode(ChatID int64, Code string) error {

	msg := tgbotapi.NewMessage(ChatID, MessageBindCode+fmt.Sprintf("\n`%v`\n", Code)+MessageConfirmCodePost)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) SendDone(ChatID int64) error {

	msg := tgbotapi.NewMessage(ChatID, MessageSuccessBind)

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}*/

func (b *Bot) HandleCommand(message *tgbotapi.Message) error {

	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandGetCode),
		),
	)
	if message.Command() == CommandStart {

		msg := tgbotapi.NewMessage(message.From.ID, "Здарова, ты попал в квест чувааааааак")
		msg.ReplyMarkup = numericKeyboard
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		if _, err := b.bot.Send(msg); err != nil {
			return err

		}
	}
	if message.Command() == "vvv" {

		msg := tgbotapi.NewMessage(message.From.ID, "www")
		msg.ReplyMarkup = numericKeyboard
		msg.ParseMode = tgbotapi.ModeMarkdownV2

		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
		b.msgs <- message.Command()
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
			}
			continue
		}

		if update.Message.Text == CommandGetCode {

			msg := tgbotapi.NewMessage(update.Message.From.ID, "MessageAlreadyExist")
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			if _, err := b.bot.Send(msg); err != nil {
			}
			continue

		}
		if update.Message.Text == "пук" {

			msg := tgbotapi.NewMessage(update.Message.From.ID, "Сам ты пук")
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			if _, err := b.bot.Send(msg); err != nil {
			}
			continue

		}

	}
}
