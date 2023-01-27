package commands

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Comander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	words := strings.Split(args, " ")

	arg, err := strconv.Atoi(words[0])
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	c.productService.Edit(arg-1, words[1])
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Продукт изменен")
	c.bot.Send(msg)
}
