package tgmanager

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"tgbot/internal/model"
)

const token = "6694897672:AAEkkL38aHyei2_YkeOYN47D12bgwMkGIHA"
const chatId = -4081879081

var TelegramApi *CoreTgApi

type CoreTgApi struct {
	bot    *tgbotapi.BotAPI
	ChatID int64
	Msgs   chan string
	SendTo chan model.SendTo
}

func InitTgApi(msgs chan string, sendTo chan model.SendTo) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = false

	TelegramApi = &CoreTgApi{
		bot:    bot,
		ChatID: chatId,
		Msgs:   msgs,
		SendTo: sendTo,
	}

	go TelegramApi.Start()

	return nil
}

func (b *CoreTgApi) Start() {

	go func() {
		for msg := range b.Msgs {
			b.SendMsg(b.ChatID, msg)
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

func (b *CoreTgApi) HandleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			text := strings.Split(update.Message.Text, " ")
			if len(text) >= 3 && text[0] == "send" {
				alias := text[1]
				msg := strings.Join(text[2:], " ")

				if alias == "all" {
					usersIds := model.PullUsers.GetAllUserIds()
					if len(usersIds) == 0 {
						err := b.SendMsg(update.Message.From.ID, "can not send to "+alias+": not in pulls")
						if err != nil {
							log.Println("Can not send msg to chat")
						}
						continue
					}
					for _, cid := range usersIds {
						b.SendTo <- model.SendTo{
							ChatId: cid,
							Msg:    msg,
						}
					}
					continue
				}

				userChatId, err := model.PullUsers.GetUser(alias)
				if err != nil {
					err = b.SendMsg(update.Message.From.ID, "can not send to "+alias+": not in pulls")
					if err != nil {
						log.Println("Can not send msg to chat")
					}
					continue
				}

				b.SendTo <- model.SendTo{
					ChatId: userChatId,
					Msg:    msg,
				}
			}
		}
	}
}

func (b *CoreTgApi) SendMsg(ChatID int64, input string) error {

	msg := tgbotapi.NewMessage(ChatID, input)
	//msg.ParseMode = tgbotapi.ModeMarkdownV2

	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
