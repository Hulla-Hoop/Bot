package main

import (
	"log"

	"os"

	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shamil/bot/internal/app/commands"
	"github.com/shamil/bot/internal/service/product"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(filepath.Join(".env"))
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	Comander := commands.NewComander(bot, productService)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		switch update.Message.Command() {
		case "edit":
			Comander.Edit(update.Message)
		case "delete":
			Comander.Delete(update.Message)
		case "new":
			Comander.New(update.Message)
		case "get":
			Comander.Get(update.Message)
		case "help":
			Comander.Help(update.Message)
		case "list":
			Comander.List(update.Message)
		default:
			Comander.Default(update.Message)
		}
	}
}
