package main

//
//import (
//	"github.com/catsworld/qq-bot-api"
//	"log"
//	"net/http"
//)
//
//func main() {
//	bot, err := qqbotapi.NewBotAPI("", "http://cqhttp:5700", "")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	bot.Debug = true
//
//	u := qqbotapi.NewWebhook("/webhook_endpoint")
//	u.PreloadUserInfo = true
//
//	// use WebHook as event method
//	updates := bot.ListenForWebhook(u)
//	// or if you love WebSocket Reverse
//	// updates := bot.ListenForWebSocket(u)
//	go http.ListenAndServe("0.0.0.0:8085", nil)
//
//	for update := range updates {
//		if update.Message == nil {
//			continue
//		}
//
//		log.Printf("[%s] %s", update.Message.From.String(), update.Message.Text)
//
//		msg := qqbotapi.NewMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)
//		bot.Send(msg)
//	}
//}
