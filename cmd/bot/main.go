package main

import (
	"log"
	"os"

	"github.com/Feokrat/demo-tg-bot/internal/app/commands"
	"github.com/Feokrat/demo-tg-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "love":
			commander.Love(update.Message)
		case "list":
			commander.List(update.Message)
		case "get":
			commander.Get(update.Message)
		default:
			commander.Default(update.Message)
		}
	}
}
