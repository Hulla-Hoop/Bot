package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	c.productService.Delete(arg - 1)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Продукт удален")
	c.bot.Send(msg)
}
