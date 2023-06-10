package commands

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) Default(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("User %v wrote message: %v", message.Chat.UserName, message.Text))
	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		args := strings.Split(update.CallbackQuery.Data, "_")

		// args := json.Unmarshal([]byte(update.CallbackQuery.Data))
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Command: %s\n", args[0])+
				fmt.Sprintf("Offset: %s", args[1]),
		)
		c.bot.Send(msg)
	}

	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "love":
		c.Love(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
