package slack

import (
	"encoding/json"
	"fmt"

	"github.com/foresta/notice-hub/config"
	"github.com/foresta/notice-hub/util"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

type Webhook struct {
	ConfigFilePath string
}

func (w *Webhook) Notify(text string) {

	webhook_config := config.LoadSlackConfig(w.ConfigFilePath).Webhook

	message := Message{
		Username: webhook_config.Name,
		Text:     text,
		Channel:  webhook_config.Channel}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("json parse error :", err)

		return
	}

	headers := map[string]string{}
	util.PostJson(webhook_config.Url, jsonBytes, headers)
}
