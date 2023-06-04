package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("TBD"))
	c.bot.Send(msg)
}
