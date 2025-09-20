package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func SendMessage(chatID int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	sentMsg, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
	return sentMsg, err
}

func RunTelegram() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_ID"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Println("Telegram bot is listening")
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() { // ignore any non-command Messages
			msg, err := SendMessage(update.Message.Chat.ID, "You said: "+update.Message.Text)
			if err != nil {
				log.Printf("Error sending echo message: %v", err)
			}
			log.Printf("Echoed message: %s", msg.Text)
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Welcome!"
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."

		default:
			msg.Text = "I don't know that command"
		}

		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			// TODO
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
}
