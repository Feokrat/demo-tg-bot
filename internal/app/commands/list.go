package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(message *tgbotapi.Message) {
	outputMessageText := "All products: \n\n"
	products := c.productService.List()

	for _, p := range products {
		outputMessageText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, outputMessageText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "list_10"),
		),
	)

	c.bot.Send(msg)
}
