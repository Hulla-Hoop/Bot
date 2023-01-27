package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) List(inputMessage *tgbotapi.Message) {
	soobsh := "Наши продукты"
	product := c.productService.List()
	for _, p := range product {
		soobsh += p.Title
		soobsh += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, soobsh)
	c.bot.Send(msg)
}
