package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/foresta/notice-hub/config"
)

type Message struct {
	Name    string `json:"username"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

type Webhook struct {
	ConfigFilePath string
}

func (w *Webhook) Notify(text string) {

	webhook_config := config.LoadSlackConfig(w.ConfigFilePath).Webhook

	message := Message{
		Name:    webhook_config.Name,
		Text:    text,
		Channel: webhook_config.Channel}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println("json parse error :", err)

		return
	}

	postJson(webhook_config.Url, jsonBytes)
}

func postJson(url string, jsonBytes []byte) {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(jsonBytes),
	)

	if err != nil {
		fmt.Println("request create error :", err)

		return
	}

	req.Header.Set("ContentType", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request error: ", err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}
