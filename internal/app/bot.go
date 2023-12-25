package app

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	// Messages
	MessageStartHead       = "Для успешной привязки вашего аккаунта Telegram к аккаунту UniOne, вам необходимо ввести код\n"
	MessageStartPost       = "\nв соответствующую форму на платформе\\. После ввода кода, ваш аккаунт Telegram будет надежно связан с аккаунтом UniOne, что позволит вам использовать его для авторизации на платформе\\."
	MessageSuccessBind     = "Ваш аккаунт Telegram успешно привязан к аккаунту UniOne. Чтобы отвязать аккаунт Telegram заблокируйте этого бота и авторизуйтесь на платформе по коду."
	MessageConfirmCodeHead = "*Код подтверждения для входа в UniOne:*\n"
	MessageConfirmCodePost = "\n*Никому* не давайте код, даже если его требуют от имени Telegram\\! Этот код используется для входа в Ваш аккаунт в UniOne\\. Он никогда не нужен для чего\\-то еще\\. Если Вы не запрашивали код для входа, проигнорируйте это сообщение\\."
	MessageAlreadyExist    = "Ваш аккаунт Telegram уже привязан к аккаунту UniOne\\. Чтобы отвязать аккаунт Telegram заблокируйте этого бота и авторизуйтесь на платформе по коду\\."
	MessageBindCode        = "*Код подтверждения для привязки аккаунта Telegram в UniOne:*\n"

	//Command
	CommandGetCode = "Получить код"
	CommandStart   = "start"
)

type Bot struct {
	bot tgbotapi.BotAPI
}

var BotApi Bot

func InitBotApi(token string) error {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = false
	BotApi.bot = *bot

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

func (b *Bot) SendCode(ChatID int64, Code string) error {

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
}

func (b *Bot) HandleCommand(message *tgbotapi.Message) error {

	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(CommandGetCode),
		),
	)
	if message.Command() == CommandStart {
		ok := pullsTelegram.CheckExistTelegramId(message.From.ID)
		if !ok {
			msg := tgbotapi.NewMessage(message.From.ID, MessageStartHead+fmt.Sprintf("\n`%v`\n", message.Chat.ID)+MessageStartPost)
			msg.ReplyMarkup = numericKeyboard
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			if _, err := b.bot.Send(msg); err != nil {
				return err
			}
		} else {
			msg := tgbotapi.NewMessage(message.From.ID, MessageAlreadyExist)
			msg.ReplyMarkup = numericKeyboard
			msg.ParseMode = tgbotapi.ModeMarkdownV2

			if _, err := b.bot.Send(msg); err != nil {
				return err
			}
		}
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
			ok := pullsTelegram.CheckExistTelegramId(update.Message.From.ID)
			if !ok {
				msg := tgbotapi.NewMessage(update.Message.From.ID, MessageStartHead+fmt.Sprintf("\n`%v`\n", update.Message.Chat.ID)+MessageStartPost)
				msg.ParseMode = tgbotapi.ModeMarkdownV2

				if _, err := b.bot.Send(msg); err != nil {
				}
				continue
			} else {
				msg := tgbotapi.NewMessage(update.Message.From.ID, MessageAlreadyExist)
				msg.ParseMode = tgbotapi.ModeMarkdownV2

				if _, err := b.bot.Send(msg); err != nil {
				}
				continue
			}
		}
	}
}
