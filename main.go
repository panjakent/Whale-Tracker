package main

import (
	"encoding/json"
	"fmt"
	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//type Prev struct {
//	Value int64 `json:"value"`
//}
//
//type Vin struct {
//	Prev Prev `json:"prevout"`
//}
//
//type Tx struct {
//	Txid string `json:"txid"`
//	Vin  []Vin  `json:"vin"`
//}

type PrevOut struct {
	Addr  string `json:"addr"`
	Value int64  `json:"value"`
}

type Input struct {
	PrevOut PrevOut `json:"prev_out"`
}

type Txs struct {
	Inputs []Input `json:"inputs"`
	Hash   string  `json:"hash"`
}

type Tx struct {
	Txs []Txs `json:"txs"`
}

func main() {

	// get api

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
			for range time.Tick(time.Second * 1) {
				resp, err := http.Get("https://blockchain.info/unconfirmed-transactions?format=json") // Ganti dengan URL API yang valid
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

				var transaction Tx

				err = json.Unmarshal(body, &transaction)
				if err != nil {
					log.Fatalf("Error unmarshaling response: %v", err)
				}

				//log.Printf("Response body: %s", string(body))
				//log.Printf("transaksi: %v sebesar: %v", transaction.Txid)
				//log.Printf("transaksi: %v sebesar: %v", transaction.Txid, transaction.Vin[0].Prev.Value)
				//log.Printf("transaksi: %v sebesar: %v", transaction.Txs[0].Inputs[0].Addr)
				if transaction.Txs[0].Inputs[0].PrevOut.Value >= 1000000000 {
					//fmt.Printf("Txid: %s\n", transaction.Txid, "Value: %s\n", transaction.Val)
					hash := transaction.Txs[0].Hash
					addr := transaction.Txs[0].Inputs[0].PrevOut.Addr
					value := transaction.Txs[0].Inputs[0].PrevOut.Value
					txLink := fmt.Sprintf("https://btcscan.org/tx/%s", hash)
					//msg := telegrambot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("transaksi: %v sebesar: %v", transaction.Txid, transaction.Vin[0].Prev.Value))

					msg := telegrambot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Txid: %s\nAddress: %s\nValue: %v\nLink: %s", hash, addr, value, txLink))

					log.Printf("Authorized on account %s", update.Message.Chat.ID)
					log.Printf("Authorized on account %s", update.Message.Text)
					log.Printf("Authorized on account %s", update.Message.MessageID)

					msg.ReplyToMessageID = update.Message.MessageID

					bot.Send(msg)
				}

			}
		}
	}

	//if len(transaction.Txs) > 0 && len(transaction.Txs[0].Inputs) > 0 {
	//	addr := transaction.Txs[0].Inputs[0].PrevOut.Addr
	//	log.Printf("Address: %s", addr)
	//} else {
	//	log.Printf("No data available")
	//}

}
