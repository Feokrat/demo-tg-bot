package main

import (
	"fmt"
	"log"
	"os"

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

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "help":
			processHelp(bot, update.Message)
		case "love":
			processLove(bot, update.Message)
		case "list":
			processList(bot, update.Message, productService)
		default:
			defaultBehaviour(bot, update.Message)
		}

	}
}

func processHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(
		"/help - help\n"+
			"/list - list something\n"+
			"/love - get love"))
	bot.Send(msg)
}

func processLove(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("I love you, %v", message.Chat.UserName))
	bot.Send(msg)
}

func processList(bot *tgbotapi.BotAPI, message *tgbotapi.Message, productService *product.Service) {
	outputMessageText := "All products: \n\n"
	products := productService.List()

	for _, p := range products {
		outputMessageText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, outputMessageText)
	bot.Send(msg)
}

func defaultBehaviour(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("User %v wrote message: %v", message.Chat.UserName, message.Text))
	bot.Send(msg)
}
