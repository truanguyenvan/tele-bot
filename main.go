package main

import (
	"telegram-bot/handlers"
	"telegram-bot/router"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json") // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.SetLevel(log.INFO)
	err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// set daily report reminder
	go handlers.SetDailyReportReminder(viper.GetString("telegram.domain"), viper.GetString("telegram.bot_token"), viper.GetString("telegram.default_chat_id"))
	// create a new echo instance
	e := router.New()
	e.Logger.Fatal(e.Start(":5000"))
}
