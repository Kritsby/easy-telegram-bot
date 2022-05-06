package telegram

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const commandStart = "start"
const commandHistory = "history"
const commandHelp = "help"

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввёл команду /start"
		_, err := b.bot.Send(msg)
		return err
	case commandHistory:
		msg.Text = "Ты ввёл команду /history"
		_, err := b.bot.Send(msg)
		return err
	case commandHelp:
		msg.Text = "Доступные команды:\n /help\n /start\n /history"
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	_, err := b.bot.Send(msg)
	return err
}
