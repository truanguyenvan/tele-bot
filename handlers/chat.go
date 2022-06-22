package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"telegram-bot/models"
	"telegram-bot/utils"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func removeMessage(domain, token, chatId, messageId string) error {
	httpRequest := fmt.Sprintf(`%s/bot%s/deleteMessage?chat_id=%s&message_id=%s`, domain, token, chatId, messageId)

	resp, err := http.Get(httpRequest)

	if err != nil {
		log.Debug("handlers -- ", err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Debug("handlers -- ", err)
		return err
	}
	bodyJson := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &bodyJson)
	if err != nil {
		log.Debug("handlers -- ", err)
		return err
	}

	if !bodyJson["ok"].(bool) {
		err = fmt.Errorf(fmt.Sprint(bodyJson["description"]))
		log.Debug("handlers -- ", err)
		return err
	}
	return nil
}

func sendMessage2Chat(domain, token, chatId, header, text, emoij string) (interface{}, error) {
	text = utils.Text2Emoij(emoij) + utils.GenerateHeader(header) + text
	text = strings.ReplaceAll(text, "NT", "\n")
	text = strings.ReplaceAll(text, "TT", "\t")

	httpRequest := fmt.Sprintf(`%s/bot%s/sendMessage?chat_id=%s&parse_mode=Markdown&text=`, domain, token, chatId) + url.QueryEscape(text)

	resp, err := http.Get(httpRequest)

	if err != nil {
		log.Debug("handlers -- ", err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Debug("handlers -- ", err)
		return nil, err
	}
	bodyJson := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &bodyJson)
	if err != nil {
		log.Debug("handlers -- ", err)
		return nil, err
	}
	if !bodyJson["ok"].(bool) {
		err = fmt.Errorf(fmt.Sprint(bodyJson["description"]))
		log.Debug("handlers -- ", err)
		return nil, err
	}

	resultJson, oke := bodyJson["result"].(map[string]interface{})
	if !oke {
		return nil, err
	}

	messageId, oke := resultJson["message_id"].(float64)
	if oke {
		go func() {
			a := viper.GetInt64("telegram.auto_remove_time")
			time.Sleep(time.Minute * time.Duration(a))
			removeMessage(domain, token, chatId, fmt.Sprintf("%f", messageId))
		}()
	}

	return bodyJson["result"], nil
}

func SendMessageHandler(sendData *models.SendMessage) (interface{}, error) {
	domain := viper.GetString("telegram.domain")
	token := viper.GetString("telegram.bot_token")
	resp, err := sendMessage2Chat(domain, token, sendData.Chat_id, sendData.Header, sendData.Text, sendData.Emoji)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
