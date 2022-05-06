package main

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/spf13/viper"

	"github.com/kritsby/telegram-bot/pkg/telegram"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs %s", err.Error())
	}
	bot, err := tgbotapi.NewBotAPI(viper.GetString("token"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
