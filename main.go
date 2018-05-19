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

func main() {

	app := App()
	app.Action = func(c *cli.Context) error {

		msg := c.String("message")
		config := c.String("config")

		if config == "None" {
			help()
			os.Exit(1)
		}

		webhook.Notify(msg, config)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
