package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if IsValidUUID(update.Message.Text) {
			values := map[string]string{"key": update.Message.Text}
			_, err := json.Marshal(values)

			if err != nil {
				log.Println(err)
				continue
			}

			resp, err := http.Get("http://192.168.1.69:8080/test")

			if err != nil {
				log.Println(err)
				continue
			}

			var res map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&res)

			if err != nil {
				log.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "valid uuid")
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
				continue
			}
		}
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)

	return err == nil
}
