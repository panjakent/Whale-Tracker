package main

import (
	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// get api
	resp, err := http.Get("https://btcscan.org/api/tx/066af0bd204d3a4e923d6a94e720738020a288a564d74661dee3fbd54dbc91f7") // Ganti dengan URL API yang valid
	if err != nil {
		//log.Fatalf("Error during GET request: %v", err)
		log.Printf("Error during GET request: %v", err)
	}
	defer resp.Body.Close()

	// Membaca body dari response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	log.Printf("Response body: %s", string(body))

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
