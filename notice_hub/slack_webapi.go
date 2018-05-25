package notice_hub

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
	AsUser   bool   `json:"as_user"`
}

type SlackWebAPI struct {
	ConfigFilePath string
}

func (w *SlackWebAPI) Notify(text string) {

	url := "https://slack.com/api/chat.postMessage"

	webapi_config := LoadSlackConfig(w.ConfigFilePath).WebAPI

	message := message{
		Text:     text,
		Channel:  webapi_config.Channel,
		Username: webapi_config.Name,
		AsUser:   false,
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("json parse error :", err)
	}

	headers := map[string]string{"Authorization": "Bearer " + webapi_config.Token}
	PostJson(url, jsonBytes, headers)
}
