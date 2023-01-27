package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	c.productService.New(args)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Продукт добавлен")
	c.bot.Send(msg)
}
