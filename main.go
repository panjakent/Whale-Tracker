package main

import (
	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	// get api

	//log
	log.Printf("s")

	// Initialize the TelegramBot
	bot, err := telegrambot.NewBotAPI("7453307755:AAHTcuZ4qc9ADGFfQmh8t-JI3Hl02Htg9N8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegrambot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := telegrambot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
