package main

import (
	"log"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// accepts command and array of arguments
// returns the output of execution
// example: ExecCmd("ls", [])
func ExecCmd(cmd string, args []string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// main function
func main() {
    // change the token
	const TOKEN string = "SECRET-TOKEN"

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

    // verbose bot username
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
        // ignore any non-message updates
		if update.Message == nil { 
			continue
		}

		var cmds []string
		if strings.ContainsAny(update.Message.Text, "\n") {
            // if msg contains new-lines then split it by `\n` separator
			cmds = strings.Split(update.Message.Text, "\n")
		} else {
            // otherwise split it by space separator
			cmds = strings.Split(update.Message.Text, " ")
		}

        // verbose command and args
		log.Printf("ExecCmd(%s, %v)", cmds[0], cmds[1:])

        // Execute cmd by ExecCmd function
		out, err := ExecCmd(cmds[0], cmds[1:])
		if err != nil {
			out = "[!] ERROR:\n" + err.Error()
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
		msg.ReplyToMessageID = update.Message.MessageID
        // send cmd-output to user
		bot.Send(msg)
	}
}

