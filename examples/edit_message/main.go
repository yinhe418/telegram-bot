package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/yinhe418/telegram-bot"
	"github.com/yinhe418/telegram-bot/models"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	m, errSend := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if errSend != nil {
		fmt.Printf("error sending message: %v\n", errSend)
		return
	}

	time.Sleep(time.Second * 2)

	_, errEdit := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    m.Chat.ID,
		MessageID: m.ID,
		Text:      "New Message!",
	})
	if errEdit != nil {
		fmt.Printf("error edit message: %v\n", errEdit)
		return
	}
}
