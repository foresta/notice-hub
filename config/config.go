package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Slack SlackConfig
}

type SlackConfig struct {
	Webhook WebhookConfig
	WebAPI  WebAPIConfig
}

type WebhookConfig struct {
	Url     string
	Name    string
	Channel string
}

type WebAPIConfig struct {
	Token   string
	Name    string
	Channel string
}

func LoadSlackConfig(filepath string) SlackConfig {
	config := load(filepath)
	return config.Slack
}

func load(filepath string) Config {
	var config Config

	_, err := toml.DecodeFile(filepath, &config)
	if err != nil {
		fmt.Println("config load error: ", err)
	}

	return config
}
