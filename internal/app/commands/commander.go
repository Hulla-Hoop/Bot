package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shamil/bot/internal/service/product"
)

type Comander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewComander(bot *tgbotapi.BotAPI, productService *product.Service) *Comander {
	return &Comander{
		bot:            bot,
		productService: productService,
	}
}
