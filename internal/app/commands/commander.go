package commands

import (
	"github.com/Feokrat/demo-tg-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot             *tgbotapi.BotAPI
	produuctService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:             bot,
		produuctService: productService,
	}
}
