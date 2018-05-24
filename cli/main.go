package main

import (
	"log"
	"os"

	"github.com/foresta/notice-hub/slack"
	"github.com/urfave/cli"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "notice-hub-cli"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "message, m",
			Value: "hello",
			Usage: "send message",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "None",
			Usage: "load config file",
		},
	}

	return app
}

func help() {
	app := App()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func callWebhook(msg string, config string) {
	webhook := &slack.Webhook{
		ConfigFilePath: config,
	}
	webhook.Notify(msg)
}

func callWebAPI(msg string, config string) {
	webapi := &slack.WebAPI{
		ConfigFilePath: config,
	}
	webapi.Notify(msg)
}

func main() {

	app := App()
	app.Action = func(c *cli.Context) error {

		msg := c.String("message")
		config := c.String("config")

		if config == "None" {
			help()
			os.Exit(1)
		}

		callWebAPI(msg, config)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
