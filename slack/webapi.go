package slack

import (
	"encoding/json"
	"fmt"

	"github.com/foresta/notice-hub/config"
	"github.com/foresta/notice-hub/util"
)

type message struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
	AsUser   bool   `json:"as_user"`
}

type WebAPI struct {
	ConfigFilePath string
}

func (w *WebAPI) Notify(text string) {

	url := "https://slack.com/api/chat.postMessage"

	webapi_config := config.LoadSlackConfig(w.ConfigFilePath).WebAPI

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
	util.PostJson(url, jsonBytes, headers)
}
