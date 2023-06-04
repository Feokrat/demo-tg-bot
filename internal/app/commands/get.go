package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrongs args", id)
		return
	}

	product, err := c.productService.Get(id)
	if err != nil {
		log.Println(err)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("product of id %d is %s", id, product.Title))
	c.bot.Send(msg)
}
