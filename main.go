package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/imankulov/pt-numbers/converter"
)

func main() {

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("Set valid BOT_TOKEN variable")
	}

	bot, err := api.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal("Unable to create a bot: ", err)
	}
	bot.Debug = true

	// webhookURL is a full URL to webhook, starting with https, etc.
	webhookURL := os.Getenv("WEBHOOK_URL")
	webhookURLObj, err := url.Parse(webhookURL)
	if err != nil {
		log.Fatal("Set valid WEBHOOK_URL environment variable: ", err)
	}

	_, err = bot.SetWebhook(api.NewWebhook(webhookURL))
	if err != nil {
		log.Fatal("Unable to set a webhook: ", err)
	}

	// Start listening for updates
	updates := bot.ListenForWebhook(webhookURLObj.Path)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "9090"
	}

	go http.ListenAndServe(fmt.Sprintf(":%s", httpPort), nil)

	for update := range updates {
		text := update.Message.Text

		msg := api.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyToMessageID = update.Message.MessageID

		if strings.HasPrefix(text, "/start") || strings.HasPrefix(text, "/help") {
			msg.Text = "Send a raw numeric value to the bot, and it sends back a string " +
				"representation of the number in Portuguese."
		} else {
			n, err := strconv.Atoi(text)
			if err != nil {
				msg.Text = "Not a number"
			} else {
				msg.Text = converter.Do(n)
			}
		}
		bot.Send(msg)
	}
}
