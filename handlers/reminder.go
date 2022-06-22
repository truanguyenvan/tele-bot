package handlers

import (
	"telegram-bot/utils"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron"
)

func PushNotiToChat(domain, token, chatId string) {
	log.Info("handlers -- ", "Remind Daily Report")
	header := "DAILY REPORT REMINDER"
	text := "Anh/Chị nhớ điền daily report và logwork trước khi về nhé!\nHave a good evening."
	sendMessage2Chat(domain, token, chatId, header, text, utils.Text2Emoij("info"))
}

func SetDailyReportReminder(domain, token, chatId string) {
	log.Info("handlers -- ", "Set Daily Report Reminder")
	c := cron.New()
	c.AddFunc("0 50 16 * * 1-5", func() { PushNotiToChat(domain, token, chatId) })
	c.Start()
}
