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
}

type WebhookConfig struct {
	Name    string
	Url     string
	Channel string
}

func LoadSlackConfig() SlackConfig {
	config := load()
	return config.Slack
}

func load() Config {
	var config Config

	_, err := toml.DecodeFile("./config/config.toml", &config)
	if err != nil {
		fmt.Println("config load error: ", err)
	}

	return config
}
