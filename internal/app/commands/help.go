package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"/help - help\n"+
			"/list - list something\n"+
			"/love - get love"))
	c.bot.Send(msg)
}
