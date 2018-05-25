package notice_hub

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

type SlackWebhook struct {
	ConfigFilePath string
}

func (w *SlackWebhook) Notify(text string) {

	webhook_config := LoadSlackConfig(w.ConfigFilePath).Webhook

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
	PostJson(webhook_config.Url, jsonBytes, headers)
}
